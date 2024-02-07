package utils

import (
	"fmt"
	"reflect"

	"github.com/Eclalang/Ecla/interpreter/eclaType"
)

// GoToEclaType converts a go type to an eclaType.
func GoToEclaType(arg any) eclaType.Type {
	//TODO: Refactor the use of Reflect because it may be slow
	switch reflect.TypeOf(arg).Kind() {
	case reflect.Int:
		return eclaType.Int(arg.(int))
	case reflect.Float64:
		return eclaType.Float(float32(arg.(float64)))
	case reflect.String:
		return eclaType.String(arg.(string))
	case reflect.Bool:
		return eclaType.Bool(arg.(bool))
	case reflect.Int32:
		return eclaType.Char(arg.(rune))
	case reflect.Slice:
		//TODO: Refactor the use of Reflect because it may be slow
		var types []eclaType.Type
		newList := reflect.ValueOf(arg)
		for i := 0; i < newList.Len(); i++ {
			types = append(types, GoToEclaType(reflect.ValueOf(newList.Index(i).Interface()).Interface()))
		}
		return &eclaType.List{Value: types, Typ: fmt.Sprint(reflect.TypeOf(arg))}
	case reflect.Map:
		//TODO: Refactor the use of Reflect because it may be slow
		var keys []eclaType.Type
		var values []eclaType.Type
		newMap := reflect.ValueOf(arg)
		iter := newMap.MapRange() // Get the map iterator
		for iter.Next() {         // Get each key and value from iterator until exhausted
			keys = append(keys, GoToEclaType(iter.Key().Interface()))       // Convert the key to eclaType
			values = append(values, GoToEclaType(iter.Value().Interface())) // Convert the value to eclaType
		}
		return &eclaType.Map{Keys: keys, Values: values, Typ: fmt.Sprint(reflect.TypeOf(arg)), TypKey: reflect.TypeOf(arg).Key().String(), TypVal: reflect.TypeOf(arg).Elem().String()}
	default:
		return eclaType.Null{}
	}
}
