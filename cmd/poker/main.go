package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Scullder/poker/internal/ws"
	"github.com/gorilla/mux"
)

const PORT string = "localhost:8080"

func RootHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "web/static/index.html")
}

func setupAPI() {
	manager := ws.NewManager()

	router := mux.NewRouter()
	router.HandleFunc("/", RootHandler)
	router.HandleFunc("/ws", manager.ServerWs)

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
