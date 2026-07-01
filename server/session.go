package server

import (
	"context"
	"fmt"
	"net"

	"github.com/johnmaddison/ocpp-go"
	"github.com/johnmaddison/ocpp-go/ocpp16"
	"github.com/johnmaddison/ocpp-go/ocpp21"
)

// Session represents one currently connected charge point.
type Session struct {
	chargePointID string
	protocol      string
	remoteAddr    net.Addr
	localAddr     net.Addr
	ocpp16Context *ocpp16.Context
	ocpp21Context *ocpp21.Context
}

// Session returns the currently addressable session for a charge point ID.
func (s *Server) Session(chargePointID string) (*Session, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	session, ok := s.sessions[chargePointID]
	return session, ok
}

// ChargePointID returns the charge point ID associated with the session.
func (s *Session) ChargePointID() string {
	return s.chargePointID
}

// Protocol returns the negotiated OCPP websocket subprotocol.
func (s *Session) Protocol() string {
	return s.protocol
}

// RemoteAddr returns the remote websocket address for the session.
func (s *Session) RemoteAddr() net.Addr {
	return s.remoteAddr
}

// LocalAddr returns the local websocket address for the session.
func (s *Session) LocalAddr() net.Addr {
	return s.localAddr
}

// Context16 returns the OCPP 1.6 context for this session, if applicable.
func (s *Session) Context16() (*ocpp16.Context, bool) {
	return s.ocpp16Context, s.ocpp16Context != nil
}

// Context21 returns the OCPP 2.1 context for this session, if applicable.
func (s *Session) Context21() (*ocpp21.Context, bool) {
	return s.ocpp21Context, s.ocpp21Context != nil
}

// Send16 sends an OCPP 1.6 CALL to the connected charge point.
func (s *Session) Send16(call ocpp.Call) (*ocpp16.ResultOrError, error) {
	if s.ocpp16Context == nil {
		return nil, s.wrongProtocolError("ocpp1.6")
	}
	return s.ocpp16Context.Send(call)
}

// Send16Call sends a typed OCPP 1.6 CALL to the connected charge point.
func (s *Session) Send16Call(action ocpp16.Action, payload any) (*ocpp16.ResultOrError, error) {
	if s.ocpp16Context == nil {
		return nil, s.wrongProtocolError("ocpp1.6")
	}
	return s.ocpp16Context.SendCall(action, payload)
}

// Send16WithContext sends an OCPP 1.6 CALL with context cancellation.
func (s *Session) Send16WithContext(ctx context.Context, call ocpp.Call) (*ocpp16.ResultOrError, error) {
	if s.ocpp16Context == nil {
		return nil, s.wrongProtocolError("ocpp1.6")
	}
	return s.ocpp16Context.SendWithContext(ctx, call)
}

// Send16CallWithContext sends a typed OCPP 1.6 CALL with context cancellation.
func (s *Session) Send16CallWithContext(ctx context.Context, action ocpp16.Action, payload any) (*ocpp16.ResultOrError, error) {
	if s.ocpp16Context == nil {
		return nil, s.wrongProtocolError("ocpp1.6")
	}
	return s.ocpp16Context.SendCallWithContext(ctx, action, payload)
}

// Send21 sends an OCPP 2.1 CALL to the connected charge point.
func (s *Session) Send21(call ocpp.Call) (*ocpp21.ResultOrError, error) {
	if s.ocpp21Context == nil {
		return nil, s.wrongProtocolError("ocpp2.1")
	}
	return s.ocpp21Context.Send(call)
}

// Send21Call sends a typed OCPP 2.1 CALL to the connected charge point.
func (s *Session) Send21Call(action ocpp21.Action, payload any) (*ocpp21.ResultOrError, error) {
	if s.ocpp21Context == nil {
		return nil, s.wrongProtocolError("ocpp2.1")
	}
	return s.ocpp21Context.SendCall(action, payload)
}

// Send21WithContext sends an OCPP 2.1 CALL with context cancellation.
func (s *Session) Send21WithContext(ctx context.Context, call ocpp.Call) (*ocpp21.ResultOrError, error) {
	if s.ocpp21Context == nil {
		return nil, s.wrongProtocolError("ocpp2.1")
	}
	return s.ocpp21Context.SendWithContext(ctx, call)
}

// Send21CallWithContext sends a typed OCPP 2.1 CALL with context cancellation.
func (s *Session) Send21CallWithContext(ctx context.Context, action ocpp21.Action, payload any) (*ocpp21.ResultOrError, error) {
	if s.ocpp21Context == nil {
		return nil, s.wrongProtocolError("ocpp2.1")
	}
	return s.ocpp21Context.SendCallWithContext(ctx, action, payload)
}

func (s *Session) wrongProtocolError(want string) error {
	return fmt.Errorf("session %s uses %s, not %s", s.chargePointID, s.protocol, want)
}

func (s *Server) registerSession(session *Session) {
	s.mu.Lock()
	if s.sessions == nil {
		s.sessions = make(map[string]*Session)
	}
	s.sessions[session.chargePointID] = session
	s.mu.Unlock()
}

func (s *Server) unregisterSession(session *Session) {
	s.mu.Lock()
	if s.sessions[session.chargePointID] == session {
		delete(s.sessions, session.chargePointID)
	}
	s.mu.Unlock()
}
