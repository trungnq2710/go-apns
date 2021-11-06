// Created at 10/21/2021 4:50 PM
// Developer: trungnq2710 (trungnq2710@gmail.com)

package go_apns

type ApsAlert struct {
	Title           string   `json:"title,omitempty"`
	Subtitle        string   `json:"subtitle,omitempty"`
	Body            string   `json:"body,omitempty"`
	LaunchImage     string   `json:"launch-image,omitempty"`
	TitleLocKey     string   `json:"title-loc-key,omitempty"`
	TitleLocArgs    []string `json:"title-loc-args,omitempty"`
	SubtitleLocKey  string   `json:"subtitle-loc-key,omitempty"`
	SubtitleLocArgs []string `json:"subtitle-loc-args,omitempty"`
	LocKey          string   `json:"loc-key,omitempty"`
	LocArgs         []string `json:"loc-args,omitempty"`
}

func NewApsAlert() *ApsAlert {
	return &ApsAlert{}
}

// The title of the notification. Apple Watch displays this string in the short look notification
// interface. Specify a string that’s quickly understood by the user
func (a *ApsAlert) SetTitle(i string) *ApsAlert {
	a.Title = i
	return a
}

// Additional information that explains the purpose of the notification
func (a *ApsAlert) SetSubtitle(i string) *ApsAlert {
	a.Subtitle = i
	return a
}

// The content of the alert message
func (a *ApsAlert) SetBody(i string) *ApsAlert {
	a.Body = i
	return a
}

// The name of the launch image file to display. If the user chooses to launch your app, the
// contents of the specified image or storyboard file are displayed instead of your app’s
// normal launch image
func (a *ApsAlert) SetLaunchImage(i string) *ApsAlert {
	a.LaunchImage = i
	return a
}

// The key for a localized title string. Specify this key instead of the title key to retrieve
// the title from your app’s Localizable.strings files. The value must contain the name of a
// key in your strings file
func (a *ApsAlert) SetTitleLocKey(i string) *ApsAlert {
	a.TitleLocKey = i
	return a
}

// An array of strings containing replacement values for variables in your title string.
// Each %@ character in the string specified by the title-loc-key is replaced by a value
// from this array. The first item in the array replaces the first instance of the %@ character
// in the string, the second item replaces the second instance, and so on
func (a *ApsAlert) SetTitleLocArgs(i []string) *ApsAlert {
	a.TitleLocArgs = i
	return a
}

// The key for a localized subtitle string. Use this key, instead of the subtitle key, to retrieve
// the subtitle from your app’s Localizable.strings file. The value must contain the name of a key
// in your strings file
func (a *ApsAlert) SetSubtitleLocKey(i string) *ApsAlert {
	a.SubtitleLocKey = i
	return a
}

// An array of strings containing replacement values for variables in your title string.
// Each %@ character in the string specified by subtitle-loc-key is replaced by a value from this
// array. The first item in the array replaces the first instance of the %@ character in the string,
// the second item replaces the second instance, and so on.
func (a *ApsAlert) SetSubtitleLocArgs(i []string) *ApsAlert {
	a.SubtitleLocArgs = i
	return a
}

// The key for a localized message string. Use this key, instead of the body key, to retrieve the
// message text from your app’s Localizable.strings file. The value must contain the name of a key
// in your strings file.
func (a *ApsAlert) SetLocKey(i string) *ApsAlert {
	a.LocKey = i
	return a
}

// An array of strings containing replacement values for variables in your message text. Each %@
// character in the string specified by loc-key is replaced by a value from this array. The first
// item in the array replaces the first instance of the %@ character in the string, the second
// item replaces the second instance, and so on
func (a *ApsAlert) SetLocArgs(i []string) *ApsAlert {
	a.LocArgs = i
	return a
}
