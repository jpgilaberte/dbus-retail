package agent

import (
	"bytes"
	"github.com/jpgilaberte/cervantes/dbus"
	"io/ioutil"
	"os"
)

type Status int

const(
	statusUnknow Status = iota
	statusStarting
	statusStoped
	statusRunning
	statusRemoved
	statusError
)

type Unit struct{
	Version int
	Name string
	Path string
	Status Status
}

func (u *Unit) startUnit()(error){
	conn := dbus.NewDbusStruct()
	return conn.StartUnit(u.Name)
}

func (u *Unit) CreateLinkUnitFile()(error){
	os.Remove("/etc/systemd/system/example.service")
	return os.Symlink("/tmp/example", "/etc/systemd/system/example.service")
}

func (u *Unit) CreateUnitFile()(error){
	os.Remove("/tmp/example")
	return ioutil.WriteFile("/tmp/example", u.test(), 0755)
}

func (u *Unit) test() []byte{

	return bytes.NewBufferString(
`[Unit]
	Description=example
	Requires=docker.service
	After=docker.service

[Service]
	Restart=always
	ExecStop=/usr/bin/docker stop -t 2 test
	ExecStopPost=/usr/bin/docker rm -f test
	ExecStart=/usr/bin/docker run --env KAFKA_ZOOKEEPER_CONNECT=10.200.1.221:2181 --env KAFKA_ADVERTISED_LISTENERS=PLAINTEXT://0.0.0.0:9092 --name test 122d8c8b14e2`).Bytes()
}