package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

var Friends []Friend

func showMyFriends(c echo.Context) error {
	myID := c.FormValue("my_id")

	// データベースのハンドルを取得する
	db, err := sql.Open("mysql", db_state)
	if err != nil {
		log.Fatal(err)
		return err // エラーを返す
	}
	defer db.Close()

	// SQLの実行
	rows, err := db.Query("SELECT DISTINCT * FROM Friends WHERE my_id=?", myID)
	if err != nil {
		log.Fatal(err)
		return err // エラーを返す
	}
	defer rows.Close()

	// Friendsスライスをクリア
	Friends = nil

	// データベースから人物を取得
	for rows.Next() {
		var f Friend

		err := rows.Scan(&f.My_id, &f.Your_id)
		if err != nil {
			log.Fatal(err)
			return err // エラーを返す
		}

		Friends = append(Friends, f)
	}

	return c.JSON(http.StatusOK, Friends)
}
