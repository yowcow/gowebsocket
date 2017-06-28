package handler

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type WsMessage struct {
	Text string `json:"text"`
}

var count int = 0
var mux = &sync.Mutex{}

func WebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	mux.Lock()
	count += 1
	mux.Unlock()

	conn.WriteJSON(&WsMessage{fmt.Sprintf("Hi, now I have %d connection(s)", count)})

	for {
		input := &WsMessage{}
		err = conn.ReadJSON(input)
		if err != nil {
			fmt.Println("Failed reading JSON", err)
			mux.Lock()
			count -= 1
			mux.Unlock()
			break
		}

		fmt.Println("Got a message from client", input.Text)

		output := &WsMessage{
			Text: "You said '" + input.Text + "'",
		}
		conn.WriteJSON(output)
	}
}
