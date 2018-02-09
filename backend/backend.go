package backend

type Message struct {
	Direction string
	Object string
	Path string
	InterfaceName string
	MethodName string
	Args []map[string]interface{}
	ReturnValue map[string]interface{}
}

type Backend interface{
	Call(*Message)(interface{}, error)
}
