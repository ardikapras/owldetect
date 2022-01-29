package constant

type ContextKey string

const (
	ContextAction        ContextKey = "request-action"
	ContextBirthTime     ContextKey = "birth-time"
	ContextUUID          ContextKey = "request-uuid"
	ContextUserAgent     ContextKey = "client-user-agent"
	ContextUserIP        ContextKey = "client-net-ip"
	ContextDeviceID      ContextKey = "client-device-id"
	ContextRouterParam   ContextKey = "router-params"
	ContextMessageID     ContextKey = "nsq-message-id"
	ContextReferenceUUID ContextKey = "reference-request-uuid"

	// for response purpose
	ContextMetricName ContextKey = "metric-name"
)
