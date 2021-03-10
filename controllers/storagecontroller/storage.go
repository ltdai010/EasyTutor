package storagecontroller

import (
	"EasyTutor/data/data"
	"EasyTutor/data/rest/responses"
	"EasyTutor/usecase/storageusecase"
	"EasyTutor/utils/myerror"
	"github.com/beego/beego/v2/server/web"
)

// Operations about storage
type StorageController struct {
	web.Controller
}

// @Title Post
// @Description create User
// @Param	token		header		string	true		"the token"
// @Param	file		formData 	file	true		"The file content"
// @Success 200 {string} id
// @Failure 403 body is empty
// @router / [post]
func (o *StorageController) Post() {
	file, _, err := o.GetFile("file")
	if err != nil {
		o.Ctx.Output.SetStatus(400)
		return
	}
	path, err := storageusecase.SaveFile(file)
	if myerror.IsError(err) {
		o.Ctx.Output.SetStatus(data.MapErrorCode[err])
		return
	}
	o.Data["json"] = responses.ResponseCommonSingle{
		Data: path,
	}
	o.ServeJSON()
}
