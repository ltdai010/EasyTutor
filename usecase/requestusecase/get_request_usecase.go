package requestusecase

import (
	"EasyTutor/data/data"
	"EasyTutor/data/rest/responses"
	"EasyTutor/models"
)

func (t *RequestHandler) GetOne(id string) (*responses.Request, error) {
	request := &models.Request{}
	request.ID = id
	err := request.Get()
	if err != nil {
		return nil, data.NotExisted
	}
	return &responses.Request{
		ID:      id,
		Request: request.Request,
	}, data.Success
}

func (t *RequestHandler) GetPage(pageNumber int, pageSize int) ([]*responses.Request, int, error) {
	request := &models.Request{}
	listRequest, total, err := request.GetPage(pageNumber - 1, pageSize)
	if err != nil {
		return nil, 0, data.NotMore
	}
	return listRequest, total, data.Success
}
