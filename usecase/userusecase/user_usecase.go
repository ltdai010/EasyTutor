package userusecase

import (
	"EasyTutor/data/data"
	"EasyTutor/data/rest/requests"
	"EasyTutor/data/rest/responses"
)

type UserInterface interface {
	CreateOne(post requests.UserPost) (string, error)
	GetPage(pageNumber int, pageSize int) ([]*responses.User, int, error)
	GetOne(id string) (*responses.User, error)
	UpdateOne(id string, put requests.UserPut) error
	RemoveOne(id string) error
	Login(login data.LoginInfo) (string, error)
}

func GetUserUseCase() UserInterface {
	return &UserHandler{}
}

type UserHandler struct {}