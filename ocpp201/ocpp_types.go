// Package ocpp201 contains the data structures for OCPP 2.0.1 messages.
package ocpp201

import "time"

type AdditionalInfo struct {
	AdditionalIDToken string `json:"additionalIdToken" validate:"required,max=36"`
	Type              string `json:"type" validate:"required,max=50"`
}

type APN struct {
	APN                     string                `json:"apn" validate:"required,max=512"`
	APNAuthentication       APNAuthenticationType `json:"apnAuthentication" validate:"required"`
	APNPassword             *string               `json:"apnPassword,omitempty" validate:"max=20"`
	APNUserName             *string               `json:"apnUserName,omitempty" validate:"max=20"`
	PreferredNetwork        *string               `json:"preferredNetwork,omitempty" validate:"max=6"`
	UseOnlyPreferredNetwork bool                  `json:"useOnlyPreferredNetwork,omitempty"`
}

type CertificateHashData struct {
	HashAlgorithm  HashAlgorithm `json:"hashAlgorithm" validate:"required"`
	IssuerNameHash string        `json:"issuerNameHash" validate:"required,max=128"`
	IssuerKeyHash  string        `json:"issuerKeyHash" validate:"required,max=128"`
	SerialNumber   string        `json:"serialNumber" validate:"required,max=40"`
}

type CertificateHashDataChain struct {
	CertificateHashData      CertificateHashData   `json:"certificateHashData" validate:"required"`
	CertificateType          GetCertificateIdUse   `json:"certificateType" validate:"required"`
	ChildCertificateHashData []CertificateHashData `json:"childCertificateHashData,omitempty"`
}

type ChargingLimit struct {
	ChargingLimitSource string `json:"chargingLimitSource" validate:"required,max=20"`
	IsGridCritical      *bool  `json:"isGridCritical,omitempty"`
}

type ChargingNeeds struct {
	ACChargingParameters    *ACChargingParameters `json:"acChargingParameters,omitempty"`
	DCChargingParameters    *DCChargingParameters `json:"dcChargingParameters,omitempty"`
	RequestedEnergyTransfer EnergyTransferMode    `json:"requestedEnergyTransfer" validate:"required"`
	DepartureTime           *time.Time            `json:"departureTime,omitempty"`
}

type ACChargingParameters struct {
	EnergyAmount int `json:"energyAmount" validate:"required"`
	EVMinCurrent int `json:"evMinCurrent" validate:"required"`
	EVMaxCurrent int `json:"evMaxCurrent" validate:"required"`
	EVMaxVoltage int `json:"evMaxVoltage" validate:"required"`
}

type DCChargingParameters struct {
	EVMaxCurrent     int  `json:"evMaxCurrent" validate:"required"`
	EVMaxVoltage     int  `json:"evMaxVoltage" validate:"required"`
	EnergyAmount     *int `json:"energyAmount,omitempty"`
	EVMaxPower       *int `json:"evMaxPower,omitempty"`
	StateOfCharge    *int `json:"stateOfCharge,omitempty"`
	EVEnergyCapacity *int `json:"evEnergyCapacity,omitempty"`
	FullSoC          *int `json:"fullSoC,omitempty"`
	BulkSoC          *int `json:"bulkSoC,omitempty"`
}

type ChargingProfile struct {
	ID                     int                    `json:"id" validate:"required"`
	StackLevel             int                    `json:"stackLevel" validate:"required"`
	ChargingProfilePurpose ChargingProfilePurpose `json:"chargingProfilePurpose" validate:"required"`
	ChargingProfileKind    ChargingProfileKind    `json:"chargingProfileKind" validate:"required"`
	RecurrencyKind         *RecurrencyKind        `json:"recurrencyKind,omitempty"`
	ValidFrom              *time.Time             `json:"validFrom,omitempty"`
	ValidTo                *time.Time             `json:"validTo,omitempty"`
	ChargingSchedule       []ChargingSchedule     `json:"chargingSchedule" validate:"required,min=1"`
	TransactionID          *string                `json:"transactionId,omitempty" validate:"max=36"`
}

