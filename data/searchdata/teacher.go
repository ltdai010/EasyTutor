package searchdata

import "EasyTutor/data/data"

type TeacherSearch struct {
	ObjectID string	`json:"objectID"`
	data.TeacherInfo
	data.Schedule
}

