package model

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq" // importするとinit()を呼び出してDBのドライバとして自分を登録する
)

// DB はDBへの接続のプール
var DB *sql.DB

func init() {
	var err error
	DB, err = sql.Open("postgres", os.Getenv("DATA_SOURCE_NAME"))
	if err != nil {
		panic(err)
	}
}
