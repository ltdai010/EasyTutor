package notificationcontroller

import (
	"EasyTutor/data/data"
	"EasyTutor/data/rest/responses"
	"EasyTutor/usecase/notificationusecase"
	"EasyTutor/utils/myerror"
	"github.com/beego/beego/v2/server/web"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

type NotificationController struct {
	web.Controller
}

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

// @Title Join
// @Description WebSocket, connect only, don't write anything
// @router /
func (w *NotificationController) Join() {
	ws, err := upgrader.Upgrade(w.Ctx.ResponseWriter, w.Ctx.Request, nil)
	if err != nil {
		log.Println(err)
		w.Ctx.Output.SetStatus(200)
		return
	}
	defer ws.Close()

	err = notificationusecase.GetNotificationHandler().Connect(ws)
	if myerror.IsError(err) {
		w.Ctx.Output.SetStatus(data.MapErrorCode[err])
		return
	}

	w.Ctx.Output.SetStatus(200)
}

// @Title Get
// @Description find object by objectid
// @Param	token			header	string	true		"the token"
// @Param	page_number		query 	int		true		"the page number you want to get"
// @Param	page_size		query 	int		true		"the page size you want to get"
// @Success 200 {object} responses.Teacher
// @Failure 403 :teacher_id is empty
// @router /list [get]
func (o *NotificationController) GetNotification() {
	pageNumber, _ := o.GetInt("page_number")
	pageSize, _ := o.GetInt("page_size")
	username := o.Ctx.Input.Header("username")
	userType := o.Ctx.Input.Header("user_type")
	ob, total, err := notificationusecase.GetNotificationHandler().GetListNotification(pageNumber, pageSize, username, userType)
	if myerror.IsError(err) {
		o.Ctx.Output.SetStatus(data.MapErrorCode[err])
		return
	} else {
		o.Data["json"] = responses.ResponseCommonArray{
			Data: ob,
			TotalCount: total,
		}
		o.ServeJSON()
	}
}
