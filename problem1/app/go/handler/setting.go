package handler

const (
	user_id_query_key string = "ID"
	page_limit_query_key string = "limit"
	page_query_key string = "page"
)

type UserData struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type FriendList struct {
	Friends []*UserData `json:"friends"`
}

type pagingQuery struct {
	user_id int
	limit int
	page int
}
