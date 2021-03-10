package requestusecase

import (
	"EasyTutor/data/data"
	"EasyTutor/data/rest/requests"
	"EasyTutor/models"
	"EasyTutor/utils/myerror"
)

func (t *RequestHandler) UpdateOne(username, id string, put requests.RequestPut) error {
	request := &models.Request{}
	request.ID = id
	err := request.Get()
	if myerror.IsError(err) {
		return data.NotExisted
	}

	request.Username = username

	if !data.DataRequestIsValid(put.RequestInfo) {
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
