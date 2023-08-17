package consts

const (
	AppName     = "botfb"
	BotVersion  = "v1.0"
	VerifyToken = "botfbkkk"
	HubMode     = "subscribe"

	// event from other app
	EventTopic = "myapp"
)

type EventType int

const (
	EventAddToCart EventType = iota
	EventRemoveFromCart
	EventCheckout
	EventPaymentSuccess
	EventPaymentFailure
	EventOrderPlaced
	EventOrderCancelled
	EventOrderShipped
)
