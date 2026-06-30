// Package ocpp21 contains the data structures for OCPP 2.1 messages.
package ocpp21

import "time"

// CustomData contains vendor-specific extension data.
type CustomData map[string]any

type ACChargingParameters struct {
	CustomData   CustomData `json:"customData,omitempty"`
	EnergyAmount float64    `json:"energyAmount" validate:"required"`
	EVMaxCurrent float64    `json:"evMaxCurrent" validate:"required"`
	EVMaxVoltage float64    `json:"evMaxVoltage" validate:"required"`
	EVMinCurrent float64    `json:"evMinCurrent" validate:"required"`
}

type APN struct {
	APN                     string                `json:"apn" validate:"required,max=2000"`
	APNAuthentication       APNAuthenticationEnum `json:"apnAuthentication" validate:"required"`
	APNPassword             *string               `json:"apnPassword,omitempty" validate:"max=64"`
	APNUserName             *string               `json:"apnUserName,omitempty" validate:"max=50"`
	CustomData              CustomData            `json:"customData,omitempty"`
	PreferredNetwork        *string               `json:"preferredNetwork,omitempty" validate:"max=6"`
	SIMPin                  *int                  `json:"simPin,omitempty"`
	UseOnlyPreferredNetwork *bool                 `json:"useOnlyPreferredNetwork,omitempty"`
}

type AbsolutePriceSchedule struct {
	AdditionalSelectedServices []AdditionalSelectedServices `json:"additionalSelectedServices,omitempty" validate:"min=1,max=5"`
	Currency                   string                       `json:"currency" validate:"required,max=3"`
	CustomData                 CustomData                   `json:"customData,omitempty"`
	Language                   string                       `json:"language" validate:"required,max=8"`
	MaximumCost                *RationalNumber              `json:"maximumCost,omitempty"`
	MinimumCost                *RationalNumber              `json:"minimumCost,omitempty"`
	OverstayRuleList           *OverstayRuleList            `json:"overstayRuleList,omitempty"`
	PriceAlgorithm             string                       `json:"priceAlgorithm" validate:"required,max=2000"`
	PriceRuleStacks            []PriceRuleStack             `json:"priceRuleStacks" validate:"required,min=1,max=1024"`
	PriceScheduleDescription   *string                      `json:"priceScheduleDescription,omitempty" validate:"max=160"`
	PriceScheduleID            int                          `json:"priceScheduleID" validate:"required,gte=0.0"`
	TaxRules                   []TaxRule                    `json:"taxRules,omitempty" validate:"min=1,max=10"`
	TimeAnchor                 time.Time                    `json:"timeAnchor" validate:"required"`
}

type AdditionalInfo struct {
	AdditionalIDToken string     `json:"additionalIdToken" validate:"required,max=255"`
	CustomData        CustomData `json:"customData,omitempty"`
	TypeValue         string     `json:"type" validate:"required,max=50"`
}

type AdditionalSelectedServices struct {
	CustomData  CustomData     `json:"customData,omitempty"`
	ServiceFee  RationalNumber `json:"serviceFee" validate:"required"`
	ServiceName string         `json:"serviceName" validate:"required,max=80"`
}

type Address struct {
	Address1   string     `json:"address1" validate:"required,max=100"`
	Address2   *string    `json:"address2,omitempty" validate:"max=100"`
	City       string     `json:"city" validate:"required,max=100"`
	Country    string     `json:"country" validate:"required,max=50"`
	CustomData CustomData `json:"customData,omitempty"`
	Name       string     `json:"name" validate:"required,max=50"`
	PostalCode *string    `json:"postalCode,omitempty" validate:"max=20"`
}

type AuthorizationData struct {
	CustomData  CustomData   `json:"customData,omitempty"`
	IDToken     IDToken      `json:"idToken" validate:"required"`
	IDTokenInfo *IDTokenInfo `json:"idTokenInfo,omitempty"`
}

type BatteryData struct {
	CustomData     CustomData `json:"customData,omitempty"`
	EVSEID         int        `json:"evseId" validate:"required,gte=0.0"`
	ProductionDate *time.Time `json:"productionDate,omitempty"`
	SerialNumber   string     `json:"serialNumber" validate:"required,max=50"`
	SoC            float64    `json:"soC" validate:"required,gte=0.0,lte=100.0"`
	SoH            float64    `json:"soH" validate:"required,gte=0.0,lte=100.0"`
	VendorInfo     *string    `json:"vendorInfo,omitempty" validate:"max=500"`
}

type CertificateHashData struct {
	CustomData     CustomData        `json:"customData,omitempty"`
	HashAlgorithm  HashAlgorithmEnum `json:"hashAlgorithm" validate:"required"`
	IssuerKeyHash  string            `json:"issuerKeyHash" validate:"required,max=128"`
	IssuerNameHash string            `json:"issuerNameHash" validate:"required,max=128"`
	SerialNumber   string            `json:"serialNumber" validate:"required,max=40"`
}

type CertificateHashDataChain struct {
	CertificateHashData      CertificateHashData     `json:"certificateHashData" validate:"required"`
	CertificateType          GetCertificateIDUseEnum `json:"certificateType" validate:"required"`
	ChildCertificateHashData []CertificateHashData   `json:"childCertificateHashData,omitempty" validate:"min=1,max=4"`
	CustomData               CustomData              `json:"customData,omitempty"`
}

type CertificateStatus struct {
	CertificateHashData CertificateHashData         `json:"certificateHashData" validate:"required"`
	CustomData          CustomData                  `json:"customData,omitempty"`
	NextUpdate          time.Time                   `json:"nextUpdate" validate:"required"`
	Source              CertificateStatusSourceEnum `json:"source" validate:"required"`
	Status              CertificateStatusEnum       `json:"status" validate:"required"`
}

type CertificateStatusRequestInfo struct {
	CertificateHashData CertificateHashData         `json:"certificateHashData" validate:"required"`
	CustomData          CustomData                  `json:"customData,omitempty"`
	Source              CertificateStatusSourceEnum `json:"source" validate:"required"`
	Urls                []string                    `json:"urls" validate:"required,min=1,max=5"`
}

type ChargingLimit struct {
	ChargingLimitSource string     `json:"chargingLimitSource" validate:"required,max=20"`
	CustomData          CustomData `json:"customData,omitempty"`
	IsGridCritical      *bool      `json:"isGridCritical,omitempty"`
	IsLocalGeneration   *bool      `json:"isLocalGeneration,omitempty"`
}

type ChargingNeeds struct {
	ACChargingParameters    *ACChargingParameters    `json:"acChargingParameters,omitempty"`
	AvailableEnergyTransfer []EnergyTransferModeEnum `json:"availableEnergyTransfer,omitempty" validate:"min=1"`
	ControlMode             *ControlModeEnum         `json:"controlMode,omitempty"`
	CustomData              CustomData               `json:"customData,omitempty"`
	DCChargingParameters    *DCChargingParameters    `json:"dcChargingParameters,omitempty"`
	DepartureTime           *time.Time               `json:"departureTime,omitempty"`
	DERChargingParameters   *DERChargingParameters   `json:"derChargingParameters,omitempty"`
	EVEnergyOffer           *EVEnergyOffer           `json:"evEnergyOffer,omitempty"`
	MobilityNeedsMode       *MobilityNeedsModeEnum   `json:"mobilityNeedsMode,omitempty"`
	RequestedEnergyTransfer EnergyTransferModeEnum   `json:"requestedEnergyTransfer" validate:"required"`
	V2XChargingParameters   *V2XChargingParameters   `json:"v2xChargingParameters,omitempty"`
}

type ChargingPeriod struct {
	CustomData  CustomData      `json:"customData,omitempty"`
	Dimensions  []CostDimension `json:"dimensions,omitempty" validate:"min=1"`
	StartPeriod time.Time       `json:"startPeriod" validate:"required"`
	TariffID    *string         `json:"tariffId,omitempty" validate:"max=60"`
}

