package notificationusecase

import (
	"EasyTutor/data/data"
	"EasyTutor/middleware"
	"EasyTutor/models"
	"EasyTutor/utils/logger"
	"fmt"
	"github.com/gorilla/websocket"
	"time"
)

const (
	writeWait = 10 * time.Second

	pongWait = 60 * time.Second

	pingPeriod = (pongWait * 9)/10

)

type NotificationInterface interface {
	Connect(ws *websocket.Conn) error
}

type notificationHandler struct {}

func GetNotificationHandler() NotificationInterface {
	return &notificationHandler{}
}

func init() {
	go listenMessage(models.GetHub())
}

func listenMessage(hub *models.Hub) {
	logger.Info("[ListenMessage] Start listen message websocket")
	defer func() {
		logger.Info("[StopListenMessage] End websocket listen websocket")
		if r := recover(); r != nil {
			fmt.Println("usecase/notificationusecase/notification_usecase.go:41 ", r)
		}
	}()

	for {
		select {
		case ms := <-hub.GetBcChannel():
			if ms.UserType == "user" {
				wss := hub.GetUserClients(ms.Username)
				wss.Range(func(key, value interface{}) bool {
					ws := key.(*websocket.Conn)
					err := writeJson(ws, ms)
					if err != nil {
						hub.GetErrChannel(ws) <- err
					}
					return true
				})
			} else if ms.UserType == "teacher" {
				wss := hub.GetTeacherClients(ms.Username)
				wss.Range(func(key, value interface{}) bool {
					ws := key.(*websocket.Conn)
					err := writeJson(ws, ms)
					if err != nil {
						hub.GetErrChannel(ws) <- err
					}
					return true
				})
			}
		}
	}
}

func (n *notificationHandler) Connect(ws *websocket.Conn) error {
	ticker := time.NewTicker(pingPeriod)

	userID := ""
	userType := ""
	joinMsg := data.GetInMessage{}

	defer func() {
		ticker.Stop()
		logger.Info("[End Connection] end connection userID = %v userType = %v", userID, userType)
		if r := recover(); r != nil {
			fmt.Println("Recovered in function", r)
		}
	}()

	//read first message
	err := ws.ReadJSON(&joinMsg)
	if err != nil {
		return data.BadRequest
	}

	userID, userType = middleware.ValidateToken(joinMsg.Token)

	if userType == "user" {
		models.GetHub().GetUserClients(userID).Store(ws, true)
		defer models.GetHub().GetUserClients(userID).Delete(ws)
	} else if userType == "teacher" {
		models.GetHub().GetTeacherClients(userID).Store(ws, true)
		defer models.GetHub().GetTeacherClients(userID).Delete(ws)
	} else {
		return data.BadRequest
	}

	logger.Info("[Connect] open connection for user id = %v user type = %v", userID, userType)
	writeJson(ws, data.NewErr(data.Success))

	for {
		select {
		case err = <- models.GetHub().GetErrChannel(ws):
			return data.ErrUnknown
		case <-ticker.C:
			return data.ErrUnknown
		}
	}
}

func write(ws *websocket.Conn, mt int, payload []byte) error {
	err := ws.SetWriteDeadline(time.Now().Add(writeWait))
	if err != nil {
		return err
	}
	return ws.WriteMessage(mt, payload)
}

func writeJson(ws *websocket.Conn, v interface{}) error {
	err := ws.SetWriteDeadline(time.Now().Add(writeWait))
	if err != nil {
		return err
	}
	return ws.WriteJSON(v)
}