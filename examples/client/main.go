package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/JohnMaddison/ocpp-go"
	"github.com/JohnMaddison/ocpp-go/client"
	"github.com/JohnMaddison/ocpp-go/ocpp16"
	"github.com/google/uuid"
)

// var debug = flag.Bool("debug", false, "enable debug logging")
var profiler = flag.Bool("profile", false, "enable profiler")

var model = "Simulator"
var vendor = "Maddison"

func main() {
	flag.Parse()

	//setupLogging()

	// Handle graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigChan
		log.Print("Received interrupt signal, shutting down gracefully...")
		// Cancel context to signal shutdown
		cancel()
	}()

	if *profiler {
		go func() {
			log.Println("Starting pprof server on :6060")
			log.Println(http.ListenAndServe("localhost:6060", nil))
		}()
	}

	ChargePointID := "CP_JOHN"
	//Create client
	cp := client.NewExtensibleOcppClient(ChargePointID, "ws://127.0.0.1:9001/ocpp",
		ocpp16.OcppCallbacks{GetConfiguration: getConfiguarationHandler},
		ocpp.SocketCallbacks{Connected: ConnectedHandler, Disconnect: DisconnectHandler}).
		WithWebsocketKeepalive(30*time.Second, 45*time.Second).
		EnableKeepaliveLogging(true).
		EnableTrafficLogging(true)

	// Loop websocket connect untill we succedd
	for true {

		err := cp.Connect()
		if err != nil {
			log.Printf("Failed to connect to server %v", err)
			time.Sleep(10 * time.Second)
		} else {
			break
		}

	}

	// Send bootnotification untill we get accepted
	for true {

		time.Sleep(2 * time.Second)
		c := ocpp.Call{MessageType: ocpp.MessageTypeCall, MessageID: uuid.New().String(), Action: "BootNotification", Payload: ocpp16.BootNotificationRequest{ChargePointModel: model, ChargePointVendor: vendor}}
		request, err := cp.Ocppcontext.Send(c)
		if err != nil {
			log.Printf("Failed to send bootnotification %s", err)
		}

		if request.CallError != nil {
			log.Printf("Received callerror: %s %+v\n", ChargePointID, request.CallError)
		}
		if request.CallResult != nil {

			result, _ := request.CallResult.Payload.(ocpp16.BootNotificationResponse)

			if result.Status == ocpp16.RegistrationStatusAccepted {
				log.Printf("Recieved accepted boot response")
				break
			} else if result.Status == ocpp16.RegistrationStatusRejected {
				log.Printf("Recieved rejected boot response")
				time.Sleep(time.Duration(result.Interval) * time.Second)
			}

		}
	}

	statusNotificationCall := ocpp.Call{MessageType: ocpp.MessageTypeCall, MessageID: uuid.New().String(), Action: string(ocpp16.ActionStatusNotification), Payload: ocpp16.StatusNotificationRequest{ConnectorId: 0, ErrorCode: ocpp16.ChargePointErrorCodeNoError, Status: ocpp16.ChargePointStatusAvailable}}
	cp.Ocppcontext.Send(statusNotificationCall)
	statusNotificationCall = ocpp.Call{MessageType: ocpp.MessageTypeCall, MessageID: uuid.New().String(), Action: string(ocpp16.ActionStatusNotification), Payload: ocpp16.StatusNotificationRequest{ConnectorId: 1, ErrorCode: ocpp16.ChargePointErrorCodeNoError, Status: ocpp16.ChargePointStatusAvailable}}
	cp.Ocppcontext.Send(statusNotificationCall)
	statusNotificationCall = ocpp.Call{MessageType: ocpp.MessageTypeCall, MessageID: uuid.New().String(), Action: string(ocpp16.ActionStatusNotification), Payload: ocpp16.StatusNotificationRequest{ConnectorId: 2, ErrorCode: ocpp16.ChargePointErrorCodeNoError, Status: ocpp16.ChargePointStatusAvailable}}
	cp.Ocppcontext.Send(statusNotificationCall)

	cp.SendCall(ocpp16.ActionHeartbeat, ocpp16.EmptyPayload{})
	<-ctx.Done()
	log.Print("Shutdown complete")

}

func ConnectedHandler() {
	//TODO ADD ARGUMENTS
	log.Printf("Connected")
}

func DisconnectHandler() {
	//TODO ADD ARGUMENTS
	log.Printf("Disconnected")
}

func getConfiguarationHandler(ctx *ocpp16.OcppContext, request ocpp16.GetConfigurationRequest) (*ocpp16.GetConfigurationResponse, *ocpp16.OcppError) {
	log.Printf("GetConfiguration received from server: %s %+v\n", ctx.ChargePointID, request)

	vendor := "Simulator"
	conf := []ocpp16.KeyValue{
		{
			Key:      "ChargePointVendor",
			Value:    &vendor,
			Readonly: true,
		},
		{
			Key:      "ChargePointModel",
			Value:    &model,
			Readonly: true,
		},
	}
	return &ocpp16.GetConfigurationResponse{ConfigurationKey: conf}, nil
}