type ChargingProfile struct {
	ChargingProfileKind         ChargingProfileKindEnum    `json:"chargingProfileKind" validate:"required"`
	ChargingProfilePurpose      ChargingProfilePurposeEnum `json:"chargingProfilePurpose" validate:"required"`
	ChargingSchedule            []ChargingSchedule         `json:"chargingSchedule" validate:"required,min=1,max=3"`
	CustomData                  CustomData                 `json:"customData,omitempty"`
	DynUpdateInterval           *int                       `json:"dynUpdateInterval,omitempty"`
	DynUpdateTime               *time.Time                 `json:"dynUpdateTime,omitempty"`
	ID                          int                        `json:"id" validate:"required"`
	InvalidAfterOfflineDuration *bool                      `json:"invalidAfterOfflineDuration,omitempty"`
	MaxOfflineDuration          *int                       `json:"maxOfflineDuration,omitempty"`
	PriceScheduleSignature      *string                    `json:"priceScheduleSignature,omitempty" validate:"max=256"`
	RecurrencyKind              *RecurrencyKindEnum        `json:"recurrencyKind,omitempty"`
	StackLevel                  int                        `json:"stackLevel" validate:"required,gte=0.0"`
	TransactionID               *string                    `json:"transactionId,omitempty" validate:"max=36"`
	ValidFrom                   *time.Time                 `json:"validFrom,omitempty"`
	ValidTo                     *time.Time                 `json:"validTo,omitempty"`
}

type ChargingProfileCriterion struct {
	ChargingLimitSource    []string                    `json:"chargingLimitSource,omitempty" validate:"min=1,max=4"`
	ChargingProfileID      []int                       `json:"chargingProfileId,omitempty" validate:"min=1"`
	ChargingProfilePurpose *ChargingProfilePurposeEnum `json:"chargingProfilePurpose,omitempty"`
	CustomData             CustomData                  `json:"customData,omitempty"`
	StackLevel             *int                        `json:"stackLevel,omitempty" validate:"gte=0.0"`
}

type ChargingSchedule struct {
	AbsolutePriceSchedule  *AbsolutePriceSchedule   `json:"absolutePriceSchedule,omitempty"`
	ChargingRateUnit       ChargingRateUnitEnum     `json:"chargingRateUnit" validate:"required"`
	ChargingSchedulePeriod []ChargingSchedulePeriod `json:"chargingSchedulePeriod" validate:"required,min=1,max=1024"`
	CustomData             CustomData               `json:"customData,omitempty"`
	DigestValue            *string                  `json:"digestValue,omitempty" validate:"max=88"`
	Duration               *int                     `json:"duration,omitempty"`
	ID                     int                      `json:"id" validate:"required"`
	LimitAtSoC             *LimitAtSoC              `json:"limitAtSoC,omitempty"`
	MinChargingRate        *float64                 `json:"minChargingRate,omitempty"`
	PowerTolerance         *float64                 `json:"powerTolerance,omitempty"`
	PriceLevelSchedule     *PriceLevelSchedule      `json:"priceLevelSchedule,omitempty"`
	RandomizedDelay        *int                     `json:"randomizedDelay,omitempty" validate:"gte=0.0"`
	SalesTariff            *SalesTariff             `json:"salesTariff,omitempty"`
	SignatureID            *int                     `json:"signatureId,omitempty" validate:"gte=0.0"`
	StartSchedule          *time.Time               `json:"startSchedule,omitempty"`
	UseLocalTime           *bool                    `json:"useLocalTime,omitempty"`
}

type ChargingSchedulePeriod struct {
	CustomData             CustomData           `json:"customData,omitempty"`
	DischargeLimit         *float64             `json:"dischargeLimit,omitempty" validate:"lte=0.0"`
	DischargeLimitL2       *float64             `json:"dischargeLimit_L2,omitempty" validate:"lte=0.0"`
	DischargeLimitL3       *float64             `json:"dischargeLimit_L3,omitempty" validate:"lte=0.0"`
	EVSESleep              *bool                `json:"evseSleep,omitempty"`
	Limit                  *float64             `json:"limit,omitempty"`
	LimitL2                *float64             `json:"limit_L2,omitempty"`
	LimitL3                *float64             `json:"limit_L3,omitempty"`
	NumberPhases           *int                 `json:"numberPhases,omitempty" validate:"gte=0.0,lte=3.0"`
	OperationMode          *OperationModeEnum   `json:"operationMode,omitempty"`
	PhaseToUse             *int                 `json:"phaseToUse,omitempty" validate:"gte=0.0,lte=3.0"`
	PreconditioningRequest *bool                `json:"preconditioningRequest,omitempty"`
	Setpoint               *float64             `json:"setpoint,omitempty"`
	SetpointReactive       *float64             `json:"setpointReactive,omitempty"`
	SetpointReactiveL2     *float64             `json:"setpointReactive_L2,omitempty"`
	SetpointReactiveL3     *float64             `json:"setpointReactive_L3,omitempty"`
	SetpointL2             *float64             `json:"setpoint_L2,omitempty"`
	SetpointL3             *float64             `json:"setpoint_L3,omitempty"`
	StartPeriod            int                  `json:"startPeriod" validate:"required"`
	V2XBaseline            *float64             `json:"v2xBaseline,omitempty"`
	V2XFreqWattCurve       []V2XFreqWattPoint   `json:"v2xFreqWattCurve,omitempty" validate:"min=1,max=20"`
	V2XSignalWattCurve     []V2XSignalWattPoint `json:"v2xSignalWattCurve,omitempty" validate:"min=1,max=20"`
}

type ChargingScheduleUpdate struct {
	CustomData         CustomData `json:"customData,omitempty"`
	DischargeLimit     *float64   `json:"dischargeLimit,omitempty" validate:"lte=0.0"`
	DischargeLimitL2   *float64   `json:"dischargeLimit_L2,omitempty" validate:"lte=0.0"`
	DischargeLimitL3   *float64   `json:"dischargeLimit_L3,omitempty" validate:"lte=0.0"`
	Limit              *float64   `json:"limit,omitempty"`
	LimitL2            *float64   `json:"limit_L2,omitempty"`
	LimitL3            *float64   `json:"limit_L3,omitempty"`
	Setpoint           *float64   `json:"setpoint,omitempty"`
	SetpointReactive   *float64   `json:"setpointReactive,omitempty"`
	SetpointReactiveL2 *float64   `json:"setpointReactive_L2,omitempty"`
	SetpointReactiveL3 *float64   `json:"setpointReactive_L3,omitempty"`
	SetpointL2         *float64   `json:"setpoint_L2,omitempty"`
	SetpointL3         *float64   `json:"setpoint_L3,omitempty"`
}

type ChargingStation struct {
	CustomData      CustomData `json:"customData,omitempty"`
	FirmwareVersion *string    `json:"firmwareVersion,omitempty" validate:"max=50"`
	Model           string     `json:"model" validate:"required,max=20"`
	Modem           *Modem     `json:"modem,omitempty"`
	SerialNumber    *string    `json:"serialNumber,omitempty" validate:"max=25"`
	VendorName      string     `json:"vendorName" validate:"required,max=50"`
}

type ClearChargingProfile struct {
	ChargingProfilePurpose *ChargingProfilePurposeEnum `json:"chargingProfilePurpose,omitempty"`
	CustomData             CustomData                  `json:"customData,omitempty"`
	EVSEID                 *int                        `json:"evseId,omitempty" validate:"gte=0.0"`
	StackLevel             *int                        `json:"stackLevel,omitempty" validate:"gte=0.0"`
}

type ClearMonitoringResult struct {
	CustomData CustomData                `json:"customData,omitempty"`
	ID         int                       `json:"id" validate:"required,gte=0.0"`
	Status     ClearMonitoringStatusEnum `json:"status" validate:"required"`
	StatusInfo *StatusInfo               `json:"statusInfo,omitempty"`
}

type ClearTariffsResult struct {
	CustomData CustomData            `json:"customData,omitempty"`
	Status     TariffClearStatusEnum `json:"status" validate:"required"`
	StatusInfo *StatusInfo           `json:"statusInfo,omitempty"`
	TariffID   *string               `json:"tariffId,omitempty" validate:"max=60"`
}

