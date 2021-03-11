package requestusecase

import (
	"EasyTutor/data/data"
	"EasyTutor/models"
	"EasyTutor/utils/myerror"
)

func (t *requestHandler) RemoveOne(username, id string) error {
	request := &models.Request{}
	request.ID = id
	err := request.Get()
	if err != nil {
		return data.NotExisted
	}

	if request.Username != username {
		return data.NotPermission
	}
	
	err = request.Delete()
	if myerror.IsError(err) {
		return data.ErrSystem
	}
	return data.Success
}
