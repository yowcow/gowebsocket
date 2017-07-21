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

type WSMessage struct {
	Text string `json:"text"`
}

type WSEnvelope struct {
	message *WSMessage
	from    *websocket.Conn
}

type WSHub struct {
	register   chan *websocket.Conn
	unregister chan *websocket.Conn
	conns      map[*websocket.Conn]bool
	broadcast  chan *WSEnvelope
}

func (hub *WSHub) Start() {
	for {
		select {
		case conn := <-hub.register:
			hub.RegisterConnection(conn)
		case conn := <-hub.unregister:
			hub.UnregisterConnection(conn)
		case envelope := <-hub.broadcast:
			hub.BroadcastEnvelope(envelope)
		}
	}
}

func (hub *WSHub) RegisterConnection(conn *websocket.Conn) {
	if _, ok := hub.conns[conn]; ok == false {
		hub.conns[conn] = true
		fmt.Println("connection registered to hub")
	}

	message := &WSMessage{
		Text: fmt.Sprintf("Got a new connection. Now we have %d connection(s).", len(hub.conns)),
	}
	envelope := &WSEnvelope{message, nil}

	hub.BroadcastEnvelope(envelope)
}

func (hub *WSHub) UnregisterConnection(conn *websocket.Conn) {
	if _, ok := hub.conns[conn]; ok == true {
		if err := conn.Close(); err != nil {
			fmt.Println(err)
		}
		delete(hub.conns, conn)
		fmt.Println("connection unregistered from hub")
	}

	message := &WSMessage{
		Text: fmt.Sprintf("Got a connection disconnected. Now we have %d connection(s).", len(hub.conns)),
	}
	envelope := &WSEnvelope{message, nil}

	hub.BroadcastEnvelope(envelope)
}

func (hub *WSHub) BroadcastEnvelope(envelope *WSEnvelope) {
	for conn, _ := range hub.conns {
		if conn == envelope.from {
			continue
		}
		if err := conn.WriteJSON(envelope.message); err != nil {
			fmt.Println(err)
		}
	}
	fmt.Println("message broadcasted")
}

func PrepareWSHub() http.HandlerFunc {
	hub := &WSHub{
		register:   make(chan *websocket.Conn),
		unregister: make(chan *websocket.Conn),
		conns:      map[*websocket.Conn]bool{},
		broadcast:  make(chan *WSEnvelope),
	}
	go hub.Start()

	return func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			fmt.Println(err)
			return
		}

		hub.register <- conn

		for {
			input := &WSMessage{}
			err = conn.ReadJSON(input)

			if err != nil {
				fmt.Println("failed reading JSON", err)
				hub.unregister <- conn
				break
			}

			fmt.Println("got a message from client", input.Text)

			message := &WSMessage{
				Text: "Someone said '" + input.Text + "'",
			}
			envelope := &WSEnvelope{message, conn}

			hub.broadcast <- envelope
		}
	}
}
