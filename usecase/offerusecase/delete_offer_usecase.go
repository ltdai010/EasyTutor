package offerusecase

import (
	"EasyTutor/data/data"
	"EasyTutor/models"
	"EasyTutor/utils/myerror"
)

func (t *offerHandler) RemoveOne(username, id string) error {
	offer := &models.Offer{}
	request := &models.Request{}

	offer.ID = id
	err := offer.Get()
	if err != nil {
		return data.NotExisted
	}

	request.ID = offer.RequestID
	err = request.Get()
	if err != nil {
		return data.NotExisted
	}

	if request.AcceptOffer == id {
		request.AcceptOffer = ""
		err = request.Update()
		if err != nil {
			return data.ErrSystem
		}
	}

	if offer.TeacherID != username {
		return data.NotPermission
	}

	err = offer.Delete()
	if myerror.IsError(err) {
		return data.ErrSystem
	}
	return data.Success
}

