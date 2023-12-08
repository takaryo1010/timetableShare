package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func removeClass(e echo.Context) error {
	class_id := e.FormValue("class_id")
	

	// データベースのハンドルを取得する
	db, err := sql.Open("mysql", db_state)
	if err != nil {
		log.Fatal(err)
		return e.JSON(http.StatusCreated, err) // エラーを返す
	}
	defer db.Close()

	
	// SQLの準備（Personからnameに一致するidを取得する）
	db.QueryRow("DELETE FROM Class WHERE class_id = ? ", class_id)

	// データベースから全ての時間割を取得
	rows, err := db.Query("SELECT * FROM Class")
	if err != nil {
		log.Fatal(err)
		return e.JSON(http.StatusCreated, err) // エラーを返す
	}
	defer rows.Close()

	
	

	// classInfosスライスをクリア
	classInfos := []classInfo{}

	// データベースから授業情報を取得
	for rows.Next() {
		var cl classInfo
		err := rows.Scan(&cl.Class_id, &cl.Name, &cl.Day, &cl.Period, &cl.Unit, &cl.Must, &cl.Teacher, &cl.Room, &cl.Term, &cl.Department)
		if err != nil {
			log.Fatal(err)
			return err // エラーを返す
		}
		classInfos = append(classInfos, cl)
	}

	return e.JSON(http.StatusOK, classInfos)
}
