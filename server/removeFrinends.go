package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func removeFriends(e echo.Context) error {

	my_name := e.FormValue("my_name")
	your_name := e.FormValue("your_name")

	// データベースのハンドルを取得する
	db, err := sql.Open("mysql", db_state)
	if err != nil {
		log.Fatal(err)
		return e.JSON(http.StatusCreated, err) // エラーを返す
	}
	defer db.Close()

	// SQLの準備（Personからnameに一致するidを取得する）
	query1 := "SELECT id FROM Person WHERE name = ?"
	var my_id int
	err = db.QueryRow(query1, my_name).Scan(&my_id)
	if err != nil {
		
		return e.JSON(http.StatusCreated, err) // エラーを返す
	}
	// SQLの準備（Personからnameに一致するidを取得する）
	query2 := "SELECT id FROM Person WHERE name = ?"
	var your_id int
	err = db.QueryRow(query2, your_name).Scan(&your_id)
	if err != nil {
		
		return e.JSON(http.StatusCreated, err) // エラーを返す
	}

	// SQLの準備（Personからnameに一致するidを取得する）
	db.QueryRow("DELETE FROM Friends WHERE my_id = ? AND your_id = ?", my_id, your_id)

	// データベースから全ての時間割を取得
	rows, err := db.Query("SELECT * FROM Friends")
	if err != nil {
		log.Fatal(err)
		return e.JSON(http.StatusCreated, err) // エラーを返す
	}
	defer rows.Close()

	// friendInfosスライスをクリア
	friends = nil

	// データベースから個人の時間割を取得
	for rows.Next() {
		var c friend

		err := rows.Scan(&c.My_id, &c.Your_id)

		if err != nil {
			log.Fatal(err)
			return e.JSON(http.StatusCreated, err) // エラーを返す
		}

		friends = append(friends, c)
	}

	return e.JSON(http.StatusCreated, friends)
}
