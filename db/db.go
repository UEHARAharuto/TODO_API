package db

import (
	"database/sql" // 標準ライブラリ
	"log"

	_ "github.com/mattn/go-sqlite3" // ドライバをインポート
)

var DB *sql.DB // 型が *gorm.DB から *sql.DB に変わる

func Init() {
	var err error
	// データベースファイルを開く。存在しなければ作成される。
	DB, err = sql.Open("sqlite3", "./test.db")
	if err != nil {
		log.Fatal(err)
	}

	// テーブルが存在しない場合のみ作成するSQL文
	const createTableSQL = `
	CREATE TABLE IF NOT EXISTS todos (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`

	// SQLを実行してテーブルを作成
	if _, err := DB.Exec(createTableSQL); err != nil {
		log.Fatal(err)
	}
}
