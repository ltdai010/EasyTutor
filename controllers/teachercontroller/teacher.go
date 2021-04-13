package teachercontroller

import (
	"EasyTutor/data/data"
	"EasyTutor/data/rest/requests"
	"EasyTutor/data/rest/responses"
	"EasyTutor/usecase/commentusecase"
	"EasyTutor/usecase/offerusecase"
	"EasyTutor/usecase/teacherusecase"
	"EasyTutor/utils/myerror"
	"encoding/json"
	"github.com/beego/beego/v2/server/web"
)

// Operations about user
type TeacherController struct {
	web.Controller
}

// @Title Post
// @Description create teacher
// @Param	body		body 	requests.TeacherPost	true		"The object content"
// @Success 200 {string} id
// @Failure 403 body is empty
// @router / [post]
func (o *TeacherController) Post() {
	var ob requests.TeacherPost
	err := json.Unmarshal(o.Ctx.Input.RequestBody, &ob)
	if err != nil {
		o.Ctx.Output.SetStatus(400)
		return
	}
	id, err := teacherusecase.GetTeacherUseCase().CreateOne(ob)
	if myerror.IsError(err) {
		o.Ctx.Output.SetStatus(data.MapErrorCode[err])
		return
	}
	o.Data["json"] = responses.ResponseCommonSingle{
		Data: id,
	}
	o.ServeJSON()
}

// @Title ForgotPass
// @Description create teacher
// @Param	teacher_id		query 	string	true		"The object content"
// @Success 200 {string} id
// @Failure 403 body is empty
// @router /forgot-password [post]
func (o *TeacherController) ForgotPass() {
	teacherID := o.GetString("teacher_id")
	err := teacherusecase.GetTeacherUseCase().ForgotPassword(teacherID)
	if myerror.IsError(err) {
		o.Ctx.Output.SetStatus(data.MapErrorCode[err])
		return
	}
	o.Data["json"] = responses.ResponseBool{
		Data: true,
	}
	o.ServeJSON()
}

// @Title ForgotPass
// @Description create teacher
// @Param	body		body 	requests.ResetPass	true		"The object content"
// @Success 200 {string} id
// @Failure 403 body is empty
// @router /reset-pass [post]
func (o *TeacherController) ResetPass() {
	ob := requests.ResetPass{}
	err := json.Unmarshal(o.Ctx.Input.RequestBody, &ob)
	if err != nil {
		o.Ctx.Output.SetStatus(400)
		return
	}
	err = teacherusecase.GetTeacherUseCase().ValidateResetCode(ob)
	if myerror.IsError(err) {
		o.Ctx.Output.SetStatus(data.MapErrorCode[err])
		return
	}
	o.Data["json"] = responses.ResponseBool{
		Data: true,
	}
	o.ServeJSON()
}

// @Title Get
// @Description find object by objectid
// @Param	teacher_id		path 	string	true		"the teacher_id you want to get"
// @Success 200 {object} responses.Teacher
// @Failure 403 :teacher_id is empty
// @router /:teacher_id [get]
func (o *TeacherController) Get() {
	teacherID := o.Ctx.Input.Param(":teacher_id")
	if teacherID == "" {
		o.Ctx.Output.SetStatus(400)
		return
	}
	ob, err := teacherusecase.GetTeacherUseCase().GetOne(teacherID)
	if myerror.IsError(err) {
		o.Ctx.Output.SetStatus(data.MapErrorCode[err])
		return
	} else {
		o.Data["json"] = responses.ResponseCommonSingle{Data: ob}
		o.ServeJSON()
	}
}

// @Title GetTeacherInfo
// @Description find object by teacherID
// @Param	token	header	string	true	"token"
// @Success 200 {object} responses.Teacher
// @Failure 403 :teacher_id is empty
// @router /data/profile/ [get]
func (o *TeacherController) GetTeacherInfo() {
	teacherID := o.Ctx.Input.Header("teacher_id")
	if teacherID == "" {
		o.Ctx.Output.SetStatus(400)
		return
	}
	ob, err := teacherusecase.GetTeacherUseCase().Profile(teacherID)
	if myerror.IsError(err) {
		o.Ctx.Output.SetStatus(data.MapErrorCode[err])
		return
	} else {
		o.Data["json"] = responses.ResponseCommonSingle{Data: ob}
		o.ServeJSON()
	}
}

// @Title GetOffer
// @Description find object by objectid
// @Param	token	header	string	true	"token"
// @Success 200 {object} responses.Offer
// @Failure 403 :teacher_id is empty
// @router /data/my-offer/ [get]
func (o *TeacherController) GetOffer() {
	teacherID := o.Ctx.Input.Header("teacher_id")
	if teacherID == "" {
		o.Ctx.Output.SetStatus(400)
		return
	}
	ob, err := offerusecase.GetOfferUseCase().GetOfferOfTeacher(teacherID)
	if myerror.IsError(err) {
		o.Ctx.Output.SetStatus(data.MapErrorCode[err])
		return
	} else {
		o.Data["json"] = responses.ResponseCommonSingle{Data: ob}
		o.ServeJSON()
	}
}

