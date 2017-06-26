package handler

import (
	"net/http/httptest"
	"testing"
)

func TestWebSocket(t *testing.T) {
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/ws", nil)
	req.Header.Add("upgrade", "websocket")
	req.Header.Add("connection", "Upgrade")
	req.Header.Add("sec-websocket-version", "13")
	req.Header.Add("sec-websocket-key", "hogefugafoobar")
	WebSocket(w, req)
}
