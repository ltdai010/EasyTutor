package teacherusecase

import (
	"EasyTutor/data/data"
	"EasyTutor/models"
	"EasyTutor/utils/myerror"
)

func (t *teacherHandler) RemoveOne(username string) error {
	teacher := &models.Teacher{}
	teacher.Username = username
	err := teacher.Get()
	if err != nil {
		return data.NotExisted
	}

	err = teacher.Delete()
	if myerror.IsError(err) {
		return data.ErrSystem
	}
	return data.Success
}
