package searchdata

import "EasyTutor/data/data"

type RequestSearch struct {
	ObjectID string	`json:"objectID"`
	data.Request
}
