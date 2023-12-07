package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

type (
	friendInfo struct {
		My_id   int `json:"my_id"`
		Your_id int `json:"your_id"`
	}
)

var friendInfos []friendInfo

func showFriends(c echo.Context) error {
	name := c.FormValue("my_name")

	// データベースのハンドルを取得する
	db, err := sql.Open("mysql", db_state)
	if err != nil {
		log.Fatal(err)
		return err // エラーを返す
	}
	defer db.Close()

	// SQLの実行
	query := "SELECT f.my_id, f.your_id FROM Friends f JOIN Person p ON f.my_id = p.id WHERE p.name = ?;"
	rows, err := db.Query(query, name)
	if err != nil {
		log.Fatal(err)
		return err // エラーを返す
	}
	defer rows.Close()

	// Friendsスライスをクリア
	friends = nil

	// データベースから人物を取得
	for rows.Next() {
		var f friend

		err := rows.Scan(&f.My_id, &f.Your_id)
		if err != nil {
			log.Fatal(err)
			return err // エラーを返す
		}

		friends = append(friends, f)
	}

	return c.JSON(http.StatusOK, friends)
}
