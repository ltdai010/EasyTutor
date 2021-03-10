package offerusecase

import (
	"EasyTutor/data/data"
	"EasyTutor/data/rest/requests"
	"EasyTutor/models"
	"EasyTutor/utils/myerror"
)

func (t *OfferHandler) UpdateOne(username, id string, put requests.OfferPut) error {
	offer := &models.Offer{}
	offer.ID = id
	err := offer.Get()
	if myerror.IsError(err) {
		return data.NotExisted
	}

	if offer.TeacherID != username {
		return data.NotPermission
	}

	offer.OfferInfo = put.OfferInfo

	err = offer.Update()
	if myerror.IsError(err) {
		return data.ErrSystem
	}
	return data.Success
}