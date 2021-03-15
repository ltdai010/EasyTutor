package data

import (
	"EasyTutor/utils/datastruct"
	"EasyTutor/utils/logger"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type TeacherInfo struct {
	Name            string   `json:"name"`
	Email		    string	 `json:"email"`
	ListSubject     []string `json:"list_subject"`
	Location        int      `json:"location"`
	ListMethod      []string `json:"list_method"`
	Description     string   `json:"description"`
	Achievement     string   `json:"achievement"`
	Topic           []string `json:"topic"`

	BackgroundImage string   `json:"background_image"`
	Avatar          string   `json:"avatar"`
	Gender          string   `json:"gender"`
	DateOfBirth     int64    `json:"date_of_birth"`
	Graduation      string   `json:"graduation"`
}

type Teacher struct {
	LoginInfo
	Active          bool     `json:"active"`
	User		    []string `json:"user"`
	TeacherInfo
	Schedule
}

type TeacherUpdate struct {
	Teacher
	UpdateTime	int64	`json:"update_time"`
}

func SetDataTeacher(request TeacherInfo) (TeacherInfo, error) {
	file, err := ioutil.ReadFile("tinh_tp.json")
	if err != nil {
		logger.Error("[SetDataTeacher] error read json file err = %v", err)
		return TeacherInfo{}, ErrSystem
	}
	var data map[string]interface{}
	err = json.Unmarshal(file, &data)
	if err != nil {
		logger.Error("[SetDataTeacher] error unmarshal json file err = %v", err)
		return TeacherInfo{}, ErrSystem
	}
	if _, ok := data[fmt.Sprintf("%02d", request.Location)]; !ok {
		return TeacherInfo{}, BadRequest
	}
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
		Email:   		 request.Email,
	}
	if NewGender(request.Gender) == "" || request.Gender == All{
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
		teacher.ListSubject = datastruct.MapIStringToArray(mapData)
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
		teacher.ListMethod = datastruct.MapIStringToArray(mapData)
	}

	//set topic
	if len(request.Topic) > 0 {
		teacher.Topic = request.Topic
	}

	return teacher, Success
}
