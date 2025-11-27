package jsonobj

import (
	"encoding/json"
	"github.com/oliveagle/jsonpath"
	"github.com/simonalong/gole/util"
	"io"
)

type ObjectOfJson struct {
	dataJson interface{}
}

func LoadReader(ioReader io.Reader) (*ObjectOfJson, error) {
	dataReader, err := io.ReadAll(ioReader)
	if err != nil {
		return nil, err
	}

	var dataJson interface{}
	if err := json.Unmarshal(dataReader, &dataJson); err != nil {
		return nil, err
	}

	return &ObjectOfJson{dataJson: dataJson}, nil
}

func Load(jsonContent string) (*ObjectOfJson, error) {
	var dataJson interface{}
	if err := json.Unmarshal([]byte(jsonContent), &dataJson); err != nil {
		return nil, err
	}

	return &ObjectOfJson{dataJson: dataJson}, nil
}

func (obj *ObjectOfJson) Get(key string) any {
	result, err := jsonpath.JsonPathLookup(obj.dataJson, key)
	if err != nil {
		return nil
	}
	return result
}

func (obj *ObjectOfJson) GetString(key string) string {
	return util.ToString(obj.Get(key))
}

func (obj *ObjectOfJson) GetInt(key string) int {
	return util.ToInt(obj.Get(key))
}

func (obj *ObjectOfJson) GetInt8(key string) int8 {
	return util.ToInt8(obj.Get(key))
}

func (obj *ObjectOfJson) GetInt16(key string) int16 {
	return util.ToInt16(obj.Get(key))
}

func (obj *ObjectOfJson) GetInt32(key string) int32 {
	return util.ToInt32(obj.Get(key))
}

func (obj *ObjectOfJson) GetInt64(key string) int64 {
	return util.ToInt64(obj.Get(key))
}

func (obj *ObjectOfJson) GetUInt(key string) uint {
	return util.ToUInt(obj.Get(key))
}

func (obj *ObjectOfJson) GetUInt8(key string) uint8 {
	return util.ToUInt8(obj.Get(key))
}

func (obj *ObjectOfJson) GetUInt16(key string) uint16 {
	return util.ToUInt16(obj.Get(key))
}

func (obj *ObjectOfJson) GetUInt32(key string) uint32 {
	return util.ToUInt32(obj.Get(key))
}

func (obj *ObjectOfJson) GetUInt64(key string) uint64 {
	return util.ToUInt64(obj.Get(key))
}

func (obj *ObjectOfJson) GetFloat32(key string) float32 {
	return util.ToFloat32(obj.Get(key))
}

func (obj *ObjectOfJson) GetFloat64(key string) float64 {
	return util.ToFloat64(obj.Get(key))
}

func (obj *ObjectOfJson) GetBool(key string) bool {
	return util.ToBool(obj.Get(key))
}

func (obj *ObjectOfJson) GetObject(key string, targetPtrObj any) error {
	data := obj.Get(key)
	_, err := util.DataToEntity(data, targetPtrObj)
	if err != nil {
		return err
	}
	return nil
}
