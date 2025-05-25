package db

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB // 他ファイルからも使えるグローバル変数

// InitDB はPostgreSQLへの接続を初期化する
func InitDB() {
    dsn := fmt.Sprintf(
        "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        getEnv("DB_HOST", "localhost"),
        getEnv("DB_PORT", "5432"),
        getEnv("DB_USER", "thready"),
        getEnv("DB_PASSWORD", "secret"),
        getEnv("DB_NAME", "thready"),
    )

    var err error
    DB, err = sqlx.Connect("postgres", dsn)
    if err != nil {
        log.Fatalf("❌ DB接続に失敗しました: %v", err)
    }

    log.Println("✅ DB接続成功")
}

// getEnv は環境変数が未設定ならデフォルト値を返す
func getEnv(key, fallback string) string {
    if value := os.Getenv(key); value != "" {
        return value
    }
    return fallback
}
