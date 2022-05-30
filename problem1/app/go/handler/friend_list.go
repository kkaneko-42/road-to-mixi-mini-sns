package handler

import (
	"fmt"
	"net/http"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

const user_id_query_key string = "ID"

func HandleFriendListGet(db *sql.DB) echo.HandlerFunc {
	return func (c echo.Context) error {
		user_id := c.QueryParam(user_id_query_key)
		rows, err := db.Query(
			"SELECT friend_id, friend_name FROM friend_link " +
			"INNER JOIN users ON friend_id = user_id " +
			"WHERE user_id = ?",
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

func createFriendList(rows *sql.Rows) (*FriendList, error) {
	var (
		friend_data UserData
		friend_list []*UserData
	)

	for rows.Next() {
		err := rows.Scan(
			&friend_data.Id,
			&friend_data.Name,
		)
		if err != nil {
			return nil, err
		}
		friend_list = append(friend_list, &friend_data)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &FriendList{Friends: friend_list}, nil
}

type UserData struct {
	Id int `json:"id"`
	Name string  `json:"name"`
}

type FriendList struct {
	Friends []*UserData `json:"friends"`
}

