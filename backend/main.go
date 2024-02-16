package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

//define the buffer size of Read and Write
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

//define the Reader to read the message from internet
func reader(conn *websocket.Conn) {
	for {
        //receive the message, and check if there are any errors
		meassageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
        //print the message
		fmt.Println(string(p))


		if err := conn.WriteMessage(meassageType, p); err != nil {
            log.Println(err)
            return
		}
	}
}

func serveWs(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r.Host)

  // upgrade this connection to a WebSocket
  // connection
    ws, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Println(err)
  }
  // listen indefinitely for new messages coming
  // through on our WebSocket connection
    reader(ws)
}

func setupRoutes() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Simple Server")
	})
}

func main() {
    fmt.Println("Chat App v0.01")
	setupRoutes()
	http.ListenAndServe(":8080", nil)
}
