package admincontroller

import (
	"EasyTutor/data/data"
	"EasyTutor/data/rest/responses"
	"EasyTutor/usecase/adminusecase"
	"EasyTutor/utils/myerror"
	"encoding/json"
	"github.com/beego/beego/v2/server/web"
)

// Operations about admin
type AdminController struct {
	web.Controller
}

// @Title Login
// @Description login teacher
// @Param	body		body 	data.LoginInfo	true		"The object content"
// @Success 200 {string} id
// @Failure 403 body is empty
// @router /login [post]
func (o *AdminController) Login() {
	var ob data.LoginInfo
	err := json.Unmarshal(o.Ctx.Input.RequestBody, &ob)
	if err != nil {
		o.Ctx.Output.SetStatus(400)
		return
	}
	id, err := adminusecase.GetAdminUseCase().Login(ob)
	if myerror.IsError(err) {
		o.Ctx.Output.SetStatus(data.MapErrorCode[err])
		return
	}
	o.Data["json"] = responses.ResponseCommonSingle{
		Data: id,
	}
	o.ServeJSON()
}


// @Title GetPage
// @Description get all objects
// @Param	token		header	string  true	"admin"
// @Param	page_number	query	int	true	"page number"
// @Param	page_length	query	int	true	"page length"
// @Success 200 {object} responses.Teacher
// @Failure 403 is empty
// @router /unactivated-teacher [get]
func (o *AdminController) GetPageTeacher() {
	pageNumber, _ := o.GetInt("page_number", 0)
	pageLength, _ := o.GetInt("page_length", 0)
	obs, total, err := adminusecase.GetAdminUseCase().GetListUnActiveTeacher(pageNumber, pageLength)
	if myerror.IsError(err) {
		o.Ctx.Output.SetStatus(data.MapErrorCode[err])
		return
	}
	o.Data["json"] = responses.ResponseCommonArray{
		Data:       obs,
		TotalCount: total,
	}
	o.ServeJSON()
}

// @Title GetPageRequest
// @Description get all request
// @Param	token		header	string  true	"admin"
// @Param	page_number	query	int	true	"page number"
// @Param	page_length	query	int	true	"page length"
// @Success 200 {object} responses.Request
// @Failure 403 is empty
// @router /unactivated-request [get]
func (o *AdminController) GetPageRequest() {
	pageNumber, _ := o.GetInt("page_number", 0)
	pageLength, _ := o.GetInt("page_length", 0)
	obs, total, err := adminusecase.GetAdminUseCase().GetListUnActiveRequest(pageNumber, pageLength)
	if myerror.IsError(err) {
		o.Ctx.Output.SetStatus(data.MapErrorCode[err])
		return
	}
	o.Data["json"] = responses.ResponseCommonArray{
		Data:       obs,
		TotalCount: total,
	}
	o.ServeJSON()
}

// @Title GetComment
// @Description get all comment
// @Param	token		header	string  true	"admin"
// @Success 200 {object} responses.TeacherComment
// @Failure 403 is empty
// @router /unactivated-comment [get]
func (o *AdminController) GetComment() {
	obs, err := adminusecase.GetAdminUseCase().GetAllUnActiveComment()
	if myerror.IsError(err) {
		o.Ctx.Output.SetStatus(data.MapErrorCode[err])
		return
	}
	o.Data["json"] = responses.ResponseCommonArray{
		Data:       obs,
		TotalCount: len(obs),
	}
	o.ServeJSON()
}

// @Title ActiveUser
// @Description get all objects
// @Param	token		header	string  true	"admin"
// @Param	teacher_id	path	string	true	"path"
// @Success 200 {object} responses.Teacher
// @Failure 403 is empty
// @router /unactivated-teacher/:teacher_id [put]
func (o *AdminController) ActiveUser() {
	teacherID := o.GetString(":teacher_id")
	err := adminusecase.GetAdminUseCase().ValidateTeacher(teacherID)
	if myerror.IsError(err) {
		o.Ctx.Output.SetStatus(data.MapErrorCode[err])
		return
	}
	o.Data["json"] = responses.ResponseBool{Data: true}
	o.ServeJSON()
}

// @Title ActiveRequest
// @Description get all objects
// @Param	token		header	string  true	"admin"
// @Param	request_id	path	string	true	"path"
// @Success 200 {object} responses.Request
// @Failure 403 is empty
// @router /unactivated-request/:request_id [put]
func (o *AdminController) ActiveRequest() {
	requestID := o.GetString(":request_id")
	err := adminusecase.GetAdminUseCase().ValidateRequest(requestID)
	if myerror.IsError(err) {
		o.Ctx.Output.SetStatus(data.MapErrorCode[err])
		return
	}
	o.Data["json"] = responses.ResponseBool{Data: true}
	o.ServeJSON()
}

// @Title ActiveComment
// @Description get all objects
// @Param	token		header	string  true	"admin"
// @Param	comment_id	path	string	true	"path"
// @Success 200 {object} responses.Comment
// @Failure 403 is empty
// @router /unactivated-comment/:comment_id [put]
func (o *AdminController) ActiveComment() {
	commentID := o.GetString(":comment_id")
	err := adminusecase.GetAdminUseCase().ValidateComment(commentID)
	if myerror.IsError(err) {
		o.Ctx.Output.SetStatus(data.MapErrorCode[err])
		return
	}
	o.Data["json"] = responses.ResponseBool{Data: true}
	o.ServeJSON()
}


// @Title UpdateUser
// @Description get all objects
// @Param	token		header	string  true	"admin"
// @Param	teacher_id	path	string	true	"path"
// @Success 200 {object} responses.Teacher
// @Failure 403 is empty
// @router /updating-teacher/:teacher_id [put]
func (o *AdminController) UpdateUser() {
	teacherID := o.GetString(":teacher_id")
	err := adminusecase.GetAdminUseCase().ValidateUpdateTeacher(teacherID)
	if myerror.IsError(err) {
		o.Ctx.Output.SetStatus(data.MapErrorCode[err])
		return
	}
	o.Data["json"] = responses.ResponseBool{Data: true}
	o.ServeJSON()
}