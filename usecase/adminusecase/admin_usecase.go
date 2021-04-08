package adminusecase

import (
	"EasyTutor/data/data"
	"EasyTutor/data/rest/responses"
)

type AdminInterface interface {
	Login(login data.LoginInfo) (string, error)
	ValidateTeacher(teacherID string) error
	ValidateRequest(requestID string) error
	ValidateComment(commentID string) error
	GetListUnActiveTeacher(pageNumber, pageSize int) ([]*responses.Teacher, int, error)
	GetListUnActiveRequest(pageNumber int, pageSize int) ([]*responses.Request, int, error)
	GetAllUnActiveComment() ([]*responses.TeacherComment, error)
	ValidateUpdateTeacher(teacherID string) error
}

func GetAdminUseCase() AdminInterface {
	return &adminHandler{}
}

type adminHandler struct {}
