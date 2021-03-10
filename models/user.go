package models

import (
	"EasyTutor/consts"
	"EasyTutor/data/data"
	"EasyTutor/data/rest/responses"
	"EasyTutor/drivers"
	"cloud.google.com/go/firestore"
	"context"
)

type User struct {
	data.User
}

func (t *User) GetCollectionKey() string {
	return consts.USER
}

func (t *User) GetCollection() *firestore.CollectionRef {
	return drivers.GetDriver().GetCloudStore().Collection(t.GetCollectionKey())
}

func (t *User) Add() (string, error) {
	_, err := t.GetCollection().Doc(t.Username).Set(context.Background(), t.User)
	if err != nil {
		return "", err
	}
	return t.Username, nil
}

func (t *User) Get() error {
	doc, err := t.GetCollection().Doc(t.Username).Get(context.Background())
	if err != nil {
		return err
	}
	err = doc.DataTo(&t.User)
	if err != nil {
		return err
	}
	return nil
}

func (t *User) Update() error {
	_, err := t.GetCollection().Doc(t.Username).Set(context.Background(), t.User)
	if err != nil {
		return err
	}
	return nil
}

func (t *User) Delete() error {
	_, err := t.GetCollection().Doc(t.Username).Delete(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (t *User) GetPage(pageNumber int, pageSize int) ([]*responses.User, int, error) {
	res := []*responses.User{}
	start := pageNumber * pageSize
	end := (pageNumber + 1) * pageSize
	list, err := t.GetCollection().DocumentRefs(context.Background()).GetAll()
	if err != nil {
		return nil, 0, err
	}
	total := len(list)
	if start > total {
		return res, 0, data.NotMore
	}
	if end > total {
		end = total
	}
	for _, ref := range list[start:end] {
		user := responses.User{}
		doc, err := t.GetCollection().Doc(ref.ID).Get(context.Background())
		if err != nil {
			continue
		}
		err = doc.DataTo(&user)
		if err != nil {
			continue
		}
		user.Username = ref.ID
		res = append(res, &user)
	}
	return res, total, err
}