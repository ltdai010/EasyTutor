package models

import (
	"EasyTutor/consts"
	"EasyTutor/data/data"
	"EasyTutor/drivers"
	"cloud.google.com/go/firestore"
	"context"
	"log"
	"time"
)

type Notification struct {
	ID string	`json:"id"`
	data.Notification
}

func (n *Notification) GetCollectionKey() string {
	return consts.NOTIFICATION
}

func (n *Notification) GetCollection() *firestore.CollectionRef {
	return drivers.GetDriver().GetCloudStore().Collection(n.GetCollectionKey())
}

func (n *Notification) Add() (string, error) {
	ref, _, err := n.GetCollection().Add(context.Background(), n.Notification)
	if err != nil {
		return "", err
	}
	return ref.ID, nil
}

func (n *Notification) Get() error {
	doc, err := n.GetCollection().Doc(n.ID).Get(context.Background())
	if err != nil {
		return err
	}
	err = doc.DataTo(&n.Notification)
	return err
}

func (n *Notification) GetRecent(pageNumber, pageSize int, username string, userType string) ([]*data.Notification, int, error)  {
	res := []*data.Notification{}
	start := pageNumber * pageSize
	end := (pageNumber + 1) * pageSize
	list, err := n.GetCollection().Where("Username", "==", username).Where("UserType", "==", userType).
		OrderBy("CreateTime", firestore.Desc).Documents(context.Background()).GetAll()
	if err != nil {
		log.Println(err)
		return nil, 0, err
	}
	total := len(list)
	if start > total {
		return res, 0, data.NotMore
	}
	if end > total {
		end = total
	}
	for _, i := range list[start:end] {
		noti := data.Notification{}
		err = i.DataTo(&noti)
		if err != nil {
			continue
		}
		res = append(res, &noti)
	}
	return res, total, nil
}

func (n *Notification) GetRecentResetPassCode(username, userType string) (error) {
	noti, err := n.GetCollection().Where("Username", "==", username).Where("UserType", "==", userType).
		Where("NotifyType", "==", data.ResetPassword).OrderBy("CreateTime", firestore.Desc).Limit(1).
		Documents(context.Background()).GetAll()
	if err != nil || len(noti) == 0{
		return err
	}
	f := data.Notification{}
	err = noti[0].DataTo(&f)
	if err != nil || f.CreateTime + 120 < time.Now().Unix(){
		return data.BadRequest
	}
	n.ID = noti[0].Ref.ID
	n.Notification = f
	return nil
}

func (n *Notification) Delete() error {
	_, err := n.GetCollection().Doc(n.ID).Delete(context.Background())
	return err
}