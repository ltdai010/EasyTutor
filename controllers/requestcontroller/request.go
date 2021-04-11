package requestcontroller

import (
	"EasyTutor/data/data"
	"EasyTutor/data/rest/requests"
	"EasyTutor/data/rest/responses"
	"EasyTutor/usecase/offerusecase"
	"EasyTutor/usecase/requestusecase"
	"EasyTutor/utils/myerror"
	"encoding/json"
	"github.com/beego/beego/v2/server/web"
)

// Operations about request
type RequestController struct {
	web.Controller
}


// @Title Post
// @Description create teacher
// @Param	token		header	string				true		"token"
// @Param	request_id	path	string				true		"request"
// @Param	body		body 	requests.OfferPost	true		"The object content"
// @Success 200 {string} id
// @Failure 403 body is empty
// @router /:request_id/offer [post]
func (o *RequestController) Post() {
	username := o.Ctx.Input.Header("teacher_id")
	requestID := o.Ctx.Input.Param(":request_id")
	var ob requests.OfferPost
	err := json.Unmarshal(o.Ctx.Input.RequestBody, &ob)
	if err != nil {
		o.Ctx.Output.SetStatus(400)
		return
	}
	id, err := offerusecase.GetOfferUseCase().CreateOne(username, requestID, ob)
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
// @Param	request_id	path	string	true	"request id"
// @Success 200 {object} responses.Offer
// @Failure 403 is empty
// @router /:request_id/offer [get]
func (o *RequestController) GetOffer() {
	requestID := o.GetString(":request_id")
	obs, err := offerusecase.GetOfferUseCase().GetOfferOfRequest(requestID)
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

// @Title GetAvailableTeacher
// @Description get all objects
// @Param	request_id	path	string	true	"request id"
// @Success 200 {object} responses.TeacherSearch
// @Failure 403 is empty
// @router /:request_id/available-teacher [get]
func (o *RequestController) GetAvailableTeacher() {
	requestID := o.GetString(":request_id")
	obs, err := requestusecase.GetRequestUseCase().FindAvailableTeacher(requestID)
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

// @Title PostRequest
// @Description create teacher
// @Param	token		header	string					true		"string"
// @Param	body		body 	requests.RequestPost	true		"The object content"
// @Success 200 {string} id
// @Failure 403 body is empty
// @router / [post]
func (o *RequestController) PostRequest() {
	username := o.Ctx.Input.Header("username")
	var ob requests.RequestPost
	err := json.Unmarshal(o.Ctx.Input.RequestBody, &ob)
	if err != nil {
		o.Ctx.Output.SetStatus(400)
		return
	}
	id, err := requestusecase.GetRequestUseCase().CreateOne(username, ob)
	if myerror.IsError(err) {
		o.Ctx.Output.SetStatus(data.MapErrorCode[err])
		return
	}
	o.Data["json"] = responses.ResponseCommonSingle{
		Data: id,
	}
	o.ServeJSON()
}

// @Title AcceptOffer
// @Description create teacher
// @Param	token		header	string	true		"string"
// @Param	offer_id	path 	string	true		"The offer id"
// @Success 200 {string} id
// @Failure 403 body is empty
// @router /accepted-offer/:offer_id [put]
func (o *RequestController) AcceptOffer() {
	username := o.Ctx.Input.Header("username")
	offerID := o.GetString(":offer_id")
	err := requestusecase.GetRequestUseCase().AcceptOffer(username, offerID)
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
// @Description create teacher
// @Param	page_number	query	int	true	"page number"
// @Param	page_length	query	int	true	"page length"
// @Success 200 {string} id
// @Failure 403 body is empty
// @router / [get]
func (o *RequestController) GetPage() {
	pageNumber, _ := o.GetInt("page_number", 0)
	pageLength, _ := o.GetInt("page_length", 0)
	obs, total, err := requestusecase.GetRequestUseCase().GetPageActive(pageNumber, pageLength)
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
// @Param	request_id	path	string				true		"the request_id"
// @Param	body		body 	requests.RequestPut	true		"The body"
// @Success 200 {string} success
// @Failure 403  is empty
// @router /:request_id [put]
func (o *RequestController) Put() {
	username := o.Ctx.Input.Header("username")
	requestID := o.GetString(":request_id")
	body := requests.RequestPut{}

	err := json.Unmarshal(o.Ctx.Input.RequestBody, &body)
	if err != nil {
		o.Ctx.Output.SetStatus(400)
		return
	}
	err = requestusecase.GetRequestUseCase().UpdateOne(username, requestID, body)
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
// @Description open the request
// @Param	token		header 	string				true		"The token"
// @Param	request_id	path	string				true		"the request_id"
// @Success 200 {string} success
// @Failure 403  is empty
// @router /:request_id/open [put]
func (o *RequestController) Open() {
	username := o.Ctx.Input.Header("username")
	requestID := o.GetString(":request_id")
	body := requests.RequestPut{}

	err := json.Unmarshal(o.Ctx.Input.RequestBody, &body)
	if err != nil {
		o.Ctx.Output.SetStatus(400)
		return
	}
	err = requestusecase.GetRequestUseCase().OpenRequest(username, requestID)
	if myerror.IsError(err) {
		o.Ctx.Output.SetStatus(data.MapErrorCode[err])
		return
	}
	o.Data["json"] = responses.ResponseBool{
		Data:       true,
	}
	o.ServeJSON()
}

// @Title Close
// @Description update the object
// @Param	token		header 	string				true		"The token"
// @Param	request_id	path	string				true		"the request_id"
// @Success 200 {string} success
// @Failure 403  is empty
// @router /:request_id/close [put]
func (o *RequestController) Close() {
	username := o.Ctx.Input.Header("username")
	requestID := o.GetString(":request_id")
	body := requests.RequestPut{}

	err := json.Unmarshal(o.Ctx.Input.RequestBody, &body)
	if err != nil {
		o.Ctx.Output.SetStatus(400)
		return
	}
	err = requestusecase.GetRequestUseCase().CloseRequest(username, requestID)
	if myerror.IsError(err) {
		o.Ctx.Output.SetStatus(data.MapErrorCode[err])
		return
	}
	o.Data["json"] = responses.ResponseBool{
		Data:       true,
	}
	o.ServeJSON()
}

// @Title Get
// @Description update the object
// @Param	token		header 	string		true		"The token"
// @Param	request_id	path	string		true		"the request_id"
// @Success 200 {string} success
// @Failure 403  is empty
// @router /:request_id [get]
func (o *RequestController) Get() {
	requestID := o.GetString(":request_id")

	ob, err := requestusecase.GetRequestUseCase().GetOne(requestID)
	if myerror.IsError(err) {
		o.Ctx.Output.SetStatus(data.MapErrorCode[err])
		return
	}
	o.Data["json"] = responses.ResponseCommonSingle{
		Data:       ob,
	}
	o.ServeJSON()
}

// @Title Get
// @Description get user's requests
// @Param	token		header 	string		true		"The token"
// @Success 200 {string} success
// @Failure 403  is empty
// @router /user/request [get]
func (o *RequestController) GetUserRequest() {
	username := o.Ctx.Input.Header("username")

	ob, err := requestusecase.GetRequestUseCase().GetAllRequestOfUser(username)
	if myerror.IsError(err) {
		o.Ctx.Output.SetStatus(data.MapErrorCode[err])
		return
	}
	o.Data["json"] = responses.ResponseCommonSingle{
		Data:       ob,
	}
	o.ServeJSON()
}

// @Title DeleteOffer
// @Description create teacher
// @Param	token		header	string	true		"string"
// @Param	offer_id	path 	string	true		"The offer id"
// @Success 200 {string} id
// @Failure 403 body is empty
// @router /offer/:offer_id [delete]
func (o *RequestController) DeleteOffer() {
	username := o.Ctx.Input.Header("username")
	offerID := o.GetString(":offer_id")
	err := requestusecase.GetRequestUseCase().DeleteOffer(username, offerID)
	if myerror.IsError(err) {
		o.Ctx.Output.SetStatus(data.MapErrorCode[err])
		return
	}
	o.Data["json"] = responses.ResponseBool{
		Data: true,
	}
	o.ServeJSON()
}

// @Title Delete
// @Description delete the object
// @Param	token		header 	string				true		"The token"
// @Param	request_id	path	string				true		"the request_id"
// @Success 200 {string} success
// @Failure 403  is empty
// @router /:request_id [delete]
func (o *RequestController) Delete() {
	username := o.Ctx.Input.Header("username")
	requestID := o.GetString(":request_id")

	if username == "" || requestID == "" {
		o.Ctx.Output.SetStatus(400)
		return
	}

	err := requestusecase.GetRequestUseCase().RemoveOne(username, requestID)
	if myerror.IsError(err) {
		o.Ctx.Output.SetStatus(data.MapErrorCode[err])
		return
	}
	o.Data["json"] = responses.ResponseBool{
		Data:       true,
	}
	o.ServeJSON()
}

