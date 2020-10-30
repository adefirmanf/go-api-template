package handlers

import (
	"go-api-template/internal/httpserver/response"
	"net/http"
)

// EchoMessage .
type EchoMessage struct{}

// GetHelloMessage .
func (e *EchoMessage) GetHelloMessage(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	message, ok := ctx.Value(contextMessageKey).(string)
	if !ok {
		response.TransformErrorAsJSON(w, http.StatusUnprocessableEntity)
		return
	}
	response.TransformSuccessAsJSON(w, http.StatusOK, message)
}

// NewEchoMessageHandler .
func NewEchoMessageHandler() *EchoMessage {
	return &EchoMessage{}
}
