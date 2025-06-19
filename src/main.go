package main

import (
	"net/http"
	"thready/src/db"
	"thready/src/handler"
    "log"
)

func main() {
    // データベースの初期化
    db.InitDB()
    // ホームページのハンドラ
    http.HandleFunc("/ping", handler.HandlePing)
    http.HandleFunc("/", handler.HandleHome)
    // User関連のハンドラ
    http.HandleFunc("/signup", handler.HandleSignUp)
    http.HandleFunc("/login", handler.HandleLogin)
    http.HandleFunc("/logout", handler.HandleLogout)
    http.HandleFunc("/mypage", handler.HandleMyPage)
    // Thread関連のハンドラ
    http.HandleFunc("/threads", handler.HandleThreads)
    http.HandleFunc("/threads/new", handler.HandleThreadNew)
    http.HandleFunc("/threads/", handler.HandleThreadShowOrPost)

    log.Println("🚀 Starting Thready at http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}