type Component struct {
	CustomData CustomData `json:"customData,omitempty"`
	EVSE       *EVSE      `json:"evse,omitempty"`
	Instance   *string    `json:"instance,omitempty" validate:"max=50"`
	Name       string     `json:"name" validate:"required,max=50"`
}

type ComponentVariable struct {
	Component  Component  `json:"component" validate:"required"`
	CustomData CustomData `json:"customData,omitempty"`
	Variable   *Variable  `json:"variable,omitempty"`
}

type CompositeSchedule struct {
	ChargingRateUnit       ChargingRateUnitEnum     `json:"chargingRateUnit" validate:"required"`
	ChargingSchedulePeriod []ChargingSchedulePeriod `json:"chargingSchedulePeriod" validate:"required,min=1"`
	CustomData             CustomData               `json:"customData,omitempty"`
	Duration               int                      `json:"duration" validate:"required"`
	EVSEID                 int                      `json:"evseId" validate:"required,gte=0.0"`
	ScheduleStart          time.Time                `json:"scheduleStart" validate:"required"`
}

type ConstantStreamData struct {
	CustomData           CustomData                `json:"customData,omitempty"`
	ID                   int                       `json:"id" validate:"required,gte=0.0"`
	Params               PeriodicEventStreamParams `json:"params" validate:"required"`
	VariableMonitoringID int                       `json:"variableMonitoringId" validate:"required,gte=0.0"`
}

type ConsumptionCost struct {
	Cost       []Cost     `json:"cost" validate:"required,min=1,max=3"`
	CustomData CustomData `json:"customData,omitempty"`
	StartValue float64    `json:"startValue" validate:"required"`
}

type Cost struct {
	Amount           int          `json:"amount" validate:"required"`
	AmountMultiplier *int         `json:"amountMultiplier,omitempty"`
	CostKind         CostKindEnum `json:"costKind" validate:"required"`
	CustomData       CustomData   `json:"customData,omitempty"`
}

type CostDetails struct {
	ChargingPeriods    []ChargingPeriod `json:"chargingPeriods,omitempty" validate:"min=1"`
	CustomData         CustomData       `json:"customData,omitempty"`
	FailureReason      *string          `json:"failureReason,omitempty" validate:"max=500"`
	FailureToCalculate *bool            `json:"failureToCalculate,omitempty"`
	TotalCost          TotalCost        `json:"totalCost" validate:"required"`
	TotalUsage         TotalUsage       `json:"totalUsage" validate:"required"`
}

type CostDimension struct {
	CustomData CustomData        `json:"customData,omitempty"`
	TypeValue  CostDimensionEnum `json:"type" validate:"required"`
	Volume     float64           `json:"volume" validate:"required"`
}

type DCChargingParameters struct {
	BulkSoC          *int       `json:"bulkSoC,omitempty" validate:"gte=0.0,lte=100.0"`
	CustomData       CustomData `json:"customData,omitempty"`
	EnergyAmount     *float64   `json:"energyAmount,omitempty"`
	EVEnergyCapacity *float64   `json:"evEnergyCapacity,omitempty"`
	EVMaxCurrent     float64    `json:"evMaxCurrent" validate:"required"`
	EVMaxPower       *float64   `json:"evMaxPower,omitempty"`
	EVMaxVoltage     float64    `json:"evMaxVoltage" validate:"required"`
	FullSoC          *int       `json:"fullSoC,omitempty" validate:"gte=0.0,lte=100.0"`
	StateOfCharge    *int       `json:"stateOfCharge,omitempty" validate:"gte=0.0,lte=100.0"`
}

type DERChargingParameters struct {
	CustomData                             CustomData               `json:"customData,omitempty"`
	EVDurationLevel1DCInjection            *float64                 `json:"evDurationLevel1DCInjection,omitempty"`
	EVDurationLevel2DCInjection            *float64                 `json:"evDurationLevel2DCInjection,omitempty"`
	EVInverterHwVersion                    *string                  `json:"evInverterHwVersion,omitempty" validate:"max=50"`
	EVInverterManufacturer                 *string                  `json:"evInverterManufacturer,omitempty" validate:"max=50"`
	EVInverterModel                        *string                  `json:"evInverterModel,omitempty" validate:"max=50"`
	EVInverterSerialNumber                 *string                  `json:"evInverterSerialNumber,omitempty" validate:"max=50"`
	EVInverterSwVersion                    *string                  `json:"evInverterSwVersion,omitempty" validate:"max=50"`
	EVIslandingDetectionMethod             []IslandingDetectionEnum `json:"evIslandingDetectionMethod,omitempty" validate:"min=1"`
	EVIslandingTripTime                    *float64                 `json:"evIslandingTripTime,omitempty"`
	EVMaximumLevel1DCInjection             *float64                 `json:"evMaximumLevel1DCInjection,omitempty"`
	EVMaximumLevel2DCInjection             *float64                 `json:"evMaximumLevel2DCInjection,omitempty"`
	EVOverExcitedMaxDischargePower         *float64                 `json:"evOverExcitedMaxDischargePower,omitempty"`
	EVOverExcitedPowerFactor               *float64                 `json:"evOverExcitedPowerFactor,omitempty"`
	EVReactiveSusceptance                  *float64                 `json:"evReactiveSusceptance,omitempty"`
	EVSessionTotalDischargeEnergyAvailable *float64                 `json:"evSessionTotalDischargeEnergyAvailable,omitempty"`
	EVSupportedDERControl                  []DERControlEnum         `json:"evSupportedDERControl,omitempty" validate:"min=1"`
	EVUnderExcitedMaxDischargePower        *float64                 `json:"evUnderExcitedMaxDischargePower,omitempty"`
	EVUnderExcitedPowerFactor              *float64                 `json:"evUnderExcitedPowerFactor,omitempty"`
	MaxApparentPower                       *float64                 `json:"maxApparentPower,omitempty"`
	MaxChargeApparentPower                 *float64                 `json:"maxChargeApparentPower,omitempty"`
	MaxChargeApparentPowerL2               *float64                 `json:"maxChargeApparentPower_L2,omitempty"`
	MaxChargeApparentPowerL3               *float64                 `json:"maxChargeApparentPower_L3,omitempty"`
	MaxChargeReactivePower                 *float64                 `json:"maxChargeReactivePower,omitempty"`
	MaxChargeReactivePowerL2               *float64                 `json:"maxChargeReactivePower_L2,omitempty"`
	MaxChargeReactivePowerL3               *float64                 `json:"maxChargeReactivePower_L3,omitempty"`
	MaxDischargeApparentPower              *float64                 `json:"maxDischargeApparentPower,omitempty"`
	MaxDischargeApparentPowerL2            *float64                 `json:"maxDischargeApparentPower_L2,omitempty"`
	MaxDischargeApparentPowerL3            *float64                 `json:"maxDischargeApparentPower_L3,omitempty"`
	MaxDischargeReactivePower              *float64                 `json:"maxDischargeReactivePower,omitempty"`
	MaxDischargeReactivePowerL2            *float64                 `json:"maxDischargeReactivePower_L2,omitempty"`
	MaxDischargeReactivePowerL3            *float64                 `json:"maxDischargeReactivePower_L3,omitempty"`
	MaxNominalVoltage                      *float64                 `json:"maxNominalVoltage,omitempty"`
	MinChargeReactivePower                 *float64                 `json:"minChargeReactivePower,omitempty"`
	MinChargeReactivePowerL2               *float64                 `json:"minChargeReactivePower_L2,omitempty"`
	MinChargeReactivePowerL3               *float64                 `json:"minChargeReactivePower_L3,omitempty"`
	MinDischargeReactivePower              *float64                 `json:"minDischargeReactivePower,omitempty"`
	MinDischargeReactivePowerL2            *float64                 `json:"minDischargeReactivePower_L2,omitempty"`
	MinDischargeReactivePowerL3            *float64                 `json:"minDischargeReactivePower_L3,omitempty"`
	MinNominalVoltage                      *float64                 `json:"minNominalVoltage,omitempty"`
	NominalVoltage                         *float64                 `json:"nominalVoltage,omitempty"`
	NominalVoltageOffset                   *float64                 `json:"nominalVoltageOffset,omitempty"`
}