type ChargingSchedule struct {
	ID                     int                      `json:"id" validate:"required"`
	StartSchedule          *time.Time               `json:"startSchedule,omitempty"`
	Duration               *int                     `json:"duration,omitempty"`
	ChargingRateUnit       ChargingRateUnit         `json:"chargingRateUnit" validate:"required"`
	ChargingSchedulePeriod []ChargingSchedulePeriod `json:"chargingSchedulePeriod" validate:"required,min=1"`
	MinChargingRate        *float64                 `json:"minChargingRate,omitempty"`
	SalesTariff            *SalesTariff             `json:"salesTariff,omitempty"`
}

type ChargingSchedulePeriod struct {
	StartPeriod  int     `json:"startPeriod" validate:"required"`
	Limit        float64 `json:"limit" validate:"required"`
	NumberPhases *int    `json:"numberPhases,omitempty"`
	PhaseToUse   *int    `json:"phaseToUse,omitempty"`
}

type ChargingStation struct {
	Model           string  `json:"model" validate:"required,max=20"`
	VendorName      string  `json:"vendorName" validate:"required,max=50"`
	SerialNumber    *string `json:"serialNumber,omitempty" validate:"max=25"`
	FirmwareVersion *string `json:"firmwareVersion,omitempty" validate:"max=50"`
	Modem           *Modem  `json:"modem,omitempty"`
}

type ClearMonitoringResult struct {
	Status     ClearMonitoringStatus `json:"status" validate:"required"`
	ID         int                   `json:"id" validate:"required"`
	StatusInfo *StatusInfo           `json:"statusInfo,omitempty"`
}

type Component struct {
	Name     string  `json:"name" validate:"required,max=50"`
	Instance *string `json:"instance,omitempty" validate:"max=50"`
	Evse     *EVSE   `json:"evse,omitempty"`
}

type ComponentVariable struct {
	Component Component `json:"component" validate:"required"`
	Variable  *Variable `json:"variable,omitempty"`
}

type CompositeSchedule struct {
	ChargingSchedulePeriod []ChargingSchedulePeriod `json:"chargingSchedulePeriod" validate:"required,min=1"`
	EvseID                 int                      `json:"evseId" validate:"required"`
	Duration               int                      `json:"duration" validate:"required"`
	ScheduleStart          time.Time                `json:"scheduleStart" validate:"required"`
	ChargingRateUnit       ChargingRateUnit         `json:"chargingRateUnit" validate:"required"`
}

type ConsumptionCost struct {
	StartValue float64 `json:"startValue" validate:"required"`
	Cost       []Cost  `json:"cost" validate:"required,min=1"`
}

type Cost struct {
	CostKind         CostKind `json:"costKind" validate:"required"`
	Amount           int      `json:"amount" validate:"required"`
	AmountMultiplier *int     `json:"amountMultiplier,omitempty"`
}

type EVSE struct {
	ID          int  `json:"id" validate:"required,gt=0"`
	ConnectorID *int `json:"connectorId,omitempty" validate:"gt=0"`
}

type EventData struct {
	EventID           int               `json:"eventId" validate:"required"`
	Timestamp         time.Time         `json:"timestamp" validate:"required"`
	Trigger           EventTrigger      `json:"trigger" validate:"required"`
	Cause             *int              `json:"cause,omitempty"`
	ActualValue       string            `json:"actualValue" validate:"required,max=2500"`
	TechCode          *string           `json:"techCode,omitempty" validate:"max=50"`
	TechInfo          *string           `json:"techInfo,omitempty" validate:"max=500"`
	Cleared           *bool             `json:"cleared,omitempty"`
	TransactionID     *string           `json:"transactionId,omitempty" validate:"max=36"`
	Component         Component         `json:"component" validate:"required"`
	Variable          Variable          `json:"variable" validate:"required"`
	EventNotification EventNotification `json:"eventNotificationType" validate:"required"`
}

type Firmware struct {
	Location           string     `json:"location" validate:"required,max=512"`
	RetrieveDateTime   time.Time  `json:"retrieveDateTime" validate:"required"`
	InstallDateTime    *time.Time `json:"installDateTime,omitempty"`
	SigningCertificate *string    `json:"signingCertificate,omitempty" validate:"max=5500"`
	Signature          *string    `json:"signature,omitempty" validate:"max=800"`
}

