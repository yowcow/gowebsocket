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
	handler := http.HandlerFunc(WebSocket)
	server := httptest.NewServer(handler)
	defer server.Close()

	dialer := &websocket.Dialer{}
	url := strings.Replace(server.URL, "http", "ws", 1)
	conn, _, err := dialer.Dial(url, nil)
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
