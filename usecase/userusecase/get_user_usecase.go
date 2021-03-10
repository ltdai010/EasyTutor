package userusecase

import (
	"EasyTutor/data/data"
	"EasyTutor/data/rest/responses"
	"EasyTutor/models"
)

func (t *UserHandler) GetOne(username string) (*responses.User, error) {
	user := &models.User{}
	user.Username = username
	err := user.Get()
	if err != nil {
		return nil, data.NotExisted
	}
	return &responses.User{
		Username: username,
		UserInfo: user.UserInfo,
	}, data.Success
}

func (t *UserHandler) GetPage(pageNumber int, pageSize int) ([]*responses.User, int, error) {
	user := &models.User{}
	listUser, total, err := user.GetPage(pageNumber, pageSize)
	if err != nil {
		return nil, 0, data.NotMore
	}
	return listUser, total, data.Success
}

