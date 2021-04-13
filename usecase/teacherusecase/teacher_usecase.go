package teacherusecase

import (
	"EasyTutor/data/data"
	"EasyTutor/data/rest/requests"
	"EasyTutor/data/rest/responses"
)

type TeacherInterface interface {
	CreateOne(post requests.TeacherPost) (string, error)
	GetPage(pageNumber int, pageSize int) ([]*responses.Teacher, int, error)
	GetOne(id string) (*responses.Teacher, error)
	UpdateOne(id string, put requests.TeacherPut) error
	RemoveOne(id string) error
	Login(login data.LoginInfo) (*responses.TeacherLogin, error)
	FindAvailableRequest(teacherID string) ([]*responses.RequestSearch, error)
	Profile(username string) (*responses.Teacher, error)
	UpdateSchedule(id string, put data.Schedule) error
	ForgotPassword(username string) error
	ValidateResetCode(request requests.ResetPass) error
}

func GetTeacherUseCase() TeacherInterface {
	return &teacherHandler{}
}

type teacherHandler struct {}