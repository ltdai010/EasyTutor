package offerusecase

import (
	"EasyTutor/data/data"
	"EasyTutor/models"
	"EasyTutor/utils/myerror"
)

func (t *OfferHandler) RemoveOne(username, id string) error {
	offer := &models.Offer{}
	offer.ID = id
	err := offer.Get()
	if err != nil {
		return data.NotExisted
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

