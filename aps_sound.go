// Created at 10/21/2021 4:51 PM
// Developer: trungnq2710 (trungnq2710@gmail.com)

package go_apns

type ApsSound struct {
	Critical int `json:"critical,omitempty"`
	Name string `json:"name,omitempty"`
	Volume int `json:"volume,omitempty"`
}

func NewApsSound() *ApsSound {
	return &ApsSound{}
}

// The critical alert flag. Set to 1 to enable the critical alert
func (a *ApsSound) SetCritical(i int) *ApsSound {
	a.Critical = i
	return a
}

// The name of a sound file in your app’s main bundle or in the Library/Sounds folder of your
// app’s container directory. Specify the string “default” to play the system sound.
func (a *ApsSound) SetName(i string) *ApsSound {
	a.Name = i
	return a
}

// The volume for the critical alert’s sound. Set this to a value between 0 (silent) and
// 1 (full volume)
func (a *ApsSound) SetVolume(i int) *ApsSound {
	a.Volume = i
	return a
}
