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

	// 入力データが足りているかどうかを示す。
	missingFields := []string{}
	if name == "" {
		missingFields = append(missingFields, "name")
	}
	if day == "" {
		missingFields = append(missingFields, "day")
	}
	if period == "" {
		missingFields = append(missingFields, "period")
	}
	if unit == "" {
		missingFields = append(missingFields, "unit")
	}
	if must == "" {
		missingFields = append(missingFields, "must")
	}
	if teacher == "" {
		missingFields = append(missingFields, "teacher")
	}
	if room == "" {
		missingFields = append(missingFields, "room")
	}
	if term == "" {
		missingFields = append(missingFields, "term")
	}
	if department == "" {
		missingFields = append(missingFields, "department")
	}

	if len(missingFields) > 0 {
		// 値が足りないときに足りないことを示す
		errMsg := fmt.Sprintf("入力データが足りません。\n以下は足りないデータの一覧です。\ns: %v", missingFields)
		return c.JSON(http.StatusBadRequest, errMsg)
	}

	// must の値が 0 か 1 以外のときに 0 か 1 である必要があることを示す
	if must != "0" && must != "1" {
		return c.JSON(http.StatusBadRequest, "must に 0 と 1 以外の値が入っています。")
	}

	// データベースのハンドルを取得する
	db, err := sql.Open("mysql", db_state)
	if err != nil {
		log.Fatal(err)
		return err // エラーを返す
	}
	defer db.Close()

	// 名前、先生名、セメスタが一致する授業があるか確認するクエリ
	existsQuery := "SELECT COUNT(*) FROM Class WHERE Name = ? Teacher = ? AND Term = ?"
	var count int
	err = db.QueryRow(existsQuery, name, term).Scan(&count)
	if err != nil {
		log.Fatal(err)
		return err
	}

	// 既に存在する場合
	if count > 0 {
		return c.JSON(http.StatusConflict, ("すでに登録されています。")) // ステータスコード409: Conflict
	}

	// 存在しない場合は授業を登録する
	ins, err := db.Prepare("INSERT INTO Class (Name, Day, Period, Unit, Must, Teacher, Room,Term,Department) VALUES (?, ?, ?, ?, ?, ?, ?,?,?)")
	if err != nil {
		log.Fatal(err)
		return err
	}

	// SQLの実行
	result, err := ins.Exec(name, day, period, unit, must, teacher, room, term, department)
	if err != nil {
		log.Fatal(err)
		return c.JSON(http.StatusBadRequest, "登録が間違っています") // ステータスコード400: Bad Request
	}

	// 登録が成功したかどうかを確認
	affected, _ := result.RowsAffected()
	if affected == 0 {
		return c.JSON(http.StatusBadRequest, "登録が間違っています") // ステータスコード400: Bad Request
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
