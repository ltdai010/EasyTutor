package adminusecase

import (
	"EasyTutor/data/data"
	"EasyTutor/middlerware"
	"EasyTutor/models"
	"EasyTutor/utils/logger"
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

	logger.Info("[Admin Login] admin login admin_id = %v", request.Username)
	return middlerware.GenerateToken(request.Username, "admin")
}