type DERCurve struct {
	CurveData           []DERCurvePoints     `json:"curveData" validate:"required,min=1,max=10"`
	CustomData          CustomData           `json:"customData,omitempty"`
	Duration            *float64             `json:"duration,omitempty"`
	Hysteresis          *Hysteresis          `json:"hysteresis,omitempty"`
	Priority            int                  `json:"priority" validate:"required,gte=0.0"`
	ReactivePowerParams *ReactivePowerParams `json:"reactivePowerParams,omitempty"`
	ResponseTime        *float64             `json:"responseTime,omitempty"`
	StartTime           *time.Time           `json:"startTime,omitempty"`
	VoltageParams       *VoltageParams       `json:"voltageParams,omitempty"`
	YUnit               DERUnitEnum          `json:"yUnit" validate:"required"`
}

type DERCurveGet struct {
	Curve        DERCurve       `json:"curve" validate:"required"`
	CurveType    DERControlEnum `json:"curveType" validate:"required"`
	CustomData   CustomData     `json:"customData,omitempty"`
	ID           string         `json:"id" validate:"required,max=36"`
	IsDefault    bool           `json:"isDefault" validate:"required"`
	IsSuperseded bool           `json:"isSuperseded" validate:"required"`
}

type DERCurvePoints struct {
	CustomData CustomData `json:"customData,omitempty"`
	X          float64    `json:"x" validate:"required"`
	Y          float64    `json:"y" validate:"required"`
}

type EVAbsolutePriceSchedule struct {
	Currency                       string                         `json:"currency" validate:"required,max=3"`
	CustomData                     CustomData                     `json:"customData,omitempty"`
	EVAbsolutePriceScheduleEntries []EVAbsolutePriceScheduleEntry `json:"evAbsolutePriceScheduleEntries" validate:"required,min=1,max=1024"`
	PriceAlgorithm                 string                         `json:"priceAlgorithm" validate:"required,max=2000"`
	TimeAnchor                     time.Time                      `json:"timeAnchor" validate:"required"`
}

type EVAbsolutePriceScheduleEntry struct {
	CustomData  CustomData    `json:"customData,omitempty"`
	Duration    int           `json:"duration" validate:"required"`
	EVPriceRule []EVPriceRule `json:"evPriceRule" validate:"required,min=1,max=8"`
}

type EVEnergyOffer struct {
	CustomData              CustomData               `json:"customData,omitempty"`
	EVAbsolutePriceSchedule *EVAbsolutePriceSchedule `json:"evAbsolutePriceSchedule,omitempty"`
	EVPowerSchedule         EVPowerSchedule          `json:"evPowerSchedule" validate:"required"`
}

type EVPowerSchedule struct {
	CustomData             CustomData             `json:"customData,omitempty"`
	EVPowerScheduleEntries []EVPowerScheduleEntry `json:"evPowerScheduleEntries" validate:"required,min=1,max=1024"`
	TimeAnchor             time.Time              `json:"timeAnchor" validate:"required"`
}

type EVPowerScheduleEntry struct {
	CustomData CustomData `json:"customData,omitempty"`
	Duration   int        `json:"duration" validate:"required"`
	Power      float64    `json:"power" validate:"required"`
}

type EVPriceRule struct {
	CustomData      CustomData `json:"customData,omitempty"`
	EnergyFee       float64    `json:"energyFee" validate:"required"`
	PowerRangeStart float64    `json:"powerRangeStart" validate:"required"`
}

type EVSE struct {
	ConnectorID *int       `json:"connectorId,omitempty" validate:"gte=0.0"`
	CustomData  CustomData `json:"customData,omitempty"`
	ID          int        `json:"id" validate:"required,gte=0.0"`
}

type EnterService struct {
	CustomData  CustomData `json:"customData,omitempty"`
	Delay       *float64   `json:"delay,omitempty"`
	HighFreq    float64    `json:"highFreq" validate:"required"`
	HighVoltage float64    `json:"highVoltage" validate:"required"`
	LowFreq     float64    `json:"lowFreq" validate:"required"`
	LowVoltage  float64    `json:"lowVoltage" validate:"required"`
	Priority    int        `json:"priority" validate:"required,gte=0.0"`
	RampRate    *float64   `json:"rampRate,omitempty"`
	RandomDelay *float64   `json:"randomDelay,omitempty"`
}

type EnterServiceGet struct {
	CustomData   CustomData   `json:"customData,omitempty"`
	EnterService EnterService `json:"enterService" validate:"required"`
	ID           string       `json:"id" validate:"required,max=36"`
}

type EventData struct {
	ActualValue           string                `json:"actualValue" validate:"required,max=2500"`
	Cause                 *int                  `json:"cause,omitempty" validate:"gte=0.0"`
	Cleared               *bool                 `json:"cleared,omitempty"`
	Component             Component             `json:"component" validate:"required"`
	CustomData            CustomData            `json:"customData,omitempty"`
	EventID               int                   `json:"eventId" validate:"required,gte=0.0"`
	EventNotificationType EventNotificationEnum `json:"eventNotificationType" validate:"required"`
	Severity              *int                  `json:"severity,omitempty" validate:"gte=0.0"`
	TechCode              *string               `json:"techCode,omitempty" validate:"max=50"`
	TechInfo              *string               `json:"techInfo,omitempty" validate:"max=500"`
	Timestamp             time.Time             `json:"timestamp" validate:"required"`
	TransactionID         *string               `json:"transactionId,omitempty" validate:"max=36"`
	Trigger               EventTriggerEnum      `json:"trigger" validate:"required"`
	Variable              Variable              `json:"variable" validate:"required"`
	VariableMonitoringID  *int                  `json:"variableMonitoringId,omitempty" validate:"gte=0.0"`
}

type Firmware struct {
	CustomData         CustomData `json:"customData,omitempty"`
	InstallDateTime    *time.Time `json:"installDateTime,omitempty"`
	Location           string     `json:"location" validate:"required,max=2000"`
	RetrieveDateTime   time.Time  `json:"retrieveDateTime" validate:"required"`
	Signature          *string    `json:"signature,omitempty" validate:"max=800"`
	SigningCertificate *string    `json:"signingCertificate,omitempty" validate:"max=5500"`
}

type FixedPF struct {
	CustomData   CustomData `json:"customData,omitempty"`
	Displacement float64    `json:"displacement" validate:"required"`
	Duration     *float64   `json:"duration,omitempty"`
	Excitation   bool       `json:"excitation" validate:"required"`
	Priority     int        `json:"priority" validate:"required,gte=0.0"`
	StartTime    *time.Time `json:"startTime,omitempty"`
}

type FixedPFGet struct {
	CustomData   CustomData `json:"customData,omitempty"`
	FixedPF      FixedPF    `json:"fixedPF" validate:"required"`
	ID           string     `json:"id" validate:"required,max=36"`
	IsDefault    bool       `json:"isDefault" validate:"required"`
	IsSuperseded bool       `json:"isSuperseded" validate:"required"`
}

type FixedVAR struct {
	CustomData CustomData  `json:"customData,omitempty"`
	Duration   *float64    `json:"duration,omitempty"`
	Priority   int         `json:"priority" validate:"required,gte=0.0"`
	Setpoint   float64     `json:"setpoint" validate:"required"`
	StartTime  *time.Time  `json:"startTime,omitempty"`
	Unit       DERUnitEnum `json:"unit" validate:"required"`
}

type FixedVARGet struct {
	CustomData   CustomData `json:"customData,omitempty"`
	FixedVAR     FixedVAR   `json:"fixedVar" validate:"required"`
	ID           string     `json:"id" validate:"required,max=36"`
	IsDefault    bool       `json:"isDefault" validate:"required"`
	IsSuperseded bool       `json:"isSuperseded" validate:"required"`
}

