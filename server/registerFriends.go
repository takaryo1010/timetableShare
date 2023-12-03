package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type (
	friend struct {
		My_id   int `json:"my_id"`
		Your_id int `json:"your_id"`
	}
)

var friends []friend

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

	// データベースから全ての友達情報を取得
	query3 := "SELECT * FROM Course"
	rows, err := db.Query(query3)
	if err != nil {
		log.Fatal(err)
		return c.JSON(http.StatusCreated, err) // エラーを返す
	}
	defer rows.Close()

	// friendsスライスをクリア
	friends = nil

	// データベースから個人の友達情報を取得
	for rows.Next() {
		var f friend

		err := rows.Scan(&f.My_id, &f.Your_id)

		if err != nil {
			log.Fatal(err)
			return c.JSON(http.StatusCreated, err) // エラーを返す
		}

		friends = append(friends, f)
	}

	// friendsスライスが空でない場合、最後の個人の友達情報（f）を取得して返す
	if len(friends) > 0 {
		lastCourse := friends[len(friends)-1]
		return c.JSON(http.StatusCreated, lastCourse)
	}

	return c.JSON(http.StatusCreated, err)
}
