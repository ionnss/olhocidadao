package main

import (
	"fmt"
	"log"
	"olhocidadao/db"
)

func main() {
	fmt.Println("Olho Cidadao Aberto")

	// Connect to the database
	conn, err := db.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer conn.Close()
}
