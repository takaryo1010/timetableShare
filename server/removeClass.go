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

	Unit9- Whaling is not permitted around the world, so if you don't keep it, you look like you're breaking the rules. But I like the whale meat. And I want to keep Japanese culutre. We have eaten whale meat in Japan. So I want to experience people it in the future.

Unit11- We can reserch person in the town if we use facial recognition. But big databases can generates face picture. So I think afraid.

Unit12- The most problem is overproduction. If overproduction we thorw away it. We should produce enough to eat

Unit13- We should get more place to live. So we should explore to space.For example we may live to Mars. So we can get more place to live. But in the case, we must live in the room.We can't  go out door. So I don't live in marz.// データベースから授業情報を取得
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
