package main

import (
	"fmt"
	"log"
	"net/http"

	pokerWebsocket "github.com/Scullder/poker/internal/websocket"
	"github.com/gorilla/mux"
)

const PORT string = "localhost:8080"

func RootHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "web/static/index.html")
}

func setupAPI() {
	handlers := pokerWebsocket.GetHandlers()
	manager := pokerWebsocket.NewManager(handlers)

	router := mux.NewRouter()
	router.HandleFunc("/", RootHandler)
	router.HandleFunc("/websocket", manager.ServerWs)

	http.Handle("/web/", http.StripPrefix("/web/", http.FileServer(http.Dir("web"))))
	http.Handle("/", router)
}

func main() {
	setupAPI()

	fmt.Printf("Server is running: %v\n", PORT)
	err := http.ListenAndServe(PORT, nil)
	if err != nil {
		log.Fatal(err)
	}
}
