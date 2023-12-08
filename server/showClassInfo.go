package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type (
	classInfo struct {
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

		err := rows.Scan(&cl.Class_id, &cl.Name, &cl.Day, &cl.Period, &cl.Unit, &cl.Must, &cl.Teacher, &cl.Room, &cl.Term, &cl.Department)

		if err != nil {
			log.Fatal(err)
			return err // エラーを返す
		}

		classInfos = append(classInfos, cl)
	}

	return c.JSON(http.StatusCreated, classInfos)
}

func showClassInfoTimeSpecification(c echo.Context) error {
	// 入力データの解析

	// 任意の条件を取得
	conditions := make(map[string]interface{})
	conditions["unit"] = c.FormValue("unit")
	conditions["must"] = c.FormValue("must")
	conditions["teacher"] = c.FormValue("teacher")
	conditions["room"] = c.FormValue("room")
	conditions["term"] = c.FormValue("term")
	conditions["Department"] = c.FormValue("department")
	conditions["day"] = c.FormValue("day")
	conditions["period"] = c.FormValue("period")

	// データベースのハンドルを取得する
	db, err := sql.Open("mysql", db_state)
	if err != nil {
		log.Fatal(err)
		return err // エラーを返す
	}
	defer db.Close()

	// SQLクエリの構築
	query := "SELECT * FROM Class"

	// 任意の条件が指定されている場合、クエリに追加
	whereClause := false
	args := []interface{}{}

	for key, value := range conditions {
		if value != "" {
			if !whereClause {
				query += " WHERE"
				whereClause = true
			} else {
				query += " AND"
			}
			query += fmt.Sprintf(" %s = ?", key)
			args = append(args, value)
		}
	}

	// SQLの準備
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err)
		return err // エラーを返す
	}

	// SQLの実行
	rows, err := stmt.Query(args...)
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
		err := rows.Scan(&cl.Class_id, &cl.Name, &cl.Day, &cl.Period, &cl.Unit, &cl.Must, &cl.Teacher, &cl.Room, &cl.Term, &cl.Department)
		if err != nil {
			log.Fatal(err)
			return err // エラーを返す
		}
		classInfos = append(classInfos, cl)
	}

	return c.JSON(http.StatusOK, classInfos)
}

func showMyClassInfo(c echo.Context) error {
	// 入力データの解析
	name := c.FormValue("name")

	// データベースのハンドルを取得する
	db, err := sql.Open("mysql", db_state)
	if err != nil {
		log.Fatal(err)
		return err // エラーを返す
	}
	defer db.Close()

	// SQLクエリの構築（名前を条件に追加）
	query := "SELECT Class.* FROM Person JOIN Course ON Person.id = Course.person_id JOIN Class ON Course.class_id = Class.class_id WHERE Person.name = ?"

	// SQLの準備
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err)
		return err // エラーを返す
	}
	defer stmt.Close() // ステートメントを閉じる

	// SQLの実行
	rows, err := stmt.Query(name)
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
		err := rows.Scan(&cl.Class_id, &cl.Name, &cl.Day, &cl.Period, &cl.Unit, &cl.Must, &cl.Teacher, &cl.Room, &cl.Term, &cl.Department)
		if err != nil {
			log.Fatal(err)
			return err // エラーを返す
		}
		classInfos = append(classInfos, cl)
	}

	return c.JSON(http.StatusOK, classInfos)
}
