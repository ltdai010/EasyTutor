package models

import (
	"EasyTutor/consts"
	"EasyTutor/data/data"
	"EasyTutor/data/rest/responses"
	"EasyTutor/drivers"
	"cloud.google.com/go/firestore"
	"context"
)

type Comment struct {
	ID string	`json:"id"`
	data.Comment
}

func (t *Comment) GetCollectionKey() string {
	return consts.COMMENT
}

func (t *Comment) GetCollection() *firestore.CollectionRef {
	return drivers.GetDriver().GetCloudStore().Collection(t.GetCollectionKey())
}

func (t *Comment) Add() (string, error) {
	ref, _, err := t.GetCollection().Add(context.Background(), t.Comment)
	if err != nil {
		return "", err
	}
	return ref.ID, nil
}

func (t *Comment) Get() error {
	doc, err := t.GetCollection().Doc(t.ID).Get(context.Background())
	if err != nil {
		return err
	}
	err = doc.DataTo(&t.Comment)
	if err != nil {
		return err
	}
	return nil
}

func (t *Comment) Update() error {
	_, err := t.GetCollection().Doc(t.ID).Set(context.Background(), t.Comment)
	if err != nil {
		return err
	}
	return nil
}

func (t *Comment) Delete() error {
	_, err := t.GetCollection().Doc(t.ID).Delete(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (t *Comment) GetCommentOfTeacher(teacherID string) ([]*responses.Comment, error) {
	res := []*responses.Comment{}
	docs, err := t.GetCollection().Where("TeacherID", "==", teacherID).Documents(context.Background()).GetAll()
	if err != nil {
		return res, err
	}
	for _, doc := range docs {
		comment := responses.Comment{}
		err = doc.DataTo(&comment)
		if err != nil {
			continue
		}
		comment.ID = doc.Ref.ID
		res = append(res, &comment)
	}
	return res, nil
}