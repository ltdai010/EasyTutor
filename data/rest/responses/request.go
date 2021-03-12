package responses

import "EasyTutor/data/data"

type Request struct {
	ID string `json:"id"`
	data.Request
}

type RequestSearch struct {
	ID string `json:"id"`
	data.RequestInfo
	data.Schedule
}
