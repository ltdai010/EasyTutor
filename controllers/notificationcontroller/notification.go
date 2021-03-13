package notificationcontroller

import (
	"EasyTutor/data/data"
	"EasyTutor/usecase/notificationusecase"
	"EasyTutor/utils/myerror"
	"github.com/beego/beego/v2/server/web"
	"log"
	"net/http"
	"github.com/gorilla/websocket"
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
func (w *NotificationController) JoinTeacher() {
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