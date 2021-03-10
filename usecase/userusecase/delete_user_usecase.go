package userusecase

import (
	"EasyTutor/data/data"
	"EasyTutor/models"
	"EasyTutor/utils/myerror"
)

func (t *UserHandler) RemoveOne(username string) error {
	user := &models.User{}
	user.Username = username
	err := user.Get()
	if err != nil {
		return data.NotExisted
	}

	err = user.Delete()
	if myerror.IsError(err) {
		return data.ErrSystem
	}
	return data.Success
}

