package dbus

import (
	"github.com/jpgilaberte/dbus-retail/backend"
	"github.com/Sirupsen/logrus"
)

func Call(message *backend.Message)(interface{}, error){
	i, _ := NewDbusComodoStruct()
	a, _ := i.MethodInterfaceReturnInterface(message.Object, message.Path, message.InterfaceName, message.MethodName, message.Args)
	logrus.Infof( "%v ------- ", a)
	return nil, nil
}
