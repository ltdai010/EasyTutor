package usercontroller

import (
	"EasyTutor/data/data"
	"EasyTutor/data/rest/requests"
	"EasyTutor/data/rest/responses"
	"EasyTutor/usecase/userusecase"
	"EasyTutor/utils/myerror"
	"encoding/json"

	beego "github.com/beego/beego/v2/server/web"
)

// Operations about Users
type UserController struct {
	beego.Controller
}

// @Title Post
// @Description create User
// @Param	body		body 	requests.UserPost	true		"The object content"
// @Success 200 {string} id
// @Failure 403 body is empty
// @router / [post]
func (o *UserController) Post() {
	var ob requests.UserPost
	err := json.Unmarshal(o.Ctx.Input.RequestBody, &ob)
	if err != nil {
		o.Ctx.Output.SetStatus(400)
		return
	}
	id, err := userusecase.GetUserUseCase().CreateOne(ob)
	if myerror.IsError(err) {
		o.Ctx.Output.SetStatus(data.MapErrorCode[err])
		return
	}
	o.Data["json"] = responses.ResponseCommonSingle{
		Data: id,
	}
	o.ServeJSON()
}

// @Title Get
// @Description find object by objectid
// @Param	user_id		path 	string	true		"the user_id you want to get"
// @Success 200 {object} responses.User
// @Failure 403 :user_id is empty
// @router /:user_id [get]
func (o *UserController) Get() {
	UserID := o.Ctx.Input.Param(":user_id")
	if UserID == "" {
		o.Ctx.Output.SetStatus(400)
		return
	}
	ob, err := userusecase.GetUserUseCase().GetOne(UserID)
	if myerror.IsError(err) {
		o.Ctx.Output.SetStatus(data.MapErrorCode[err])
		return
	} else {
		o.Data["json"] = responses.ResponseCommonSingle{Data: ob}
		o.ServeJSON()
	}
}

// @Title ForgotPass
// @Description create teacher
// @Param	username		query 	string	true		"The object content"
// @Success 200 {string} id
// @Failure 403 body is empty
// @router /forgot-password [post]
func (o *UserController) ForgotPass() {
	teacherID := o.GetString("username")
	err := userusecase.GetUserUseCase().ForgotPassword(teacherID)
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
// @Description reset pass
// @Param	body		body 	requests.ResetPass	true		"The object content"
// @Success 200 {string} id
// @Failure 403 body is empty
// @router /reset-pass [post]
func (o *UserController) ResetPass() {
	ob := requests.ResetPass{}
	err := json.Unmarshal(o.Ctx.Input.RequestBody, &ob)
	if err != nil {
		o.Ctx.Output.SetStatus(400)
		return
	}
	err = userusecase.GetUserUseCase().ValidateResetCode(ob)
	if myerror.IsError(err) {
		o.Ctx.Output.SetStatus(data.MapErrorCode[err])
		return
	}
	o.Data["json"] = responses.ResponseBool{
		Data: true,
	}
	o.ServeJSON()
}

// @Title GetPage
// @Description get all objects
// @Param	page_number	query	int	true	"page number"
// @Param	page_length	query	int	true	"page size"
// @Success 200 {object} responses.User
// @Failure 403 is empty
// @router / [get]
func (o *UserController) GetPage() {
	pageNumber, _ := o.GetInt("page_number", 0)
	pageSize, _ := o.GetInt("page_size", 0)
	obs, total, err := userusecase.GetUserUseCase().GetPage(pageNumber, pageSize)
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
// @Param	body		body 	requests.UserPut	true		"The body"
// @Success 200 {string} success
// @Failure 403  is empty
// @router / [put]
func (o *UserController) Put() {
	username := o.Ctx.Input.Header("username")

	body := requests.UserPut{}

	err := json.Unmarshal(o.Ctx.Input.RequestBody, &body)
	if err != nil {
		o.Ctx.Output.SetStatus(400)
		return
	}
	err = userusecase.GetUserUseCase().UpdateOne(username, body)
	if myerror.IsError(err) {
		o.Ctx.Output.SetStatus(data.MapErrorCode[err])
		return
	}
	o.Data["json"] = responses.ResponseBool{
		Data:       true,
	}
	o.ServeJSON()
}

// @Title Post
// @Description login user
// @Param	body		body 	data.LoginInfo	true		"The object content"
// @Success 200 {string} id
// @Failure 403 body is empty
// @router /login [post]
func (o *UserController) Login() {
	var ob data.LoginInfo
	err := json.Unmarshal(o.Ctx.Input.RequestBody, &ob)
	if err != nil {
		o.Ctx.Output.SetStatus(400)
		return
	}
	id, err := userusecase.GetUserUseCase().Login(ob)
	if myerror.IsError(err) {
		o.Ctx.Output.SetStatus(data.MapErrorCode[err])
		return
	}
	o.Data["json"] = responses.ResponseCommonSingle{
		Data: id,
	}
	o.ServeJSON()
}