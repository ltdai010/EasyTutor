package offerusecase

import (
	"EasyTutor/data/data"
	"EasyTutor/data/rest/requests"
	"EasyTutor/models"
	"EasyTutor/utils/logger"
	"EasyTutor/utils/myerror"
)

func (t *offerHandler) UpdateOne(username, id string, put requests.OfferPut) error {
	offer := &models.Offer{}
	request := &models.Request{}

	offer.ID = id
	err := offer.Get()
	if myerror.IsError(err) {
		return data.NotExisted
	}
	request.ID = offer.RequestID
	err = request.Get()
	if err != nil {
		return data.NotExisted
	}

	if request.AcceptOffer == id {
		return data.NotPermission
	}

	if offer.TeacherID != username {
		return data.NotPermission
	}

	offer.OfferInfo = put.OfferInfo

	err = offer.Update()
	if myerror.IsError(err) {
		logger.Error("[Error Decline Offer] Request error offer id = %v err = %v", offer.ID, err)
		return data.ErrSystem
	}
	return data.Success
}