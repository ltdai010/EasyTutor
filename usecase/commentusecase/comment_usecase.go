package commentusecase

import (
	"EasyTutor/data/rest/requests"
	"EasyTutor/data/rest/responses"
)

type CommentInterface interface {
	CreateOne(username, teacherID string, post requests.CommentPost) (string, error)
	GetCommentOfTeacher(requestID string) ([]*responses.Comment, error)
	GetOne(id string) (*responses.Comment, error)
	UpdateOne(username, id string, put requests.CommentPut) error
	RemoveOne(username, id string) error
}

func GetCommentUseCase() CommentInterface {
	return &commentHandler{}
}

type commentHandler struct {}