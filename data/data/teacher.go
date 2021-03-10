package data

import (
	"EasyTutor/utils/structdata"
)

type TeacherInfo struct {
	Name            string           `json:"name"`
	ListSubject     []string		 `json:"list_subject"`
	Location        int              `json:"location"`
	ListMethod      []string		 `json:"list_method"`
	Description     string           `json:"description"`
	Achievement     string           `json:"achievement"`
	Topic           []string         `json:"topic"`
	BackgroundImage string           `json:"background_image"`
	Avatar          string           `json:"avatar"`
	Gender          string           `json:"gender"`
	DateOfBirth     int64            `json:"date_of_birth"`
	Graduation      string	         `json:"graduation"`
	Active		    bool			 `json:"active"`
}

type Teacher struct {
	LoginInfo
	TeacherInfo
	Schedule
}

func SetDataTeacher(request TeacherInfo) (TeacherInfo, error) {
	teacher := TeacherInfo{
		Name:            request.Name,
		ListSubject:     []string{},
		Location:        request.Location,
		ListMethod:      []string{},
		Description:     request.Description,
		Achievement:     request.Achievement,
		Topic:           []string{},
		BackgroundImage: request.BackgroundImage,
		Avatar:          request.Avatar,
		Gender:          request.Gender,
		DateOfBirth:     request.DateOfBirth,
		Graduation:      request.Graduation,
		Active:          false,
	}
	if NewGender(request.Gender) == "" {
		return TeacherInfo{}, BadRequest
	}

	if NewGraduation(request.Graduation) == "" {
		return TeacherInfo{}, BadRequest
	}

	//put data list subject
	if len(request.ListSubject) > 0 {
		mapData := map[string]bool{}
		for _, i := range request.ListSubject {
			if NewSubject(i) == "" {
				return TeacherInfo{}, BadRequest
			}
			mapData[i] = true
		}
		teacher.ListSubject = structdata.MapIStringToArray(mapData)
	}

	//put data list method
	if len(request.ListMethod) > 0 {
		mapData := map[string]bool{}
		for _, i := range request.ListMethod {
			if NewMethod(i) == "" {
				return TeacherInfo{}, BadRequest
			}
			mapData[i] = true
		}
		teacher.ListMethod = structdata.MapIStringToArray(mapData)
	}

	//set topic
	if len(request.Topic) > 0 {
		teacher.Topic = request.Topic
	}

	return teacher, nil
}