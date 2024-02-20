package LibraryController

import (
	"errors"
	"fmt"
	"github.com/Eclalang/Ecla/interpreter"
	"reflect"

	"github.com/Eclalang/Ecla/interpreter/eclaType"
	"github.com/Eclalang/LibraryController/utils"
	"github.com/Eclalang/json"
)

type Json struct {
	functionMap map[string]interface{}
}

func NewJson() *Json {
	return &Json{
		functionMap: map[string]interface{}{
			"marshal":   nil,
			"unmarshal": nil,
		},
	}
}

func (j *Json) Call(name string, args []eclaType.Type) ([]eclaType.Type, error) {
	newArgs := make([]any, len(args))
	for k, arg := range args {
		newArgs[k] = utils.EclaTypeToGo(arg)
	}
	if _, ok := j.functionMap[name]; !ok {
		return nil, errors.New(fmt.Sprintf("Method %s not found in package json", name))
	}
	switch name {
	case "marshal":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.Map {
			content, err := json.Marshal(newArgs[0].(map[string]string))
			return []eclaType.Type{utils.GoToEclaType(content)}, err
		} else if reflect.TypeOf(newArgs[0]).Kind() == reflect.Slice {
			content, err := json.Marshal(newArgs[0].([]map[string]string))
			return []eclaType.Type{utils.GoToEclaType(content)}, err
		}
	case "unmarshal":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.String {
			content, err := json.Unmarshal(newArgs[0].(string))
			return []eclaType.Type{utils.GoToEclaType(content)}, err
		}
	default:
		return nil, errors.New(fmt.Sprintf("Method %s not found in package json", name))
	}

	return []eclaType.Type{eclaType.Null{}}, nil
}

func (j *Json) GetScope() *interpreter.Scope {
	return nil
}
