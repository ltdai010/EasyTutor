package adminusecase

import (
	"EasyTutor/data/data"
	"EasyTutor/middlerware"
	"EasyTutor/models"
	"golang.org/x/crypto/bcrypt"
)

func (a *adminHandler) Login(request data.LoginInfo) (string, error) {
	user := &models.Admin{}
	user.Username = request.Username
	err := user.Get()
	if err != nil {
		return "", data.BadRequest
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		return "", data.NotPermission
	}

	return middlerware.GenerateToken(request.Username, "admin")
}

func (a *adminHandler) ValidateTeacher(teacherID string) error {
	teacher := &models.Teacher{}
	teacher.Username = teacherID
	err := teacher.Get()
	if err != nil {
		return data.NotExisted
	}

	teacher.Active = true
	err = teacher.Update()
	if err != nil {
		return data.ErrSystem
	}
	return data.Success
}