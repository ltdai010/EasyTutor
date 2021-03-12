package requestusecase

import (
	"EasyTutor/data/data"
	"EasyTutor/data/rest/responses"
	"EasyTutor/models"
)

func (t *requestHandler) GetOne(id string) (*responses.Request, error) {
	request := &models.Request{}
	request.ID = id
	err := request.Get()
	if err != nil {
		return nil, data.NotExisted
	}
	return &responses.Request{
		ID:      id,
		Request: request.Request,
	}, data.Success
}

func (t *requestHandler) GetPage(pageNumber int, pageSize int) ([]*responses.Request, int, error) {
	request := &models.Request{}
	listRequest, total, err := request.GetPage(pageNumber - 1, pageSize)
	if err != nil {
		return nil, 0, data.NotMore
	}
	return listRequest, total, data.Success
}

func (t *requestHandler) FindAvailableTeacher(requestID string) ([]*responses.Teacher, error) {
	teacher := &models.Teacher{}
	request := &models.Request{}

	request.ID = requestID
	err := request.Get()
	if err != nil {
		return nil, data.NotExisted
	}

	result, err := teacher.Search("", request.Location, 0, 0,
		data.NewGraduation(request.Graduation), data.NewSubject(request.Subject), data.NewGender(request.Gender),
		data.NewMethod(request.Method))
	if err != nil {
		return nil, data.NotMore
	}
	return filterSchedule(result, request.Schedule)
}


func filterSchedule(teachers []*responses.TeacherSearch, schedule data.Schedule) ([]*responses.Teacher, error) {
	res := []*responses.Teacher{}
	for _, teacher := range teachers {
		if checkSchedule(schedule, teacher.Schedule) {
			res = append(res, &responses.Teacher{
				Username:    teacher.Username,
				TeacherInfo: teacher.TeacherInfo,
				Schedule:    teacher.Schedule,
			})
		}
	}

	return res, data.Success
}

func checkSchedule(desSchedule, srcSchedule data.Schedule) bool {
	if srcSchedule.MonMorning && !desSchedule.MonMorning {
		return false
	}
	if srcSchedule.MonAfternoon && !desSchedule.MonAfternoon {
		return false
	}
	if srcSchedule.MonNight && !desSchedule.MonNight {
		return false
	}
	if srcSchedule.TueMorning && !desSchedule.TueMorning {
		return false
	}
	if srcSchedule.TueAfternoon && !desSchedule.TueAfternoon {
		return false
	}
	if srcSchedule.TueNight && !desSchedule.TueNight {
		return false
	}
	if srcSchedule.WedMorning && !desSchedule.WedMorning {
		return false
	}
	if srcSchedule.WedAfternoon && !desSchedule.WedAfternoon {
		return false
	}
	if srcSchedule.WedNight && !desSchedule.WedNight {
		return false
	}
	if srcSchedule.ThuMorning && !desSchedule.ThuMorning {
		return false
	}
	if srcSchedule.ThuAfternoon && !desSchedule.ThuAfternoon {
		return false
	}
	if srcSchedule.ThuNight && !desSchedule.ThuNight {
		return false
	}
	if srcSchedule.FriMorning && !desSchedule.FriMorning {
		return false
	}
	if srcSchedule.FriAfternoon && !desSchedule.FriAfternoon {
		return false
	}
	if srcSchedule.FriNight && !desSchedule.FriNight {
		return false
	}
	if srcSchedule.SatMorning && !desSchedule.SatMorning {
		return false
	}
	if srcSchedule.SatAfternoon && !desSchedule.SatAfternoon {
		return false
	}
	if srcSchedule.SatNight && !desSchedule.SatNight {
		return false
	}
	if srcSchedule.SunMorning && !desSchedule.SunMorning {
		return false
	}
	if srcSchedule.SatAfternoon && !desSchedule.SatAfternoon {
		return false
	}
	if srcSchedule.SunNight && !desSchedule.SunNight {
		return false
	}
	return true
}