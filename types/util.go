package types

type ResponseMessage struct {
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
}
