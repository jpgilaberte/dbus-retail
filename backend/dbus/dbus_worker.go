package dbus

import (
	"github.com/jpgilaberte/dbus-retail/backend"
	"fmt"
)

var(
	DbusWorkerInstance *DbusWorker
)

type MsgDbusWorkerResponse struct{
	Err             chan error
	ResponseUnit    *backend.Message
}

type MsgDbusWorkerRequest struct{
	Res         *MsgDbusWorkerResponse
	CurrentUnit *backend.Message
}

type DbusWorker struct{
  Request chan *MsgDbusWorkerRequest
}

func NewDbusWorker() *DbusWorker{
	if DbusWorkerInstance == nil{
		DbusWorkerInstance = new(DbusWorker)
		DbusWorkerInstance.Request = make(chan *MsgDbusWorkerRequest)

		go DbusWorkerInstance.Run()
	}
	return DbusWorkerInstance
}

func (w *DbusWorker) Run(){
	for{
		select {
		case m := <-w.Request:
			fmt.Print(m)
		}
	}
}
