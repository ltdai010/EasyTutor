package models

import (
	"EasyTutor/consts"
	"EasyTutor/data/data"
	"EasyTutor/drivers"
	"cloud.google.com/go/firestore"
	"context"
)

type Admin struct {
	data.LoginInfo
}

func (t *Admin) GetCollectionKey() string {
	return consts.ADMIN
}

func (t *Admin) GetCollection() *firestore.CollectionRef {
	return drivers.GetDriver().GetCloudStore().Collection(t.GetCollectionKey())
}

func (a *Admin) Get() error {
	doc, err := a.GetCollection().Doc(a.Username).Get(context.Background())
	if err != nil {
		return err
	}
	err = doc.DataTo(&a.LoginInfo)
	return err
}