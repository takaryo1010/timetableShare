package main

import (
	"net/http"

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
	e.GET("/", connect_check) // ローカル環境の場合、http://localhost:80/をGETするとDBと接続できたか返す
	e.POST("/registerPerson", registerPerson)
	e.POST("/registerCourse", registerCourse)
	e.POST("/registerClass", registerClass)
	e.POST("/registerFriends", registerFriends)
	e.POST("/showPeopleInfoAll", showPeopleInfoAll)
	e.POST("/showClassInfoAll", showClassInfoAll)
	e.POST("/showMyClassInfo", showMyClassInfo)
	e.POST("/showClassInfoTimeSpecification", showClassInfoTimeSpecification)
	e.POST("/showMyFriends", showMyFriends)

	// サーバーをポート番号80で起動
	e.Logger.Fatal(e.Start(":80"))
}

// ハンドラーを定義
func connect_check(c echo.Context) error {
	res := connectOnly()
	return c.String(http.StatusOK, res)
}

func insert_sample(c echo.Context) error {
	// curl -d "name=hoge"  http://localhost:80/registPerson
	name := c.FormValue("name")
	sqlInsert(name)
	return c.String(http.StatusOK, "inserted")
}
