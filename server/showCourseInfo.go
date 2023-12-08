package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type (
	courseInfo struct {
		CourseId int `json:"course_id"`
		PersonId int `json:"person_id"`
		ClassId  int `json:"class_id"`
	}
)

var courseInfos []courseInfo

func showCourseInfoAll(c echo.Context) error {
	// DB上にある個人の情報をすべてだす
	// データベースのハンドルを取得する
	db, err := sql.Open("mysql", db_state)
	if err != nil {
		log.Fatal(err)
		return err // エラーを返す
	}
	defer db.Close()

	// SQLの準備
	query1 := "SELECT * FROM Course "
	ins, err := db.Prepare(query1)
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

	// データベースから全ての登録情報を取得
	rows, err := db.Query(query1)
	if err != nil {
		log.Fatal(err)
		return err // エラーを返す
	}
	defer rows.Close()

	// classInfosスライスをクリア
	personInfos = nil

	// データベースから登録情報を取得
	for rows.Next() {
		var cl courseInfo

		err := rows.Scan(&cl.CourseId, &cl.PersonId, &cl.ClassId)

		if err != nil {
			log.Fatal(err)
			return err // エラーを返す
		}

		courseInfos = append(courseInfos, cl)
	}

	return c.JSON(http.StatusCreated, courseInfos)
}
