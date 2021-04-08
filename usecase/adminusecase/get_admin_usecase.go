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

func (t *adminHandler) GetListUnActiveRequest(pageNumber int, pageSize int) ([]*responses.Request, int, error) {
	request := &models.Request{}
	listRequest, total, err := request.GetPageUnActive(pageNumber - 1, pageSize)
	if err != nil {
		return nil, 0, data.NotMore
	}
	return listRequest, total, data.Success
}

func (t *adminHandler) GetAllUnActiveComment() ([]*responses.TeacherComment, error) {
	comment := &models.Comment{}
	teacherComment, err := comment.GetUnActiveCommentOfAll()
	if err != nil {
		return nil, data.NotMore
	}
	return teacherComment, data.Success
}