package model

// DefaultHandlerResponse is the default implementation of HandlerResponse.
type DefaultHandlerResponse struct {
	Status  int         `json:"status"    dc:"Status code"`
	Code    int         `json:"code"    dc:"Error code"`
	Message string      `json:"msg" dc:"Error message"`
	Data    interface{} `json:"data"    dc:"Result data for certain request according API definition"`
}