type GetVariableData struct {
	Component     Component      `json:"component" validate:"required"`
	Variable      Variable       `json:"variable" validate:"required"`
	AttributeType *AttributeType `json:"attributeType,omitempty"`
}

type IdToken struct {
	IDToken        string           `json:"idToken" validate:"required,max=36"`
	Type           IdTokenType      `json:"type" validate:"required"`
	AdditionalInfo []AdditionalInfo `json:"additionalInfo,omitempty"`
}

type IdTokenInfo struct {
	Status              AuthorizationStatus `json:"status" validate:"required"`
	CacheExpiryDateTime *time.Time          `json:"cacheExpiryDateTime,omitempty"`
	ChargingPriority    *int                `json:"chargingPriority,omitempty"`
	Language1           *string             `json:"language1,omitempty" validate:"max=8"`
	EvseID              []int               `json:"evseId,omitempty"`
	GroupIDToken        *IdToken            `json:"groupIdToken,omitempty"`
	Language2           *string             `json:"language2,omitempty" validate:"max=8"`
	PersonalMessage     *MessageContent     `json:"personalMessage,omitempty"`
}

type LogParameters struct {
	RemoteLocation  string     `json:"remoteLocation" validate:"required,max=512"`
	OldestTimestamp *time.Time `json:"oldestTimestamp,omitempty"`
	LatestTimestamp *time.Time `json:"latestTimestamp,omitempty"`
}

type MessageContent struct {
	Format   MessageFormat `json:"format" validate:"required"`
	Language *string       `json:"language,omitempty" validate:"max=8"`
	Content  string        `json:"content" validate:"required,max=512"`
}

type MessageInfo struct {
	ID            int             `json:"id" validate:"required"`
	Priority      MessagePriority `json:"priority" validate:"required"`
	State         *MessageState   `json:"state,omitempty"`
	StartDateTime *time.Time      `json:"startDateTime,omitempty"`
	EndDateTime   *time.Time      `json:"endDateTime,omitempty"`
	TransactionID *string         `json:"transactionId,omitempty" validate:"max=36"`
	Message       MessageContent  `json:"message" validate:"required"`
	Display       *Component      `json:"display,omitempty"`
}

type MeterValue struct {
	Timestamp    time.Time      `json:"timestamp" validate:"required"`
	SampledValue []SampledValue `json:"sampledValue" validate:"required,min=1"`
}

type Modem struct {
	Iccid *string `json:"iccid,omitempty" validate:"max=20"`
	Imsi  *string `json:"imsi,omitempty" validate:"max=20"`
}

type MonitoringData struct {
	Component          Component            `json:"component" validate:"required"`
	Variable           Variable             `json:"variable" validate:"required"`
	VariableMonitoring []VariableMonitoring `json:"variableMonitoring" validate:"required,min=1"`
}

type NetworkConnectionProfile struct {
	OCPPVersion     string `json:"ocppVersion" validate:"required,max=12"`
	OCPPTransport   string `json:"ocppTransport" validate:"required,max=12"`
	OCPPCSMSURL     string `json:"ocppCsmsUrl" validate:"required,max=512"`
	MessageTimeout  int    `json:"messageTimeout" validate:"required"`
	SecurityProfile int    `json:"securityProfile" validate:"required"`
	OCPPInterface   string `json:"ocppInterface" validate:"required,max=8"`
	VPN             *VPN   `json:"vpn,omitempty"`
	APN             *APN   `json:"apn,omitempty"`
}

type OCSPRequestData struct {
	HashAlgorithm  HashAlgorithm `json:"hashAlgorithm" validate:"required"`
	IssuerNameHash string        `json:"issuerNameHash" validate:"required,max=128"`
	IssuerKeyHash  string        `json:"issuerKeyHash" validate:"required,max=128"`
	SerialNumber   string        `json:"serialNumber" validate:"required,max=40"`
	ResponderURL   string        `json:"responderURL" validate:"required,max=512"`
}

type RelativeTimeInterval struct {
	Start    int  `json:"start" validate:"required"`
	Duration *int `json:"duration,omitempty"`
}

