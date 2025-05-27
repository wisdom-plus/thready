package main

import (
	"net/http"
	"thready/src/db"
	"thready/src/handler"
    "log"
)

func main() {
    db.InitDB()

    http.HandleFunc("/", handler.HandleHome)
    http.HandleFunc("/threads", handler.HandleThreads)
    http.HandleFunc("/threads/new", handler.HandleThreadNew)
    http.HandleFunc("/threads/", handler.HandleThreadShowOrPost)

    log.Println("🚀 Starting Thready at http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}
