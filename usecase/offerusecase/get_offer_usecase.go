package offerusecase

import (
	"EasyTutor/data/data"
	"EasyTutor/data/rest/responses"
	"EasyTutor/models"
)

func (t *offerHandler) GetOne(id string) (*responses.Offer, error) {
	offer := &models.Offer{}
	offer.ID = id
	err := offer.Get()
	if err != nil {
		return nil, data.NotExisted
	}
	return &responses.Offer{
		ID:      id,
		Offer: offer.Offer,
	}, data.Success
}

func (t *offerHandler) GetOfferOfRequest(requestID string) ([]*responses.Offer, error) {
	offer := &models.Offer{}
	listOffer, err := offer.GetOfferOfRequest(requestID)
	if err != nil {
		return nil, data.NotMore
	}
	return listOffer, data.Success
}
