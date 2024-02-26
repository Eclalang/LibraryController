package LibraryController

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/Eclalang/Ecla/interpreter/eclaType"
	"github.com/Eclalang/LibraryController/utils"
	enc "github.com/Eclalang/encoding"
)

type Encoding struct {
	functionMap map[string]interface{}
}

func NewEncoding() *Encoding {
	return &Encoding{
		functionMap: map[string]interface{}{
			"asciiToString": nil,
			"decodeBase64":  nil,
			"decodeGob":     nil,
			"decodeHex":     nil,
			"encodeBase64":  nil,
			"encodeGob":     nil,
			"encodeHex":     nil,
			"stringToAscii": nil,
		},
	}
}

func (e *Encoding) Call(name string, args []eclaType.Type) ([]eclaType.Type, error) {
	newArgs := make([]any, len(args))
	for k, arg := range args {
		newArgs[k] = utils.EclaTypeToGo(arg)
	}
	if _, ok := e.functionMap[name]; !ok {
		return nil, errors.New(fmt.Sprintf("Method %s not found in package encoding", name))
	}
	switch name {
	case "asciiToString":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.Slice && len(newArgs) > 1 {
			return []eclaType.Type{utils.GoToEclaType(enc.AsciiToString(newArgs[0].([]int)))}, nil
		}
	case "decodeBase64":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.String && len(newArgs) > 1 {
			return []eclaType.Type{utils.GoToEclaType(enc.DecodeBase64(newArgs[0].(string)))}, nil
		}
	case "decodeGob":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.Slice && len(newArgs) > 1 {
			return []eclaType.Type{utils.GoToEclaType(enc.DecodeGob(newArgs[0].([]int)))}, nil
		}
	case "decodeHex":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.String && len(newArgs) > 1 {
			return []eclaType.Type{utils.GoToEclaType(enc.DecodeHex(newArgs[0].(string)))}, nil
		}
	case "encodeBase64":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.Slice && len(newArgs) > 1 {
			return []eclaType.Type{utils.GoToEclaType(enc.EncodeBase64(newArgs[0].([]int)))}, nil
		}
	case "encodeGob":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.String && len(newArgs) > 1 {
			return []eclaType.Type{utils.GoToEclaType(enc.EncodeGob(newArgs[0].(string)))}, nil
		}
	case "encodeHex":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.Slice && len(newArgs) > 1 {
			return []eclaType.Type{utils.GoToEclaType(enc.EncodeHex(newArgs[0].([]int)))}, nil
		}
	case "stringToAscii":
		return []eclaType.Type{utils.GoToEclaType(enc.StringToAscii(newArgs[0].(string)))}, nil
	default:
		return nil, errors.New(fmt.Sprintf("Method %s not found in package encoding", name))
	}
	return []eclaType.Type{eclaType.Null{}}, nil
}
