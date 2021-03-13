package searchcontroller

import (
	"EasyTutor/data/data"
	"EasyTutor/data/rest/responses"
	"EasyTutor/usecase/searchusecase"
	"EasyTutor/utils/myerror"
	"github.com/beego/beego/v2/server/web"
)

// Operations about search
type SearchController struct {
	web.Controller
}

// @Title SearchTeacher
// @Description get all objects
// @Param	page_number	query	int	true	"page number"
// @Param	page_length	query	int	true	"page length"
// @Param	key			query	string	false	"key"
// @Param	gender		query	string  false	"gender"
// @Param	location	query	int	false	"location"
// @Param	graduation	query	string	false	"graduation"
// @Param	subject		query	string	false	"subject"
// @Param	method		query	string	false	"method"
// @Success 200 {object} responses.Teacher
// @Failure 403 is empty
// @router /teacher [get]
func (o *SearchController) SearchTeacher() {
	pageNumber, _ := o.GetInt("page_number", 0)
	pageLength, _ := o.GetInt("page_length", 0)
	key := o.GetString("key", "")
	gender := o.GetString("gender", "")
	location, _ := o.GetInt("location", 0)
	graduation := o.GetString("graduation", "")
	subject := o.GetString("subject", "")
	method := o.GetString("method", "")
	obs, err := searchusecase.GetSearchUseCase().SearchTeacher(key, pageNumber, pageLength,
		location, graduation, subject, gender, method)
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

// @Title SearchRequest
// @Description get all request
// @Param	page_number	query	int	true	"page number"
// @Param	page_length	query	int	true	"page length"
// @Param	key			query	string	false	"key"
// @Param	gender		query	string  false	"gender"
// @Param	location	query	int	false	"location"
// @Param	subject		query	string	false	"subject"
// @Param	method		query	string	false	"method"
// @Success 200 {object} responses.Teacher
// @Failure 403 is empty
// @router /request [get]
func (o *SearchController) SearchRequest() {
	pageNumber, _ := o.GetInt("page_number", 0)
	pageLength, _ := o.GetInt("page_length", 0)
	key := o.GetString("key", "")
	gender := o.GetString("gender", "")
	location, _ := o.GetInt("location", 0)
	subject := o.GetString("subject", "")
	method := o.GetString("method", "")
	obs, err := searchusecase.GetSearchUseCase().SearchRequest(key, pageNumber, pageLength,
		location, subject, gender, method)
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
