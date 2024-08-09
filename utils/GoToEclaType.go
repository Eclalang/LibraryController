package utils

// If you have tried to optimize this code but failed add the time wasted to the counter below
// Time wasted optimizing : 1h

import (
	"fmt"
	"github.com/Eclalang/Ecla/interpreter/eclaDecl"
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
	case reflect.Struct:
		structFields := GetFields(arg)
		structName := reflect.TypeOf(arg).Name()

		structDefinition := new(eclaDecl.StructDecl)
		structDefinition.Name = structName
		structDefinition.Fields = make(map[string]string)

		struc := eclaType.Struct{
			Fields:     make(map[string]*eclaType.Type),
			Typ:        structName,
			Definition: structDefinition,
		}
		for _, sField := range structFields {
			structDefinition.Order = append(structDefinition.Order, sField.rField.Name)
			structDefinition.Fields[sField.rField.Name] = ReflectTypeToString(sField.rField.Type)
			var val = GoToEclaType(sField.rValue.Interface())
			struc.Fields[sField.rField.Name] = &val
		}
		return &struc
	default:
		return eclaType.Null{}
	}
}

func ReflectTypeToString(typ reflect.Type) string {
	switch typ.Kind() {
	case reflect.String:
		return "string"
	case reflect.Int:
		return "int"
	case reflect.Float64:
		return "float"
	case reflect.Bool:
		return "bool"
	case reflect.Int32:
		return "char"
	case reflect.Slice:
		return "[]" + ReflectTypeToString(typ.Elem())
	case reflect.Map:
		return "map[" + ReflectTypeToString(typ.Key()) + "]" + ReflectTypeToString(typ.Elem())
	case reflect.Struct:
		return typ.Name()
	default:
		return ""
	}
}
