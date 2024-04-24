package utils

import (
	"encoding/json"
	"errors"
	"net/url"
	"reflect"
	"strconv"
)

func StructToUrlValues(transStruct interface{}) (url.Values, error) {
	var err error
	parameters := make(url.Values)
	if transStruct != nil {
		reqValue := reflect.ValueOf(transStruct)
		if reqValue.Kind() == reflect.Pointer {
			reqValue = reqValue.Elem()
		}
		if reqValue.Kind() != reflect.Struct {
			err = errors.New("要求req为struct")
			return nil, err
		}
		//for i := 0; i < reqValue.NumField(); i++ {
		//	if reqValue.Field(i).Kind() == reflect.Slice {
		//		fieldSlice, ok := reqValue.Field(i).Interface().([]string)
		//		if !ok {
		//			err = errors.New("字段不是string类型")
		//			return nil, err
		//		}
		//		for _, fieldValue := range fieldSlice {
		//			parameters.Add(reqValue.Type().Field(i).Tag.Get("json"), fieldValue)
		//		}
		//	}
		//	fieldString, ok := reqValue.Field(i).Interface().(string)
		//	if !ok {
		//		err = errors.New("字段不是string类型")
		//		return nil, err
		//	}
		//	parameters.Set(reqValue.Type().Field(i).Tag.Get("json"), fieldString)
		//}
		data, err := json.Marshal(transStruct)
		if err != nil {
			err = errors.New("marshal结构体失败")
			return nil, err
		}
		structMap := make(map[string]interface{})
		err = json.Unmarshal(data, &structMap)
		if err != nil {
			err = errors.New("unMarshal结构体失败")
			return nil, err
		}
		for k, v := range structMap {
			switch v.(type) {
			case string:
				parameters.Set(k, v.(string))
			case []interface{}:
				for _, valueInterface := range v.([]interface{}) {
					value, ok := valueInterface.(string)
					if !ok {
						err = errors.New("要求结构体切片字段为[]string类型")
						return nil, err
					}
					parameters.Add(k, value)
				}
			default:
				err = errors.New("要求结构体字段为string类型")
				return nil, err
			}
		}
	}
	return parameters, nil
}

func StructToMap(transStruct interface{}) (map[string]string, error) {
	var stringMap map[string]string = make(map[string]string)
	reqValue := reflect.ValueOf(transStruct)
	if reqValue.Kind() == reflect.Pointer {
		reqValue = reqValue.Elem()
	}
	for i := 0; i < reqValue.NumField(); i++ {
		switch reqValue.Field(i).Kind() {
		case reflect.Map:
			fieldMap, ok := reqValue.Field(i).Interface().(map[string]string)
			if !ok {
				err := errors.New("字段不是map[string]string类型")
				return nil, err
			}
			for k, v := range fieldMap {
				k = reqValue.Type().Field(i).Tag.Get("json") + "." + k
				stringMap[k] = v
			}
		case reflect.String:
			value := reqValue.Field(i).String()
			keyName := reqValue.Type().Field(i).Tag.Get("json")
			stringMap[keyName] = value
		case reflect.Bool:
			value := reqValue.Field(i).Bool()
			keyName := reqValue.Type().Field(i).Tag.Get("json")
			stringMap[keyName] = strconv.FormatBool(value)
		case reflect.Float64:
			value := reqValue.Field(i).Float()
			keyName := reqValue.Type().Field(i).Tag.Get("json")
			stringMap[keyName] = strconv.FormatFloat(value, 'f', -1, 64)
		default:
			//TODO
		}
	}
	return stringMap, nil
}
