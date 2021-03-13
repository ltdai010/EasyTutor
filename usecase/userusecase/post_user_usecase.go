package userusecase

import (
	"EasyTutor/data/data"
	"EasyTutor/data/rest/requests"
	"EasyTutor/middleware"
	"EasyTutor/models"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
	"math/rand"
	"time"
)

func (t *userHandler) CreateOne(request requests.UserPost) (string, error) {
	user := &models.User{}

	user.Username = request.Username
	err := user.Get()
	if err == nil {
		return "", data.Existed
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", data.BadRequest
	}

	user.LoginInfo = data.LoginInfo{
		Username: request.Username,
		Password: string(hashed),
	}

	user.UserInfo = request.UserInfo
	if len(user.FavoriteTeacher) == 0 {
		user.FavoriteTeacher = []string{}
	}

	id, err := user.Add()
	if err != nil {
		return "", data.ErrSystem
	}
	return id, data.Success
}

func (t *userHandler) Login(login data.LoginInfo) (string, error) {
	user := &models.User{}
	log.Println(login)
	user.Username = login.Username

	err := user.Get()
	if err != nil {
		return "", data.NotExisted
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password)); err != nil {
		return "", data.BadRequest
	}

	token, err := middleware.GenerateToken(login.Username, "user")
	if err != nil {
		log.Println(err, " usecase/userusecase/post_user_usecase.go:58")
		return "", data.ErrLogin
	}
	return token, data.Success
}

func (t *userHandler) ForgotPassword(username string) error {
	user := &models.User{}
	mail := &models.Mail{}
	user.Username = username
	err := user.Get()
	if err != nil {
		return data.NotExisted
	}

	rand.Seed(time.Now().Unix())
	code := rand.Intn(999999)

	models.GetHub().BroadcastMessage(data.Notification{
		NotificationInfo: data.NotificationInfo{
			Username:   username,
			UserType:   "user",
			NotifyType: data.ResetPassword,
			Message:    data.ForgotPassword{CheckCode: fmt.Sprintf("%06d", code)},
		},
		CreateTime: time.Now().Unix(),
	})
	go mail.Send(user.Email)

	return data.Success
}

func (t *userHandler) ValidateResetCode(request requests.ResetPass) error {
	user := &models.User{}
	notification := &models.Notification{}
	user.Username = request.Username
	err := user.Get()
	if err != nil {
		return data.NotExisted
	}
	err = notification.GetRecentResetPassCode(request.Username, "user")
	if err != nil {
		return data.BadRequest
	}
	if notification.Message.(map[string]interface{})["CheckCode"] == request.Code {
		hashed, err := bcrypt.GenerateFromPassword([]byte(request.NewPass), bcrypt.DefaultCost)
		if err != nil {
			return data.ErrSystem
		}
		user.Password = string(hashed)
		err = user.Update()
		if err != nil {
			return data.ErrSystem
		}
		err = notification.Delete()
		if err != nil {
			return data.ErrSystem
		}
		return data.Success
	}
	return data.BadRequest
}