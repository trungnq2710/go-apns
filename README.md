<h1 align="center">go-apns</h1>
<p align="center">Simple apns client library for Go</p>


[TOC]

## Features

- Support APNs HTTP/2 connection
- Suport voip

- Support read certificate from .pem and .p12 file

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

## Notification

- Create notification with builder

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

