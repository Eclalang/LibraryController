package utils

import (
	"fmt"
	"github.com/Eclalang/Ecla/interpreter/eclaType"
	"reflect"
)

type field struct {
	rField reflect.StructField
	rValue reflect.Value
}

func (f field) String() string {
	return fmt.Sprintf("Field: %v, Value: %v", f.rField, f.rValue)
}

func FieldsToGenericStruct(fields []field) (reflect.Value, error) {
	// Create a new instance of the destination type
	var fieldsStruct []reflect.StructField
	var fieldsValues []reflect.Value
	for _, f := range fields {
		fieldsStruct = append(fieldsStruct, f.rField)
		fieldsValues = append(fieldsValues, f.rValue)
	}
	// Create a new instance of the destination type
	ogStructOf := reflect.StructOf(fieldsStruct)
	inter := reflect.New(ogStructOf).Elem().Interface()

	// Get the value of the source struct
	srcVal := reflect.ValueOf(inter)

	// Get the type of the source struct
	srcType := srcVal.Type()

	// Create a new instance of the destination type
	dst := reflect.New(srcType).Elem()

	// dont do any checks since the destination type is a generic struct
	// Copy fields with their values from the source struct to the destination struct
	for i := 0; i < srcVal.NumField(); i++ {
		dstField := dst.FieldByName(srcType.Field(i).Name)
		if dstField.IsValid() && dstField.CanSet() {
			dstField.Set(fieldsValues[i])
		}
	}

	return dst, nil
}

func EclaStructFieldsToFilds(eclaStruct *eclaType.Struct) []field {
	var fields []field
	for key, val := range eclaStruct.Fields {
		var f field
		var derefValGoType = EclaTypeToGo(*val)
		if derefValGoType == nil {
			// for now, we will just ignore the field if it is nil since it is not possible to get the type of nil except when using a function in the EclaTypeToGo function
			continue
		}
		f.rField = reflect.StructField{
			Name: key,
			Type: reflect.TypeOf(derefValGoType),
		}
		f.rValue = reflect.ValueOf(derefValGoType)
		fields = append(fields, f)
	}
	return fields
}

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
	case *eclaType.Struct:
		eclaStruct := arg.(*eclaType.Struct)
		out, err := FieldsToGenericStruct(EclaStructFieldsToFilds(eclaStruct))
		if err != nil {
			return nil
		}
		return out.Interface()
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
