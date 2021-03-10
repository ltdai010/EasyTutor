package responses

import "EasyTutor/data/data"

type User struct {
	Username string	`json:"username"`
	data.UserInfo
}
