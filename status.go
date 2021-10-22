// Created at 10/21/2021 2:33 PM
// Developer: trungnq2710 (trungnq2710@gmail.com)

package go_apns

const (
	StatusCodeSuccess                 = 200
	StatusCodeBadRequest              = 400
	StatusCodeErrorCertificateOrToken = 403
	StatusCodeInvalidPath             = 404
	StatusCodeMethodUnsupported       = 405
	StatusCodeErrorDeviceToken        = 410
	StatusCodePayloadWasTooLarge      = 413
	StatusCodeTooManyRequests         = 429
	StatusCodeInternalServerError     = 500
	StatusCodeServerShuttingDown      = 503
)

var statusText = map[int]string{
	StatusCodeSuccess:                 "Success",
	StatusCodeBadRequest:              "Bad request",
	StatusCodeErrorCertificateOrToken: "There was an error with the certificate or with the provider’s authentication token",
	StatusCodeInvalidPath:             "The request contained an invalid :path value",
	StatusCodeMethodUnsupported:       "The request used an invalid :method value. Only POST requests are supported",
	StatusCodeErrorDeviceToken:        "The device token is no longer active for the topic",
	StatusCodePayloadWasTooLarge:      "The notification payload was too large",
	StatusCodeTooManyRequests:         "The server received too many requests for the same device token",
	StatusCodeInternalServerError:     "Internal server error",
	StatusCodeServerShuttingDown:      "The server is shutting down and unavailable",
}

func StatusText(code int) string {
	return statusText[code]
}
