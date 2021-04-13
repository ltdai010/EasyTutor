package responses

import "EasyTutor/data/data"

type Teacher struct {
	Username string	`json:"username"`
	data.TeacherInfo
	data.Schedule
}

type TeacherSearch struct {
	Username string	`json:"username"`
	data.TeacherInfo
	data.Schedule
}

type TeacherLogin struct {
	Token	string	`json:"token"`
	Username string	`json:"username"`
	data.TeacherInfo
}