type FreqDroop struct {
	CustomData   CustomData `json:"customData,omitempty"`
	Duration     *float64   `json:"duration,omitempty"`
	OverDroop    float64    `json:"overDroop" validate:"required"`
	OverFreq     float64    `json:"overFreq" validate:"required"`
	Priority     int        `json:"priority" validate:"required,gte=0.0"`
	ResponseTime float64    `json:"responseTime" validate:"required"`
	StartTime    *time.Time `json:"startTime,omitempty"`
	UnderDroop   float64    `json:"underDroop" validate:"required"`
	UnderFreq    float64    `json:"underFreq" validate:"required"`
}

type FreqDroopGet struct {
	CustomData   CustomData `json:"customData,omitempty"`
	FreqDroop    FreqDroop  `json:"freqDroop" validate:"required"`
	ID           string     `json:"id" validate:"required,max=36"`
	IsDefault    bool       `json:"isDefault" validate:"required"`
	IsSuperseded bool       `json:"isSuperseded" validate:"required"`
}

type GetVariableData struct {
	AttributeType *AttributeEnum `json:"attributeType,omitempty"`
	Component     Component      `json:"component" validate:"required"`
	CustomData    CustomData     `json:"customData,omitempty"`
	Variable      Variable       `json:"variable" validate:"required"`
}

type GetVariableResult struct {
	AttributeStatus     GetVariableStatusEnum `json:"attributeStatus" validate:"required"`
	AttributeStatusInfo *StatusInfo           `json:"attributeStatusInfo,omitempty"`
	AttributeType       *AttributeEnum        `json:"attributeType,omitempty"`
	AttributeValue      *string               `json:"attributeValue,omitempty" validate:"max=2500"`
	Component           Component             `json:"component" validate:"required"`
	CustomData          CustomData            `json:"customData,omitempty"`
	Variable            Variable              `json:"variable" validate:"required"`
}

type Gradient struct {
	CustomData   CustomData `json:"customData,omitempty"`
	Gradient     float64    `json:"gradient" validate:"required"`
	Priority     int        `json:"priority" validate:"required,gte=0.0"`
	SoftGradient float64    `json:"softGradient" validate:"required"`
}

type GradientGet struct {
	CustomData CustomData `json:"customData,omitempty"`
	Gradient   Gradient   `json:"gradient" validate:"required"`
	ID         string     `json:"id" validate:"required,max=36"`
}

type Hysteresis struct {
	CustomData         CustomData `json:"customData,omitempty"`
	HysteresisDelay    *float64   `json:"hysteresisDelay,omitempty"`
	HysteresisGradient *float64   `json:"hysteresisGradient,omitempty"`
	HysteresisHigh     *float64   `json:"hysteresisHigh,omitempty"`
	HysteresisLow      *float64   `json:"hysteresisLow,omitempty"`
}

type IDToken struct {
	AdditionalInfo []AdditionalInfo `json:"additionalInfo,omitempty" validate:"min=1"`
	CustomData     CustomData       `json:"customData,omitempty"`
	IDToken        string           `json:"idToken" validate:"required,max=255"`
	TypeValue      string           `json:"type" validate:"required,max=20"`
}

type IDTokenInfo struct {
	CacheExpiryDateTime *time.Time              `json:"cacheExpiryDateTime,omitempty"`
	ChargingPriority    *int                    `json:"chargingPriority,omitempty"`
	CustomData          CustomData              `json:"customData,omitempty"`
	EVSEID              []int                   `json:"evseId,omitempty" validate:"min=1"`
	GroupIDToken        *IDToken                `json:"groupIdToken,omitempty"`
	Language1           *string                 `json:"language1,omitempty" validate:"max=8"`
	Language2           *string                 `json:"language2,omitempty" validate:"max=8"`
	PersonalMessage     *MessageContent         `json:"personalMessage,omitempty"`
	Status              AuthorizationStatusEnum `json:"status" validate:"required"`
}

type LimitAtSoC struct {
	CustomData CustomData `json:"customData,omitempty"`
	Limit      float64    `json:"limit" validate:"required"`
	SoC        int        `json:"soc" validate:"required,gte=0.0,lte=100.0"`
}

type LimitMaxDischarge struct {
	CustomData              CustomData `json:"customData,omitempty"`
	Duration                *float64   `json:"duration,omitempty"`
	PctMaxDischargePower    *float64   `json:"pctMaxDischargePower,omitempty"`
	PowerMonitoringMustTrip *DERCurve  `json:"powerMonitoringMustTrip,omitempty"`
	Priority                int        `json:"priority" validate:"required,gte=0.0"`
	StartTime               *time.Time `json:"startTime,omitempty"`
}

type LimitMaxDischargeGet struct {
	CustomData        CustomData        `json:"customData,omitempty"`
	ID                string            `json:"id" validate:"required,max=36"`
	IsDefault         bool              `json:"isDefault" validate:"required"`
	IsSuperseded      bool              `json:"isSuperseded" validate:"required"`
	LimitMaxDischarge LimitMaxDischarge `json:"limitMaxDischarge" validate:"required"`
}

type LogParameters struct {
	CustomData      CustomData `json:"customData,omitempty"`
	LatestTimestamp *time.Time `json:"latestTimestamp,omitempty"`
	OldestTimestamp *time.Time `json:"oldestTimestamp,omitempty"`
	RemoteLocation  string     `json:"remoteLocation" validate:"required,max=2000"`
}

type MessageContent struct {
	Content    string            `json:"content" validate:"required,max=1024"`
	CustomData CustomData        `json:"customData,omitempty"`
	Format     MessageFormatEnum `json:"format" validate:"required"`
	Language   *string           `json:"language,omitempty" validate:"max=8"`
}

type MessageInfo struct {
	CustomData    CustomData          `json:"customData,omitempty"`
	Display       *Component          `json:"display,omitempty"`
	EndDateTime   *time.Time          `json:"endDateTime,omitempty"`
	ID            int                 `json:"id" validate:"required,gte=0.0"`
	Message       MessageContent      `json:"message" validate:"required"`
	MessageExtra  []MessageContent    `json:"messageExtra,omitempty" validate:"min=1,max=4"`
	Priority      MessagePriorityEnum `json:"priority" validate:"required"`
	StartDateTime *time.Time          `json:"startDateTime,omitempty"`
	State         *MessageStateEnum   `json:"state,omitempty"`
	TransactionID *string             `json:"transactionId,omitempty" validate:"max=36"`
}

type MeterValue struct {
	CustomData   CustomData     `json:"customData,omitempty"`
	SampledValue []SampledValue `json:"sampledValue" validate:"required,min=1"`
	Timestamp    time.Time      `json:"timestamp" validate:"required"`
}

type Modem struct {
	CustomData CustomData `json:"customData,omitempty"`
	ICCID      *string    `json:"iccid,omitempty" validate:"max=20"`
	IMSI       *string    `json:"imsi,omitempty" validate:"max=20"`
}

type MonitoringData struct {
	Component          Component            `json:"component" validate:"required"`
	CustomData         CustomData           `json:"customData,omitempty"`
	Variable           Variable             `json:"variable" validate:"required"`
	VariableMonitoring []VariableMonitoring `json:"variableMonitoring" validate:"required,min=1"`
}

type NetworkConnectionProfile struct {
	APN               *APN              `json:"apn,omitempty"`
	BasicAuthPassword *string           `json:"basicAuthPassword,omitempty" validate:"max=64"`
	CustomData        CustomData        `json:"customData,omitempty"`
	Identity          *string           `json:"identity,omitempty" validate:"max=48"`
	MessageTimeout    int               `json:"messageTimeout" validate:"required"`
	OCPPCSMSURL       string            `json:"ocppCsmsUrl" validate:"required,max=2000"`
	OCPPInterface     OCPPInterfaceEnum `json:"ocppInterface" validate:"required"`
	OCPPTransport     OCPPTransportEnum `json:"ocppTransport" validate:"required"`
	OCPPVersion       *OCPPVersionEnum  `json:"ocppVersion,omitempty"`
	SecurityProfile   int               `json:"securityProfile" validate:"required,gte=0.0"`
	VPN               *VPN              `json:"vpn,omitempty"`
}

