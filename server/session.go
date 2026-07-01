package server

import (
	"context"
	"fmt"
	"net"

	"github.com/JohnMaddison/ocpp-go"
	"github.com/JohnMaddison/ocpp-go/ocpp16"
	"github.com/JohnMaddison/ocpp-go/ocpp21"
)

// Session represents one currently connected charge point.
type Session struct {
	chargePointID string
	protocol      string
	remoteAddr    net.Addr
	localAddr     net.Addr
	ocpp16Context *ocpp16.OCPPContext
	ocpp21Context *ocpp21.OCPPContext
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

// OCPP16Context returns the OCPP 1.6 context for this session, if applicable.
func (s *Session) OCPP16Context() (*ocpp16.OCPPContext, bool) {
	return s.ocpp16Context, s.ocpp16Context != nil
}

// OCPP21Context returns the OCPP 2.1 context for this session, if applicable.
func (s *Session) OCPP21Context() (*ocpp21.OCPPContext, bool) {
	return s.ocpp21Context, s.ocpp21Context != nil
}

// SendOCPP16 sends an OCPP 1.6 CALL to the connected charge point.
func (s *Session) SendOCPP16(call ocpp.Call) (*ocpp16.ResultOrError, error) {
	if s.ocpp16Context == nil {
		return nil, s.wrongProtocolError("ocpp1.6")
	}
	return s.ocpp16Context.Send(call)
}

// SendOCPP16Call sends a typed OCPP 1.6 CALL to the connected charge point.
func (s *Session) SendOCPP16Call(action ocpp16.Action, payload any) (*ocpp16.ResultOrError, error) {
	if s.ocpp16Context == nil {
		return nil, s.wrongProtocolError("ocpp1.6")
	}
	return s.ocpp16Context.SendCall(action, payload)
}

// SendOCPP16WithContext sends an OCPP 1.6 CALL with context cancellation.
func (s *Session) SendOCPP16WithContext(ctx context.Context, call ocpp.Call) (*ocpp16.ResultOrError, error) {
	if s.ocpp16Context == nil {
		return nil, s.wrongProtocolError("ocpp1.6")
	}
	return s.ocpp16Context.SendWithContext(ctx, call)
}

// SendOCPP16CallWithContext sends a typed OCPP 1.6 CALL with context cancellation.
func (s *Session) SendOCPP16CallWithContext(ctx context.Context, action ocpp16.Action, payload any) (*ocpp16.ResultOrError, error) {
	if s.ocpp16Context == nil {
		return nil, s.wrongProtocolError("ocpp1.6")
	}
	return s.ocpp16Context.SendCallWithContext(ctx, action, payload)
}

// SendOCPP21 sends an OCPP 2.1 CALL to the connected charge point.
func (s *Session) SendOCPP21(call ocpp.Call) (*ocpp21.ResultOrError, error) {
	if s.ocpp21Context == nil {
		return nil, s.wrongProtocolError("ocpp2.1")
	}
	return s.ocpp21Context.Send(call)
}

// SendOCPP21Call sends a typed OCPP 2.1 CALL to the connected charge point.
func (s *Session) SendOCPP21Call(action ocpp21.Action, payload any) (*ocpp21.ResultOrError, error) {
	if s.ocpp21Context == nil {
		return nil, s.wrongProtocolError("ocpp2.1")
	}
	return s.ocpp21Context.SendCall(action, payload)
}

// SendOCPP21WithContext sends an OCPP 2.1 CALL with context cancellation.
func (s *Session) SendOCPP21WithContext(ctx context.Context, call ocpp.Call) (*ocpp21.ResultOrError, error) {
	if s.ocpp21Context == nil {
		return nil, s.wrongProtocolError("ocpp2.1")
	}
	return s.ocpp21Context.SendWithContext(ctx, call)
}

// SendOCPP21CallWithContext sends a typed OCPP 2.1 CALL with context cancellation.
func (s *Session) SendOCPP21CallWithContext(ctx context.Context, action ocpp21.Action, payload any) (*ocpp21.ResultOrError, error) {
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
