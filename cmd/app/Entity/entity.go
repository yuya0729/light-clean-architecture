package entity

// gatewayから呼ばれる
// entity

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Task struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Name  string `json:"name"`
}
