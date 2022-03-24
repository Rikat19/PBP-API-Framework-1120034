package model

type User struct {
	UID      int    `json:"uid"`
	Username string `json:"username"`
	Age      int    `json:"age"`
}
