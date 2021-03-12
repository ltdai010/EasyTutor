package requestusecase

import (
	"EasyTutor/data/data"
	"EasyTutor/data/rest/requests"
	"EasyTutor/models"
	"EasyTutor/utils/datastruct"
	"EasyTutor/utils/logger"
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
		logger.Error("[Error UpdateOne] Update Request error request id = %v err = %v", id, err)
		return data.ErrSystem
	}
	return data.Success
}

func (t *requestHandler) AcceptOffer(username, offerID string) error {
	//var
	offer := &models.Offer{}
	request := &models.Request{}
	teacher := &models.Teacher{}

	offer.ID = offerID
	err := offer.Get()
	if err != nil {
		return data.NotExisted
	}

	teacher.Username = offer.TeacherID
	err = teacher.Get()
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
		logger.Error("[Error Accept Offer] Update error request id = %v err = %v", request.ID, err)
		return data.ErrSystem
	}

	mapUser := datastruct.ArrayToMapIString(teacher.User)
	mapUser[username] = true
	teacher.User = datastruct.MapIStringToArray(mapUser)

	err = teacher.Update()
	if err != nil {
		logger.Error("[Error Accept Offer] Update error teacher id = %v err = %v", offer.TeacherID, err)
		return data.ErrSystem
	}

	return data.Success
}
