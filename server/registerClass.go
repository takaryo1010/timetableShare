package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type (
	class struct {
		Class_id int    `json:"class_id"`
		Name     string `json:"name"`
		Day      string `json:"day"`
		Period   string `json:"period"`
		Unit     int    `json:"unit"`
		Must     bool   `json:"must"`
		Teacher  string `json:"teacher"`
		Room     string `json:"room"`
	}
)

var classes []class

func registerClass(e echo.Context) error {

	name := e.FormValue("name")
	day := e.FormValue("day")
	period := e.FormValue("period")
	unit := e.FormValue("unit")
	must := e.FormValue("must")
	teacher := e.FormValue("teacher")
	room := e.FormValue("room")

	// データベースのハンドルを取得する
	db, err := sql.Open("mysql", db_state)
	if err != nil {
		log.Fatal(err)
		return err // エラーを返す
	}
	defer db.Close()

	// SQLの準備
	ins, err := db.Prepare("INSERT INTO Class (Name, Day, Period, Unit, Must, Teacher, Room) VALUES (?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
		return err // エラーを返す
	}

	// SQLの実行
	_, err = ins.Exec(name, day, period, unit, must, teacher, room)
	if err != nil {
		log.Fatal(err)
		return err // エラーを返す
	}

	// データベースから全ての授業を取得
	rows, err := db.Query("SELECT * FROM Class")
	if err != nil {
		log.Fatal(err)
		return err // エラーを返す
	}
	defer rows.Close()

	// classesスライスをクリア
	classes = nil

	// データベースから授業を取得
	for rows.Next() {
		var c class

		err := rows.Scan(&c.Class_id, &c.Name, &c.Day, &c.Period,
			&c.Unit, &c.Must, &c.Teacher, &c.Room)

		if err != nil {
			log.Fatal(err)
			return err // エラーを返す
		}

		classes = append(classes, c)
	}

	// classesスライスが空でない場合、最後の授業（c）を取得して返す
	if len(classes) > 0 {
		lastClass := classes[len(classes)-1]
		return e.JSON(http.StatusCreated, lastClass)
	}

	return e.JSON(http.StatusCreated, nil)

}
