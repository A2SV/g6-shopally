package contextkeys

// Package contextkeys provides shared context keys for passing request-scoped values.
// Use exported vars of an unexported type to avoid collisions as per context.WithValue docs.

type key string

// Exported variables to be used as context keys across packages.
var (
	RespLang     = key("Accept-Language")
	RespCurrency = key("resp_currency")
	DeviceID     = key("X-Device-ID")
)
