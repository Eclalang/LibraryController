package LibraryController

import (
	"errors"
	"fmt"
	"github.com/Eclalang/Ecla/interpreter"
	"github.com/Eclalang/Ecla/interpreter/eclaType"
	"github.com/Eclalang/LibraryController/utils"
	"github.com/Eclalang/console"
)

type Console struct {
	functionMap map[string]interface{}
}

func NewConsole() *Console {
	return &Console{
		functionMap: map[string]interface{}{
			"clear":        nil,
			"input":        nil,
			"inputFloat":   nil,
			"inputInt":     nil,
			"print":        nil,
			"printf":       nil,
			"printInColor": nil,
			"println":      nil,
		},
	}
}

func (c *Console) Call(name string, args []eclaType.Type) ([]eclaType.Type, error) {
	newArgs := make([]any, len(args))
	for k, arg := range args {
		newArgs[k] = utils.EclaTypeToGo(arg)
	}
	if _, ok := c.functionMap[name]; !ok {
		return nil, errors.New(fmt.Sprintf("Method %s not found in package console", name))
	}
	switch name {
	case "clear":
		console.Clear()
	case "input":
		return []eclaType.Type{utils.GoToEclaType(console.Input())}, nil
	case "inputFloat":
		input, err := console.InputFloat()
		return []eclaType.Type{utils.GoToEclaType(input)}, err
	case "inputInt":
		input, err := console.InputInt()
		return []eclaType.Type{utils.GoToEclaType(input)}, err
	case "print":
		console.Print(newArgs...)
	case "printf":
		// TODO: refactor this line
		console.Printf(newArgs[0].(string), newArgs[1:]...)
	case "printInColor":
		console.PrintInColor(newArgs[0].(string), newArgs[1:]...)
		// To add later
		//case "printlnInColor":
		//	console.PrintlnInColor(newArgs[0].(string), newArgs[1:]...)
	case "println":
		console.Println(newArgs...)
	default:
		return nil, errors.New(fmt.Sprintf("Method %s not found in package console", name))
	}

	return []eclaType.Type{eclaType.Null{}}, nil
}

func (c *Console) GetScope() *interpreter.Scope {
	return nil
}
