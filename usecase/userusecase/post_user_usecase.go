package userusecase

import (
	"EasyTutor/data/data"
	"EasyTutor/data/rest/requests"
	"EasyTutor/data/rest/responses"
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

func (t *userHandler) Login(login data.LoginInfo) (*responses.UserLogin, error) {
	user := &models.User{}
	user.Username = login.Username

	err := user.Get()
	if err != nil {
		return nil, data.NotExisted
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password)); err != nil {
		return nil, data.BadRequest
	}

	token, err := middleware.GenerateToken(login.Username, "user")
	if err != nil {
		log.Println(err, " usecase/userusecase/post_user_usecase.go:58")
		return nil, data.ErrLogin
	}
	return &responses.UserLogin{
		Token: token,
		Username: user.Username,
		UserInfo: user.UserInfo,
	}, data.Success
}

func (t *userHandler) ForgotPassword(username string) error {
	user := &models.User{}
	mail := &models.Mail{}
	reset := &models.ResetCode{}
	user.Username = username
	err := user.Get()
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
		To:      user.Email,
		Subject: "[Easy-Tutor] Reset mật khẩu tài khoản người dùng",
		Msg:     "Mã Reset mật khẩu tài khoản người dùng " + username + " : " + fmt.Sprintf("%06d", code),
	}
	err = reset.Add()
	if err != nil {
		return data.ErrSystem
	}

	go mail.Send(user.Email)

	return data.Success
}

func (t *userHandler) ValidateResetCode(request requests.ResetPass) error {
	user := &models.User{}
	reset := &models.ResetCode{}
	user.Username = request.Username
	err := user.Get()
	if err != nil {
		return data.NotExisted
	}
	reset.Username = request.Username
	err = reset.Get()
	if err != nil {
		return data.NotExisted
	}
	if reset.Code == request.Code {
		hashed, err := bcrypt.GenerateFromPassword([]byte(request.NewPass), bcrypt.DefaultCost)
		if err != nil {
			return data.ErrSystem
		}
		user.Password = string(hashed)
		err = user.Update()
		if err != nil {
			return data.ErrSystem
		}
		return data.Success
	}
	return data.BadRequest
}