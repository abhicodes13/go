package main

import (
	"log"
	"net/http"
)

func main() {
    router := http.NewServeMux()

    router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte(`{"status":"ok"}`))
    })

    log.Println("server running on :8080")

    err := http.ListenAndServe(":8080", router)

    if err != nil {
        log.Fatal(err)
    }
}