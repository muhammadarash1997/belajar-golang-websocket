package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// Upgrader specifies parameters for upgrading
// an HTTP connection to a WebSocket connection.
var upgrader = websocket.Upgrader{
	// I/O buffering The process of temporarily storing data that is passing between a processor and a peripheral.
	// The usual purpose is to smooth out the difference in rates at which the two devices can handle data.
	ReadBufferSize: 1024,	// specify I/O buffer sizes in bytes
	WriteBufferSize: 1024,	// specify I/O buffer sizes in bytes
}

func homePageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page")
}

func webSocketHandler(w http.ResponseWriter, r *http.Request) {
	// CheckOrigin if for validating the request origin to prevent cross-site request forgery.
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}

	// Upgrade upgrades the HTTP server connection to the WebSocket protocol
	// or in other words this is to make websocket connection
	wsConn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	log.Println("Client successfully connected")

	for {
		// messageType can be either BinaryMessage or TextMessage
		// if data sent is like json then better use TextMessage
		// but data sent is like image or audio then use BinaryMessage
		messageType, p, err := wsConn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		log.Println(string(p))

		err = wsConn.WriteMessage(messageType, p)
		if err != nil {
			log.Println(err)
			return
		}
	}
}

func setupRoutes() {
	http.HandleFunc("/", homePageHandler)
	http.HandleFunc("/ws", webSocketHandler)
}

func main() {
	setupRoutes()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
