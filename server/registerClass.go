package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func errorByFmtErrorF() error {
	return fmt.Errorf("Error: %s", "fmt.Errorf")
}

type (
	class struct {
		Class_id   int    `json:"class_id"`
		Name       string `json:"name"`
		Day        string `json:"day"`
		Period     string `json:"period"`
		Unit       int    `json:"unit"`
		Must       bool   `json:"must"`
		Teacher    string `json:"teacher"`
		Room       string `json:"room"`
		Term       string `json:"term"`
		Department string `json:"department"`
	}
)

var classes []class

func registerClass(c echo.Context) error {

	name := c.FormValue("name")
	day := c.FormValue("day")
	period := c.FormValue("period")
	unit := c.FormValue("unit")
	must := c.FormValue("must")
	teacher := c.FormValue("teacher")
	room := c.FormValue("room")
	term := c.FormValue("term")
	department := c.FormValue("department")

	// データベースのハンドルを取得する
	db, err := sql.Open("mysql", db_state)
	if err != nil {
		log.Fatal(err)
		return c.JSON(http.StatusCreated, err) // エラーを返す
	}
	defer db.Close()

	// 名前と期間が既に存在するか確認するクエリ
	existsQuery := "SELECT COUNT(*) FROM Class WHERE Name = ? AND Term = ?"
	var count int
	err = db.QueryRow(existsQuery, name, term).Scan(&count)
	if err != nil {
		log.Fatal(err)
		return c.JSON(http.StatusCreated, err) // エラーを返す
	}

	// 既に存在する場合
	if count > 0 {
		return c.JSON(http.StatusConflict, "Already Registered") // ステータスコード409: Conflict
	}

	// 存在しない場合は授業を登録する
	ins, err := db.Prepare("INSERT INTO Class (Name, Day, Period, Unit, Must, Teacher, Room,Term,Department) VALUES (?, ?, ?, ?, ?, ?, ?,?,?)")
	if err != nil {
		log.Fatal(err)
		return c.JSON(http.StatusCreated, err) // エラーを返す
	}

	// SQLの実行
	result, err := ins.Exec(name, day, period, unit, must, teacher, room, term, department)
	if err != nil {
		log.Fatal(err)
		return c.JSON(http.StatusBadRequest, "Incorrect Registered") // ステータスコード400: Bad Request
	}

	// 登録が成功したかどうかを確認
	affected, _ := result.RowsAffected()
	if affected == 0 {
		return c.JSON(http.StatusBadRequest, "Incorrect Registered") // ステータスコード400: Bad Request
	}

	// データベースから全ての授業を取得
	rows, err := db.Query("SELECT * FROM Class")
	if err != nil {
		log.Fatal(err)
		return c.JSON(http.StatusCreated, err) // エラーを返す
	}
	defer rows.Close()

	// classesスライスをクリア
	classes = nil

	// データベースから授業を取得
	for rows.Next() {
		var cc class

		err := rows.Scan(&cc.Class_id, &cc.Name, &cc.Day, &cc.Period,
			&cc.Unit, &cc.Must, &cc.Teacher, &cc.Room, &cc.Term, &cc.Department)

		if err != nil {
			log.Fatal(err)
			return c.JSON(http.StatusCreated, err) // エラーを返す
		}

		classes = append(classes, cc)
	}

	// classesスライスが空でない場合、最後の授業（c）を取得して返す
	if len(classes) > 0 {
		lastClass := classes[len(classes)-1]
		return c.JSON(http.StatusCreated, lastClass)
	}

	return c.JSON(http.StatusCreated, nil)

}
