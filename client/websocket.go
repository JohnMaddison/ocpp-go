package client

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/gorilla/websocket"
	"github.com/johnmaddison/ocpp-go/internal/ws"
	"github.com/johnmaddison/ocpp-go/ocpp16"
	"github.com/johnmaddison/ocpp-go/ocpp21"
)

// Connect opens the websocket connection and starts the read/write pumps.
func (c *Client) Connect() error {
	return c.ConnectContext(context.Background())
}

// ConnectContext opens the websocket connection and starts the read/write
// pumps. The context controls the websocket handshake. Disconnect may be used
// concurrently to cancel the handshake or close an established connection.
func (c *Client) ConnectContext(ctx context.Context) error {
	if ctx == nil {
		ctx = context.Background()
	}

	c.mu.Lock()
	if c.connection != nil {
		c.mu.Unlock()
		return ErrAlreadyConnected
	}

	runtime, err := c.runtime()
	if err != nil {
		c.mu.Unlock()
		return err
	}
	dialCtx, cancel := context.WithCancel(ctx)
	state := &connectionState{cancel: cancel, done: make(chan struct{})}
	c.connection = state
	address := c.address
	chargePointID := c.chargePointID
	username := c.username
	password := c.password
	subprotocol := c.subprotocol
	logTraffic := c.logTraffic
	logKeepalive := c.logKeepalive
	pingInterval := c.pingInterval
	pongTimeout := c.pongTimeout
	socketCallbacks := c.socketCallbacks
	c.mu.Unlock()

	websocketURL, err := appendChargePointID(address, chargePointID)
	if err != nil {
		c.finishConnection(state)
		return err
	}

	headers := http.Header{}
	if username != "" || password != "" {
		auth := username + ":" + password
		authHeader := "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))
		headers.Add("Authorization", authHeader)
	}

	dialer := websocket.Dialer{
		Subprotocols: []string{subprotocol}, EnableCompression: true, ReadBufferSize: 2048, WriteBufferSize: 2048, HandshakeTimeout: 10 * time.Second,
	}
	netDialer := &net.Dialer{}
	dialer.NetDialContext = func(ctx context.Context, network, address string) (net.Conn, error) {
		conn, err := netDialer.DialContext(ctx, network, address)
		if err != nil {
			return nil, err
		}
		c.mu.Lock()
		if c.connection != state || ctx.Err() != nil {
			c.mu.Unlock()
			_ = conn.Close()
			if err := ctx.Err(); err != nil {
				return nil, err
			}
			return nil, errors.New("ocpp client connection was canceled")
		}
		state.dialConn = conn
		c.mu.Unlock()
		return conn, nil
	}
	dialFinished := make(chan struct{})
	go func() {
		select {
		case <-dialCtx.Done():
			c.closeDialConnection(state)
		case <-dialFinished:
		}
	}()

	conn, resp, err := dialer.DialContext(dialCtx, websocketURL, headers)
	close(dialFinished)
	if err != nil {
		c.finishConnection(state)
		if resp != nil {
			return fmt.Errorf("dial failed: %v, HTTP status: %v", err, resp.Status)
		}
		return fmt.Errorf("dial failed: %v", err)
	}

	c.mu.Lock()
	if c.connection != state || dialCtx.Err() != nil {
		c.mu.Unlock()
		_ = conn.Close()
		c.finishConnection(state)
		if err := dialCtx.Err(); err != nil {
			return fmt.Errorf("dial canceled: %w", err)
		}
		return errors.New("ocpp client connection was canceled")
	}
	state.conn = conn
	state.dialConn = nil
	c.mu.Unlock()

	go func() {
		ws.Run(conn, runtime, socketCallbacks, &ws.Options{
			LogSent:      logTraffic,
			LogKeepalive: logKeepalive,
			PingInterval: pingInterval,
			PongTimeout:  pongTimeout,
		})
		c.finishConnection(state)
	}()
	return nil
}

