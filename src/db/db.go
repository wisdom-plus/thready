package db

import (
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB // 他ファイルからも使えるグローバル変数

// InitDB はPostgreSQLへの接続を初期化する
func InitDB() {
    dsn := os.Getenv("DATABASE_URL")
    if dsn == "" {
        log.Fatal("❌ DATABASE_URL が設定されていません")
    }

    var err error
    DB, err = sqlx.Connect("postgres", dsn)
    if err != nil {
        log.Fatalf("❌ DB接続に失敗しました: %v", err)
    }

    log.Println("✅ DB接続成功")
}
