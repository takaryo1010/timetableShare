package main

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type (
	classInfo struct {
		Class_id int    `json:"class_id"`
		Name     string `json:"name"`
		Day      string `json:"day"`
		Period   string `json:"period"`
		Unit     int    `json:"unit"`
		Must     bool   `json:"must"`
		Teacher  string `json:"teacher"`
		Room     string `json:"room"`
		Term	 string `json:"term"`
		Department string `json:"department"`

	}
)

var classInfos []classInfo

func showClassInfoAll(c echo.Context) error {
	// DB上にある授業の情報をすべてだす
	// データベースのハンドルを取得する
	db, err := sql.Open("mysql", db_state)
	if err != nil {
		log.Fatal(err)
		return err // エラーを返す
	}
	defer db.Close()

	// SQLの準備
	ins, err := db.Prepare("SELECT * FROM Class ")
	if err != nil {
		log.Fatal(err)
		return err // エラーを返す
	}

	// SQLの実行
	_, err = ins.Exec()
	if err != nil {
		log.Fatal(err)
		return err // エラーを返す
	}

	// データベースから全ての人物を取得
	rows, err := db.Query("SELECT * FROM Class")
	if err != nil {
		log.Fatal(err)
		return err // エラーを返す
	}
	defer rows.Close()

	// classInfosスライスをクリア
	classInfos = nil

	// データベースから人物を取得
	for rows.Next() {
		var cl classInfo

		err := rows.Scan(&cl.Class_id, &cl.Name, &cl.Day, &cl.Period, &cl.Unit, &cl.Must, &cl.Teacher, &cl.Room,&cl.Term,&cl.Department)

		if err != nil {
			log.Fatal(err)
			return err // エラーを返す
		}

		classInfos = append(classInfos, cl)
	}

	return c.JSON(http.StatusCreated, classInfos)
}

func showClassInfoTimeSpecification(c echo.Context) error {
	// 指定した曜日時間の授業を返す
	day := c.FormValue("day")
	period,err := strconv.Atoi((c.FormValue("period")))
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

	// SQLの準備
	stmt, err := db.Prepare("SELECT * FROM Class WHERE day = ? AND period = ?")
	if err != nil {
		log.Fatal(err)
		return err // エラーを返す
	}

	// SQLの実行
	rows, err := stmt.Query(day, period)
	if err != nil {
		log.Fatal(err)
		return err // エラーを返す
	}
	defer rows.Close()

	// classInfosスライスをクリア
	classInfos := []classInfo{}

	// データベースから授業情報を取得
	for rows.Next() {
		var cl classInfo
		err := rows.Scan(&cl.Class_id, &cl.Name, &cl.Day, &cl.Period, &cl.Unit, &cl.Must, &cl.Teacher, &cl.Room)
		if err != nil {
			log.Fatal(err)
			return err // エラーを返す
		}
		classInfos = append(classInfos, cl)
	}

	return c.JSON(http.StatusOK, classInfos)
}
