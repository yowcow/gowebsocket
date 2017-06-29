package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
)

func TestWebSocket_counts_conn(t *testing.T) {
	handler := http.HandlerFunc(WebSocket)
	server := httptest.NewServer(handler)
	defer server.Close()

	dialer := &websocket.Dialer{}
	message := &WsMessage{}
	url := strings.Replace(server.URL, "http", "ws", 1)

	conn1, _, _ := dialer.Dial(url, nil)
	conn1.ReadJSON(message)
	assert.Equal(t, "Hi, now I have 1 connection(s)", message.Text)

	conn2, _, _ := dialer.Dial(url, nil)
	conn2.ReadJSON(message)
	assert.Equal(t, "Hi, now I have 2 connection(s)", message.Text)
	conn2.Close()

	time.Sleep(time.Second)

	conn3, _, _ := dialer.Dial(url, nil)
	conn3.ReadJSON(message)
	assert.Equal(t, "Hi, now I have 2 connection(s)", message.Text)
}

func TestWebSocket_echos(t *testing.T) {
	handler := http.HandlerFunc(WebSocket)
	server := httptest.NewServer(handler)
	defer server.Close()

	dialer := &websocket.Dialer{}
	message := &WsMessage{}
	url := strings.Replace(server.URL, "http", "ws", 1)

	conn1, _, _ := dialer.Dial(url, nil)
	conn1.ReadJSON(message)

	message.Text = "hogehoge"
	conn1.WriteJSON(message)
	conn1.ReadJSON(message)
	assert.Equal(t, "You said 'hogehoge'", message.Text)
}
