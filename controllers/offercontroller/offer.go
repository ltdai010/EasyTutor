package offercontroller

import (
	"EasyTutor/data/data"
	"EasyTutor/data/rest/requests"
	"EasyTutor/data/rest/responses"
	"EasyTutor/usecase/offerusecase"
	"EasyTutor/utils/myerror"
	"encoding/json"
	"github.com/beego/beego/v2/server/web"
)

// Operations about offer
type OfferController struct {
	web.Controller
}



// @Title Get
// @Description find object by objectid
// @Param	offer_id		path 	string	true		"the offerID you want to get"
// @Success 200 {object} responses.Offer
// @Failure 403 :offer_id is empty
// @router /:offer_id [get]
func (o *OfferController) Get() {
	offerID := o.Ctx.Input.Param(":offer_id")
	if offerID == "" {
		o.Ctx.Output.SetStatus(400)
		return
	}
	ob, err := offerusecase.GetOfferUseCase().GetOne(offerID)
	if myerror.IsError(err) {
		o.Ctx.Output.SetStatus(data.MapErrorCode[err])
		return
	} else {
		o.Data["json"] = responses.ResponseCommonSingle{Data: ob}
		o.ServeJSON()
	}
}

// @Title Update
// @Description update the object
// @Param	token		header 	string				true		"The token"
// @Param	offer_id	path	string				true		"the offer"
// @Param	body		body 	requests.OfferPut	true		"The body"
// @Success 200 {string} success
// @Failure 403  is empty
// @router /:offer_id [put]
func (o *OfferController) Put() {
	username := o.Ctx.Input.Header("teacher_id")
	offerID := o.GetString(":offer_id")
	body := requests.OfferPut{}

	err := json.Unmarshal(o.Ctx.Input.RequestBody, &body)
	if err != nil {
		o.Ctx.Output.SetStatus(400)
		return
	}
	err = offerusecase.GetOfferUseCase().UpdateOne(username, offerID, body)
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
// @Param	offer_id	path	string				true		"the offer_id"
// @Success 200 {string} success
// @Failure 403  is empty
// @router /:offer_id [delete]
func (o *OfferController) Delete() {
	username := o.Ctx.Input.Header("teacher_id")
	offerID := o.GetString(":offer_id")

	if username == "" || offerID == "" {
		o.Ctx.Output.SetStatus(400)
		return
	}

	err := offerusecase.GetOfferUseCase().RemoveOne(username, offerID)
	if myerror.IsError(err) {
		o.Ctx.Output.SetStatus(data.MapErrorCode[err])
		return
	}
	o.Data["json"] = responses.ResponseBool{
		Data:       true,
	}
	o.ServeJSON()
}