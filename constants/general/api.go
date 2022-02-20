package general

import (
	"time"
)

const (
	APITimeDuration2s  time.Duration = 2 * time.Second
	APITimeDuration5s  time.Duration = 5 * time.Second
	APITimeDuration10s time.Duration = 10 * time.Second
	APITimeDuration1m  time.Duration = 1 * time.Minute
)

const (
	APIHeaderContentType   string = "Content-Type"
	APIHeaderBorzoToken    string = "X-DV-Auth-Token"
	APIHeaderJetClientKey  string = "clientkey"
	APIHeaderAuthorization string = "Authorization"
)

const (
	APIHeaderContentTypeJSon           string = "application/json"
	APIHeaderContentTypeFormURLEncoded string = "application/x-www-form-urlencoded"
)

const (
	HandlerErrorAuthInvalid              string = "authorization invalid"
	HandlerErrorResponseKeyIDEmpty       string = "key id cannot be empty"
	HandlerErrorRequestDataNotValid      string = "request data not valid"
	HandlerErrorRequestDataEmpty         string = "request data empty"
	HandlerErrorRequestDataFormatInvalid string = "request data format invalid"
	HandlerErrorCookiesEmpty             string = "key data cannot be empty"
	HandlerErrorCookiesInvalid           string = "key data invalid"
	HandlerErrorKeyIDInvalid             string = "key id invalid"
	HandlerErrorImageSizeTooLarge        string = "image too large, max size 1 Mb"
	HandlerErrorImageDataInvalid         string = "image data invalid"
	HandlerErrorImageDataEmpty           string = "image data cannot be empty"
	HandlerErrorFileSizeTooLarge         string = "file too large, max size 1 Mb"
	HandlerErrorFileDataInvalid          string = "file data invalid"
	HandlerErrorFileDataEmpty            string = "file data cannot be empty"
)
