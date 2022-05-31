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
			"SELECT f2.user_id, f2.name FROM friend_link u "+
				"INNER JOIN friend_link f1 ON u.friend_id = f1.user_id "+
				"INNER JOIN users f2 ON f1.friend_id = f2.user_id "+
				"WHERE u.user_id = ? "+
				"AND f2.user_id NOT IN ("+
					"SELECT object_id FROM block_list "+
					"WHERE subject_id = ?);", user_id, user_id)
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
