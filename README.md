<h1 align="center">go-apns</h1>
<p align="center">Simple apns client library for Go</p>


[TOC]


## Features

- Support APNs HTTP/2 connection
- Suport voip

- Support read certificate from .pem and .p12 file

## Example voip

```go
import (
	go_apns "github.com/trungnq2710/go-apns"
	"log"
)

func Voip() {
	// read cert
	cert, err := go_apns.CertificateFromPemFile("cert.pem")
	if err != nil {
		log.Fatal.(err)
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
		log.Fatal.(err)
	}

	log.Println(res)
}
```

