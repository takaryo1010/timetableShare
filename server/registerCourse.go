package main

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type (
	course struct {
		Course_id int `json:"course_id"`
		Person_id int `json:"person_id"`
		Class_id  int `json:"class_id"`
	}
)

var courses []course

func registerCourse(e echo.Context) error {
	name := e.FormValue("person_name")
	class_id, err := strconv.Atoi(e.FormValue("class_id"))
	if err != nil {
		log.Fatal(err)
		return err // エラーを返す
	}

	// データベースのハンドルを取得する
	db, err := sql.Open("mysql", db_state)
	if err != nil {
		log.Fatal(err)
		return err // エラーを返す
	}
	defer db.Close()

	// SQLの準備（Personからnameに一致するidを取得する）
	row := db.QueryRow("SELECT id FROM Person WHERE name = ?", name)
	var id int
	err = row.Scan(&id)
	if err != nil {
		log.Fatal(err)
		return err // エラーを返す
	}

	// INSERT INTO Course ステートメントの準備
	ins, err := db.Prepare("INSERT INTO Course (person_id, class_id) VALUES (?, ?)")
	if err != nil {
		log.Fatal(err)
		return err // エラーを返す
	}

	// SQLの実行（Courseへの挿入）
	_, err = ins.Exec(id, class_id)
	if err != nil {
		log.Fatal(err)
		return err // エラーを返す
	}

	// データベースから全ての時間割を取得
	rows, err := db.Query("SELECT * FROM Course")
	if err != nil {
		log.Fatal(err)
		return err // エラーを返す
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
			return err // エラーを返す
		}

		courses = append(courses, c)
	}

	// coursesスライスが空でない場合、最後の個人の時間割（c）を取得して返す
	if len(courses) > 0 {
		lastCourse := courses[len(courses)-1]
		return e.JSON(http.StatusCreated, lastCourse)
	}

	return e.JSON(http.StatusCreated, nil)
}
