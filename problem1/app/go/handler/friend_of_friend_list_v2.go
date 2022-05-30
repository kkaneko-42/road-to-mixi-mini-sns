package handler

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

func HandleFriendOfFriendListGetV2(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		user_id := c.QueryParam(user_id_query_key)
		rows, err := db.Query(
			"SELECT u.user_id, u.name FROM friend_link f0 "+
				"INNER JOIN friend_link f1 ON f0.friend_id = f1.user_id "+
				"INNER JOIN users u ON f1.friend_id = u.user_id "+
				"WHERE f0.user_id = ?;",
			user_id)
		if err != nil {
			return c.String(http.StatusInternalServerError,
				fmt.Sprintf("%s", err))
		}

		response, err := createFriendList(rows)
		if err != nil {
			return c.String(http.StatusInternalServerError,
				fmt.Sprintf("%s", err))
		}
		return c.JSON(http.StatusOK, response)
	}
}
