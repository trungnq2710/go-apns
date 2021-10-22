// Created at 10/21/2021 1:25 AM
// Developer: trungnq2710 (trungnq2710@gmail.com)

package go_apns

import (
	"time"
)

// TODO impl token

type Notification struct {
	// authorization string
	pushType   PushType
	apnsID     string
	expiration time.Time
	priority   Priority
	topic      string
	collapseID string

	deviceToken string
	payload     *Payload
}

func NewNotification() *Notification {
	return &Notification{
		pushType: PushTypeAlert,
	}
}

// (Required for token-based authentication) The value of this header is bearer <provider_token>,
// where <provider_token> is the encrypted token that authorizes you to send notifications for
// the specified topic. APNs ignores this header if you use certificate-based authentication
//func (n *Notification) Authorization(i string) *Notification {
//	n.authorization = i
//	return n
//}

// (Required for watchOS 6 and later; recommended for macOS, iOS, tvOS, and iPadOS)
// The value of this header must accurately reflect the contents of your notification’s payload.
// If there’s a mismatch, or if the header is missing on required systems, APNs may return an error,
// delay the delivery of the notification, or drop it altogether.
func (n *Notification) PushType(i PushType) *Notification {
	n.pushType = i
	return n
}

// A canonical UUID that is the unique ID for the notification. If an error occurs when sending
// the notification, APNs includes this value when reporting the error to your server. Canonical
// UUIDs are 32 lowercase hexadecimal digits, displayed in five groups separated by hyphens in
// the form 8-4-4-4-12. For example: 123e4567-e89b-12d3-a456-4266554400a0. If you omit this header,
// APNs creates a UUID for you and returns it in its response.
func (n *Notification) ApnsID(i string) *Notification {
	n.apnsID = i
	return n
}

// The date at which the notification is no longer valid. This value is a UNIX epoch expressed in
// seconds (UTC). If the value is nonzero, APNs stores the notification and tries to deliver it at
// least once, repeating the attempt as needed until the specified date. If the value is 0, APNs
// attempts to deliver the notification only once and doesn’t store it.
//
// A single APNs attempt may involve retries over multiple network interfaces and connections of
// the destination device. Often these retries span over some time period, depending on the network
// characteristics. In addition, a push notification may take some time on the network after APNs
// sends it to the device. APNs uses best efforts to honor the expiry date without any guarantee.
// If the value is nonzero, the notification may be delivered after the mentioned date. If the value
// is 0, the notification may be delivered with some delay.
func (n *Notification) Expiration(i time.Time) *Notification {
	n.expiration = i
	return n
}

// The priority of the notification. If you omit this header, APNs sets the notification priority to 10.
// Specify 10 to send the notification immediately.
// Specify 5 to send the notification based on power considerations on the user’s device.
func (n *Notification) Priority(i Priority) *Notification {
	n.priority = i
	return n
}

// The topic for the notification. In general, the topic is your app’s bundle ID/app ID.
// It can have a suffix based on the type of push notification. If you’re using a certificate that
// supports PushKit VoIP or watchOS complication notifications, you must include this header with
// bundle ID of you app and if applicable, the proper suffix. If you’re using token-based authentication
// with APNs, you must include this header with the correct bundle ID and suffix combination
func (n *Notification) Topic(i string) *Notification {
	n.topic = i
	return n
}

// An identifier you use to coalesce multiple notifications into a single notification for the
// user. Typically, each notification request causes a new notification to be displayed on the
// user’s device. When sending the same notification more than once, use the same value in this
// header to coalesce the requests. The value of this key must not exceed 64 bytes.
func (n *Notification) CollapseID(i string) *Notification {
	n.collapseID = i
	return n
}

func (n *Notification) DeviceToken(i string) *Notification {
	n.deviceToken = i
	return n
}
func (n *Notification) Payload(i *Payload) *Notification {
	n.payload = i
	return n
}
