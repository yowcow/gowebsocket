package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/yowcow/gowebsocket/handler"
)

func main() {
	port := "8888"

	flag.StringVar(&port, "port", "8888", "Port to listen")
	flag.Parse()

	servemux := http.NewServeMux()
	servemux.HandleFunc("/", handler.Html)
	servemux.HandleFunc("/ws", handler.PrepareWSHandler())

	fmt.Println("Booting app listening to port", port)

	server := &http.Server{
		Addr:    ":" + port,
		Handler: servemux,
	}

	server.ListenAndServe()
}
