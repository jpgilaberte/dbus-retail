package dbus

import (
	log "github.com/Sirupsen/logrus"
	"github.com/godbus/dbus"
	"github.com/pkg/errors"
	"regexp"
)

const (
	service                 = "org.freedesktop.systemd1"
	object                  = "/org/freedesktop/systemd1"
	getMachineId            = "org.freedesktop.DBus.Peer.GetMachineId"
	getUnit                 = "org.freedesktop.systemd1.Manager.GetUnit"
	startUnit               = "org.freedesktop.systemd1.Manager.StartUnit"
	stopUnit                = "org.freedesktop.systemd1.Manager.StopUnit"
	killUnit                = "org.freedesktop.systemd1.Manager.KillUnit"
	getUnitproperties       = "org.freedesktop.systemd1.Service.MainPID"
	stopMode                = "replace"
	killMode                = "main"
	startMode               = "replace"
	killSignall       int32 = 9
	systemdPathRegex        = `^/[a-zA-Z0-9/_\-.]*$`
)

var (
	DbusInstance *DSbusStruct
)

type DSbusStruct struct {
	c *dbus.Conn
	o dbus.BusObject
}

func NewDbusStruct() (*DSbusStruct){
	if DbusInstance == nil{
		DbusInstance = new(DSbusStruct)
		DbusInstance.startSystemDBus()
	}
	return DbusInstance
}

func (d *DSbusStruct) startSystemDBus() (err error) {
	DbusInstance.c, err = dbus.SystemBus()
	if err != nil {
		log.Errorf("dbus.NewDbus %v", err)
	} else {
		DbusInstance.o = DbusInstance.c.Object(service, object)
	}
	return
}

func (d *DSbusStruct) GetMachineId() (r string, err error) {
	err = d.o.Call(getMachineId, 0).Store(&r)
	if err != nil {
		log.Errorf("dbus.GetMachineId %v", err.Error())
	}
	return
}

// replace, fail, isolate, ignore-dependencies, ignore-requirements
func (d *DSbusStruct) StopUnit(unitName string) (err error) {
	var path dbus.ObjectPath
	err = d.o.Call(stopUnit, 0, unitName, stopMode).Store(&path)
	if err != nil {
		log.Infof("dbus.StopUnit %v", err.Error())
	}
	return
}

func (d *DSbusStruct) KillUnit(unitName string) (err error) {
	err = d.o.Call(killUnit, 0, unitName, killMode, killSignall).Store()
	if err != nil {
		log.Errorf("dbus.KillUnit %v", err.Error())
	}
	return
}

// replace, fail, isolate, ignore-dependencies, ignore-requirements
func (d *DSbusStruct) StartUnit(unitName string) (err error) {
	var path dbus.ObjectPath
	if err = d.o.Call(startUnit, 0, unitName, startMode).Store(&path); err != nil {
		log.Infof("dbus.StartUnit %v", err)
	}
	return
}

func (d *DSbusStruct) GetUnit(unitName string) (res string, err error) {
	var unitPath dbus.ObjectPath
	conn := d.o.Call(getUnit, 0, unitName)
	err = conn.Store(&unitPath)
	if err != nil {
		log.Infof("dbus.GetUnit %v", err.Error())
	} else {
		res = string(unitPath)
	}
	return
}

func (d *DSbusStruct) GetUnitPid(unitPath string) (res uint32, err error) {
	log.Debug("dbus.GetUnitPid")
	var variant dbus.Variant
	var validPath = regexp.MustCompile(systemdPathRegex)
	if b := validPath.Match([]byte(unitPath)); b {
		obj := d.c.Object(service, dbus.ObjectPath(unitPath))
		variant, err = obj.GetProperty(getUnitproperties)
		if err != nil {
			log.Infof("dbus.GetUnitPid %v", err.Error())
			return res, err
		} else {
			res = variant.Value().(uint32)
		}
	} else {
		log.Infof("dbus.GetUnitPid - ERROR: Unit path can not compile with regex.")
		err = errors.New("dbus.GetUnitPid - ERROR: Invalid unit path: " + unitPath + ". Can not compile with validation expression.")
	}
	log.Debugf("dbus.GetUnitPid - RES: %v", res)
	return
}
