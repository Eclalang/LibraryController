package LibraryController

import (
	"errors"
	"fmt"
	"github.com/Eclalang/Ecla/interpreter/eclaType"
	"github.com/Eclalang/LibraryController/utils"
	"github.com/Eclalang/array"
	"reflect"
)

type Array struct {
	functionMap map[string]interface{}
}

func NewArray() *Array {
	return &Array{
		functionMap: map[string]interface{}{
			"contain":  nil,
			"find":     nil,
			"isEqual":  nil,
			"max":      nil,
			"min":      nil,
			"remove":   nil,
			"slice":    nil,
			"sortAsc":  nil,
			"sortDesc": nil,
		},
	}
}

func (a *Array) Call(name string, args []eclaType.Type) ([]eclaType.Type, error) {
	newArgs := make([]any, len(args))
	for k, arg := range args {
		newArgs[k] = utils.EclaTypeToGo(arg)
	}
	if _, ok := a.functionMap[name]; !ok {
		return nil, errors.New(fmt.Sprintf("Method %s not found in package array", name))
	}
	switch name {
	case "contain":
		if len(newArgs) == 2 && reflect.TypeOf(newArgs[0]).Kind() == reflect.Slice {
			return []eclaType.Type{utils.GoToEclaType(array.Contain(newArgs[0].([]any), newArgs[1]))}, nil
		}
	case "find":
		if len(newArgs) == 2 && reflect.TypeOf(newArgs[0]).Kind() == reflect.Slice {
			return []eclaType.Type{utils.GoToEclaType(array.Find(newArgs[0].([]any), newArgs[1]))}, nil
		}
	case "isEqual":
		if len(newArgs) == 2 && reflect.TypeOf(newArgs[0]).Kind() == reflect.Slice && reflect.TypeOf(newArgs[1]).Kind() == reflect.Slice {
			return []eclaType.Type{utils.GoToEclaType(array.IsEqual(newArgs[0].([]any), newArgs[1].([]any)))}, nil
		}
	case "max":
		if len(newArgs) == 1 && reflect.TypeOf(newArgs[0]).Kind() == reflect.Slice {
			typ := reflect.TypeOf(newArgs[0].([]any)[0]).Kind()
			switch typ {
			case reflect.Int:
				result, err := array.Max(newArgs[0].([]int))
				return []eclaType.Type{utils.GoToEclaType(result)}, err
			case reflect.Float64:
				result, err := array.Max(newArgs[0].([]float64))
				return []eclaType.Type{utils.GoToEclaType(result)}, err
			case reflect.String:
				result, err := array.Max(newArgs[0].([]string))
				return []eclaType.Type{utils.GoToEclaType(result)}, err
			default:
				return nil, errors.New(fmt.Sprintf("Type %s not supported in method max", typ))
			}
		}
	default:
		return nil, errors.New(fmt.Sprintf("Method %s not found in package array", name))
	}

	return []eclaType.Type{eclaType.Null{}}, nil
}

func (a *Array) GetVariables() map[string]eclaType.Type {
	return map[string]eclaType.Type{}
}
