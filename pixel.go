package LibraryController

import (
	"errors"
	"fmt"
	"github.com/Eclalang/Ecla/interpreter/eclaType"
	"github.com/Eclalang/LibraryController/utils"
	"github.com/faiface/mainthread"
	"reflect"
)

type Pixel struct {
	functionMap map[string]interface{}
}

func NewPixel() *Pixel {
	return &Pixel{
		functionMap: map[string]interface{}{
			"start":              nil,
			"testWindowCreation": nil,
			"clear":              nil,
			"updateWindow":       nil,
		},
	}
}

func (p *Pixel) Call(name string, args []eclaType.Type) ([]eclaType.Type, error) {
	// create a channel to sync and block this call until the realCall is done
	ch := make(chan struct{})
	ret := struct {
		result []eclaType.Type
		err    error
	}{}
	mainthread.Run(func() {
		ret.result, ret.err = realCall(p, name, args)
		close(ch)
	})
	<-ch
	return ret.result, ret.err
}

func realCall(p *Pixel, name string, args []eclaType.Type) ([]eclaType.Type, error) {
	newArgs := make([]any, len(args))
	for k, arg := range args {
		newArgs[k] = utils.EclaTypeToGo(arg)
	}
	if _, ok := p.functionMap[name]; !ok {
		return nil, errors.New(fmt.Sprintf("Method %s not found in package pixel", name))
	}
	switch name {
	case "start":
		Start()
		return nil, nil
	case "testWindowCreation":
		return []eclaType.Type{eclaType.Int(TestWindowCreation())}, nil
	case "updateWindow":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.Int && len(newArgs) == 1 {
			UpdateWindow(newArgs[0].(int))
		}
	default:
		return nil, errors.New(fmt.Sprintf("Method %s not found in package pixel", name))
	}
	return []eclaType.Type{eclaType.Null{}}, nil
}

func (p *Pixel) GetVariables() map[string]eclaType.Type {
	return map[string]eclaType.Type{}
}
