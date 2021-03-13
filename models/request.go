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
)

type Request struct {
	ID string 	`json:"id"`
	data.Request
}

func (t *Request) GetCollectionKey() string {
	return consts.REQUEST
}

func (t *Request) GetCollection() *firestore.CollectionRef {
	return drivers.GetDriver().GetCloudStore().Collection(t.GetCollectionKey())
}

func (t *Request) GetSearchIndex() search.IndexInterface {
	return drivers.GetDriver().GetSearchRequest()
}

func (t *Request) Add() (string, error) {
	ref, _, err := t.GetCollection().Add(context.Background(), t.Request)
	if err != nil {
		return "", err
	}
	_, err = t.GetSearchIndex().SaveObject(searchdata.RequestSearch{
		ObjectID: ref.ID,
		Request:  t.Request,
	})
	if err != nil {
		log.Println(err, " error save search request")
	}
	return ref.ID, nil
}

func (t *Request) Get() error {
	doc, err := t.GetCollection().Doc(t.ID).Get(context.Background())
	if err != nil {
		return err
	}
	err = doc.DataTo(&t.Request)
	if err != nil {
		return err
	}
	return nil
}

func (t *Request) Update() error {
	_, err := t.GetCollection().Doc(t.ID).Set(context.Background(), t.Request)
	if err != nil {
		return err
	}
	_, err = t.GetSearchIndex().SaveObject(searchdata.RequestSearch{
		ObjectID: t.ID,
		Request:  t.Request,
	})
	if err != nil {
		log.Println(err, " error update search index")
	}
	return nil
}

func (t *Request) Delete() error {
	_, err := t.GetCollection().Doc(t.ID).Delete(context.Background())
	if err != nil {
		return err
	}
	_, err = t.GetSearchIndex().DeleteObject(t.ID)
	if err != nil {
		log.Println(err, " error delete search index")
	}
	return nil
}

func (t *Request) GetPage(pageNumber int, pageSize int) ([]*responses.Request, int, error) {
	res := []*responses.Request{}
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
		request := responses.Request{}
		doc, err := t.GetCollection().Doc(ref.ID).Get(context.Background())
		if err != nil {
			continue
		}
		err = doc.DataTo(&request)
		if err != nil {
			continue
		}
		request.ID = ref.ID
		res = append(res, &request)
	}
	return res, total, err
}

func (t *Request) Search(key string, location, pageNumber, pageSize int,
	subjects []data.Subject, gender data.Gender, methods []data.Method) ([]*responses.RequestSearch, error) {
	res := []*responses.RequestSearch{}
	searchResult := search.QueryRes{}
	var err error
	filters := ""
	if location > 0 {
		filters += "location:'" + fmt.Sprint(location) + "'"
	}
	if len(subjects) > 0 {
		if filters != "" {
			filters += " AND ("
		}
		for i, subject := range subjects {
			if subject != "" {
				if i != 0 {
					filters += " OR subject:'" + string(subject) + "'"
				} else {
					filters += "subject:'" + string(subject) + "'"
				}
			}
		}
		filters += ")"
	}
	if gender != "" {
		if filters != "" {
			filters += " AND "
		}
		filters += "gender:'" + string(gender) + "'"
	}
	if len(methods) > 0 {
		if filters != "" {
			filters += " AND ("
		}
		for i, method := range methods {
			if method != "" {
				if i != 0 {
					filters += " OR method:'" + string(method) + "'"
				} else {
					filters += "method:'" + string(method) + "'"
				}

			}
		}
		filters += ")"
	}
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
	result := []*searchdata.RequestSearch{}
	err = searchResult.UnmarshalHits(&result)
	if err != nil {
		return nil, err
	}
	for _, i := range result {
		res = append(res, &responses.RequestSearch{
			ID:    i.ObjectID,
			Request: i.Request,
		})
	}
	return res, nil
}