package models

import (
	"EasyTutor/consts"
	"EasyTutor/data/data"
	"EasyTutor/data/rest/responses"
	"EasyTutor/data/searchdata"
	"EasyTutor/drivers"
	"cloud.google.com/go/firestore"
	"context"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/search"
	"log"
)

type Teacher struct {
	data.Teacher
}

func (t *Teacher) GetCollectionKey() string {
	return consts.TEACHER
}

func (t *Teacher) GetCollection() *firestore.CollectionRef {
	return drivers.GetDriver().GetCloudStore().Collection(t.GetCollectionKey())
}

func (t *Teacher) GetSearchIndex() search.IndexInterface {
	return drivers.GetDriver().GetSearchTeacher()
}

func (t *Teacher) Add() (string, error) {
	_, err := t.GetCollection().Doc(t.Username).Set(context.Background(), t.Teacher)
	if err != nil {
		return "", err
	}
	_, err = t.GetSearchIndex().SaveObject(searchdata.TeacherSearch{
		ObjectID:    t.Username,
		TeacherInfo: t.TeacherInfo,
	})
	if err != nil {
		log.Println(err, " error create data search teacher")
	}
	return t.Username, nil
}

func (t *Teacher) Get() error {
	doc, err := t.GetCollection().Doc(t.Username).Get(context.Background())
	if err != nil {
		return err
	}
	err = doc.DataTo(&t.Teacher)
	if err != nil {
		return err
	}
	return nil
}

func (t *Teacher) Update() error {
	_, err := t.GetCollection().Doc(t.Username).Set(context.Background(), t.Teacher)
	if err != nil {
		return err
	}
	_, err = t.GetSearchIndex().SaveObject(searchdata.TeacherSearch{
		ObjectID:    t.Username,
		TeacherInfo: t.TeacherInfo,
	})
	if err != nil {
		log.Println(err, " error update data search teacher")
	}
	return nil
}

func (t *Teacher) Delete() error {
	_, err := t.GetCollection().Doc(t.Username).Delete(context.Background())
	if err != nil {
		return err
	}
	_, err = t.GetSearchIndex().DeleteObject(t.Username)
	if err != nil {
		log.Println(err, " error delete data search teacher")
	}
	return nil
}

func (t *Teacher) GetPage(pageNumber int, pageSize int) ([]*responses.Teacher, int, error) {
	res := []*responses.Teacher{}
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
		teacher := responses.Teacher{}
		doc, err := t.GetCollection().Doc(ref.ID).Get(context.Background())
		if err != nil {
			continue
		}
		err = doc.DataTo(&teacher)
		if err != nil {
			continue
		}
		teacher.Username = ref.ID
		res = append(res, &teacher)
	}
	return res, total, err
}