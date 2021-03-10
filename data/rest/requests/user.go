package requests

import "EasyTutor/data/data"

type UserPost struct {
	data.UserInfo
	data.LoginInfo
}

type UserPut struct {
	data.UserInfo
}
