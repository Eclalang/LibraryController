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
			typ := reflect.TypeOf(newArgs[0]).Elem().Kind()
			switch typ {
			case reflect.Int:
				if reflect.TypeOf(newArgs[1]).Kind() != reflect.Int {
					return nil, errors.New(fmt.Sprintf("Type %s mismatch with array type %s", reflect.TypeOf(newArgs[1]).Kind(), typ))
				}
				return []eclaType.Type{utils.GoToEclaType(array.Contain(newArgs[0].([]int), newArgs[1].(int)))}, nil
			case reflect.Float64:
				if reflect.TypeOf(newArgs[1]).Kind() != reflect.Float64 {
					return nil, errors.New(fmt.Sprintf("Type %s mismatch with array type %s", reflect.TypeOf(newArgs[1]).Kind(), typ))
				}
				return []eclaType.Type{utils.GoToEclaType(array.Contain(newArgs[0].([]float64), newArgs[1].(float64)))}, nil
			case reflect.String:
				if reflect.TypeOf(newArgs[1]).Kind() != reflect.String {
					return nil, errors.New(fmt.Sprintf("Type %s mismatch with array type %s", reflect.TypeOf(newArgs[1]).Kind(), typ))
				}
				return []eclaType.Type{utils.GoToEclaType(array.Contain(newArgs[0].([]string), newArgs[1].(string)))}, nil
			default:
				return nil, errors.New(fmt.Sprintf("Type %s not supported in method contain", typ))
			}
		}
	case "find":
		if len(newArgs) == 2 && reflect.TypeOf(newArgs[0]).Kind() == reflect.Slice {
			typ := reflect.TypeOf(newArgs[0]).Elem().Kind()
			switch typ {
			case reflect.Int:
				if reflect.TypeOf(newArgs[1]).Kind() != reflect.Int {
					return nil, errors.New(fmt.Sprintf("Type %s mismatch with array type %s", reflect.TypeOf(newArgs[1]).Kind(), typ))
				}
				return []eclaType.Type{utils.GoToEclaType(array.Find(newArgs[0].([]int), newArgs[1].(int)))}, nil
			case reflect.Float64:
				if reflect.TypeOf(newArgs[1]).Kind() != reflect.Float64 {
					return nil, errors.New(fmt.Sprintf("Type %s mismatch with array type %s", reflect.TypeOf(newArgs[1]).Kind(), typ))
				}
				return []eclaType.Type{utils.GoToEclaType(array.Find(newArgs[0].([]float64), newArgs[1].(float64)))}, nil
			case reflect.String:
				if reflect.TypeOf(newArgs[1]).Kind() != reflect.String {
					return nil, errors.New(fmt.Sprintf("Type %s mismatch with array type %s", reflect.TypeOf(newArgs[1]).Kind(), typ))
				}
				return []eclaType.Type{utils.GoToEclaType(array.Find(newArgs[0].([]string), newArgs[1].(string)))}, nil
			default:
				return nil, errors.New(fmt.Sprintf("Type %s not supported in method find", typ))
			}
		}
	case "isEqual":
		if len(newArgs) == 2 && reflect.TypeOf(newArgs[0]).Kind() == reflect.Slice && reflect.TypeOf(newArgs[1]).Kind() == reflect.Slice {
			typ := reflect.TypeOf(newArgs[0]).Elem().Kind()
			switch typ {
			case reflect.Int:
				if reflect.TypeOf(newArgs[1]).Elem().Kind() != reflect.Int {
					return nil, errors.New(fmt.Sprintf("Array type %s mismatch with array type %s", reflect.TypeOf(newArgs[1]).Elem().Kind(), typ))
				}
				return []eclaType.Type{utils.GoToEclaType(array.IsEqual(newArgs[0].([]int), newArgs[1].([]int)))}, nil
			case reflect.Float64:
				if reflect.TypeOf(newArgs[1]).Elem().Kind() != reflect.Float64 {
					return nil, errors.New(fmt.Sprintf("Array type %s mismatch with array type %s", reflect.TypeOf(newArgs[1]).Elem().Kind(), typ))
				}
				return []eclaType.Type{utils.GoToEclaType(array.IsEqual(newArgs[0].([]float64), newArgs[1].([]float64)))}, nil
			case reflect.String:
				if reflect.TypeOf(newArgs[1]).Elem().Kind() != reflect.String {
					return nil, errors.New(fmt.Sprintf("Array type %s mismatch with array type %s", reflect.TypeOf(newArgs[1]).Elem().Kind(), typ))
				}
				return []eclaType.Type{utils.GoToEclaType(array.IsEqual(newArgs[0].([]string), newArgs[1].([]string)))}, nil
			default:
				return nil, errors.New(fmt.Sprintf("Array type %s not supported in method isEqual", typ))
			}
		}
	case "max":
		if len(newArgs) == 1 && reflect.TypeOf(newArgs[0]).Kind() == reflect.Slice {
			typ := reflect.TypeOf(newArgs[0]).Elem().Kind()
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
	case "min":
		if len(newArgs) == 1 && reflect.TypeOf(newArgs[0]).Kind() == reflect.Slice {
			typ := reflect.TypeOf(newArgs[0]).Elem().Kind()
			switch typ {
			case reflect.Int:
				result, err := array.Min(newArgs[0].([]int))
				return []eclaType.Type{utils.GoToEclaType(result)}, err
			case reflect.Float64:
				result, err := array.Min(newArgs[0].([]float64))
				return []eclaType.Type{utils.GoToEclaType(result)}, err
			case reflect.String:
				result, err := array.Min(newArgs[0].([]string))
				return []eclaType.Type{utils.GoToEclaType(result)}, err
			default:
				return nil, errors.New(fmt.Sprintf("Type %s not supported in method min", typ))
			}
		}
	case "remove":
		if len(newArgs) == 2 && reflect.TypeOf(newArgs[0]).Kind() == reflect.Slice && reflect.TypeOf(newArgs[1]).Kind() == reflect.Int {
			typ := reflect.TypeOf(newArgs[0]).Elem().Kind()
			switch typ {
			case reflect.Int:
				return []eclaType.Type{utils.GoToEclaType(array.Remove(newArgs[0].([]int), newArgs[1].(int)))}, nil
			case reflect.Float64:
				return []eclaType.Type{utils.GoToEclaType(array.Remove(newArgs[0].([]float64), newArgs[1].(int)))}, nil
			case reflect.String:
				return []eclaType.Type{utils.GoToEclaType(array.Remove(newArgs[0].([]string), newArgs[1].(int)))}, nil
			default:
				return nil, errors.New(fmt.Sprintf("Array type %s not supported in method remove", typ))
			}
		}
	case "slice":
		if len(newArgs) == 3 && reflect.TypeOf(newArgs[0]).Kind() == reflect.Slice && reflect.TypeOf(newArgs[1]).Kind() == reflect.Int && reflect.TypeOf(newArgs[2]).Kind() == reflect.Int {
			typ := reflect.TypeOf(newArgs[0]).Elem().Kind()
			switch typ {
			case reflect.Int:
				slice, err := array.Slice(newArgs[0].([]int), newArgs[1].(int), newArgs[2].(int))
				return []eclaType.Type{utils.GoToEclaType(slice)}, err
			case reflect.Float64:
				slice, err := array.Slice(newArgs[0].([]float64), newArgs[1].(int), newArgs[2].(int))
				return []eclaType.Type{utils.GoToEclaType(slice)}, err
			case reflect.String:
				slice, err := array.Slice(newArgs[0].([]string), newArgs[1].(int), newArgs[2].(int))
				return []eclaType.Type{utils.GoToEclaType(slice)}, err
			default:
				return nil, errors.New(fmt.Sprintf("Array type %s not supported in method slice", typ))
			}
		}
	case "sortAsc":
		if len(newArgs) == 1 && reflect.TypeOf(newArgs[0]).Kind() == reflect.Slice {
			typ := reflect.TypeOf(newArgs[0]).Elem().Kind()
			switch typ {
			case reflect.Int:
				array.SortAsc(newArgs[0].([]int))
				return []eclaType.Type{utils.GoToEclaType(newArgs[0].([]int))}, nil
			case reflect.Float64:
				array.SortAsc(newArgs[0].([]float64))
				return []eclaType.Type{utils.GoToEclaType(newArgs[0].([]float64))}, nil
			case reflect.String:
				array.SortAsc(newArgs[0].([]string))
				return []eclaType.Type{utils.GoToEclaType(newArgs[0].([]string))}, nil
			default:
				return nil, errors.New(fmt.Sprintf("Array type %s not supported in method sortAsc", typ))
			}
		}
	case "sortDesc":
		if len(newArgs) == 1 && reflect.TypeOf(newArgs[0]).Kind() == reflect.Slice {
			typ := reflect.TypeOf(newArgs[0]).Elem().Kind()
			switch typ {
			case reflect.Int:
				array.SortDesc(newArgs[0].([]int))
				return []eclaType.Type{utils.GoToEclaType(newArgs[0].([]int))}, nil
			case reflect.Float64:
				array.SortDesc(newArgs[0].([]float64))
				return []eclaType.Type{utils.GoToEclaType(newArgs[0].([]float64))}, nil
			case reflect.String:
				array.SortDesc(newArgs[0].([]string))
				return []eclaType.Type{utils.GoToEclaType(newArgs[0].([]string))}, nil
			default:
				return nil, errors.New(fmt.Sprintf("Array type %s not supported in method sortDesc", typ))
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
