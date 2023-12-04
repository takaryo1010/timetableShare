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
