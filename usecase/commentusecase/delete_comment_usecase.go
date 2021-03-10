package commentusecase

import (
	"EasyTutor/data/data"
	"EasyTutor/models"
	"EasyTutor/utils/myerror"
)

func (t *CommentHandler) RemoveOne(username, id string) error {
	comment := &models.Comment{}
	comment.ID = id
	err := comment.Get()
	if err != nil {
		return data.NotExisted
	}

	if comment.Username != username {
		return data.NotPermission
	}

	err = comment.Delete()
	if myerror.IsError(err) {
		return data.ErrSystem
	}
	return data.Success
}


