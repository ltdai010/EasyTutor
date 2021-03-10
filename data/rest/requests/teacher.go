package requests

import "EasyTutor/data/data"

type TeacherPost struct {
	data.LoginInfo
	data.TeacherInfo
}

type TeacherPut struct {
	data.TeacherInfo
}
