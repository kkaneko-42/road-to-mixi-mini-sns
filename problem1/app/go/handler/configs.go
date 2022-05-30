package handler

const user_id_query_key string = "ID"

type UserData struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type FriendList struct {
	Friends []*UserData `json:"friends"`
}
