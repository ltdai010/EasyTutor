package teacherusecase

import (
	"EasyTutor/data/data"
	"EasyTutor/data/rest/requests"
	"EasyTutor/models"
	"EasyTutor/utils/myerror"
)

func (t *teacherHandler) UpdateOne(username string, request requests.TeacherPut) error {
	teacher := &models.Teacher{}
	teacher.Username = username
	err := teacher.Get()
	if myerror.IsError(err) {
		return data.NotExisted
	}
	teacher.TeacherInfo, err = data.SetDataTeacher(request.TeacherInfo)
	if myerror.IsError(err) {
		return data.BadRequest
	}
	err = teacher.Update()
	if myerror.IsError(err) {
		return data.ErrSystem
	}
	return data.Success
}