type OCSPRequestData struct {
	CustomData     CustomData        `json:"customData,omitempty"`
	HashAlgorithm  HashAlgorithmEnum `json:"hashAlgorithm" validate:"required"`
	IssuerKeyHash  string            `json:"issuerKeyHash" validate:"required,max=128"`
	IssuerNameHash string            `json:"issuerNameHash" validate:"required,max=128"`
	ResponderURL   string            `json:"responderURL" validate:"required,max=2000"`
	SerialNumber   string            `json:"serialNumber" validate:"required,max=40"`
}

type OverstayRule struct {
	CustomData              CustomData     `json:"customData,omitempty"`
	OverstayFee             RationalNumber `json:"overstayFee" validate:"required"`
	OverstayFeePeriod       int            `json:"overstayFeePeriod" validate:"required"`
	OverstayRuleDescription *string        `json:"overstayRuleDescription,omitempty" validate:"max=32"`
	StartTime               int            `json:"startTime" validate:"required"`
}

type OverstayRuleList struct {
	CustomData             CustomData      `json:"customData,omitempty"`
	OverstayPowerThreshold *RationalNumber `json:"overstayPowerThreshold,omitempty"`
	OverstayRule           []OverstayRule  `json:"overstayRule" validate:"required,min=1,max=5"`
	OverstayTimeThreshold  *int            `json:"overstayTimeThreshold,omitempty"`
}

type PeriodicEventStreamParams struct {
	CustomData CustomData `json:"customData,omitempty"`
	Interval   *int       `json:"interval,omitempty" validate:"gte=0.0"`
	Values     *int       `json:"values,omitempty" validate:"gte=0.0"`
}

type Price struct {
	CustomData CustomData `json:"customData,omitempty"`
	ExclTax    *float64   `json:"exclTax,omitempty"`
	InclTax    *float64   `json:"inclTax,omitempty"`
	TaxRates   []TaxRate  `json:"taxRates,omitempty" validate:"min=1,max=5"`
}

type PriceLevelSchedule struct {
	CustomData                CustomData                `json:"customData,omitempty"`
	NumberOfPriceLevels       int                       `json:"numberOfPriceLevels" validate:"required,gte=0.0"`
	PriceLevelScheduleEntries []PriceLevelScheduleEntry `json:"priceLevelScheduleEntries" validate:"required,min=1,max=100"`
	PriceScheduleDescription  *string                   `json:"priceScheduleDescription,omitempty" validate:"max=32"`
	PriceScheduleID           int                       `json:"priceScheduleId" validate:"required,gte=0.0"`
	TimeAnchor                time.Time                 `json:"timeAnchor" validate:"required"`
}

type PriceLevelScheduleEntry struct {
	CustomData CustomData `json:"customData,omitempty"`
	Duration   int        `json:"duration" validate:"required"`
	PriceLevel int        `json:"priceLevel" validate:"required,gte=0.0"`
}

type PriceRule struct {
	CarbonDioxideEmission         *int            `json:"carbonDioxideEmission,omitempty" validate:"gte=0.0"`
	CustomData                    CustomData      `json:"customData,omitempty"`
	EnergyFee                     RationalNumber  `json:"energyFee" validate:"required"`
	ParkingFee                    *RationalNumber `json:"parkingFee,omitempty"`
	ParkingFeePeriod              *int            `json:"parkingFeePeriod,omitempty"`
	PowerRangeStart               RationalNumber  `json:"powerRangeStart" validate:"required"`
	RenewableGenerationPercentage *int            `json:"renewableGenerationPercentage,omitempty" validate:"gte=0.0,lte=100.0"`
}

type PriceRuleStack struct {
	CustomData CustomData  `json:"customData,omitempty"`
	Duration   int         `json:"duration" validate:"required"`
	PriceRule  []PriceRule `json:"priceRule" validate:"required,min=1,max=8"`
}

type RationalNumber struct {
	CustomData CustomData `json:"customData,omitempty"`
	Exponent   int        `json:"exponent" validate:"required"`
	Value      int        `json:"value" validate:"required"`
}

type ReactivePowerParams struct {
	AutonomousVRefEnable       *bool      `json:"autonomousVRefEnable,omitempty"`
	AutonomousVRefTimeConstant *float64   `json:"autonomousVRefTimeConstant,omitempty"`
	CustomData                 CustomData `json:"customData,omitempty"`
	VRef                       *float64   `json:"vRef,omitempty"`
}

type RelativeTimeInterval struct {
	CustomData CustomData `json:"customData,omitempty"`
	Duration   *int       `json:"duration,omitempty"`
	Start      int        `json:"start" validate:"required"`
}

type ReportData struct {
	Component               Component                `json:"component" validate:"required"`
	CustomData              CustomData               `json:"customData,omitempty"`
	Variable                Variable                 `json:"variable" validate:"required"`
	VariableAttribute       []VariableAttribute      `json:"variableAttribute" validate:"required,min=1,max=4"`
	VariableCharacteristics *VariableCharacteristics `json:"variableCharacteristics,omitempty"`
}

type SalesTariff struct {
	CustomData             CustomData         `json:"customData,omitempty"`
	ID                     int                `json:"id" validate:"required,gte=0.0"`
	NumEPriceLevels        *int               `json:"numEPriceLevels,omitempty" validate:"gte=0.0"`
	SalesTariffDescription *string            `json:"salesTariffDescription,omitempty" validate:"max=32"`
	SalesTariffEntry       []SalesTariffEntry `json:"salesTariffEntry" validate:"required,min=1,max=1024"`
}

type SalesTariffEntry struct {
	ConsumptionCost      []ConsumptionCost    `json:"consumptionCost,omitempty" validate:"min=1,max=3"`
	CustomData           CustomData           `json:"customData,omitempty"`
	EPriceLevel          *int                 `json:"ePriceLevel,omitempty" validate:"gte=0.0"`
	RelativeTimeInterval RelativeTimeInterval `json:"relativeTimeInterval" validate:"required"`
}

type SampledValue struct {
	Context          *ReadingContextEnum `json:"context,omitempty"`
	CustomData       CustomData          `json:"customData,omitempty"`
	Location         *LocationEnum       `json:"location,omitempty"`
	Measurand        *MeasurandEnum      `json:"measurand,omitempty"`
	Phase            *PhaseEnum          `json:"phase,omitempty"`
	SignedMeterValue *SignedMeterValue   `json:"signedMeterValue,omitempty"`
	UnitOfMeasure    *UnitOfMeasure      `json:"unitOfMeasure,omitempty"`
	Value            float64             `json:"value" validate:"required"`
}

type SetMonitoringData struct {
	Component           Component                  `json:"component" validate:"required"`
	CustomData          CustomData                 `json:"customData,omitempty"`
	ID                  *int                       `json:"id,omitempty" validate:"gte=0.0"`
	PeriodicEventStream *PeriodicEventStreamParams `json:"periodicEventStream,omitempty"`
	Severity            int                        `json:"severity" validate:"required,gte=0.0"`
	Transaction         *bool                      `json:"transaction,omitempty"`
	TypeValue           MonitorEnum                `json:"type" validate:"required"`
	Value               float64                    `json:"value" validate:"required"`
	Variable            Variable                   `json:"variable" validate:"required"`
}

type SetMonitoringResult struct {
	Component  Component               `json:"component" validate:"required"`
	CustomData CustomData              `json:"customData,omitempty"`
	ID         *int                    `json:"id,omitempty" validate:"gte=0.0"`
	Severity   int                     `json:"severity" validate:"required,gte=0.0"`
	Status     SetMonitoringStatusEnum `json:"status" validate:"required"`
	StatusInfo *StatusInfo             `json:"statusInfo,omitempty"`
	TypeValue  MonitorEnum             `json:"type" validate:"required"`
	Variable   Variable                `json:"variable" validate:"required"`
}

type SetVariableData struct {
	AttributeType  *AttributeEnum `json:"attributeType,omitempty"`
	AttributeValue string         `json:"attributeValue" validate:"required,max=2500"`
	Component      Component      `json:"component" validate:"required"`
	CustomData     CustomData     `json:"customData,omitempty"`
	Variable       Variable       `json:"variable" validate:"required"`
}

