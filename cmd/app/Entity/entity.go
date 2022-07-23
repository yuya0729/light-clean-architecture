package entity

// gatewayから呼ばれる
// entity

// TODO: Id to ID
type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Task struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Name  string `json:"name"`
}

type CreateTask struct {
	UserID int    `json:"user_id"`
	Title  string `json:"title"`
}
