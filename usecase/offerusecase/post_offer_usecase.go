package offerusecase

import (
	"EasyTutor/data/data"
	"EasyTutor/data/rest/requests"
	"EasyTutor/models"
	"EasyTutor/utils/logger"
	"time"
)

func (t *offerHandler) CreateOne(teacherID, requestID string, post requests.OfferPost) (string, error) {
	offer := &models.Offer{}
	teacher := &models.Teacher{}
	request := &models.Request{}

	teacher.Username = teacherID
	err := teacher.Get()
	if err != nil || !teacher.Active {
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
		logger.Error("[Error Create One] Create one error offer id = %v err = %v", offer.ID, err)
		return "", data.ErrSystem
	}

	models.GetHub().BroadcastMessage(data.Notification{
		NotificationInfo: data.NotificationInfo{
			Username:   request.Username,
			UserType:   "user",
			NotifyType: data.PostOffer,
			Message:    data.OfferNotify{
				Message:     "Một giáo viên đã thêm đề nghị dạy học",
				OfferID:     id,
				TeacherName: teacherID,
			},
		},
		CreateTime: time.Now().Unix(),
	})
	return id, data.Success
}

