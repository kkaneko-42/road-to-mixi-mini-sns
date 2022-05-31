package main

import (
	"database/sql"
	"net/http"
	"problem1/configs"
	"problem1/handler"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

func main() {
	conf := configs.Get()

	db, err := sql.Open(conf.DB.Driver, conf.DB.DataSource)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	e := echo.New()
	mappingURL(db, e)

	e.Logger.Fatal(e.Start(":" + strconv.Itoa(conf.Server.Port)))
}

func mappingURL(db *sql.DB, e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "minimal_sns_app")
	})
	e.GET("/get_friend_list", handler.HandleFriendListGet(db))
	e.GET("/get_friend_of_friend_list", handler.HandleFriendOfFriendListGet(db))
	e.GET("/get_friend_of_friend_list_v2", handler.HandleFriendOfFriendListGetV2(db))
	e.GET("/get_friend_of_friend_list_paging", handler.HandleFriendOfFriendListPagingGet(db))
}
