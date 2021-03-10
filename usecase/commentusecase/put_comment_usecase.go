package commentusecase

import (
	"EasyTutor/data/data"
	"EasyTutor/data/rest/requests"
	"EasyTutor/models"
	"EasyTutor/utils/myerror"
)

func (t *CommentHandler) UpdateOne(username, id string, put requests.CommentPut) error {
	comment := &models.Comment{}
	comment.ID = id
	err := comment.Get()
	if myerror.IsError(err) {
		return data.NotExisted
	}

	if comment.Username != username {
		return data.NotPermission
	}
	comment.CommentInfo = put.CommentInfo

	err = comment.Update()
	if myerror.IsError(err) {
		return data.ErrSystem
	}
	return data.Success
}
