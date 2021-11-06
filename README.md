<h1 align="center">go-apns</h1>
<p align="center">Simple apns client library for Go</p>


[TOC]

## Features

- Support APNs HTTP/2 connection
- Suport all push type
- Support read certificate from .pem and .p12 file
- Create with builder

## Certificate

- Read from .pem file

```go
import (
	"crypto/tls"
	go_apns "github.com/trungnq2710/go-apns"
	"log"
)

func ReadFromPemFile()  {
	var (
		cert *tls.Certificate
		err error
	)

	cert, err = go_apns.CertificateFromPemFile("cert.pem")
	if err != nil {
		log.Fatal(err)
	}
	
	_ = cert
}
```

- Read from .p12 file

```go
import (
	"crypto/tls"
	go_apns "github.com/trungnq2710/go-apns"
	"log"
)

func ReadFromPemFile()  {
	var (
		cert *tls.Certificate
		err error
	)

	cert, err := go_apns.CertificateFromP12File("cert.p12")
	if err != nil {
		log.Fatal(err)
	}
	
	_ = cert
}
```

## Payload

- Create payload with builder
- Builder method

| Func                                                         | Description                                 |
| ------------------------------------------------------------ | ------------------------------------------- |
| ```func NewPayload() *Payload```                             | Return pointer of Payload                   |
| ```func (p *Payload) SetAlertString(i string) *Payload```    | ```{"aps":{"alert": value}}```              |
| ```func (p *Payload) SetAlert(i *ApsAlert) *Payload```       | ```{"aps":{"alert": value}}```              |
| ```func (p *Payload) SetBadge(i int) *Payload```             | ```{"aps":{"badge": value}}```              |
| ```func (p *Payload) SetSoundString(i string) *Payload```    | ```{"aps":{"sound": value}}```              |
| ```func (p *Payload) SetSound(i *ApsSound) *Payload```       | ```{"aps":{"sound": value}}```              |
| ```func (p *Payload) SetThreadId(i string) *Payload```       | ```{"aps":{"thread-id": value}}```          |
| ```func (p *Payload) SetCategory(i string) *Payload```       | ```{"aps":{"category": value}}```           |
| ```func (p *Payload) SetContentAvailable(i int) *Payload```  | ```{"aps":{"content-available": value}}```  |
| ```func (p *Payload) SetMutableContent(i int) *Payload```    | ```{"aps":{"mutable-content": value}}```    |
| ```func (p *Payload) SetTargetContentId(i string) *Payload``` | ```{"aps":{"target-content-id": value}}```  |
| ```func (p *Payload) SetInterruptionLevel(i string) *Payload``` | ```{"aps":{"interruption-level": value}}``` |
| ```func (p *Payload) SetRelevanceScore(i int) *Payload```    | ```{"aps":{"relevance-score": value}}```    |
| ```func (p *Payload) SetCustom(k string, v interface{}) *Payload``` | ```{"aps":{}, key:value}```                 |

- Example

```go
import go_apns "github.com/trungnq2710/go-apns"

func CreatePayload() {
    payload := go_apns.NewPayload()
    	.SetBadge(2)
    	.SetCustom("custom_key", "custom_data")
    _ = payload
}
```



### ApsAlert

- Create aps alert with builder
- Builder method

| Func                                                         |     
| ------------------------------------------------------------ |
|```func NewApsAlert() *ApsAlert``` |
| ```func (a *ApsAlert) SetTitle(i string) *ApsAlert```        |      
| ```func (a *ApsAlert) SetSubtitle(i string) *ApsAlert```     |      
| ```func (a *ApsAlert) SetBody(i string) *ApsAlert```         |      
| ```func (a *ApsAlert) SetLaunchImage(i string) *ApsAlert```  |      
| ```func (a *ApsAlert) SetTitleLocKey(i string) *ApsAlert```  |      
| ```func (a *ApsAlert) SetTitleLocArgs(i []string) *ApsAlert``` |     
| ```func (a *ApsAlert) SetSubtitleLocKey(i string) *ApsAlert``` |      
| ```func (a *ApsAlert) SetSubtitleLocArgs(i []string) *ApsAlert``` |      
| ```func (a *ApsAlert) SetLocKey(i string) *ApsAlert```       |      
| ```func (a *ApsAlert) SetLocArgs(i []string) *ApsAlert```    |      

- Struct

```go
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
```

- Example

```go
import go_apns "github.com/trungnq2710/go-apns"

func CreateAlert() {
    alert := go_apns.NewApsAlert()
    	.SetTitle("title")
    	.SetButitle("subtitle")
    	.SetBody("body")
    _ = alert
}
```

### ApsSound

- Create aps sound with builder
- Builder method

| Func                                                  |
| ----------------------------------------------------- |
| ```func NewApsSound() *ApsSound```                    |
| ```func (a *ApsSound) SetCritical(i int) *ApsSound``` |
| ```func (a *ApsSound) SetName(i string) *ApsSound```  |
| ```func (a *ApsSound) SetVolume(i int) *ApsSound```   |

- Struct

```go
type ApsSound struct {
	Critical int    `json:"critical,omitempty"`
	Name     string `json:"name,omitempty"`
	Volume   int    `json:"volume,omitempty"`
}
```

## Notification

- Create notification with builder
- Builder method

| Func                                                         |
| ------------------------------------------------------------ |
| ```func NewNotification() *Notification```                   |
| ```func (n *Notification) PushType(i PushType) *Notification``` |
| ```func (n *Notification) ApnsID(i string) *Notification```  |
| ```func (n *Notification) Expiration(i time.Time) *Notification``` |
| ```func (n *Notification) Priority(i Priority) *Notification``` |
| ```func (n *Notification) Topic(i string) *Notification```   |
| ```func (n *Notification) CollapseID(i string) *Notification``` |
| ```func (n *Notification) DeviceToken(i string) *Notification``` |
| ```func (n *Notification) Payload(i *Payload) *Notification``` |

## Client

### Create client

- Use product env

```go
go_apns.NewClient(cert).Production()
```

- Use development env (default is development env)

```go
go_apns.NewClient(cert).Development()
```

### Push action

- Push

```go
client := go_apns.NewClient(cert).Production()

res, err := client.Push(notification)
if err != nil {
	log.Fatal(err)
}
```

- Push with context

```go
client := go_apns.NewClient(cert).Production()

res, err := client.PushWithCtx(context.TODO(), notification)
if err != nil {
	log.Fatal(err)
}
```



## Example

### Example voip

```go
import (
	go_apns "github.com/trungnq2710/go-apns"
	"log"
)

func Voip() {
	// read cert
	cert, err := go_apns.CertificateFromPemFile("cert.pem")
	if err != nil {
		log.Fatal(err)
	}

	// init payload
	payload := go_apns.NewPayload().
		SetCustom("user_id", "777000").
		SetCustom("call_id", "10303")

	// init notification
	notification := go_apns.NewNotification().
		PushType(go_apns.PushTypeVOIP).
		Topic("com.app.chat").
		DeviceToken("token").
		Payload(payload)

	// note: default development env
	client := go_apns.NewClient(cert).Production()

	res, err := client.Push(notification)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(res)
}
```

