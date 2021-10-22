// Created at 10/21/2021 12:48 AM
// Developer: trungnq2710 (trungnq2710@gmail.com)

package go_apns

const (
	DefaultDevelopmentEndpoint = "https://api.development.push.apple.com"
	DefaultProductionEndpoint  = "https://api.push.apple.com"
	DefaultPath                = "/3/device"
	DefaultPort                = 443
	OtherPort                  = 2197
)

type PushType string

const (
	PushTypeAlert        PushType = "alert"
	PushTypeBackground   PushType = "background"
	PushTypeLocation     PushType = "location"
	PushTypeVOIP         PushType = "voip"
	PushTypeComplication PushType = "complication"
	PushTypeFileProvider PushType = "fileprovider"
	PushTypeMDM          PushType = "mdm"
)

type Priority int

const (
	PriorityLow  Priority = 5
	PriorityHigh Priority = 10
)

const (
	apnsReasonNotForTopic  = "DeviceTokenNotForTopic"
	apnsReasonBadToken     = "BadDeviceToken"
	apnsReasonExpiredToken = "ExpiredProviderToken"
	apnsReasonInvalidToken = "InvalidProviderToken"
	apnsReasonUnregistered = "Unregistered"
)
