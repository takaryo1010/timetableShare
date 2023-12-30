package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type (
	course struct {
		Course_id int `json:"course_id"`
		Person_id int `json:"person_id"`
		Class_id  int `json:"class_id"`
	}
)

var courses []course

func registerCourse(c echo.Context) error {
	name := c.FormValue("name")
	class_id := c.FormValue("classid")

	// データベースのハンドルを取得する
	db, err := sql.Open("mysql", db_state)
	if err != nil {
		log.Fatal(err)
		return c.JSON(http.StatusCreated, err) // エラーを返す
	}
	defer db.Close()

	// SQLの準備（Personからnameに一致するidを取得する）
	query1 := "SELECT id FROM Person WHERE name = ?"
	var id int
	err = db.QueryRow(query1, name).Scan(&id)
	if err != nil {
		log.Fatal(err)
		return c.JSON(http.StatusCreated, err) // エラーを返す
	}

	// すでに授業が登録されているかどうか
	query2 := "SELECT count(*) FROM Course WHERE person_id = ? AND class_id = ?"
	var countSameCourse int
	err = db.QueryRow(query2, id, class_id).Scan(&countSameCourse)
	if err != nil {
		log.Fatal(err)
		return c.JSON(http.StatusCreated, err) // エラーを返す
	}

	if countSameCourse > 0 {
		// すでに授業が登録されている場合はエラーを返す
		return c.JSON(http.StatusConflict, "既に授業が登録されています")
	}
	// SQLの準備（class_idからdayとperiodを取得する）
	query3 := "SELECT day, period FROM Class WHERE class_id = ?"
	var day string
	var period int
	err = db.QueryRow(query3, class_id).Scan(&day, &period)
	if err != nil {
		log.Fatal(err)
		return c.JSON(http.StatusNotFound, "指定された授業が見つかりません") // エラーを返す
	}
	// SQLの準備（指定された日付と時間帯に一致する授業を取得する）
	query4 := "SELECT count(*) FROM Class WHERE day = ? AND period = ?"
	var countSameDayPeriod int
	err = db.QueryRow(query4, day, period).Scan(&countSameDayPeriod)
	if err != nil {
		log.Fatal(err)
		return c.JSON(http.StatusCreated, err) // エラーを返す
	}

	if countSameDayPeriod > 0 {
		// 指定された日付と時間帯に別の授業が登録されている場合はエラーを返す
		return c.JSON(http.StatusConflict, "既に別の授業が登録されています")
	}

	// INSERT INTO Course ステートメントの準備
	query5 := "INSERT INTO Course (person_id, class_id) VALUES (?, ?)"
	ins, err := db.Prepare(query5)
	if err != nil {
		log.Fatal(err)
		return c.JSON(http.StatusCreated, err) // エラーを返す
	}

	// SQLの実行（Courseへの挿入）
	_, err = ins.Exec(id, class_id)
	if err != nil {
		log.Fatal(err)
		return c.JSON(http.StatusBadRequest, err) // ステータスコード400: Bad Request
	}

	// データベースから全ての時間割を取得
	query6 := "SELECT * FROM Course"
	rows, err := db.Query(query6)
	if err != nil {
		log.Fatal(err)
		return c.JSON(http.StatusCreated, err) // エラーを返す
	}
	defer rows.Close()

	// coursesスライスをクリア
	courses = nil

	// データベースから個人の時間割を取得
	for rows.Next() {
		var cc course

		err := rows.Scan(&cc.Course_id, &cc.Person_id, &cc.Class_id)

		if err != nil {
			log.Fatal(err)
			return c.JSON(http.StatusCreated, err) // エラーを返す
		}

		courses = append(courses, cc)
	}

	// coursesスライスが空でない場合、最後の個人の時間割（c）を取得して返す
	if len(courses) > 0 {
		lastCourse := courses[len(courses)-1]
		return c.JSON(http.StatusCreated, lastCourse)
	}

	return c.JSON(http.StatusCreated, err)
}
