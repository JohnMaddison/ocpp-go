package ocpp16

import "time"

// CertificateHashData corresponds to the CertificateHashDataType.
type CertificateHashData struct {
	HashAlgorithm  HashAlgorithmEnumType `json:"hashAlgorithm" validate:"required,oneof=SHA256 SHA384 SHA512"`
	IssuerNameHash string                `json:"issuerNameHash" validate:"required,max=128"`
	IssuerKeyHash  string                `json:"issuerKeyHash" validate:"required,max=128"`
	SerialNumber   string                `json:"serialNumber" validate:"required,max=40"`
}

// ChargingProfile represents a charging profile.
type ChargingProfile struct {
	ChargingProfileID      int                        `json:"chargingProfileId"`
	TransactionID          *int                       `json:"transactionId,omitempty"`
	StackLevel             int                        `json:"stackLevel" validate:"gte=0"`
	ChargingProfilePurpose ChargingProfilePurposeType `json:"chargingProfilePurpose" validate:"required"`
	ChargingProfileKind    ChargingProfileKindType    `json:"chargingProfileKind" validate:"required"`
	RecurrencyKind         *RecurrencyKindType        `json:"recurrencyKind,omitempty"`
	ValidFrom              *time.Time                 `json:"validFrom,omitempty"`
	ValidTo                *time.Time                 `json:"validTo,omitempty"`
	ChargingSchedule       ChargingSchedule           `json:"chargingSchedule" validate:"required"`
}

// ChargingSchedule represents a charging schedule.
type ChargingSchedule struct {
	Duration               *int                     `json:"duration,omitempty"`
	StartSchedule          *time.Time               `json:"startSchedule,omitempty"`
	ChargingRateUnit       ChargingRateUnitType     `json:"chargingRateUnit" validate:"required"`
	ChargingSchedulePeriod []ChargingSchedulePeriod `json:"chargingSchedulePeriod" validate:"required,min=1"`
	MinChargingRate        *float64                 `json:"minChargingRate,omitempty"`
}

// ChargingSchedulePeriod represents a period in a charging schedule.
type ChargingSchedulePeriod struct {
	StartPeriod  int     `json:"startPeriod"`
	Limit        float64 `json:"limit" validate:"gte=0"`
	NumberPhases *int    `json:"numberPhases,omitempty"`
}

// IdTagInfo contains information about an identifier.
type IdTagInfo struct {
	ExpiryDate  *time.Time          `json:"expiryDate,omitempty"`
	ParentIdTag *string             `json:"parentIdTag,omitempty" validate:"max=20"`
	Status      AuthorizationStatus `json:"status" validate:"required"`
}

// KeyValue is a key-value pair for configuration settings.
type KeyValue struct {
	Key      string  `json:"key" validate:"required,max=50"`
	Readonly bool    `json:"readonly"`
	Value    *string `json:"value,omitempty" validate:"max=500"`
}

// LogParameters specifies parameters for log retrieval.
type LogParameters struct {
	RemoteLocation  string     `json:"remoteLocation" validate:"required,max=512"`
	OldestTimestamp *time.Time `json:"oldestTimestamp,omitempty"`
	LatestTimestamp *time.Time `json:"latestTimestamp,omitempty"`
}

// MeterValue represents a meter reading.
type MeterValue struct {
	Timestamp    time.Time      `json:"timestamp"`
	SampledValue []SampledValue `json:"sampledValue" validate:"required,min=1"`
}

// SampledValue is a single value from a meter.
type SampledValue struct {
	Value     string          `json:"value" validate:"required"`
	Context   *ReadingContext `json:"context,omitempty"`
	Format    *ValueFormat    `json:"format,omitempty"`
	Measurand *Measurand      `json:"measurand,omitempty"`
	Phase     *string         `json:"phase,omitempty" validate:"oneof=L1 L2 L3 N L1-N L2-N L3-N L1-L2 L2-L3 L3-L1"`
	Location  *Location       `json:"location,omitempty"`
	Unit      *string         `json:"unit,omitempty"`
}

// AuthorizationData is used in the local authorization list.
type AuthorizationData struct {
	IdTag     string     `json:"idTag" validate:"required,max=20"`
	IdTagInfo *IdTagInfo `json:"idTagInfo,omitempty"`
}
