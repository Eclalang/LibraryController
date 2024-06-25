package utils

import (
	"fmt"
	"reflect"
)

func ConvertStruct(struc any, outputType reflect.Type) (interface{}, error) {
	// get all fields of the struct
	var allFields = GetFields(struc)
	// convert the fields to the destination type
	return FieldsToTypedStruct(allFields, outputType)
}

// GetFields returns a list of all fields of a struct. It will not report unexported fields.
func GetFields(struc any) []field {
	numberOfFields := reflect.TypeOf(struc).NumField()
	var allFields []field
	var allFieldsValues = GetReflectFieldsValues(struc)
	var allFieldsStruct = GetReflectFields(struc)
	// create a list of all fields
	for i := 0; i < numberOfFields; i++ {
		allFields = append(allFields, field{rField: allFieldsStruct[i], rValue: allFieldsValues[i]})
	}
	return allFields
}

// GetReflectFieldsValues returns a list of all fields values of a struct. It will not report unexported fields.
func GetReflectFieldsValues(struc any) []reflect.Value {
	numberOfFields := reflect.TypeOf(struc).NumField()
	var allFieldsValues []reflect.Value
	// create a list of all fields values
	for i := 0; i < numberOfFields; i++ {
		if reflect.ValueOf(struc).Field(i).CanInterface() {
			allFieldsValues = append(allFieldsValues, reflect.ValueOf(struc).FieldByName(reflect.TypeOf(struc).Field(i).Name))
		}
	}
	return allFieldsValues
}

// GetReflectFields returns a list of all fields of a struct. It will not report unexported fields.
func GetReflectFields(struc any) []reflect.StructField {
	numberOfFields := reflect.TypeOf(struc).NumField()
	var allFields []reflect.StructField
	// create a list of all fields
	for i := 0; i < numberOfFields; i++ {
		if reflect.ValueOf(struc).Field(i).CanInterface() {
			allFields = append(allFields, reflect.TypeOf(struc).Field(i))
		}
	}
	return allFields
}

/*
FieldsToTypedStruct converts a list of fields to a struct of a specific type.
It will throw an error if :
  - the destination type is not a struct
  - the destination struct has unexported fields
  - the number of fields does not match the number of fields in the destination type
  - the order of fields does not match the order of fields in the destination type
  - the name of fields does not match the name of fields in the destination type
  - the type of fields does not match the type of fields in the destination type
*/
func FieldsToTypedStruct(fields []field, outputType reflect.Type) (reflect.Value, error) {
	// check if the destination type is a struct
	if outputType.Kind() != reflect.Struct {
		return reflect.Value{}, fmt.Errorf("destination type is not a struct")
	}

	var fieldsStruct []reflect.StructField
	var fieldsValues []reflect.Value
	for _, f := range fields {
		fieldsStruct = append(fieldsStruct, f.rField)
		fieldsValues = append(fieldsValues, f.rValue)
	}
	ogStructOf := reflect.StructOf(fieldsStruct)
	inter := reflect.New(ogStructOf).Elem().Interface()

	// Create a new instance of the destination type
	dst := reflect.New(outputType).Elem()

	// Get the value of the source struct
	srcVal := reflect.ValueOf(inter)

	// Get the type of the source struct
	srcType := srcVal.Type()

	// do checks to assure if the number of fields match
	if len(fields) != dst.NumField() {
		return reflect.Value{}, fmt.Errorf("incorrect number of fields : want %d, got %d", dst.NumField(), len(fields))
	}
	// do checks to assure if the order,name and type of fields match
	for i := 0; i < srcVal.NumField(); i++ {
		srcField := srcVal.Field(i)
		dstField := dst.Field(i)
		// Check if the field is unexported
		if !srcField.CanInterface() {
			return reflect.Value{}, fmt.Errorf("field %v is unexported so it can't be copied nor accessed", srcType.Field(i).Name)
		}

		if srcType.Field(i).Name != outputType.Field(i).Name {
			return reflect.Value{}, fmt.Errorf("name mismatch : want %v, got %v", outputType.Field(i).Name, srcType.Field(i).Name)
		}
		if reflect.TypeOf(srcField.Interface()) != reflect.TypeOf(dstField.Interface()) || srcField.Kind() != dstField.Kind() || srcField.Type() != dstField.Type() {
			return reflect.Value{}, fmt.Errorf("type mismatch on field %v : want %v, got %v", srcType.Field(i).Name, reflect.TypeOf(dstField.Interface()), reflect.TypeOf(srcField.Interface()))
		}
	}

	// Copy fields with their values from the source struct to the destination struct
	for i := 0; i < srcVal.NumField(); i++ {
		dstField := dst.FieldByName(srcType.Field(i).Name)
		if dstField.IsValid() && dstField.CanSet() {
			dstField.Set(fieldsValues[i])
		}
	}

	return dst, nil
}
