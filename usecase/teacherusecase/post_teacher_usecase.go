package teacherusecase

import (
	"EasyTutor/data/data"
	"EasyTutor/data/rest/requests"
	"EasyTutor/middleware"
	"EasyTutor/models"
	"EasyTutor/utils/logger"
	"EasyTutor/utils/myerror"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"time"
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
	if myerror.IsError(err) {
		return "", err
	}
	teacher.Active = false

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
		logger.Info("[CreateOne Teacher] Create Teacher error = %v", err)
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

	token, err := middleware.GenerateToken(login.Username, "teacher")
	if err != nil {
		return "", data.ErrLogin
	}
	return token, data.Success
}

func (t *teacherHandler) ForgotPassword(username string) error {
	teacher := &models.Teacher{}
	mail := &models.Mail{}
	reset := &models.ResetCode{}
	teacher.Username = username
	err := teacher.Get()
	if err != nil {
		return data.NotExisted
	}

	rand.Seed(time.Now().Unix())
	code := rand.Intn(999999)

	reset = &models.ResetCode{
		Username:   username,
		Code:       fmt.Sprintf("%06d", code),
		CreateTime: time.Now().Unix(),
	}

	mail = &models.Mail{
		To:      teacher.Email,
		Subject: "[Easy-Tutor] Tạo mới mật khẩu cho tài khoản giáo viên " + username + " của bạn",
		Msg:     "Mã xác nhận của bạn là " + fmt.Sprintf("%06d", code),
	}
	err = reset.Add()
	if err != nil {
		return data.ErrSystem
	}
	go mail.Send(teacher.Email)

	return data.Success
}

func (t *teacherHandler) ValidateResetCode(request requests.ResetPass) error {
	teacher := &models.Teacher{}
	reset := &models.ResetCode{}
	teacher.Username = request.Username
	err := teacher.Get()
	if err != nil {
		return data.NotExisted
	}
	reset.Username = request.Username
	err = reset.Get()
	if err != nil {
		return data.BadRequest
	}
	if reset.Code == request.Code {
		hashed, err := bcrypt.GenerateFromPassword([]byte(request.NewPass), bcrypt.DefaultCost)
		if err != nil {
			return data.ErrSystem
		}
		teacher.Password = string(hashed)
		err = teacher.Update()
		if err != nil {
			return data.ErrSystem
		}
		return data.Success
	}
	return data.BadRequest
}