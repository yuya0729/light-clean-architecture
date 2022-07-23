package entity

// gatewayから呼ばれる
// entity

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Name  string `json:"name"`
}

type CreateTask struct {
	UserID int    `json:"user_id"`
	Title  string `json:"title"`
}
