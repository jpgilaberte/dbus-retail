package dbusretail

import (
	log "github.com/Sirupsen/logrus"
	"time"
)

const()

var (
	Test1 = []byte (`{"Direction":"in",
					"Object":"org.freedesktop.DBus",
					"Path":"/org/freedesktop/DBus",
					"InterfaceName":"org.freedesktop.DBus.Introspectable",
					"MethodName":"Introspect",
					"Args":[],
					"ReturnValue":{"s":""}}`)
	a = make(chan int)
	)

func Run() {
	go aaa(a)
	for b := range a{
		log.Info(b)
	}

}

func aaa(a chan int){
	for {
		time.Sleep(time.Duration(1000000000))
		a <- 1
	}

}