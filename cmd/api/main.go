package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"task-api/internal/database"
	"task-api/internal/handlers"
)

func main() {
    db, err := database.Connect()
	if err != nil {
		fmt.Println("Database connection failed:", err)
		return
	}

	defer db.Close(context.Background())

	http.HandleFunc("/health", handlers.Health)
    http.Handle("/tasks", handlers.Tasks(db))


	port := os.Getenv("PORT")
    if port == ""{
        port = "8080"
    }

	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println(err)
	}
}