package myerror

import (
	"EasyTutor/data/data"
)

func IsError(err error) bool {
	if err == nil || err == data.Success {
		return false
	}
	return true
}
