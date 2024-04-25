package main

import (
	"fmt"
	"net/http"

	"github.com/TutorialEdge/realtime-chat-go-react/pkg/websocket"
)

func serveWs(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Host)

	// upgrade this connection to a WebSocket
	// connection
	ws, err := websocket.Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+V\n", err)
		go websocket.Writer(ws)
		websocket.Reader(ws)
	}
}

func setupRoutes() {
	http.HandleFunc("/ws", serveWs)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Simple Server")
	})
}

func main() {
	fmt.Println("Distributed Chat App v0.01")
	setupRoutes()

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("err:", err)
	}
}
