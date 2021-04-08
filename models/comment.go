package models

import (
	"EasyTutor/consts"
	"EasyTutor/data/data"
	"EasyTutor/data/rest/responses"
	"EasyTutor/drivers"
	"cloud.google.com/go/firestore"
	"context"
	"log"
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

func (t *Comment) GetActiveCommentOfTeacher(teacherID string) ([]*responses.Comment, error) {
	res := []*responses.Comment{}
	docs, err := t.GetCollection().Where("TeacherID", "==", teacherID).Documents(context.Background()).GetAll()
	if err != nil {
		return res, err
	}
	for _, doc := range docs {
		comment := responses.Comment{}
		err = doc.DataTo(&comment)
		if err != nil || !comment.Active {
			continue
		}
		comment.ID = doc.Ref.ID
		res = append(res, &comment)
	}
	return res, nil
}

func (t *Comment) GetUnActiveCommentOfAll() ([]*responses.TeacherComment, error) {
	result := map[string][]*responses.Comment{}
	res := []*responses.TeacherComment{}
	docs, err := t.GetCollection().Where("Active", "==", false).Documents(context.Background()).GetAll()
	if err != nil {
		log.Println(err)
		return res, err
	}
	for _, doc := range docs {
		comment := responses.Comment{}
		err = doc.DataTo(&comment)
		if err != nil {
			continue
		}
		comment.ID = doc.Ref.ID
		if result[comment.TeacherID] == nil {
			result[comment.TeacherID] = []*responses.Comment{}
		}
		result[comment.TeacherID] = append(result[comment.TeacherID], &comment)
	}
	//convert
	for teacherID, listComment := range result {
		doc, err := (&Teacher{}).GetCollection().Doc(teacherID).Get(context.Background())
		if err != nil {
			log.Println(err)
			continue
		}
		teacher := Teacher{}
		err = doc.DataTo(&teacher)
		if err != nil {
			log.Println(err)
			continue
		}
		res = append(res, &responses.TeacherComment{
			TeacherID:           teacher.Username,
			Name:                teacher.Name,
			ListUnActiveComment: listComment,
		})
	}

	return res, nil
}