type SetVariableResult struct {
	AttributeStatus     SetVariableStatusEnum `json:"attributeStatus" validate:"required"`
	AttributeStatusInfo *StatusInfo           `json:"attributeStatusInfo,omitempty"`
	AttributeType       *AttributeEnum        `json:"attributeType,omitempty"`
	Component           Component             `json:"component" validate:"required"`
	CustomData          CustomData            `json:"customData,omitempty"`
	Variable            Variable              `json:"variable" validate:"required"`
}

type SignedMeterValue struct {
	CustomData      CustomData `json:"customData,omitempty"`
	EncodingMethod  string     `json:"encodingMethod" validate:"required,max=50"`
	PublicKey       *string    `json:"publicKey,omitempty" validate:"max=2500"`
	SignedMeterData string     `json:"signedMeterData" validate:"required,max=32768"`
	SigningMethod   *string    `json:"signingMethod,omitempty" validate:"max=50"`
}

type StatusInfo struct {
	AdditionalInfo *string    `json:"additionalInfo,omitempty" validate:"max=1024"`
	CustomData     CustomData `json:"customData,omitempty"`
	ReasonCode     string     `json:"reasonCode" validate:"required,max=20"`
}

type StreamDataElement struct {
	CustomData CustomData `json:"customData,omitempty"`
	T          float64    `json:"t" validate:"required"`
	V          string     `json:"v" validate:"required,max=2500"`
}

type Tariff struct {
	ChargingTime     *TariffTime      `json:"chargingTime,omitempty"`
	Currency         string           `json:"currency" validate:"required,max=3"`
	CustomData       CustomData       `json:"customData,omitempty"`
	Description      []MessageContent `json:"description,omitempty" validate:"min=1,max=10"`
	Energy           *TariffEnergy    `json:"energy,omitempty"`
	FixedFee         *TariffFixed     `json:"fixedFee,omitempty"`
	IdleTime         *TariffTime      `json:"idleTime,omitempty"`
	MaxCost          *Price           `json:"maxCost,omitempty"`
	MinCost          *Price           `json:"minCost,omitempty"`
	ReservationFixed *TariffFixed     `json:"reservationFixed,omitempty"`
	ReservationTime  *TariffTime      `json:"reservationTime,omitempty"`
	TariffID         string           `json:"tariffId" validate:"required,max=60"`
	ValidFrom        *time.Time       `json:"validFrom,omitempty"`
}

type TariffAssignment struct {
	CustomData CustomData     `json:"customData,omitempty"`
	EVSEIds    []int          `json:"evseIds,omitempty" validate:"min=1"`
	IDTokens   []string       `json:"idTokens,omitempty" validate:"min=1"`
	TariffID   string         `json:"tariffId" validate:"required,max=60"`
	TariffKind TariffKindEnum `json:"tariffKind" validate:"required"`
	ValidFrom  *time.Time     `json:"validFrom,omitempty"`
}

type TariffConditions struct {
	CustomData      CustomData      `json:"customData,omitempty"`
	DayOfWeek       []DayOfWeekEnum `json:"dayOfWeek,omitempty" validate:"min=1,max=7"`
	EndTimeOfDay    *string         `json:"endTimeOfDay,omitempty"`
	EVSEKind        *EVSEKindEnum   `json:"evseKind,omitempty"`
	MaxChargingTime *int            `json:"maxChargingTime,omitempty"`
	MaxCurrent      *float64        `json:"maxCurrent,omitempty"`
	MaxEnergy       *float64        `json:"maxEnergy,omitempty"`
	MaxIdleTime     *int            `json:"maxIdleTime,omitempty"`
	MaxPower        *float64        `json:"maxPower,omitempty"`
	MaxTime         *int            `json:"maxTime,omitempty"`
	MinChargingTime *int            `json:"minChargingTime,omitempty"`
	MinCurrent      *float64        `json:"minCurrent,omitempty"`
	MinEnergy       *float64        `json:"minEnergy,omitempty"`
	MinIdleTime     *int            `json:"minIdleTime,omitempty"`
	MinPower        *float64        `json:"minPower,omitempty"`
	MinTime         *int            `json:"minTime,omitempty"`
	StartTimeOfDay  *string         `json:"startTimeOfDay,omitempty"`
	ValidFromDate   *string         `json:"validFromDate,omitempty"`
	ValidToDate     *string         `json:"validToDate,omitempty"`
}

type TariffConditionsFixed struct {
	CustomData         CustomData      `json:"customData,omitempty"`
	DayOfWeek          []DayOfWeekEnum `json:"dayOfWeek,omitempty" validate:"min=1,max=7"`
	EndTimeOfDay       *string         `json:"endTimeOfDay,omitempty"`
	EVSEKind           *EVSEKindEnum   `json:"evseKind,omitempty"`
	PaymentBrand       *string         `json:"paymentBrand,omitempty" validate:"max=20"`
	PaymentRecognition *string         `json:"paymentRecognition,omitempty" validate:"max=20"`
	StartTimeOfDay     *string         `json:"startTimeOfDay,omitempty"`
	ValidFromDate      *string         `json:"validFromDate,omitempty"`
	ValidToDate        *string         `json:"validToDate,omitempty"`
}

type TariffEnergy struct {
	CustomData CustomData          `json:"customData,omitempty"`
	Prices     []TariffEnergyPrice `json:"prices" validate:"required,min=1"`
	TaxRates   []TaxRate           `json:"taxRates,omitempty" validate:"min=1,max=5"`
}

type TariffEnergyPrice struct {
	Conditions *TariffConditions `json:"conditions,omitempty"`
	CustomData CustomData        `json:"customData,omitempty"`
	PriceKwh   float64           `json:"priceKwh" validate:"required"`
}

type TariffFixed struct {
	CustomData CustomData         `json:"customData,omitempty"`
	Prices     []TariffFixedPrice `json:"prices" validate:"required,min=1"`
	TaxRates   []TaxRate          `json:"taxRates,omitempty" validate:"min=1,max=5"`
}

type TariffFixedPrice struct {
	Conditions *TariffConditionsFixed `json:"conditions,omitempty"`
	CustomData CustomData             `json:"customData,omitempty"`
	PriceFixed float64                `json:"priceFixed" validate:"required"`
}

type TariffTime struct {
	CustomData CustomData        `json:"customData,omitempty"`
	Prices     []TariffTimePrice `json:"prices" validate:"required,min=1"`
	TaxRates   []TaxRate         `json:"taxRates,omitempty" validate:"min=1,max=5"`
}

type TariffTimePrice struct {
	Conditions  *TariffConditions `json:"conditions,omitempty"`
	CustomData  CustomData        `json:"customData,omitempty"`
	PriceMinute float64           `json:"priceMinute" validate:"required"`
}

type TaxRate struct {
	CustomData CustomData `json:"customData,omitempty"`
	Stack      *int       `json:"stack,omitempty" validate:"gte=0.0"`
	Tax        float64    `json:"tax" validate:"required"`
	TypeValue  string     `json:"type" validate:"required,max=20"`
}

type TaxRule struct {
	AppliesToEnergyFee          bool           `json:"appliesToEnergyFee" validate:"required"`
	AppliesToMinimumMaximumCost bool           `json:"appliesToMinimumMaximumCost" validate:"required"`
	AppliesToOverstayFee        bool           `json:"appliesToOverstayFee" validate:"required"`
	AppliesToParkingFee         bool           `json:"appliesToParkingFee" validate:"required"`
	CustomData                  CustomData     `json:"customData,omitempty"`
	TaxIncludedInPrice          *bool          `json:"taxIncludedInPrice,omitempty"`
	TaxRate                     RationalNumber `json:"taxRate" validate:"required"`
	TaxRuleID                   int            `json:"taxRuleID" validate:"required,gte=0.0"`
	TaxRuleName                 *string        `json:"taxRuleName,omitempty" validate:"max=100"`
}

