package models

import (
	"EasyTutor/consts"
	"EasyTutor/data/data"
	"EasyTutor/data/rest/responses"
	"EasyTutor/data/searchdata"
	"EasyTutor/drivers"
	"cloud.google.com/go/firestore"
	"context"
	"fmt"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/opt"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/search"
	"log"
	"time"
)

type Teacher struct {
	data.Teacher
}

func (t *Teacher) GetCollectionKey() string {
	return consts.TEACHER
}

func (t *Teacher) GetWaitUpdateCollectionKey() string {
	return consts.WAIT
}

func (t *Teacher) GetWaitUpdateCollection() *firestore.CollectionRef {
	return drivers.GetDriver().GetCloudStore().Collection(t.GetWaitUpdateCollectionKey())
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
		Schedule:	 t.Schedule,
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

func (t *Teacher) GetUpdating() error {
	doc, err := t.GetWaitUpdateCollection().Doc(t.Username).Get(context.Background())
	if err != nil {
		return err
	}
	err = doc.DataTo(&t.Teacher)
	if err != nil {
		return err
	}
	return nil
}

func (t *Teacher) UpdateToWait() error {
	_, err := t.GetWaitUpdateCollection().Doc(t.Username).Set(context.Background(), data.TeacherUpdate{
		Teacher:    t.Teacher,
		UpdateTime: time.Now().Unix(),
	})
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
		Schedule:    t.Schedule,
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

func (t *Teacher) GetPageActive(pageNumber int, pageSize int) ([]*responses.Teacher, int, error) {
	res := []*responses.Teacher{}
	start := pageNumber * pageSize
	end := (pageNumber + 1) * pageSize
	list, err := t.GetCollection().Where("Active", "==", true).Documents(context.Background()).GetAll()
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
	for _, t := range list[start:end] {
		teacher := responses.Teacher{}
		err = t.DataTo(&teacher)
		if err != nil {
			continue
		}
		teacher.Username = t.Ref.ID
		res = append(res, &teacher)
	}
	return res, total, nil
}

func (t *Teacher) GetPageUnActive(pageNumber int, pageSize int) ([]*responses.Teacher, int, error) {
	res := []*responses.Teacher{}
	start := pageNumber * pageSize
	end := (pageNumber + 1) * pageSize
	list, err := t.GetCollection().Where("Active", "==", false).Documents(context.Background()).GetAll()
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
	for _, t := range list[start:end] {
		teacher := responses.Teacher{}
		err = t.DataTo(&teacher)
		if err != nil {
			continue
		}
		teacher.Username = t.Ref.ID
		res = append(res, &teacher)
	}
	return res, total, nil
}

func (t *Teacher) Search(key string, location, pageNumber, pageSize int,
	graduation data.Graduation, subject data.Subject, gender data.Gender, method data.Method) ([]*responses.TeacherSearch, error) {
	res := []*responses.TeacherSearch{}
	searchResult := search.QueryRes{}
	var err error
	filters := ""
	if location > 0 {
		filters += "location:'" + fmt.Sprint(location) + "'"
	}
	if graduation != "" {
		if filters != "" {
			filters += " AND "
		}
		filters += "graduation:'" + string(graduation) + "'"
	}
	if subject != "" {
		if filters != "" {
			filters += " AND "
		}
		filters += "list_subject:'" + string(subject) + "'"
	}
	if gender != "" {
		if filters != "" {
			filters += " AND "
		}
		filters += "gender:'" + string(gender) + "'"
	}
	if method != "" {
		if filters != "" {
			filters += " AND "
		}
		filters += "list_method:'" + string(method) + "'"
	}

	log.Println(filters)

	if pageNumber >= 0 && pageSize > 0 {
		searchResult, err = t.GetSearchIndex().Search(key,
			opt.Filters(filters),
			opt.Page(pageNumber),
			opt.HitsPerPage(pageSize),
		)
	} else {
		searchResult, err = t.GetSearchIndex().Search(key,
			opt.Filters(filters),
		)
	}
	if err != nil || len(searchResult.Hits) == 0{
		return nil, data.NotMore
	}
	result := []*searchdata.TeacherSearch{}
	err = searchResult.UnmarshalHits(&result)
	if err != nil {
		return nil, err
	}
	for _, i := range result {
		res = append(res, &responses.TeacherSearch{
			Username:    i.ObjectID,
			TeacherInfo: i.TeacherInfo,
			Schedule: i.Schedule,
		})
	}
	return res, nil
}