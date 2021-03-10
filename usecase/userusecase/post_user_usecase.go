package userusecase

import (
	"EasyTutor/data/data"
	"EasyTutor/data/rest/requests"
	"EasyTutor/middlerware"
	"EasyTutor/models"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func (t *UserHandler) CreateOne(request requests.UserPost) (string, error) {
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

func (t *UserHandler) Login(login data.LoginInfo) (string, error) {
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

	token, err := middlerware.GenerateToken(login.Username, "user")
	if err != nil {
		log.Println(err, " usecase/userusecase/post_user_usecase.go:58")
		return "", data.ErrLogin
	}
	return token, data.Success
}
