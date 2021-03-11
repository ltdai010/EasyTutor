package adminusecase

import (
	"EasyTutor/data/data"
	"EasyTutor/data/rest/responses"
	"EasyTutor/models"
)

func (a *adminHandler) GetListUnActiveTeacher(pageNumber, pageSize int) ([]*responses.Teacher, int, error) {
	teacher := &models.Teacher{}
	res, total, err := teacher.GetPageUnActive(pageNumber - 1, pageSize)
	if err != nil {
		return nil, 0, data.NotMore
	}
	return res, total, data.Success
}