type SalesTariff struct {
	ID                     int                `json:"id" validate:"required"`
	SalesTariffDescription *string            `json:"salesTariffDescription,omitempty" validate:"max=32"`
	NumEPriceLevels        *int               `json:"numEPriceLevels,omitempty"`
	SalesTariffEntry       []SalesTariffEntry `json:"salesTariffEntry" validate:"required,min=1"`
}

type SalesTariffEntry struct {
	RelativeTimeInterval RelativeTimeInterval `json:"relativeTimeInterval" validate:"required"`
	EPriceLevel          *int                 `json:"ePriceLevel,omitempty"`
	ConsumptionCost      []ConsumptionCost    `json:"consumptionCost,omitempty,min=1"`
}

type SampledValue struct {
	Value           float64          `json:"value" validate:"required"`
	Context         *ReadingContext  `json:"context,omitempty"`
	Measurand       *Measurand       `json:"measurand,omitempty"`
	Phase           *Phase           `json:"phase,omitempty"`
	Location        *Location        `json:"location,omitempty"`
	SignedMeterData *SignedMeterData `json:"signedMeterData,omitempty"`
	UnitOfMeasure   *UnitOfMeasure   `json:"unitOfMeasure,omitempty"`
}

type SetVariableData struct {
	AttributeType  *AttributeType `json:"attributeType,omitempty"`
	AttributeValue string         `json:"attributeValue" validate:"required,max=1000"`
	Component      Component      `json:"component" validate:"required"`
	Variable       Variable       `json:"variable" validate:"required"`
}

type SignedMeterData struct {
	SignedMeterData string `json:"signedMeterData" validate:"required,max=2500"`
	SigningMethod   string `json:"signingMethod" validate:"required,max=50"`
	EncodingMethod  string `json:"encodingMethod" validate:"required,max=50"`
	PublicKey       string `json:"publicKey" validate:"required,max=2500"`
}

type StatusInfo struct {
	ReasonCode     string  `json:"reasonCode" validate:"required,max=20"`
	AdditionalInfo *string `json:"additionalInfo,omitempty" validate:"max=512"`
}

type Transaction struct {
	ID                string         `json:"id" validate:"required,max=36"`
	ChargingState     *ChargingState `json:"chargingState,omitempty"`
	TimeSpentCharging *int           `json:"timeSpentCharging,omitempty"`
	StoppedReason     *Reason        `json:"stoppedReason,omitempty"`
	RemoteStartID     *int           `json:"remoteStartId,omitempty"`
}

type UnitOfMeasure struct {
	Unit       *string `json:"unit,omitempty" validate:"max=20"`
	Multiplier *int    `json:"multiplier,omitempty"`
}

type Variable struct {
	Name     string  `json:"name" validate:"required,max=50"`
	Instance *string `json:"instance,omitempty" validate:"max=50"`
}

type VariableAttribute struct {
	Type       *AttributeType `json:"type,omitempty"`
	Value      *string        `json:"value,omitempty" validate:"max=2500"`
	Mutability string         `json:"mutability,omitempty" validate:"max=16"`
	Persistent bool           `json:"persistent,omitempty"`
	Constant   bool           `json:"constant,omitempty"`
}

type VariableCharacteristics struct {
	Unit               *string  `json:"unit,omitempty" validate:"max=16"`
	DataType           string   `json:"dataType" validate:"required,max=16"`
	MinLimit           *float64 `json:"minLimit,omitempty"`
	MaxLimit           *float64 `json:"maxLimit,omitempty"`
	ValuesList         *string  `json:"valuesList,omitempty" validate:"max=1000"`
	SupportsMonitoring bool     `json:"supportsMonitoring" validate:"required"`
}

type VariableMonitoring struct {
	ID          int     `json:"id" validate:"required"`
	Severity    int     `json:"severity" validate:"required"`
	Type        Monitor `json:"type" validate:"required"`
	Value       float64 `json:"value" validate:"required"`
	Transaction *bool   `json:"transaction,omitempty"`
}

type VPN struct {
	Server   string  `json:"server" validate:"required,max=512"`
	User     string  `json:"user" validate:"required,max=20"`
	Group    *string `json:"group,omitempty" validate:"max=20"`
	Password string  `json:"password" validate:"required,max=20"`
	Key      string  `json:"key" validate:"required,max=512"`
	Type     VPNType `json:"type" validate:"required"`
}
