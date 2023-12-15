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
		Your_name string `json:"your_name"`
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
	query := "SELECT DISTINCT Friend.name AS friend_name FROM Person AS You JOIN Friends ON You.id = Friends.my_id OR You.id = Friends.your_id JOIN Person AS Friend ON (You.id = Friends.my_id AND Friend.id = Friends.your_id) OR (You.id = Friends.your_id AND Friend.id = Friends.my_id WHERE You.name = ?';"
	rows, err := db.Query(query, name)
	if err != nil {
		log.Fatal(err)
		return err // エラーを返す
	}
	defer rows.Close()

	// Friendsスライスをクリア
	friendInfos = nil

	// データベースから人物を取得
	for rows.Next() {
		var f friendInfo

		err := rows.Scan(&f.Your_name)
		if err != nil {
			log.Fatal(err)
			return err // エラーを返す
		}

		friendInfos = append(friendInfos, f)
	}

	return c.JSON(http.StatusOK, friendInfos)
}
