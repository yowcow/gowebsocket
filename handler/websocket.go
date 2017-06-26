package handler

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func WebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	type Message struct {
		Text string `json:"text"`
	}

	for {
		input := &Message{}
		err = conn.ReadJSON(input)
		if err != nil {
			fmt.Println("Failed reading JSON", err)
			conn.Close()
			break
		}

		fmt.Println("Got a message from client", input.Text)

		output := &Message{
			Text: "You said '" + input.Text + "'",
		}
		conn.WriteJSON(output)
	}
}
