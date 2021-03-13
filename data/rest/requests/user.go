package requests

import "EasyTutor/data/data"

type UserPost struct {
	data.UserInfo
	data.LoginInfo
}

type UserPut struct {
	data.UserInfo
}

type ResetPass struct {
	Username string	`json:"username"`
	Code	string	`json:"code"`
	NewPass string	`json:"new_pass"`
}