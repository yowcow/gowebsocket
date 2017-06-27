package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
)

func TestWebSocket(t *testing.T) {
	//w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/ws", nil)
	req.Header.Add("upgrade", "websocket")
	req.Header.Add("connection", "Upgrade")
	req.Header.Add("sec-websocket-version", "13")
	req.Header.Add("sec-websocket-key", "hogefugafoobar")

	handler := http.HandlerFunc(WebSocket)
	server := httptest.NewServer(handler)
	defer server.Close()

	dialer := &websocket.Dialer{}
	conn, _, err := dialer.Dial(strings.Replace(server.URL, "http", "ws", 1), nil)
	//defer conn.Close()

	assert.Equal(t, nil, err)

	type Message struct {
		Text string `json:"text"`
	}

	input := Message{"hogehoge"}
	conn.WriteJSON(&input)

	output := Message{}
	conn.ReadJSON(&output)

	assert.Equal(t, "You said 'hogehoge'", output.Text)
}
