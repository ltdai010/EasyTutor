package teacherusecase

import (
	"EasyTutor/data/data"
	"EasyTutor/data/rest/requests"
	"EasyTutor/middlerware"
	"EasyTutor/models"
	"golang.org/x/crypto/bcrypt"
)

func (t *teacherHandler) CreateOne(request requests.TeacherPost) (string, error) {
	teacher := &models.Teacher{}

	teacher.Username = request.Username
	err := teacher.Get()
	if err == nil {
		return "", data.Existed
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", data.BadRequest
	}

	teacher.LoginInfo = data.LoginInfo{
		Username: request.Username,
		Password: string(hashed),
	}
	teacher.TeacherInfo, err = data.SetDataTeacher(request.TeacherInfo)
	teacher.Active = false
	if err != nil {
		return "", data.BadRequest
	}

	teacher.Schedule = data.Schedule{
		MonMorning:   false,
		MonAfternoon: false,
		MonNight:     false,
		TueMorning:   false,
		TueAfternoon: false,
		TueNight:     false,
		WedMorning:   false,
		WedAfternoon: false,
		WedNight:     false,
		ThuMorning:   false,
		ThuAfternoon: false,
		ThuNight:     false,
		FriMorning:   false,
		FriAfternoon: false,
		FriNight:     false,
		SatMorning:   false,
		SatAfternoon: false,
		SatNight:     false,
		SunMorning:   false,
		SunAfternoon: false,
		SunNight:     false,
	}

	id, err := teacher.Add()
	if err != nil {
		return "", data.ErrSystem
	}
	return id, data.Success
}

func (t *teacherHandler) Login(login data.LoginInfo) (string, error) {
	teacher := &models.Teacher{}
	teacher.Username = login.Username

	err := teacher.Get()
	if err != nil {
		return "", data.NotExisted
	}

	if !teacher.Active {
		return "", data.NotPermission
	}

	if err = bcrypt.CompareHashAndPassword([]byte(teacher.Password), []byte(login.Password)); err != nil {
		return "", data.BadRequest
	}

	token, err := middlerware.GenerateToken(login.Username, "teacher")
	if err != nil {
		return "", data.ErrLogin
	}
	return token, data.Success
}