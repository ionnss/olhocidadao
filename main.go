package main

import (
	"fmt"
	"log"
	"net/http"
	"olhocidadao/db"
	"olhocidadao/routes"
)

func main() {
	fmt.Println("Olho Cidadao Aberto")

	// Connect to the database
	conn, err := db.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer conn.Close()

	// Create the routes
	r := routes.CreateRoutes()

	// Start server
	log.Println("Server initialized at http://localhost:8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("Error initiating server:", err)
	}
}
