package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type (
	Friend struct {
		My_id  int    `json:"my_id"`
		Your_id int `json:"your_id"`
	}
)



func registerFriends(c echo.Context) error {
	my_id := c.FormValue("my_id")
	your_id := c.FormValue("your_id")

	// データベースのハンドルを取得する
	db, err := sql.Open("mysql", db_state)
	if err != nil {
		log.Fatal(err)
		return err // エラーを返す
	}
	defer db.Close()

	// SQLの準備
	ins, err := db.Prepare("INSERT INTO Friends (my_id,your_id) VALUES(?,?)")
	if err != nil {
		log.Fatal(err)
		return err // エラーを返す
	}

	// SQLの実行
	_, err = ins.Exec(my_id,your_id)
	if err != nil {
		log.Fatal(err)
		return err // エラーを返す
	}

	

	

	return c.String(http.StatusOK, err.Error())
}
