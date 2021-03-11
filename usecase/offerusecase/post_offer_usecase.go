package offerusecase

import (
	"EasyTutor/data/data"
	"EasyTutor/data/rest/requests"
	"EasyTutor/models"
	"time"
)

func (t *offerHandler) CreateOne(teacherID, requestID string, post requests.OfferPost) (string, error) {
	offer := &models.Offer{}
	teacher := &models.Teacher{}
	request := &models.Request{}

	teacher.Username = teacherID
	err := teacher.Get()
	if err != nil {
		return "", data.NotExisted
	}

	request.ID = requestID
	err = request.Get()
	if err != nil || request.AcceptOffer != "" {
		return "", data.NotExisted
	}

	offer.OfferInfo = post.OfferInfo
	offer.RequestID = requestID
	offer.TeacherID = teacherID
	offer.CreateTime = time.Now().Unix()

	id, err := offer.Add()
	if err != nil {
		return "", data.ErrSystem
	}
	return id, data.Success
}

