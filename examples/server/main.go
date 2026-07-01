package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/JohnMaddison/ocpp-go"
	"github.com/JohnMaddison/ocpp-go/ocpp16"
	"github.com/JohnMaddison/ocpp-go/server"
	"github.com/google/uuid"
)

// var debug = flag.Bool("debug", false, "enable debug logging")
var profiler = flag.Bool("profile", false, "enable profiler")
var addr = flag.String("addr", "0.0.0.0:9001", "listen address in host:port form")
var tlsCert = flag.String("tls-cert", "", "path to TLS certificate file (enables TLS when set)")
var tlsKey = flag.String("tls-key", "", "path to TLS private key file (enables TLS when set)")
var basicUser = flag.String("basic-user", "", "basic auth username (enables basic auth when set together with -basic-pass)")
var basicPass = flag.String("basic-pass", "", "basic auth password (enables basic auth when set together with -basic-user)")

func main() {

	flag.Parse()

	// Handle graceful shutdown
	_, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigChan
		log.Print("Received interrupt signal, shutting down gracefully...")
		// Cancel context to signal shutdown
		cancel()
		log.Print("Shutdown complete")
		os.Exit(0)
	}()

	if *profiler {
		go func() {
			log.Println("Starting pprof server on :6060")
			log.Println(http.ListenAndServe("localhost:6060", nil))
		}()
	}

	opts := []server.Option{
		server.WithPath("ocpp"),
		server.WithOCPP16AuthorizeHandler(authorizeHandler),
		server.WithOCPP16BootNotificationHandler(bootNotificationHandler),
		server.WithOCPP16HeartbeatHandler(heartbeatHandler),
		server.WithOCPP16MeterValuesHandler(meterValuesHandler),
		server.WithOCPP16StartTransactionHandler(startTransactionHandler),
		server.WithOCPP16StatusNotificationHandler(statusNotificationHandler),
		server.WithOCPP16StopTransactionHandler(stopTransactionHandler),
		server.WithConnectRequestHandler(ConnectHandler),
		server.WithConnectedHandler(ConnectedHandler),
		server.WithDisconnectHandler(DisconnectHandler),
		server.WithTrafficLogging(),
		server.WithKeepaliveLogging(),
		server.WithWebsocketKeepalive(30*time.Second, 45*time.Second),
	}
	if *tlsCert != "" && *tlsKey != "" {
		opts = append(opts, server.WithTLS(*tlsCert, *tlsKey))
	}
	if *basicUser != "" && *basicPass != "" {
		opts = append(opts, server.WithBasicAuth(*basicUser, *basicPass))
	}
	srv := server.NewServer(*addr, opts...)

	log.Printf("Starting server")
	err := srv.Serve()
	if err != nil {
		log.Printf("Failed to start server %s", err)
		return
	}
}

func ConnectHandler(ctx ocpp.ConnectRequest) bool {
	log.Printf("Allowed connection from %s, Headers: %s", ctx.R.PathValue("cpid"), ctx.R.Header)
	return true
}
func ConnectedHandler(info ocpp.ConnectionInfo) {
	log.Printf("Connected: %s", info.ChargePointID)
}

func DisconnectHandler(info ocpp.ConnectionInfo) {
	log.Printf("Disconnected: %s", info.ChargePointID)
}

func bootNotificationHandler(ctx *ocpp16.OCPPContext, request ocpp16.BootNotificationRequest) (*ocpp16.BootNotificationResponse, *ocpp16.OCPPError) {
	log.Printf("Bootnotifcation received from charge point: %s %+v\n",
		ctx.ChargePointID, request)

	go func() {
		time.Sleep(2 * time.Second)
		c := ocpp.Call{MessageType: ocpp.MessageTypeCall, MessageID: uuid.New().String(), Action: "GetConfiguration", Payload: ocpp16.GetConfigurationRequest{}}
		log.Printf("Sending getconfig: %s\n", ctx.ChargePointID)
		request, err := ctx.Send(c)
		if err != nil {
			log.Printf("Failed to retrieve config %s", err)
			return
		}

		if request.IsCallError() {
			log.Printf("Received callerror: %s %+v\n", ctx.ChargePointID, request.CallError)
			return
		}
		if request.IsCallResult() {
			log.Printf("Received callresult: %s %+v\n", ctx.ChargePointID, request.CallResult)
			return
		}
	}()
	return &ocpp16.BootNotificationResponse{
		CurrentTime: time.Now(),
		Status:      "Accepted",
		Interval:    10,
	}, nil
}

func heartbeatHandler(ctx *ocpp16.OCPPContext, request ocpp16.HeartbeatRequest) (*ocpp16.HeartbeatResponse, *ocpp16.OCPPError) {
	//log.Printf("Heartbeat received from charge point: %s %+v\n", ctx.ChargePointID, request)

	return &ocpp16.HeartbeatResponse{
		CurrentTime: time.Now().UTC(),
	}, nil
}

func statusNotificationHandler(ctx *ocpp16.OCPPContext, request ocpp16.StatusNotificationRequest) (*ocpp16.StatusNotificationResponse, *ocpp16.OCPPError) {
	//log.Printf("StatusNotification received from charge point: %s %+v\n", ctx.ChargePointID, request)
	return &ocpp16.StatusNotificationResponse{}, nil
}

func authorizeHandler(ctx *ocpp16.OCPPContext, request ocpp16.AuthorizeRequest) (*ocpp16.AuthorizeResponse, *ocpp16.OCPPError) {
	return &ocpp16.AuthorizeResponse{IDTagInfo: ocpp16.IDTagInfo{Status: ocpp16.AuthorizationStatusAccepted}}, nil
}

func startTransactionHandler(ctx *ocpp16.OCPPContext, request ocpp16.StartTransactionRequest) (*ocpp16.StartTransactionResponse, *ocpp16.OCPPError) {
	return &ocpp16.StartTransactionResponse{IDTagInfo: ocpp16.IDTagInfo{Status: ocpp16.AuthorizationStatusAccepted}, TransactionID: 10000}, nil
}

func meterValuesHandler(ctx *ocpp16.OCPPContext, request ocpp16.MeterValuesRequest) (*ocpp16.MeterValuesResponse, *ocpp16.OCPPError) {
	return &ocpp16.MeterValuesResponse{}, nil
}

func stopTransactionHandler(ctx *ocpp16.OCPPContext, request ocpp16.StopTransactionRequest) (*ocpp16.StopTransactionResponse, *ocpp16.OCPPError) {
	return &ocpp16.StopTransactionResponse{IDTagInfo: &ocpp16.IDTagInfo{Status: ocpp16.AuthorizationStatusAccepted}}, nil
}