// @Title Get
// @Description find object by objectid
// @Param	teacher_id		path 	string	true		"the teacher_id you want to get"
// @Success 200 {object} responses.Teacher
// @Failure 403 :teacher_id is empty
// @router /:teacher_id/available-request [get]
func (o *TeacherController) GetAvailableRequest() {
	teacherID := o.Ctx.Input.Param(":teacher_id")
	if teacherID == "" {
		o.Ctx.Output.SetStatus(400)
		return
	}
	ob, err := teacherusecase.GetTeacherUseCase().FindAvailableRequest(teacherID)
	if myerror.IsError(err) {
		o.Ctx.Output.SetStatus(data.MapErrorCode[err])
		return
	} else {
		o.Data["json"] = responses.ResponseCommonSingle{Data: ob}
		o.ServeJSON()
	}
}

// @Title GetPage
// @Description get all objects
// @Param	page_number	query	int	true	"page number"
// @Param	page_length	query	int	true	"page length"
// @Success 200 {object} responses.Teacher
// @Failure 403 is empty
// @router / [get]
func (o *TeacherController) GetPage() {
	pageNumber, _ := o.GetInt("page_number", 0)
	pageLength, _ := o.GetInt("page_length", 0)
	obs, total, err := teacherusecase.GetTeacherUseCase().GetPage(pageNumber, pageLength)
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

// @Title Update
// @Description update the object
// @Param	token		header 	string				true		"The token"
// @Param	body		body 	requests.TeacherPut	true		"The body"
// @Success 200 {string} success
// @Failure 403  is empty
// @router / [put]
func (o *TeacherController) Put() {
	username := o.Ctx.Input.Header("teacher_id")

	body := requests.TeacherPut{}

	err := json.Unmarshal(o.Ctx.Input.RequestBody, &body)
	if err != nil {
		o.Ctx.Output.SetStatus(400)
		return
	}
	err = teacherusecase.GetTeacherUseCase().UpdateOne(username, body)
	if myerror.IsError(err) {
		o.Ctx.Output.SetStatus(data.MapErrorCode[err])
		return
	}
	o.Data["json"] = responses.ResponseBool{
		Data:       true,
	}
	o.ServeJSON()
}

// @Title Update
// @Description update the object
// @Param	token		header 	string				true		"The token"
// @Param	body		body 	data.Schedule		true		"The body"
// @Success 200 {string} success
// @Failure 403  is empty
// @router /schedule [put]
func (o *TeacherController) PutSchedule() {
	username := o.Ctx.Input.Header("teacher_id")

	body := data.Schedule{}

	err := json.Unmarshal(o.Ctx.Input.RequestBody, &body)
	if err != nil {
		o.Ctx.Output.SetStatus(400)
		return
	}
	err = teacherusecase.GetTeacherUseCase().UpdateSchedule(username, body)
	if myerror.IsError(err) {
		o.Ctx.Output.SetStatus(data.MapErrorCode[err])
		return
	}
	o.Data["json"] = responses.ResponseBool{
		Data:       true,
	}
	o.ServeJSON()
}


// @Title GetComment
// @Description find object by objectid
// @Param	teacher_id		path 	string	true		"the teacherID you want to get"
// @Success 200 {object} responses.Comment
// @Failure 403 :teacher_id is empty
// @router /:teacher_id/comment [get]
func (o *TeacherController) GetComment() {
	teacherID := o.Ctx.Input.Param(":teacher_id")
	if teacherID == "" {
		o.Ctx.Output.SetStatus(400)
		return
	}
	ob, err := commentusecase.GetCommentUseCase().GetCommentOfTeacher(teacherID)
	if myerror.IsError(err) {
		o.Ctx.Output.SetStatus(data.MapErrorCode[err])
		return
	}
	o.Data["json"] = responses.ResponseCommonSingle{Data: ob}
	o.ServeJSON()
}

// @Title Login
// @Description login teacher
// @Param	body		body 	data.LoginInfo	true		"The object content"
// @Success 200 {string} id
// @Failure 403 body is empty
// @router /login [post]
func (o *TeacherController) Login() {
	var ob data.LoginInfo
	err := json.Unmarshal(o.Ctx.Input.RequestBody, &ob)
	if err != nil {
		o.Ctx.Output.SetStatus(400)
		return
	}
	id, err := teacherusecase.GetTeacherUseCase().Login(ob)
	if myerror.IsError(err) {
		o.Ctx.Output.SetStatus(data.MapErrorCode[err])
		return
	}
	o.Data["json"] = responses.ResponseCommonSingle{
		Data: id,
	}
	o.ServeJSON()
}