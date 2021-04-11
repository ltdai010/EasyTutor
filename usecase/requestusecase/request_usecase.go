package requestusecase

import (
	"EasyTutor/data/rest/requests"
	"EasyTutor/data/rest/responses"
)

type RequestInterface interface {
	CreateOne(username string, post requests.RequestPost) (string, error)
	GetPageActive(pageNumber int, pageSize int) ([]*responses.Request, int, error)
	GetOne(id string) (*responses.Request, error)
	GetAllRequestOfUser(username string) ([]*responses.Request, error)
	UpdateOne(username, id string, put requests.RequestPut) error
	CloseRequest(username, requestID string) error
	OpenRequest(username, requestID string) error
	RemoveOne(username, id string) error
	AcceptOffer(username, offerID string) error
	DeleteOffer(username, offerID string) error
	FindAvailableTeacher(requestID string) ([]*responses.TeacherSearch, error)
}

func GetRequestUseCase() RequestInterface {
	return &requestHandler{}
}

type requestHandler struct {}