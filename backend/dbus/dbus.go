package dbus

import (
	log "github.com/Sirupsen/logrus"
	"github.com/godbus/dbus"
	"github.com/godbus/dbus/introspect"
	"strings"
	"reflect"
	"strconv"
)

const (
	busName             = "unix:path=/var/run/dbus/comodo_socket"
	outDirectionArg 	= "out"
)

var (
	DbusComodoInstance *DSbusComodoStruct
)

var returnType = map[string]reflect.Type{
	"y": reflect.TypeOf(byte(0)),
	"b": reflect.TypeOf(false),
	"n": reflect.TypeOf(int16(0)),
	"q": reflect.TypeOf(uint16(0)),
	"i": reflect.TypeOf(int32(0)),
	"u": reflect.TypeOf(uint32(0)),
	"x": reflect.TypeOf(int64(0)),
	"t": reflect.TypeOf(uint64(0)),
	"d": reflect.TypeOf(float64(0)),
	"s": reflect.TypeOf(""),
	"g": reflect.TypeOf(dbus.Signature{}),
	"o": reflect.TypeOf(dbus.ObjectPath("")),
	"v": reflect.TypeOf(dbus.Variant{}),
	"h": reflect.TypeOf(dbus.UnixFD(0)),

	//TODO: ??? interfacesType = reflect.New(reflect.TypeOf([]interface{})).Interface()
	//TODO:???? "h": reflect.New(reflect.TypeOf(dbus.UnixFDIndex(0))).Interface(),
	"as": reflect.TypeOf([]string{}),
	//a: //array
	//r : struct
	//e : dict
}

var parameterType = map[string]func(string, *[]interface{}){
	"y": func(value string, lI *[]interface{}){*lI = append(*lI, []byte(value))},
	"b": func(value string, lI *[]interface{}){b, _:= strconv.ParseBool(value); *lI = append(*lI, b)},
	"n": func(value string, lI *[]interface{}){b, _:= strconv.ParseInt(value, 10, 16); *lI = append(*lI, b)},
	"q": func(value string, lI *[]interface{}){b, _:= strconv.ParseUint(value, 10, 16); *lI = append(*lI, b)},
	"i": func(value string, lI *[]interface{}){b, _:= strconv.ParseInt(value, 10, 32); *lI = append(*lI, b)},
	"u": func(value string, lI *[]interface{}){b, _:= strconv.ParseUint(value, 10, 32); *lI = append(*lI, b)},
	"x": func(value string, lI *[]interface{}){b, _:= strconv.ParseInt(value, 10, 64); *lI = append(*lI, b)},
	"t": func(value string, lI *[]interface{}){b, _:= strconv.ParseUint(value, 10, 64); *lI = append(*lI, b)},
	"d": func(value string, lI *[]interface{}){b, _:= strconv.ParseFloat(value, 64); *lI = append(*lI, b)},
	"s": func(value string, lI *[]interface{}){*lI = append(*lI, value)},

	//TODO: next type please
	//"g": signatureType,
	//"o": objectPathType,
	//"v": variantType,
	//"h": unixFDIndexType,
	//"as": arrayString,
}

var parameterType2 = map[string]func(interface{}, *[]interface{}){
	"y": func(value interface{}, lI *[]interface{}){*lI = append(*lI, []byte(value.([]byte)))},
	"b": func(value interface{}, lI *[]interface{}){*lI = append(*lI, value.(bool))},
	"n": func(value interface{}, lI *[]interface{}){*lI = append(*lI, int(value.(float64)))},
	"q": func(value interface{}, lI *[]interface{}){*lI = append(*lI, uint16(value.(float64)))},
	"i": func(value interface{}, lI *[]interface{}){*lI = append(*lI, int32(value.(float64)))},
	"u": func(value interface{}, lI *[]interface{}){*lI = append(*lI, uint32(value.(float64)))},
	"x": func(value interface{}, lI *[]interface{}){*lI = append(*lI, int64(value.(float64)))},
	"t": func(value interface{}, lI *[]interface{}){*lI = append(*lI, uint64(value.(float64)))},
	"d": func(value interface{}, lI *[]interface{}){*lI = append(*lI, value)},
	"s": func(value interface{}, lI *[]interface{}){*lI = append(*lI, value)},

	//TODO: next type please
	//"g": signatureType,
	//"o": objectPathType,
	//"v": variantType,
	//"h": unixFDIndexType,
	//"as": arrayString,
}

func init(){
	NewDbusComodoStruct()
}

type DSbusComodoStruct struct {
	c *dbus.Conn
	o dbus.BusObject
}

func NewDbusComodoStruct() (*DSbusComodoStruct, error){
	var err error
	if DbusComodoInstance == nil{
		DbusComodoInstance = new(DSbusComodoStruct)
		err = DbusComodoInstance.startSystemDBus()
	}
	return DbusComodoInstance, err
}

func (d *DSbusComodoStruct) startSystemDBus() (err error) {
	d.c, err = dbus.Dial(busName)
	if checkErrors(nil, err) {return}

	err = d.c.Auth(nil)
	if checkErrors(nil, err) {return}

	err = d.c.Hello()
	if checkErrors(nil, err) {return}
	return
}

func (d *DSbusComodoStruct) MethodInterfaceReturnInterface(object string, path string, interfaceName string, methodName string, args []map[string]interface{})(interface{}, error) {
	node, errIntrospect := introspect.Call(d.c.Object(object, dbus.ObjectPath(path)))
	if checkErrors(nil, errIntrospect) {return nil, errIntrospect}

	resTmp := getRetType(*node, interfaceName, methodName)

	errCall := d.c.Object(object, dbus.ObjectPath(path)).Call(strings.Join([]string{interfaceName, methodName}, "."), 0,  getParamValueStringInterface(args)...).Store(resTmp)
	if checkErrors(nil, errCall) {return nil, errCall}

	return resTmp, nil
}

func getParamValueStringInterface(args []map[string]interface{}) (listInterface []interface{}){
	if args != nil{
		for _, v := range args{
			//"Args":[{"as":["uno", "dos"]}, {"a{us}":[{1:"uno"},{2:"dos"}]}],
			for k1, v1 := range v {
				parameterType2[k1](v1, &listInterface)
			}
		}
	}
	return
}

func getRetType(node introspect.Node, interfaceName string, methodName string)(res interface{}){
	for i := range node.Interfaces {
		if  strings.Compare(node.Interfaces[i].Name, interfaceName) == 0 {
			for j := range node.Interfaces[i].Methods {
				if strings.Compare(node.Interfaces[i].Methods[j].Name, methodName) == 0 {
					for k := range node.Interfaces[i].Methods[j].Args {
						if strings.Compare(node.Interfaces[i].Methods[j].Args[k].Direction, outDirectionArg) == 0 {
							t := node.Interfaces[i].Methods[j].Args[k].Type
							//TODO: be careful with TypeFor2 it is in godbus/dbus dependency
							res = reflect.New(dbus.TypeFor2(t)).Interface()
							return
						}
					}
				}
			}
		}
	}
	return
}

/////////////////////////////////////
// utils TODO: move to other site
func checkErrors(err error, currentError error)(b bool){
	if currentError != nil {
		//TODO: remove if
		if err == nil {
			log.Errorf("", currentError)
		}else{
			log.Errorf(err.Error(), currentError)
		}
		b = true
	}
	return
}
