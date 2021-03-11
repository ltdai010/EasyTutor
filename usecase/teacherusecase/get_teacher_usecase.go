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