// Disconnect cancels an in-progress handshake or gracefully closes an active
// websocket. It is safe to call concurrently and more than once.
func (c *Client) Disconnect() error {
	c.mu.Lock()
	state := c.connection
	if state == nil {
		c.mu.Unlock()
		return nil
	}
	state.cancel()
	dialConn := state.dialConn
	conn := state.conn
	done := state.done
	c.mu.Unlock()

	var closeErr error
	if dialConn != nil {
		closeErr = dialConn.Close()
	}
	if conn != nil {
		deadline := time.Now().Add(time.Second)
		closeErr = conn.WriteControl(websocket.CloseMessage,
			websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""), deadline)
		if err := conn.Close(); closeErr == nil {
			closeErr = err
		}
	}
	<-done
	if errors.Is(closeErr, net.ErrClosed) || errors.Is(closeErr, websocket.ErrCloseSent) {
		return nil
	}
	return closeErr
}

func (c *Client) closeDialConnection(state *connectionState) {
	c.mu.Lock()
	conn := state.dialConn
	c.mu.Unlock()
	if conn != nil {
		_ = conn.Close()
	}
}

func (c *Client) finishConnection(state *connectionState) {
	state.cancel()
	c.mu.Lock()
	if c.connection == state {
		c.connection = nil
	}
	c.mu.Unlock()
	select {
	case <-state.done:
	default:
		close(state.done)
	}
}

func appendChargePointID(address, chargePointID string) (string, error) {
	parsed, err := url.Parse(address)
	if err != nil {
		return "", fmt.Errorf("invalid websocket address: %w", err)
	}
	if parsed.Scheme == "" || parsed.Host == "" {
		return "", errors.New("invalid websocket address: scheme and host are required")
	}

	escapedPath := strings.TrimRight(parsed.EscapedPath(), "/") + "/" + url.PathEscape(chargePointID)
	path, err := url.PathUnescape(escapedPath)
	if err != nil {
		return "", fmt.Errorf("invalid websocket path: %w", err)
	}
	parsed.Path = path
	parsed.RawPath = escapedPath
	return parsed.String(), nil
}

func (c *Client) runtime() (ws.Runtime, error) {
	switch c.subprotocol {
	case "ocpp1.6":
		c.Context16 = ocpp16.NewContextWithMessageIDGenerator(c.chargePointID, c.messageIDGenerator)
		return ws.Runtime{
			ChargePointID: c.Context16.ChargePointID,
			Protocol:      c.subprotocol,
			OutgoingCalls: requestChannel[ocpp16.Request](c.Context16.Queue),
			Parse: func(message []byte) ([]byte, error) {
				return c.ocppCallbacks.ParseMessage(message, c.Context16)
			},
			Serialize: func(call any) ([]byte, error) {
				request := call.(ocpp16.Request)
				return request.Call.SerializeOCPP()
			},
		}, nil
	case "ocpp2.1":
		c.Context21 = ocpp21.NewContextWithMessageIDGenerator(c.chargePointID, c.messageIDGenerator)
		return ws.Runtime{
			ChargePointID: c.Context21.ChargePointID,
			Protocol:      c.subprotocol,
			OutgoingCalls: requestChannel[ocpp21.Request](c.Context21.Queue),
			Parse: func(message []byte) ([]byte, error) {
				return c.ocpp21Callbacks.ParseMessage(message, c.Context21)
			},
			Serialize: func(call any) ([]byte, error) {
				request := call.(ocpp21.Request)
				return request.Call.SerializeOCPP()
			},
		}, nil
	default:
		return ws.Runtime{}, fmt.Errorf("unsupported OCPP subprotocol: %s", c.subprotocol)
	}
}

func requestChannel[T any](in <-chan T) func(done <-chan struct{}) <-chan any {
	return func(done <-chan struct{}) <-chan any {
		out := make(chan any)
		go func() {
			defer close(out)
			for {
				select {
				case item, ok := <-in:
					if !ok {
						return
					}
					select {
					case out <- item:
					case <-done:
						return
					}
				case <-done:
					return
				}
			}
		}()
		return out
	}
}
