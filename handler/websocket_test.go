package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
)

func newServerAndURL() (*httptest.Server, string) {
	handler := http.HandlerFunc(PrepareWSHandler())
	server := httptest.NewServer(handler)
	url := strings.Replace(server.URL, "http", "ws", 1)
	return server, url
}

func TestWebSocket_counts_conn(t *testing.T) {
	server, url := newServerAndURL()
	defer server.Close()

	dialer := &websocket.Dialer{}
	message := &WSMessage{}

	conn1, _, _ := dialer.Dial(url, nil)

	conn1.ReadJSON(message)
	assert.Equal(t, "Got a new connection.", message.Text)

	conn1.ReadJSON(message)
	assert.Equal(t, "Now we have 1 connection(s).", message.Text)

	conn2, _, _ := dialer.Dial(url, nil)

	conn2.ReadJSON(message)
	assert.Equal(t, "Got a new connection.", message.Text)

	conn2.ReadJSON(message)
	assert.Equal(t, "Now we have 2 connection(s).", message.Text)

	conn1.ReadJSON(message)
	assert.Equal(t, "Got a new connection.", message.Text)

	conn1.ReadJSON(message)
	assert.Equal(t, "Now we have 2 connection(s).", message.Text)

	conn2.Close()

	conn1.ReadJSON(message)
	assert.Equal(t, "Got a connection disconnected.", message.Text)

	conn1.ReadJSON(message)
	assert.Equal(t, "Now we have 1 connection(s).", message.Text)

	conn1.Close()
}

func TestWebSocket_broadcasts(t *testing.T) {
	server, url := newServerAndURL()
	defer server.Close()

	dialer := &websocket.Dialer{}
	message := &WSMessage{}

	conn1, _, _ := dialer.Dial(url, nil)
	conn1.ReadJSON(message)
	conn1.ReadJSON(message)

	conn2, _, _ := dialer.Dial(url, nil)
	conn2.ReadJSON(message)
	conn2.ReadJSON(message)
	conn1.ReadJSON(message)
	conn1.ReadJSON(message)

	conn3, _, _ := dialer.Dial(url, nil)
	conn3.ReadJSON(message)
	conn3.ReadJSON(message)
	conn2.ReadJSON(message)
	conn2.ReadJSON(message)
	conn1.ReadJSON(message)
	conn1.ReadJSON(message)

	conn1.WriteJSON(&WSMessage{"はろはろ"})

	conn2.ReadJSON(message)
	assert.Equal(t, "Someone said 'はろはろ'", message.Text)

	conn3.ReadJSON(message)
	assert.Equal(t, "Someone said 'はろはろ'", message.Text)

	conn1.Close()
	conn2.Close()
	conn3.Close()
}
