package requestusecase

import (
	"EasyTutor/data/rest/requests"
	"EasyTutor/data/rest/responses"
)

type RequestInterface interface {
	CreateOne(username string, post requests.RequestPost) (string, error)
	GetPage(pageNumber int, pageSize int) ([]*responses.Request, int, error)
	GetOne(id string) (*responses.Request, error)
	UpdateOne(username, id string, put requests.RequestPut) error
	RemoveOne(username, id string) error
}

func GetRequestUseCase() RequestInterface {
	return &RequestHandler{}
}

type RequestHandler struct {}