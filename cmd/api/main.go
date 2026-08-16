package main

import (
	"context"
	"fmt"
	"net/http"
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


	fmt.Println("Server running on :8080")

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}