package models

import (
	"EasyTutor/consts"
	"EasyTutor/data/data"
	"EasyTutor/data/rest/responses"
	"EasyTutor/drivers"
	"cloud.google.com/go/firestore"
	"context"
)

type Offer struct {
	ID string	`json:"id"`
	data.Offer
}


func (t *Offer) GetCollectionKey() string {
	return consts.OFFER
}

func (t *Offer) GetCollection() *firestore.CollectionRef {
	return drivers.GetDriver().GetCloudStore().Collection(t.GetCollectionKey())
}

func (t *Offer) Add() (string, error) {
	ref, _, err := t.GetCollection().Add(context.Background(), t.Offer)
	if err != nil {
		return "", err
	}
	return ref.ID, nil
}

func (t *Offer) Get() error {
	doc, err := t.GetCollection().Doc(t.ID).Get(context.Background())
	if err != nil {
		return err
	}
	err = doc.DataTo(&t.Offer)
	if err != nil {
		return err
	}
	return nil
}

func (t *Offer) Update() error {
	_, err := t.GetCollection().Doc(t.ID).Set(context.Background(), t.Offer)
	if err != nil {
		return err
	}
	return nil
}

func (t *Offer) Delete() error {
	_, err := t.GetCollection().Doc(t.ID).Delete(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (t *Offer) GetOfferOfRequest(requestID string) ([]*responses.Offer, error) {
	res := []*responses.Offer{}
	docs, err := t.GetCollection().Where("RequestID", "==", requestID).Documents(context.Background()).GetAll()
	if err != nil {
		return res, err
	}
	for _, doc := range docs {
		offer := responses.Offer{}
		err = doc.DataTo(&offer)
		if err != nil {
			continue
		}
		offer.ID = doc.Ref.ID
		res = append(res, &offer)
	}
	return res, nil
}

func (t *Offer) GetOfferOfTeacher(teacherID string) ([]*responses.Offer, error) {
	res := []*responses.Offer{}
	docs, err := t.GetCollection().Where("TeacherID", "==", teacherID).Documents(context.Background()).GetAll()
	if err != nil {
		return nil, err
	}
	for _, doc := range docs {
		offer := responses.Offer{}
		err = doc.DataTo(&offer)
		if err != nil {
			continue
		}
		offer.ID = doc.Ref.ID
		res = append(res, &offer)
	}
	return res, nil
}