type TotalCost struct {
	ChargingTime     *Price         `json:"chargingTime,omitempty"`
	Currency         string         `json:"currency" validate:"required,max=3"`
	CustomData       CustomData     `json:"customData,omitempty"`
	Energy           *Price         `json:"energy,omitempty"`
	Fixed            *Price         `json:"fixed,omitempty"`
	IdleTime         *Price         `json:"idleTime,omitempty"`
	ReservationFixed *Price         `json:"reservationFixed,omitempty"`
	ReservationTime  *Price         `json:"reservationTime,omitempty"`
	Total            TotalPrice     `json:"total" validate:"required"`
	TypeOfCost       TariffCostEnum `json:"typeOfCost" validate:"required"`
}

type TotalPrice struct {
	CustomData CustomData `json:"customData,omitempty"`
	ExclTax    *float64   `json:"exclTax,omitempty"`
	InclTax    *float64   `json:"inclTax,omitempty"`
}

type TotalUsage struct {
	ChargingTime    int        `json:"chargingTime" validate:"required"`
	CustomData      CustomData `json:"customData,omitempty"`
	Energy          float64    `json:"energy" validate:"required"`
	IdleTime        int        `json:"idleTime" validate:"required"`
	ReservationTime *int       `json:"reservationTime,omitempty"`
}

type Transaction struct {
	ChargingState     *ChargingStateEnum `json:"chargingState,omitempty"`
	CustomData        CustomData         `json:"customData,omitempty"`
	OperationMode     *OperationModeEnum `json:"operationMode,omitempty"`
	RemoteStartID     *int               `json:"remoteStartId,omitempty"`
	StoppedReason     *ReasonEnum        `json:"stoppedReason,omitempty"`
	TariffID          *string            `json:"tariffId,omitempty" validate:"max=60"`
	TimeSpentCharging *int               `json:"timeSpentCharging,omitempty"`
	TransactionID     string             `json:"transactionId" validate:"required,max=36"`
	TransactionLimit  *TransactionLimit  `json:"transactionLimit,omitempty"`
}

type TransactionLimit struct {
	CustomData CustomData `json:"customData,omitempty"`
	MaxCost    *float64   `json:"maxCost,omitempty"`
	MaxEnergy  *float64   `json:"maxEnergy,omitempty"`
	MaxSoC     *int       `json:"maxSoC,omitempty" validate:"gte=0.0,lte=100.0"`
	MaxTime    *int       `json:"maxTime,omitempty"`
}

type UnitOfMeasure struct {
	CustomData CustomData `json:"customData,omitempty"`
	Multiplier *int       `json:"multiplier,omitempty"`
	Unit       *string    `json:"unit,omitempty" validate:"max=20"`
}

type V2XChargingParameters struct {
	CustomData            CustomData `json:"customData,omitempty"`
	EVMaxEnergyRequest    *float64   `json:"evMaxEnergyRequest,omitempty"`
	EVMaxV2XEnergyRequest *float64   `json:"evMaxV2XEnergyRequest,omitempty"`
	EVMinEnergyRequest    *float64   `json:"evMinEnergyRequest,omitempty"`
	EVMinV2XEnergyRequest *float64   `json:"evMinV2XEnergyRequest,omitempty"`
	EVTargetEnergyRequest *float64   `json:"evTargetEnergyRequest,omitempty"`
	MaxChargeCurrent      *float64   `json:"maxChargeCurrent,omitempty"`
	MaxChargePower        *float64   `json:"maxChargePower,omitempty"`
	MaxChargePowerL2      *float64   `json:"maxChargePower_L2,omitempty"`
	MaxChargePowerL3      *float64   `json:"maxChargePower_L3,omitempty"`
	MaxDischargeCurrent   *float64   `json:"maxDischargeCurrent,omitempty"`
	MaxDischargePower     *float64   `json:"maxDischargePower,omitempty"`
	MaxDischargePowerL2   *float64   `json:"maxDischargePower_L2,omitempty"`
	MaxDischargePowerL3   *float64   `json:"maxDischargePower_L3,omitempty"`
	MaxVoltage            *float64   `json:"maxVoltage,omitempty"`
	MinChargeCurrent      *float64   `json:"minChargeCurrent,omitempty"`
	MinChargePower        *float64   `json:"minChargePower,omitempty"`
	MinChargePowerL2      *float64   `json:"minChargePower_L2,omitempty"`
	MinChargePowerL3      *float64   `json:"minChargePower_L3,omitempty"`
	MinDischargeCurrent   *float64   `json:"minDischargeCurrent,omitempty"`
	MinDischargePower     *float64   `json:"minDischargePower,omitempty"`
	MinDischargePowerL2   *float64   `json:"minDischargePower_L2,omitempty"`
	MinDischargePowerL3   *float64   `json:"minDischargePower_L3,omitempty"`
	MinVoltage            *float64   `json:"minVoltage,omitempty"`
	TargetSoC             *int       `json:"targetSoC,omitempty" validate:"gte=0.0,lte=100.0"`
}

type V2XFreqWattPoint struct {
	CustomData CustomData `json:"customData,omitempty"`
	Frequency  float64    `json:"frequency" validate:"required"`
	Power      float64    `json:"power" validate:"required"`
}

type V2XSignalWattPoint struct {
	CustomData CustomData `json:"customData,omitempty"`
	Power      float64    `json:"power" validate:"required"`
	Signal     int        `json:"signal" validate:"required"`
}

type VPN struct {
	CustomData CustomData `json:"customData,omitempty"`
	Group      *string    `json:"group,omitempty" validate:"max=50"`
	Key        string     `json:"key" validate:"required,max=255"`
	Password   string     `json:"password" validate:"required,max=64"`
	Server     string     `json:"server" validate:"required,max=2000"`
	TypeValue  VPNEnum    `json:"type" validate:"required"`
	User       string     `json:"user" validate:"required,max=50"`
}

type Variable struct {
	CustomData CustomData `json:"customData,omitempty"`
	Instance   *string    `json:"instance,omitempty" validate:"max=50"`
	Name       string     `json:"name" validate:"required,max=50"`
}

type VariableAttribute struct {
	Constant   *bool           `json:"constant,omitempty"`
	CustomData CustomData      `json:"customData,omitempty"`
	Mutability *MutabilityEnum `json:"mutability,omitempty"`
	Persistent *bool           `json:"persistent,omitempty"`
	TypeValue  *AttributeEnum  `json:"type,omitempty"`
	Value      *string         `json:"value,omitempty" validate:"max=2500"`
}

type VariableCharacteristics struct {
	CustomData         CustomData `json:"customData,omitempty"`
	DataType           DataEnum   `json:"dataType" validate:"required"`
	MaxElements        *int       `json:"maxElements,omitempty" validate:"gte=1.0"`
	MaxLimit           *float64   `json:"maxLimit,omitempty"`
	MinLimit           *float64   `json:"minLimit,omitempty"`
	SupportsMonitoring bool       `json:"supportsMonitoring" validate:"required"`
	Unit               *string    `json:"unit,omitempty" validate:"max=16"`
	ValuesList         *string    `json:"valuesList,omitempty" validate:"max=1000"`
}

type VariableMonitoring struct {
	CustomData            CustomData            `json:"customData,omitempty"`
	EventNotificationType EventNotificationEnum `json:"eventNotificationType" validate:"required"`
	ID                    int                   `json:"id" validate:"required,gte=0.0"`
	Severity              int                   `json:"severity" validate:"required,gte=0.0"`
	Transaction           bool                  `json:"transaction" validate:"required"`
	TypeValue             MonitorEnum           `json:"type" validate:"required"`
	Value                 float64               `json:"value" validate:"required"`
}

type VoltageParams struct {
	CustomData           CustomData                `json:"customData,omitempty"`
	Hv10MinMeanTripDelay *float64                  `json:"hv10MinMeanTripDelay,omitempty"`
	Hv10MinMeanValue     *float64                  `json:"hv10MinMeanValue,omitempty"`
	PowerDuringCessation *PowerDuringCessationEnum `json:"powerDuringCessation,omitempty"`
}
