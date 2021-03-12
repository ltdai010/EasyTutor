package searchusecase

import (
	"EasyTutor/data/data"
	"EasyTutor/data/rest/responses"
	"EasyTutor/models"
)

func (s *searchHandler) SearchTeacher(key string, pageNumber, pageSize, location int,
	graduation string, subject string, gender string, method string) ([]*responses.TeacherSearch, error) {
	teacher := &models.Teacher{}
	res, err := teacher.Search(key, location, pageNumber - 1, pageSize, data.NewGraduation(graduation),
		data.NewSubject(subject), data.NewGender(gender), data.NewMethod(method))
	if err != nil {
		return nil, data.NotMore
	}

	return res, data.Success
}

