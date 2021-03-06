package types

// sudo module event types
const (
	EventTypeRelayerAdded     = "relayer_added"
	EventTypeRelayerRemoved   = "relayer_removed"
	EventTypeThresholdUpdated = "threshold_updated"

	AttributeKeyArena            = "arena"
	AttributeKeyDenom            = "denom"
	AttributeKeyRelayer          = "relayer"
	AttributeKeyLastThreshold    = "last_threshold"
	AttributeKeyCurrentThreshold = "current_threshold"
)
