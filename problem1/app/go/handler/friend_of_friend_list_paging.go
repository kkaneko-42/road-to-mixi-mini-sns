package handler

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

func HandleFriendOfFriendListPagingGet(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		query, err := getPagingQuery(c)
		if err != nil {
			return c.String(http.StatusBadRequest,
					fmt.Sprintf("%s", err))
		}
		offset := query.limit * (query.page - 1)

		rows, err := db.Query(
			"SELECT f2.user_id, f2.name FROM friend_link u "+
				"INNER JOIN friend_link f1 ON u.friend_id = f1.user_id "+
				"INNER JOIN users f2 ON f1.friend_id = f2.user_id "+
				"WHERE u.user_id = ? "+
				"AND f2.user_id NOT IN ("+
					"SELECT object_id FROM block_list "+
					"WHERE subject_id = ?) "+
				"LIMIT ?, ?;",
			query.user_id, query.user_id, offset, query.limit)
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

func getPagingQuery(c echo.Context) (*pagingQuery, error) {
	var (
		query pagingQuery
		err error
	)

	query.user_id, err = strconv.Atoi(c.QueryParam(user_id_query_key))
	if err != nil {
		return nil, err
	}

	query.limit, err = strconv.Atoi(c.QueryParam(page_limit_query_key))
	if err != nil {
		return nil, err
	} else if query.limit < 1 {
		return nil, fmt.Errorf("Invalid limit")
	}

	query.page, err = strconv.Atoi(c.QueryParam(page_query_key))
	if err != nil {
		return nil, err
	} else if query.page < 1 {
		return nil, fmt.Errorf("Invalid page")
	}

	return &query, nil
}
