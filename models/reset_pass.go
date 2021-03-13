package models

import (
	"EasyTutor/consts"
	"EasyTutor/drivers"
	"cloud.google.com/go/firestore"
	"context"
)

type ResetCode struct {
	Username   string `json:"username"`
	Code       string `json:"code"`
	CreateTime int64  `json:"create_time"`
}

func (n *ResetCode) GetCollectionKey() string {
	return consts.RESET
}

func (n *ResetCode) GetCollection() *firestore.CollectionRef {
	return drivers.GetDriver().GetCloudStore().Collection(n.GetCollectionKey())
}

func (n *ResetCode) Add() error {
	_, err := n.GetCollection().Doc(n.Username).Set(context.Background(), n)
	return err
}

func (n *ResetCode) Get() error {
	doc, err := n.GetCollection().Doc(n.Username).Get(context.Background())
	if err != nil {
		return err
	}
	err = doc.DataTo(n)
	return err
}