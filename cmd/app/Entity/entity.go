package entity

// gatewayから呼ばれる
// entity

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
