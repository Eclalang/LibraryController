package LibraryController

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/Eclalang/Ecla/interpreter/eclaType"
	"github.com/Eclalang/LibraryController/utils"
	"github.com/Eclalang/time"
)

type Time struct {
	functionMap map[string]interface{}
}

func NewTime() *Time {
	return &Time{
		functionMap: map[string]interface{}{
			"convertRoman": nil,
			"date":         nil,
			"now":          nil,
			"sleep":        nil,
			"strftime":     nil,
		},
	}
}

func (t *Time) Call(name string, args []eclaType.Type) ([]eclaType.Type, error) {
	newArgs := make([]any, len(args))
	for k, arg := range args {
		newArgs[k] = utils.EclaTypeToGo(arg)
	}
	if _, ok := t.functionMap[name]; !ok {
		return nil, errors.New(fmt.Sprintf("Method %s not found in package time", name))
	}
	switch name {
	case "convertRoman":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.String && len(newArgs) == 1 {
			return []eclaType.Type{utils.GoToEclaType(time.ConvertRoman(newArgs[0].(string)))}, nil
		}
	case "date":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.Int && reflect.TypeOf(newArgs[1]).Kind() == reflect.Int && reflect.TypeOf(newArgs[2]).Kind() == reflect.Int && reflect.TypeOf(newArgs[3]).Kind() == reflect.Int && reflect.TypeOf(newArgs[4]).Kind() == reflect.Int && reflect.TypeOf(newArgs[5]).Kind() == reflect.Int && len(newArgs) == 6 {
			return []eclaType.Type{utils.GoToEclaType(time.Date(newArgs[0].(int), newArgs[1].(int), newArgs[2].(int), newArgs[3].(int), newArgs[4].(int), newArgs[5].(int)))}, nil
		}
	case "now":
		return []eclaType.Type{utils.GoToEclaType(time.Now())}, nil
	case "sleep":
		if len(newArgs) == 1 {
			switch reflect.TypeOf(newArgs[0]).Kind() {
			case reflect.Int:
				time.Sleep(newArgs[0].(int))
			case reflect.Float64:
				time.Sleep(newArgs[0].(float64))
			}
		}
	case "strftime":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.String && reflect.TypeOf(newArgs[1]).Kind() == reflect.String && len(newArgs) == 2 {
			return []eclaType.Type{utils.GoToEclaType(time.Strftime(newArgs[0].(string), newArgs[1].(string)))}, nil
		}
	default:
		return nil, errors.New(fmt.Sprintf("Method %s not found in package time", name))
	}

	return []eclaType.Type{eclaType.Null{}}, nil
}
