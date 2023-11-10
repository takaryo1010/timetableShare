package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	// インスタンスを作成
	e := echo.New()

	// ミドルウェアを設定
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// ルートを設定
	e.GET("/", connect_check) // ローカル環境の場合、http://localhost:1323/ にGETアクセスされるとhelloハンドラーを実行する
	e.POST("/add_person", insert_sample)
	
	// サーバーをポート番号1323で起動
	e.Logger.Fatal(e.Start(":1323"))
}

// ハンドラーを定義
func connect_check(c echo.Context) error {
	res := connectOnly()
	return c.String(http.StatusOK, res)
}

func insert_sample(c echo.Context) error {
	// curl -d "name=hoge" -d "class_id=3" http://localhost:1323/add_person
	name := c.FormValue("name")
	class_id := c.FormValue("class_id")
	id,err:=strconv.Atoi(class_id)
	if err!=nil{
		return c.String(http.StatusBadRequest,err.Error())
	}
	sqlInsert(name, id)
	return c.String(http.StatusOK, "inserted")
}
