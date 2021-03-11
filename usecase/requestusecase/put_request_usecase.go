package requestusecase

import (
	"EasyTutor/data/data"
	"EasyTutor/data/rest/requests"
	"EasyTutor/models"
	"EasyTutor/utils/myerror"
)

func (t *requestHandler) UpdateOne(username, id string, put requests.RequestPut) error {
	request := &models.Request{}
	request.ID = id
	err := request.Get()
	if myerror.IsError(err) {
		return data.NotExisted
	}

	request.Username = username

	if !data.DataRequestIsValid(put.RequestInfo) || request.AcceptOffer != ""{
		return data.BadRequest
	}

	request.RequestInfo = put.RequestInfo
	request.Schedule = put.Schedule

	err = request.Update()
	if myerror.IsError(err) {
		return data.ErrSystem
	}
	return data.Success
}

func (t *requestHandler) AcceptOffer(username, offerID string) error {
	//var
	offer := &models.Offer{}
	request := &models.Request{}

	offer.ID = offerID
	err := offer.Get()
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

	request.AcceptOffer = offerID
	err = request.Update()
	if err != nil {
		return data.ErrSystem
	}

	return data.Success
}

func (t *requestHandler) DeclineOffer(username, requestID string) error {
	//var
	request := &models.Request{}
	request.ID = requestID

	if request.Username != username {
		return data.NotPermission
	}

	request.AcceptOffer = ""
	err := request.Update()
	if err != nil {
		return data.ErrSystem
	}

	return data.Success
}