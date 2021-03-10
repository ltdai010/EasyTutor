package commentcontroller

import (
	"EasyTutor/data/data"
	"EasyTutor/data/rest/requests"
	"EasyTutor/data/rest/responses"
	"EasyTutor/usecase/commentusecase"
	"EasyTutor/utils/myerror"
	"encoding/json"
	"github.com/beego/beego/v2/server/web"
)

// Operations about comment
type CommentController struct {
	web.Controller
}


// @Title PostComment
// @Description find object by objectid
// @Param	token			header	string				 true		"token of user"
// @Param	teacher_id		path 	string				 true		"the teacher_id you want to comment"
// @Param	body			body	requests.CommentPost true		"comment body"
// @Success 200 {object} responses.Comment
// @Failure 403 :user_id is empty
// @router /teacher/:teacher_id [post]
func (o *CommentController) PostComment() {
	username := o.Ctx.Input.Header("username")
	teacherID := o.Ctx.Input.Param(":teacher_id")
	if teacherID == "" || username == ""{
		o.Ctx.Output.SetStatus(400)
		return
	}
	comment := requests.CommentPost{}
	ob, err := commentusecase.GetCommentUseCase().CreateOne(username, teacherID, comment)
	if myerror.IsError(err) {
		o.Ctx.Output.SetStatus(data.MapErrorCode[err])
		return
	}
	o.Data["json"] = responses.ResponseCommonSingle{Data: ob}
	o.ServeJSON()
}

// @Title Get
// @Description find object by objectid
// @Param	comment_id		path 	string	true		"the offerID you want to get"
// @Success 200 {object} responses.Comment
// @Failure 403 :comment_id is empty
// @router /:comment_id [get]
func (o *CommentController) Get() {
	commentID := o.Ctx.Input.Param(":comment_id")
	if commentID == "" {
		o.Ctx.Output.SetStatus(400)
		return
	}
	ob, err := commentusecase.GetCommentUseCase().GetOne(commentID)
	if myerror.IsError(err) {
		o.Ctx.Output.SetStatus(data.MapErrorCode[err])
		return
	} else {
		o.Data["json"] = responses.ResponseCommonSingle{Data: ob}
		o.ServeJSON()
	}
}



// @Title Put
// @Description update the object
// @Param	token		header 	string				true		"The token"
// @Param	comment_id	path	string				true		"the comment"
// @Param	body		body 	requests.CommentPut	true		"The body"
// @Success 200 {string} success
// @Failure 403  is empty
// @router /:comment_id [put]
func (o *CommentController) Put() {
	username := o.Ctx.Input.Header("username")
	commentID := o.GetString(":comment_id")
	body := requests.CommentPut{}

	err := json.Unmarshal(o.Ctx.Input.RequestBody, &body)
	if err != nil {
		o.Ctx.Output.SetStatus(400)
		return
	}
	err = commentusecase.GetCommentUseCase().UpdateOne(username, commentID, body)
	if myerror.IsError(err) {
		o.Ctx.Output.SetStatus(data.MapErrorCode[err])
		return
	}
	o.Data["json"] = responses.ResponseBool{
		Data:       true,
	}
	o.ServeJSON()
}

// @Title Delete
// @Description delete the object
// @Param	token		header 	string				true		"The token"
// @Param	comment_id	path	string				true		"the request_id"
// @Success 200 {string} success
// @Failure 403  is empty
// @router /:comment_id [delete]
func (o *CommentController) Delete() {
	username := o.Ctx.Input.Header("username")
	requestID := o.GetString(":comment_id")

	if username == "" || requestID == "" {
		o.Ctx.Output.SetStatus(400)
		return
	}

	err := commentusecase.GetCommentUseCase().RemoveOne(username, requestID)
	if myerror.IsError(err) {
		o.Ctx.Output.SetStatus(data.MapErrorCode[err])
		return
	}
	o.Data["json"] = responses.ResponseBool{
		Data:       true,
	}
	o.ServeJSON()
}
