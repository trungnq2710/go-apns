// Created at 10/25/2021 3:43 PM
// Developer: trungnq2710 (trungnq2710@gmail.com)

package go_apns

type aps struct {
	Alert             interface{} `json:"alert,omitempty"`
	Badge             int         `json:"badge,omitempty"`
	Sound             interface{} `json:"sound,omitempty"`
	ThreadId          string      `json:"thread-id,omitempty"`
	Category          string      `json:"category,omitempty"`
	ContentAvailable  int         `json:"content-available,omitempty"`
	MutableContent    int         `json:"mutable-content,omitempty"`
	TargetContentId   string      `json:"target-content-id,omitempty"`
	InterruptionLevel string      `json:"interruption-level,omitempty"`
	RelevanceScore    int         `json:"relevance-score,omitempty"`
}
