package commentusecase

import (
	"EasyTutor/data/data"
	"EasyTutor/data/rest/requests"
	"EasyTutor/models"
	"time"
)

func (t *commentHandler) CreateOne(username, teacherID string, post requests.CommentPost) (string, error) {
	//var
	comment := &models.Comment{}
	user := &models.User{}
	teacher := &models.Teacher{}

	user.Username = username
	err := user.Get()
	if err != nil {
		return "", data.NotExisted
	}

	teacher.Username = teacherID
	err = teacher.Get()
	if err != nil {
		return "", data.NotExisted
	}

	comment.TeacherID = teacherID
	comment.Username = username
	comment.CommentInfo = post.CommentInfo
	comment.CreateTime = time.Now().Unix()

	id, err := comment.Add()
	if err != nil {
		return "", data.ErrSystem
	}
	return id, data.Success
}
