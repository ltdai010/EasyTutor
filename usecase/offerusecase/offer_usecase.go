package offerusecase

import (
	"EasyTutor/data/rest/requests"
	"EasyTutor/data/rest/responses"
)

type OfferInterface interface {
	CreateOne(teacherID, requestID string, post requests.OfferPost) (string, error)
	GetOfferOfRequest(requestID string) ([]*responses.Offer, error)
	GetOfferOfTeacher(teacherID string) ([]*responses.Offer, error)
	GetOne(id string) (*responses.Offer, error)
	UpdateOne(username, id string, put requests.OfferPut) error
	RemoveOne(username, id string) error
}

func GetOfferUseCase() OfferInterface {
	return &offerHandler{}
}

type offerHandler struct {}