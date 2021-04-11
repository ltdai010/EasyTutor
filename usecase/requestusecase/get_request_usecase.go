package requestusecase

import (
	"EasyTutor/data/data"
	"EasyTutor/data/rest/responses"
	"EasyTutor/models"
	"log"
)

func (t *requestHandler) GetOne(id string) (*responses.Request, error) {
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

func (t *requestHandler) GetPageActive(pageNumber int, pageSize int) ([]*responses.Request, int, error) {
	request := &models.Request{}
	listRequest, total, err := request.GetPageActive(pageNumber - 1, pageSize)
	if err != nil {
		return nil, 0, data.NotMore
	}
	return listRequest, total, data.Success
}

func (t *requestHandler) FindAvailableTeacher(requestID string) ([]*responses.TeacherSearch, error) {
	teacher := &models.Teacher{}
	request := &models.Request{}

	request.ID = requestID
	err := request.Get()
	if err != nil || !request.Active || request.Closed {
		return nil, data.NotExisted
	}

	result, err := teacher.Search("", request.Location, 0, 0,
		data.NewGraduation(""), data.NewSubject(request.Subject), data.NewGender(request.Gender),
		data.NewMethod(request.Method))
	log.Println(result)
	if err != nil {
		return nil, data.NotMore
	}
	return filterSchedule(result, request.Schedule)
}


func filterSchedule(teachers []*responses.TeacherSearch, schedule data.Schedule) ([]*responses.TeacherSearch, error) {
	res := []*responses.TeacherSearch{}
	for _, teacher := range teachers {
		if data.CheckSchedule(teacher.Schedule, schedule) {
			res = append(res, &responses.TeacherSearch{
				Username:    teacher.Username,
				TeacherInfo: teacher.TeacherInfo,
				Schedule:    teacher.Schedule,
			})
		}
	}

	return res, data.Success
}

func (t *requestHandler) GetAllRequestOfUser(username string) ([]*responses.Request, error) {
	list, err := (&models.Request{}).GetAllOfUser(username)
	if err != nil {
		return nil, data.NotMore
	}
	return list, data.Success
}