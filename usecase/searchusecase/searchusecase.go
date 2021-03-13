package searchusecase

import "EasyTutor/data/rest/responses"

type SearchInterface interface {
	SearchTeacher(key string, pageNumber, pageSize, location int,
		graduation string, subject string, gender string, method string) ([]*responses.TeacherSearch, error)
	SearchRequest(key string, pageNumber, pageSize, location int,
		subject string, gender string, method string) ([]*responses.RequestSearch, error)
}

type searchHandler struct {}

func GetSearchUseCase() SearchInterface {
	return &searchHandler{}
}
