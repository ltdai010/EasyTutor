package teacherusecase

import (
	"EasyTutor/data/data"
	"EasyTutor/data/rest/responses"
	"EasyTutor/models"
)

func (t *teacherHandler) GetOne(username string) (*responses.Teacher, error) {
	teacher := &models.Teacher{}
	teacher.Username = username
	err := teacher.Get()
	if err != nil {
		return nil, data.NotExisted
	}
	return &responses.Teacher{
		TeacherInfo: teacher.TeacherInfo,
		Schedule:    teacher.Schedule,
	}, data.Success
}

func (t *teacherHandler) GetPage(pageNumber int, pageSize int) ([]*responses.Teacher, int, error) {
	teacher := &models.Teacher{}
	listTeacher, total, err := teacher.GetPageActive(pageNumber - 1, pageSize)
	if err != nil {
		return nil, 0, data.NotMore
	}
	return listTeacher, total, data.Success
}

func (t *teacherHandler) FindAvailableRequest(teacherID string) ([]*responses.RequestSearch, error) {
	request := &models.Request{}
	teacher := &models.Teacher{}
	teacher.Username = teacherID
	err := teacher.Get()
	if err != nil {
		return nil, data.NotExisted
	}
	result, err := request.Search("", teacher.Location, 0, 0,
		data.ListStringToSubjects(teacher.ListSubject), data.NewGender(teacher.Gender),
		data.ListStringToMethod(teacher.ListMethod))
	if err != nil {
		return nil, data.NotMore
	}

	return filterSchedule(result, teacher.Schedule)
}

func filterSchedule(requests []*responses.RequestSearch, schedule data.Schedule) ([]*responses.RequestSearch, error) {
	res := []*responses.RequestSearch{}
	for _, request := range requests {
		if data.CheckSchedule(schedule, request.Schedule) {
			res = append(res, request)
		}
	}
	return res, data.Success
}

func (t *teacherHandler) Profile(username string) (*responses.Teacher, error) {
	teacher := &models.Teacher{}
	teacher.Username = username
	err := teacher.Get()
	if err != nil {
		return nil, data.NotExisted
	}
	return &responses.Teacher{
		Username: username,
		TeacherInfo: teacher.TeacherInfo,
	}, data.Success
}