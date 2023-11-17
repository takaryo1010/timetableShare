package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type (
	person struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}
)

var people []person

func registerPerson(c echo.Context) error {
	name := c.FormValue("name")

	// データベースのハンドルを取得する
	db, err := sql.Open("mysql", db_state)
	if err != nil {
		log.Fatal(err)
		return err // エラーを返す
	}
	defer db.Close()

	// SQLの準備
	ins, err := db.Prepare("INSERT INTO Person (Name) VALUES(?)")
	if err != nil {
		log.Fatal(err)
		return err // エラーを返す
	}

	// SQLの実行
	_, err = ins.Exec(name)
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

	// peopleスライスをクリア
	people = nil

	// データベースから人物を取得
	for rows.Next() {
		var p person

		err := rows.Scan(&p.ID, &p.Name)

		if err != nil {
			log.Fatal(err)
			return err // エラーを返す
		}

		people = append(people, p)
	}

	// peopleスライスが空でない場合、最後の人物（p）を取得して返す
	if len(people) > 0 {
		lastPerson := people[len(people)-1]
		return c.JSON(http.StatusCreated, lastPerson)
	}

	return c.JSON(http.StatusCreated, nil)
}
