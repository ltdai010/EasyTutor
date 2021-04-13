package responses

import "EasyTutor/data/data"

type User struct {
	Username string	`json:"username"`
	data.UserInfo
}

type UserLogin struct {
	Token string `json:"token"`
	Username string `json:"username"`
	data.UserInfo
}