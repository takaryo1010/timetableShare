package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type (
	Friend struct {
		My_name   string `json:"my_name"`
		Your_name string `json:"your_name"`
	}
)

func registerFriends(c echo.Context) error {
	my_name := c.FormValue("my_name")
	your_name := c.FormValue("your_name")

	// データベースのハンドルを取得する
	db, err := sql.Open("mysql", db_state)
	if err != nil {
		log.Fatal(err)
		return err // エラーを返す
	}
	defer db.Close()

	// SQLの準備
	ins, err := db.Prepare("INSERT INTO Friends (my_name,your_name) VALUES(?,?)")
	if err != nil {
		log.Fatal(err)
		return err // エラーを返す
	}

	// SQLの実行
	_, err = ins.Exec(my_name, your_name)
	if err != nil {
		log.Fatal(err)
		return err // エラーを返す
	}

	return c.String(http.StatusOK, err.Error())
}
