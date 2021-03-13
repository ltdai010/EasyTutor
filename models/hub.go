package models

import (
	"EasyTutor/data/data"
	"github.com/gorilla/websocket"
	"sync"
)

var (
	doOnce = &sync.Once{}
	hub = &Hub{}
)

type Hub struct {
	//channel to broadcast data to user
	bcChannel    chan data.Notification
	//sync map user clients map[id] websocket connect
	userClients   *sync.Map
	//sync map user teacher map[id] websocket connect
	teacherClients *sync.Map
	//error channel to close ws when caught error
	errChanel    *sync.Map
}

func GetHub() *Hub {
	doOnce.Do(func() {
		hub.Init()
	})
	return hub
}

func (h *Hub) Init() {
	h.bcChannel = make(chan data.Notification, 1000)
	h.userClients = &sync.Map{}
	h.teacherClients = &sync.Map{}
	h.errChanel = &sync.Map{}
}

func (h *Hub) GetBcChannel() chan data.Notification {
	if h.bcChannel == nil {
		h.bcChannel = make(chan data.Notification, 1000)
	}
	return h.bcChannel
}

func (h *Hub) GetUserClients(username string) *sync.Map {
	if h.userClients == nil {
		h.userClients = &sync.Map{}
	}
	ws, ok := h.userClients.Load(username)
	if !ok {
		h.userClients.Store(username, &sync.Map{})
		i, _ := h.userClients.Load(username)
		return i.(*sync.Map)
	}
	return ws.(*sync.Map)
}

func (h *Hub) GetTeacherClients(teacherID string) *sync.Map {
	if h.teacherClients == nil {
		h.teacherClients = &sync.Map{}
	}
	ws, ok := h.teacherClients.Load(teacherID)
	if !ok {
		h.teacherClients.Store(teacherID, &sync.Map{})
		i, _ := h.teacherClients.Load(teacherID)
		return i.(*sync.Map)
	}
	return ws.(*sync.Map)
}

func (h *Hub) GetErrChannel(ws *websocket.Conn) chan error {
	_, ok := h.errChanel.Load(ws)
	if !ok {
		h.errChanel.Store(ws, make(chan error))
	}
	i, ok := h.errChanel.Load(ws)
	return i.(chan error)
}


func (h *Hub) BroadcastMessage(msg data.Notification) {
	h.GetBcChannel() <- msg
	notify := &Notification{}
	notify.Notification = msg
	go notify.Add()
}