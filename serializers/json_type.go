package serializers

import "encoding/json"

var(
)

type JsonType struct {}

func (jst *JsonType) Serialize(msgStruct interface{})([]byte, error){
	msgByte, err := json.Marshal(msgStruct)
	if err != nil{
		return nil, err
	}
	return msgByte, nil
}

func (jst *JsonType) Deserialize(msgByte []byte, msgStruct interface{})(interface{}, error){
	err := json.Unmarshal(msgByte, msgStruct)
	if err != nil{
		return nil, err
	}
	return msgStruct, nil
}


