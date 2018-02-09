package routes

import (
	"github.com/jpgilaberte/dbus-retail/serializers"
	"github.com/jpgilaberte/dbus-retail/backend"
	"github.com/jpgilaberte/dbus-retail/backend/dbus"
)

const()
var(
	serial Serializer = new(serializers.JsonType)
	api backend.Backend = nil
)


type Serializer interface{
	Serialize(interface{})([]byte, error)
	Deserialize([]byte, interface{})(interface{}, error)
}


func Route(msgByte []byte, f func([]byte))([]byte, error){
	msgStruct := new(backend.Message)
	_, err := serial.Deserialize(msgByte, msgStruct)
	if err != nil{
		//log
		return nil, err
	}

	res, err := dbus.Call(msgStruct)
	if err != nil{
		//log
		return nil, err
	}

	res, err = serial.Serialize(res)

	msgStruct.Direction = "out"
	msgStruct.ReturnValue["s"] = res
	return nil, nil
}