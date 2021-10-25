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
	// Use the alert push type for notifications that trigger a user interaction—for example,
	// an alert, badge, or sound. If you set this push type, the apns-topic header field must
	// use your app’s bundle ID as the topic.
	//
	// If the notification requires immediate action from the user,
	// set notification priority to 10; otherwise use 5.
	//
	// The alert push type is required on watchOS 6 and later.
	// It is recommended on macOS, iOS, tvOS, and iPadOS.
	PushTypeAlert PushType = "alert"

	// Use the background push type for notifications that deliver content in the background,
	// and don’t trigger any user interactions. If you set this push type, the apns-topic header
	// field must use your app’s bundle ID as the topic. Always use priority 5.
	// Using priority 10 is an error.
	//
	// The background push type is required on watchOS 6 and later.
	// It is recommended on macOS, iOS, tvOS, and iPadOS.
	PushTypeBackground PushType = "background"

	// Use the location push type for notifications that request a user’s location. If you set this
	// push type, the apns-topic header field must use your app’s bundle ID with .location-query
	// appended to the end.
	//
	// The location push type is recommended for iOS and iPadOS.
	// It isn’t available on macOS, tvOS, and watchOS.
	//
	// If the location query requires an immediate response from the Location Push Service Extension,
	// set notification apns-priority to 10; otherwise, use 5.
	//
	// The location push type supports only token-based authentication.
	PushTypeLocation PushType = "location"

	// Use the voip push type for notifications that provide information about an incoming
	// Voice-over-IP (VoIP) call. For more information, see Responding to VoIP Notifications from PushKit.
	//
	// If you set this push type, the apns-topic header field must use your app’s bundle ID with
	// .voip appended to the end. If you’re using certificate-based authentication, you must also
	// register the certificate for VoIP services. The topic is then part of the
	// 1.2.840.113635.100.6.3.4 or 1.2.840.113635.100.6.3.6 extension.
	//
	// The voip push type is not available on watchOS.
	// It is recommended on macOS, iOS, tvOS, and iPadOS.
	PushTypeVOIP PushType = "voip"

	// Use the complication push type for notifications that contain update information for a
	// watchOS app’s complications. For more information, see Keeping Your Complications Up to Date.
	//
	// If you set this push type, the apns-topic header field must use your app’s bundle ID
	// with .complication appended to the end. If you’re using certificate-based authentication,
	// you must also register the certificate for WatchKit services. The topic is then part of
	// the 1.2.840.113635.100.6.3.6 extension.
	//
	// The complication push type is recommended for watchOS and iOS.
	// It is not available on macOS, tvOS, and iPadOS.
	PushTypeComplication PushType = "complication"

	// Use the fileprovider push type to signal changes to a File Provider extension. If you set
	// this push type, the apns-topic header field must use your app’s bundle ID with
	// .pushkit.fileprovider appended to the end.
	//
	// The fileprovider push type is not available on watchOS.
	// It is recommended on macOS, iOS, tvOS, and iPadOS.
	PushTypeFileProvider PushType = "fileprovider"

	// Use the mdm push type for notifications that tell managed devices to contact the MDM server.
	// If you set this push type, you must use the topic from the UID attribute in the subject of
	// your MDM push certificate. For more information, see Device Management.
	//
	// The mdm push type is not available on watchOS.
	// It is recommended on macOS, iOS, tvOS, and iPadOS.
	PushTypeMDM PushType = "mdm"
)

type Priority int

const (
	PriorityLow  Priority = 5
	PriorityHigh Priority = 10
)
