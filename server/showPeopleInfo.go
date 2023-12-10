package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type (
	personInfo struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	}
)

var personInfos []personInfo

func showPeopleInfoAll(c echo.Context) error {
	// DB上にある個人の情報をすべてだす
	// データベースのハンドルを取得する
	db, err := sql.Open("mysql", db_state)
	if err != nil {
		log.Fatal(err)
		return err // エラーを返す
	}
	defer db.Close()

	// SQLの準備
	ins, err := db.Prepare("SELECT * FROM Person ")
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
	rows, err := db.Query("SELECT * FROM Person")
	if err != nil {
		log.Fatal(err)
		return err // エラーを返す
	}
	defer rows.Close()

	// classInfosスライスをクリア
	personInfos = nil

	// データベースから人物を取得
	for rows.Next() {
		var pl personInfo

		err := rows.Scan(&pl.Id, &pl.Name)

		if err != nil {
			log.Fatal(err)
			return err // エラーを返す
		}

		personInfos = append(personInfos, pl)
	}

	return c.JSON(http.StatusCreated, personInfos)
}

func showPeopleTakingSpecificClasses(c echo.Context) error {
	// 入力データの解析
	class_id := c.FormValue("class_id")

	// データベースのハンドルを取得する
	db, err := sql.Open("mysql", db_state)
	if err != nil {
		log.Fatal(err)
		return err // エラーを返す
	}
	defer db.Close()

	// SQLクエリの構築（名前を条件に追加）
	query1 := "SELECT p.name FROM Person p JOIN Course c ON p.id = c.person_id WHERE c.class_id = ?"

	// SQLの準備
	var name string
	err = db.QueryRow(query1, class_id).Scan(&name)
	if err != nil {
		log.Fatal(err)
		return c.JSON(http.StatusCreated, err) // エラーを返す
	}

	return c.JSON(http.StatusOK, name)
}
