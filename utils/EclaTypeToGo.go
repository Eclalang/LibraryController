package utils

import (
	"github.com/Eclalang/Ecla/interpreter/eclaType"
	"reflect"
)

// EclaTypeToGo converts an eclaType to a go type.
func EclaTypeToGo(arg eclaType.Type) any {
	switch arg.(type) {
	case eclaType.Int:
		return int(arg.(eclaType.Int))
	case eclaType.Float:
		return float64(arg.(eclaType.Float))
	case eclaType.String:
		return string(arg.(eclaType.String))
	case eclaType.Bool:
		return bool(arg.(eclaType.Bool))
	case eclaType.Char:
		return rune(arg.(eclaType.Char))
	case *eclaType.List:
		// TODO : base the type of the array on the type of the ecla list using eclaType.List.GetType()
		arrType := reflect.SliceOf(TypeStringToReflectType(arg.(*eclaType.List).GetValueType()))
		arr := reflect.MakeSlice(arrType, len(arg.(*eclaType.List).Value), len(arg.(*eclaType.List).Value))
		for i, val := range arg.(*eclaType.List).Value {
			arr.Index(i).Set(reflect.ValueOf(EclaTypeToGo(val)))
		}
		return arr.Interface()
	case *eclaType.Map:
		var types = make(map[any]any)
		for i := 0; i < len(arg.(*eclaType.Map).Keys); i++ {
			types[EclaTypeToGo(arg.(*eclaType.Map).Keys[i])] = EclaTypeToGo(arg.(*eclaType.Map).Values[i])
		}
		mapType := reflect.MapOf(TypeStringToReflectType(arg.(*eclaType.Map).GetKeyTypes()), TypeStringToReflectType(arg.(*eclaType.Map).GetValueTypes()))
		mapVal := reflect.MakeMap(mapType)
		for k, v := range types {
			mapVal.SetMapIndex(reflect.ValueOf(k), reflect.ValueOf(v))
		}
		return mapVal.Interface()
	default:
		return nil
	}
}

func TypeStringToReflectType(typ string) reflect.Type {
	// check if the type is a list
	// start by checking if typ is at least 2 characters long
	if len(typ) > 2 {
		if typ[0] == '[' && typ[1] == ']' {
			// get the rest of the string
			typ = typ[2:]
			// recursively call this function to get the type of the list
			return reflect.SliceOf(TypeStringToReflectType(typ))
		}
	}
	// check if the type is a map
	// start by checking if typ is at least 4 characters long
	if len(typ) > 4 {
		if typ[0] == 'm' && typ[1] == 'a' && typ[2] == 'p' && typ[3] == '[' {
			// get the rest of the string
			typ = typ[3:]
			// go to the next ] considering that there might be nested maps
			var i int
			var nested int
			for i = 0; i < len(typ); i++ {
				if typ[i] == '[' {
					nested++
				}
				if typ[i] == ']' {
					nested--
				}
				if nested == 0 {
					break
				}
			}
			keyType := typ[1:i]
			valueType := typ[i+1:]
			// recursively call this function to get the type of the map
			return reflect.MapOf(TypeStringToReflectType(keyType), TypeStringToReflectType(valueType))
		}
	}
	if typ == "string" {
		return reflect.TypeOf(string(""))
	}
	if typ == "int" {
		return reflect.TypeOf(int(0))
	}
	if typ == "float" {
		return reflect.TypeOf(float64(0))
	}
	if typ == "char" {
		return reflect.TypeOf(rune(0))
	}
	if typ == "bool" {
		return reflect.TypeOf(bool(false))
	}
	return nil
}
