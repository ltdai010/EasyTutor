package searchusecase

import (
	"EasyTutor/data/data"
	"EasyTutor/data/rest/responses"
	"EasyTutor/models"
)

func (s *searchHandler) SearchRequest(key string, pageNumber, pageSize, location int,
	subject string, gender string, method string) ([]*responses.RequestSearch, error) {
	request := &models.Request{}
	res, err := request.Search(key, location, pageNumber - 1, pageSize,
		data.ListStringToSubjects([]string{subject}), data.NewGender(gender), data.ListStringToMethod([]string{method}))
	if err != nil {
		return nil, data.NotMore
	}
	return res, data.Success
}