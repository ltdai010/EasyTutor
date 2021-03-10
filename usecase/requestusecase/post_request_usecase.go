package requestusecase

import (
	"EasyTutor/data/data"
	"EasyTutor/data/rest/requests"
	"EasyTutor/models"
	"time"
)

func (t *RequestHandler) CreateOne(username string, post requests.RequestPost) (string, error) {
	request := &models.Request{}

	if !data.DataRequestIsValid(post.RequestInfo) {
		return "", data.BadRequest
	}

	request.Username = username
	request.AcceptOffer = ""
	request.RequestInfo = post.RequestInfo
	request.CreateTime = time.Now().Unix()
	request.Schedule = post.Schedule

	id, err := request.Add()
	if err != nil {
		return "", data.ErrSystem
	}
	return id, data.Success
}


