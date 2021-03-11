package userusecase

import (
	"EasyTutor/data/data"
	"EasyTutor/data/rest/requests"
	"EasyTutor/models"
	"EasyTutor/utils/myerror"
)

func (t *userHandler) UpdateOne(username string, request requests.UserPut) error {
	user := &models.User{}
	user.Username = username
	err := user.Get()
	if myerror.IsError(err) {
		return data.NotExisted
	}
	user.UserInfo = request.UserInfo
	err = user.Update()
	if myerror.IsError(err) {
		return data.ErrSystem
	}
	return data.Success
}
