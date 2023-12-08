package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func removeCourse(e echo.Context) error {
	class_id := e.FormValue("class_id")
	name := e.FormValue("name")

	// データベースのハンドルを取得する
	db, err := sql.Open("mysql", db_state)
	if err != nil {
		log.Fatal(err)
		return e.JSON(http.StatusCreated, err) // エラーを返す
	}
	defer db.Close()

	// SQLの準備（Personからnameに一致するidを取得する）
	query1 := "SELECT id FROM Person WHERE name = ?"
	var id int
	err = db.QueryRow(query1, name).Scan(&id)
	if err != nil {
		log.Fatal(err)
		return e.JSON(http.StatusCreated, err) // エラーを返す
	}
	// SQLの準備（Personからnameに一致するidを取得する）
	db.QueryRow("DELETE FROM Course WHERE class_id = ? AND person_id = ?", class_id, id)

	// データベースから全ての時間割を取得
	rows, err := db.Query("SELECT * FROM Course")
	if err != nil {
		log.Fatal(err)
		return e.JSON(http.StatusCreated, err) // エラーを返す
	}
	defer rows.Close()

	// coursesスライスをクリア
	courses = nil

	// データベースから個人の時間割を取得
	for rows.Next() {
		var c course

		err := rows.Scan(&c.Course_id, &c.Person_id, &c.Class_id)

		if err != nil {
			log.Fatal(err)
			return e.JSON(http.StatusCreated, err) // エラーを返す
		}

		courses = append(courses, c)
	}

	

	return e.JSON(http.StatusCreated, courses)
}
