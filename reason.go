// Created at 10/21/2021 2:36 PM
// Developer: trungnq2710 (trungnq2710@gmail.com)

package go_apns

const (
	ReasonBadCollapseId               = "BadCollapseId"
	ReasonBadDeviceToken              = "BadDeviceToken"
	ReasonBadExpirationDate           = "BadExpirationDate"
	ReasonBadMessageId                = "BadMessageId"
	ReasonBadPriority                 = "BadPriority"
	ReasonBadTopic                    = "BadTopic"
	ReasonDeviceTokenNotForTopic      = "DeviceTokenNotForTopic"
	ReasonDuplicateHeaders            = "DuplicateHeaders"
	ReasonIdleTimeout                 = "IdleTimeout"
	ReasonInvalidPushType             = "InvalidPushType"
	ReasonMissingDeviceToken          = "MissingDeviceToken"
	ReasonMissingTopic                = "MissingTopic"
	ReasonPayloadEmpty                = "PayloadEmpty"
	ReasonTopicDisallowed             = "TopicDisallowed"
	ReasonBadCertificate              = "BadCertificate"
	ReasonBadCertificateEnvironment   = "BadCertificateEnvironment"
	ReasonExpiredProviderToken        = "ExpiredProviderToken"
	ReasonForbidden                   = "Forbidden"
	ReasonInvalidProviderToken        = "InvalidProviderToken"
	ReasonMissingProviderToken        = "MissingProviderToken"
	ReasonBadPath                     = "BadPath"
	ReasonMethodNotAllowed            = "MethodNotAllowed"
	ReasonUnregistered                = "Unregistered"
	ReasonPayloadTooLarge             = "PayloadTooLarge"
	ReasonTooManyProviderTokenUpdates = "TooManyProviderTokenUpdates"
	ReasonTooManyRequests             = "TooManyRequests"
	ReasonInternalServerError         = "InternalServerError"
	ReasonServiceUnavailable          = "ServiceUnavailable"
	ReasonShutdown                    = "Shutdown"
)

var reasonDescription = map[string]string{
	ReasonBadCollapseId:               "The collapse identifier exceeds the maximum allowed size",
	ReasonBadDeviceToken:              "The specified device token is invalid. Verify that the request contains a valid token and that the token matches the environment",
	ReasonBadExpirationDate:           "The apns-expiration value is invalid",
	ReasonBadMessageId:                "The apns-id value is invalid",
	ReasonBadPriority:                 "The apns-priority value is invalid",
	ReasonBadTopic:                    "The apns-topic value is invalid",
	ReasonDeviceTokenNotForTopic:      "The device token doesn’t match the specified topic",
	ReasonDuplicateHeaders:            "One or more headers are repeated",
	ReasonIdleTimeout:                 "Idle timeout",
	ReasonInvalidPushType:             "The apns-push-type value is invalid",
	ReasonMissingDeviceToken:          "The device token isn’t specified in the request :path. Verify that the :path header contains the device token",
	ReasonMissingTopic:                "The apns-topic header of the request isn’t specified and is required. The apns-topic header is mandatory when the client is connected using a certificate that supports multiple topics",
	ReasonPayloadEmpty:                "The message payload is empty",
	ReasonTopicDisallowed:             "Pushing to this topic is not allowed",
	ReasonBadCertificate:              "The certificate is invalid",
	ReasonBadCertificateEnvironment:   "The client certificate is for the wrong environment",
	ReasonExpiredProviderToken:        "The provider token is stale and a new token should be generated",
	ReasonForbidden:                   "The specified action is not allowed",
	ReasonInvalidProviderToken:        "The provider token is not valid, or the token signature can't be verified",
	ReasonMissingProviderToken:        "No provider certificate was used to connect to APNs, and the authorization header is missing or no provider token is specified",
	ReasonBadPath:                     "The request contained an invalid :path value",
	ReasonMethodNotAllowed:            "The specified :method value isn’t POST",
	ReasonUnregistered:                "The device token is inactive for the specified topic. There is no need to send further pushes to the same device token, unless your application retrieves the same device token, see Registering Your App with APNs",
	ReasonPayloadTooLarge:             "The message payload is too large. For information about the allowed payload size, see Create and Send a POST Request to APNs",
	ReasonTooManyProviderTokenUpdates: "The provider’s authentication token is being updated too often. Update the authentication token no more than once every 20 minutes",
	ReasonTooManyRequests:             "Too many requests were made consecutively to the same device token",
	ReasonInternalServerError:         "An internal server error occurred",
	ReasonServiceUnavailable:          "The service is unavailable",
	ReasonShutdown:                    "The APNs server is shutting down",
}

func ReasonDescription(reason string) string {
	return reasonDescription[reason]
}
