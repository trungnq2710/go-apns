// Created at 10/21/2021 3:18 PM
// Developer: trungnq2710 (trungnq2710@gmail.com)

package go_apns

import "encoding/json"

// Put the JSON payload with the notification’s content into the body of your request.
// The JSON payload must not be compressed and is limited to a maximum size of 4 KB (4096 bytes).
// For a Voice over Internet Protocol (VoIP) notification, the maximum size is 5 KB (5120 bytes).

type Payload struct {
	values map[string]interface{}
}

func NewPayload() *Payload {
	return &Payload{
		values: map[string]interface{}{
			"aps": &aps{},
		},
	}
}

// The information for displaying an alert. A dictionary is recommended.
// If you specify a string, the alert displays your string as the body text
//
// {"aps":{"alert": value}}
//
func (p *Payload) SetAlertString(i string) *Payload {
	p.values["aps"].(*aps).Alert = i
	return p
}

// The information for displaying an alert. A dictionary is recommended.
// If you specify a string, the alert displays your string as the body text
//
// {"aps":{"alert": value}}
//
func (p *Payload) SetAlert(i *ApsAlert) *Payload {
	p.values["aps"].(*aps).Alert = i
	return p
}

// The number to display in a badge on your app’s icon.
// Specify 0 to remove the current badge, if any.
//
// {"aps":{"badge": value}}
//
func (p *Payload) SetBadge(i int) *Payload {
	p.values["aps"].(*aps).Badge = i
	return p
}

// The name of a sound file in your app’s main bundle or in the Library/Sounds folder of your
// app’s container directory. Specify the string “default” to play the system sound. Use this
// key for regular notifications. For critical alerts, use the sound dictionary instead.
//
// {"aps":{"sound": value}}
//
func (p *Payload) SetSoundString(i string) *Payload {
	p.values["aps"].(*aps).Sound = i
	return p
}

// A dictionary that contains sound information for critical alerts.
// For regular notifications, use the sound string instead
//
// {"aps":{"sound": value}}
//
func (p *Payload) SetSound(i *ApsSound) *Payload {
	p.values["aps"].(*aps).Sound = i
	return p
}

// An app-specific identifier for grouping related notifications. This value corresponds to
// the threadIdentifier property in the UNNotificationContent object.
//
// {"aps":{"thread-id": value}}
//
func (p *Payload) SetThreadId(i string) *Payload {
	p.values["aps"].(*aps).ThreadId = i
	return p
}

// The notification’s type. This string must correspond to the identifier of one of the
// UNNotificationCategory objects you register at launch time
//
// {"aps":{"category": value}}
//
func (p *Payload) SetCategory(i string) *Payload {
	p.values["aps"].(*aps).Category = i
	return p
}

// The background notification flag. To perform a silent background update, specify the value
// 1 and don’t include the alert, badge, or sound keys in your payload.
//
// {"aps":{"content-available": value}}
//
func (p *Payload) SetContentAvailable(i int) *Payload {
	p.values["aps"].(*aps).ContentAvailable = i
	return p
}

// The notification service app extension flag. If the value is 1, the system passes the
// notification to your notification service app extension before delivery. Use your
// extension to modify the notification’s content.
//
// {"aps":{"mutable-content": value}}
//
func (p *Payload) SetMutableContent(i int) *Payload {
	p.values["aps"].(*aps).MutableContent = i
	return p
}

// The identifier of the window brought forward. The value of this key will be populated on the
// UNNotificationContent object created from the push payload. Access the value using the
// UNNotificationContent object’s targetContentIdentifier property.
//
// {"aps":{"target-content-id": value}}
//
func (p *Payload) SetTargetContentId(i string) *Payload {
	p.values["aps"].(*aps).TargetContentId = i
	return p
}

// A string that indicates the importance and delivery timing of a notification. The string
// values “passive”, “active”, “time-sensitive”, or “critical” correspond to the
// UNNotificationInterruptionLevel enumeration cases
//
// {"aps":{"interruption-level": value}}
//
func (p *Payload) SetInterruptionLevel(i string) *Payload {
	p.values["aps"].(*aps).InterruptionLevel = i
	return p
}

// The relevance score, a number between 0 and 1, that the system uses to sort the
// notifications from your app. The highest score gets featured in the notification summary.
//
// {"aps":{"relevance-score": value}}
//
func (p *Payload) SetRelevanceScore(i int) *Payload {
	p.values["aps"].(*aps).RelevanceScore = i
	return p
}

// Custom sets a custom key and value on the payload.
// This will add custom key/value data to the notification payload at root level.
//
// {"aps":{}, key:value}
//
func (p *Payload) SetCustom(k string, v interface{}) *Payload {
	p.values[k] = v
	return p
}

// marsahl payload to []byte
func (p *Payload) marshalJSON() ([]byte, error) {
	return json.Marshal(p.values)
}
