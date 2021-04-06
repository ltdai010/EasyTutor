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

func (t *requestHandler) DeleteOffer(username, offerID string) error {
	//var
	offer := &models.Offer{}
	user := &models.User{}
	request := &models.Request{}

	user.Username = username
	err := user.Get()
	if err != nil {
		return data.NotExisted
	}

	offer.ID = offerID
	err = offer.Get()
	if err != nil {
		return data.NotExisted
	}

	request.ID = offer.RequestID
	err = request.Get()
	if err != nil {
		return data.NotExisted
	}

	if request.Username != username {
		return data.NotPermission
	}

	if request.AcceptOffer == offerID {
		request.AcceptOffer = ""
		err = request.Update()
		if err != nil {
			return data.ErrSystem
		}
	}

	err = offer.Delete()
	if err != nil {
		return data.ErrSystem
	}

	return data.Success
}