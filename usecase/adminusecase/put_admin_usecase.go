package adminusecase

import (
	"EasyTutor/data/data"
	"EasyTutor/models"
	"EasyTutor/utils/logger"
)

func (a *adminHandler) ValidateTeacher(teacherID string) error {
	teacher := &models.Teacher{}
	teacher.Username = teacherID
	err := teacher.Get()
	if err != nil {
		return data.NotExisted
	}

	teacher.Active = true
	err = teacher.Update()
	if err != nil {
		logger.Error("[Error validate] Teacher update teacherID = %v error = %v", teacherID, err)
		return data.ErrSystem
	}
	mail := &models.Mail{
		To:      teacher.Email,
		Subject: "[Easy-Tutor] Đăng ký thành công",
		Msg:     "Easy-Tutor Thông báo \n" +
			"Xin chào " + teacher.Name + "\n" +
			"Chúc mừng bạn đã đăng ký thành công tài khoản giáo viên. Xin mời đăng nhập vào hệ thống\n" +
			"Cảm ơn bạn đã tham gia với Easy-Tutor, chúc bạn thành công trong công việc!",
	}
	go mail.Send(teacher.Email)
	return data.Success
}

func (a *adminHandler) ValidateUpdateTeacher(teacherID string) error {
	teacher := &models.Teacher{}
	teacher.Username = teacherID
	err := teacher.GetUpdating()
	if err != nil {
		return data.NotExisted
	}
	err = teacher.Update()
	if err != nil {
		return data.ErrSystem
	}
	return data.Success
}

func (a *adminHandler) ValidateRequest(requestID string) error {
	request := &models.Request{}
	request.ID = requestID
	err := request.Get()
	if err != nil {
		return data.NotExisted
	}
	request.Active = true
	err = request.Update()
	if err != nil {
		return data.ErrSystem
	}

	return data.Success
}