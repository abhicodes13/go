package main

import (
	"fmt"
	"net/http"
	"task-api/internal/handlers"
)

func main() {
    router := http.NewServeMux()

    router.HandleFunc("/health", handlers.Health)

    fmt.Println("server running on :8080")

    err := http.ListenAndServe(":8080", router)

    if err != nil {
        fmt.Println(err)
    }
}