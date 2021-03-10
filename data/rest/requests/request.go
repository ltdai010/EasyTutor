package requests

import "EasyTutor/data/data"

type RequestPost struct {
	data.RequestInfo
	data.Schedule
}

type RequestPut struct {
	data.RequestInfo
	data.Schedule
}
