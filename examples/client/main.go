package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/JohnMaddison/ocpp-go"
	"github.com/JohnMaddison/ocpp-go/client"
	"github.com/JohnMaddison/ocpp-go/ocpp16"
)

// var debug = flag.Bool("debug", false, "enable debug logging")
var profiler = flag.Bool("profile", false, "enable profiler")

var model = "Simulator"
var vendor = "OCPP-GO"

func main() {
	flag.Parse()

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

	ChargePointID := "CP_000001"
	//Create client
	cp := client.NewExtensibleOcppClient(ChargePointID, "ws://127.0.0.1:9001/ocpp",
		ocpp16.OcppCallbacks{GetConfiguration: getConfiguarationHandler},
		ocpp.SocketCallbacks{Connected: ConnectedHandler, Disconnect: DisconnectHandler}).
		WithWebsocketKeepalive(30*time.Second, 45*time.Second).
		EnableKeepaliveLogging().
		EnableTrafficLogging()

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

	bootNotificationResponseInterval := 10

	// Send bootnotification untill we get accepted
	for true {
		request, err := cp.SendCall(ocpp16.ActionBootNotification, ocpp16.BootNotificationRequest{ChargePointModel: model, ChargePointVendor: vendor})

		if err != nil {
			log.Printf("Failed to send bootnotification %s", err)
		}

		if request.IsCallError() {
			log.Printf("Received callerror: %s %+v\n", ChargePointID, request.CallError)
			time.Sleep(time.Duration(10) * time.Second)
		} else if request.IsCallResult() {

			bootNotificationResponse, _ := request.GetPayload().(ocpp16.BootNotificationResponse)

			if bootNotificationResponse.Status == ocpp16.RegistrationStatusAccepted {
				log.Printf("Recieved accepted boot response")
				bootNotificationResponseInterval = bootNotificationResponse.Interval
				break
			} else {
				log.Printf("Recieved non accepted boot response")
				time.Sleep(time.Duration(bootNotificationResponse.Interval) * time.Second)
			}

		}
	}

	cp.SendCall(ocpp16.ActionStatusNotification, ocpp16.StatusNotificationRequest{ConnectorId: 0, ErrorCode: ocpp16.ChargePointErrorCodeNoError, Status: ocpp16.ChargePointStatusAvailable})
	cp.SendCall(ocpp16.ActionStatusNotification, ocpp16.StatusNotificationRequest{ConnectorId: 1, ErrorCode: ocpp16.ChargePointErrorCodeNoError, Status: ocpp16.ChargePointStatusAvailable})
	cp.SendCall(ocpp16.ActionStatusNotification, ocpp16.StatusNotificationRequest{ConnectorId: 2, ErrorCode: ocpp16.ChargePointErrorCodeNoError, Status: ocpp16.ChargePointStatusAvailable})

	chargingConnectorId := 1
	cp.SendCall(ocpp16.ActionStatusNotification, ocpp16.StatusNotificationRequest{ConnectorId: chargingConnectorId, ErrorCode: ocpp16.ChargePointErrorCodeNoError, Status: ocpp16.ChargePointStatusPreparing})

	tagIdentifier := "0000-0000-0001"
	authorizeRequest, err := cp.SendCall(ocpp16.ActionAuthorize, ocpp16.AuthorizeRequest{IdTag: tagIdentifier})

	if err != nil {
		log.Printf("Failed to send authorize request %s", err)
		return
	}

	if authorizeRequest.IsCallError() {
		log.Printf("Received callerror: %s %+v\n", ChargePointID, authorizeRequest.CallError)
	} else if authorizeRequest.IsCallResult() {

		authorizeResponse, _ := authorizeRequest.GetPayload().(ocpp16.AuthorizeResponse)

		if authorizeResponse.IdTagInfo.Status == ocpp16.AuthorizationStatusAccepted {
			//Tag is valid

			meterStart := 0
			startTransactionRequest, err := cp.SendCall(ocpp16.ActionStartTransaction, ocpp16.StartTransactionRequest{ConnectorId: chargingConnectorId, IdTag: tagIdentifier, MeterStart: meterStart, Timestamp: time.Now()})

			if err != nil {
				log.Printf("Failed to send startTransaction request %s", err)
				return
			}

			if startTransactionRequest.IsCallResult() {

				startTransactionResponse, _ := startTransactionRequest.GetPayload().(ocpp16.StartTransactionResponse)

				if startTransactionResponse.IdTagInfo.Status == ocpp16.AuthorizationStatusAccepted {

					cp.SendCall(ocpp16.ActionStatusNotification, ocpp16.StatusNotificationRequest{ConnectorId: chargingConnectorId, ErrorCode: ocpp16.ChargePointErrorCodeNoError, Status: ocpp16.ChargePointStatusCharging})

					meterValuesTicker := time.NewTicker(5 * time.Second)
					defer meterValuesTicker.Stop()

				Outer:
					for {
						select {
						case <-meterValuesTicker.C:
							context := ocpp16.ReadingContextSamplePeriodic
							measurand := ocpp16.MeasurandEnergyActiveImportRegister
							unit := "Wh"

							cp.SendCall(ocpp16.ActionMeterValues, ocpp16.MeterValuesRequest{
								ConnectorId:   chargingConnectorId,
								TransactionId: &startTransactionResponse.TransactionId,
								MeterValue: []ocpp16.MeterValue{
									{
										Timestamp: time.Now(),
										SampledValue: []ocpp16.SampledValue{
											{
												Value:     strconv.Itoa(meterStart),
												Context:   &context,
												Measurand: &measurand,
												Unit:      &unit,
											},
										},
									},
								},
							})

							meterStart = meterStart + 10

						case <-ctx.Done():
							log.Println("Stopping MeterValues loop")
							break Outer
						}
					}
				}
				reason := ocpp16.ReasonEVDisconnected
				_, err := cp.SendCall(ocpp16.ActionStopTransaction, ocpp16.StopTransactionRequest{IdTag: &tagIdentifier, MeterStop: meterStart, Timestamp: time.Now(), TransactionId: startTransactionResponse.TransactionId, Reason: &reason})

				if err != nil {
					log.Printf("Failed to send stopTransaction request %s", err)
				}

				cp.SendCall(ocpp16.ActionStatusNotification, ocpp16.StatusNotificationRequest{ConnectorId: chargingConnectorId, ErrorCode: ocpp16.ChargePointErrorCodeNoError, Status: ocpp16.ChargePointStatusFinishing})
				cp.SendCall(ocpp16.ActionStatusNotification, ocpp16.StatusNotificationRequest{ConnectorId: chargingConnectorId, ErrorCode: ocpp16.ChargePointErrorCodeNoError, Status: ocpp16.ChargePointStatusAvailable})
			}

		}

	}

	heartbeatTicker := time.NewTicker(time.Duration(bootNotificationResponseInterval) * time.Second)
	defer heartbeatTicker.Stop()

	for {
		select {
		case <-heartbeatTicker.C:
			cp.SendCall(ocpp16.ActionHeartbeat, ocpp16.EmptyPayload{})

		case <-ctx.Done():
			log.Print("Shutdown complete")
			return
		}
	}

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
	log.Printf("GetConfiguration request received from server: %s %+v\n", ctx.ChargePointID, request)

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
