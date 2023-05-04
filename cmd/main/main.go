package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"

	"game_go/internal/database"
	"game_go/internal/utils"
)

func main() {

	// Initialize database connection
	db.Init()

	// Initialize models
	//models.Init()

	// Initialize router
	router := mux.NewRouter()

	// Define routes
	router = utils.Init(router)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	addr := fmt.Sprintf(":%s", port)
	log.Printf("Server listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, router))

}
