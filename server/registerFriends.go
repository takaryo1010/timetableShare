package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type (
	Friend struct {
		My_name   int `json:"my_name"`
		Your_name int `json:"your_name"`
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

	// my_nameに対応するmy_idを取得するクエリ
	query := "SELECT id FROM Person WHERE name = ?"
	var my_id, your_id int
	err = db.QueryRow(query, my_name).Scan(&my_id)
	if err != nil {
		log.Fatal(err)
		return err // エラーを返す
	}

	// your_nameに対応するyour_idを取得するクエリ
	err = db.QueryRow(query, your_name).Scan(&your_id)
	if err != nil {
		log.Fatal(err)
		return err // エラーを返す
	}

	// SQLの準備
	ins, err := db.Prepare("INSERT INTO Friends (my_id, your_id) VALUES (?, ?)")
	if err != nil {
		log.Fatal(err)
		return err // エラーを返す
	}

	// SQLの実行
	_, err = ins.Exec(my_id, your_id)
	if err != nil {
		log.Fatal(err)
		return err // エラーを返す
	}

	return c.String(http.StatusOK, err.Error())
}
