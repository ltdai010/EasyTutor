package requestusecase

import (
	"EasyTutor/data/data"
	"EasyTutor/data/rest/requests"
	"EasyTutor/models"
	"time"
)

func (t *requestHandler) CreateOne(username string, post requests.RequestPost) (string, error) {
	request := &models.Request{}

	if !data.DataRequestIsValid(post.RequestInfo) {
		return "", data.BadRequest
	}

	request.Username = username
	request.AcceptOffer = ""
	request.Active = false
	request.RequestInfo = post.RequestInfo
	request.CreateTime = time.Now().Unix()
	request.Schedule = post.Schedule
	request.Closed = false

	id, err := request.Add()
	if err != nil {
		return "", data.ErrSystem
	}
	return id, data.Success
}


