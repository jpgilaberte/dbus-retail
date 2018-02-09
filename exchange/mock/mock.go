package mock

import (
	"context"
	"github.com/Sirupsen/logrus"


	_ "time"
	"time"
)

var (
	Test1 = []byte (`{"Direction":"in",
					"Object":"org.freedesktop.DBus",
					"Path":"/org/freedesktop/DBus",
					"InterfaceName":"org.freedesktop.DBus.Introspectable",
					"MethodName":"Introspect",
					"Args":[],
					"ReturnValue":{"s":""}}`)

	Test2 = []byte (`{"Direction":"in",
					"Object":"org.freedesktop.DBus",
					"Path":"/org/freedesktop/DBus",
					"InterfaceName":"org.freedesktop.DBus",
					"MethodName":"ListNames",
					"Args":[],
					"ReturnValue":{"as":""}}`)

)

func Run(ctx context.Context){
	a := 0
	time.Now()
	for{
		/*
		select{

		case <- ctx.Done():
			return
		}
		*/

		a++
		logrus.Infof("aaaaaaaaaaaa %v", a)
		// go routes.Route(Test1, a, f)
		if 1000  == a{
			time.Sleep(10000000000000)
		}

	}
}