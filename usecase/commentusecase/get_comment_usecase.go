package commentusecase

import (
	"EasyTutor/data/data"
	"EasyTutor/data/rest/responses"
	"EasyTutor/models"
)

func (t *commentHandler) GetOne(id string) (*responses.Comment, error) {
	comment := &models.Comment{}
	comment.ID = id
	err := comment.Get()
	if err != nil {
		return nil, data.NotExisted
	}
	return &responses.Comment{
		ID:      id,
		Comment: comment.Comment,
	}, data.Success
}

func (t *commentHandler) GetCommentOfTeacher(teacherID string) ([]*responses.Comment, error) {
	comment := &models.Comment{}
	listComment, err := comment.GetCommentOfTeacher(teacherID)
	if err != nil {
		return nil, data.NotMore
	}
	return listComment, data.Success
}
