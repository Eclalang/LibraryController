package LibraryController

import (
	"errors"
	"fmt"

	"github.com/Eclalang/Ecla/interpreter/eclaType"
	"github.com/Eclalang/LibraryController/utils"
)

type DebugKingdom struct {
	functionMap map[string]interface{}
}

func NewDebugKingdom() *DebugKingdom {
	return &DebugKingdom{
		functionMap: map[string]interface{}{
			"clear":        nil,
			"returnStruct": nil,
		},
	}
}

type DebugStruct struct {
	Field1 int
	Field2 string
	Field3 bool
	Field4 float64
	Field5 []int
	Field6 map[string]int
}

func (d *DebugKingdom) Call(name string, args []eclaType.Type) ([]eclaType.Type, error) {
	newArgs := make([]any, len(args))
	for k, arg := range args {
		newArgs[k] = utils.EclaTypeToGo(arg)
	}
	if _, ok := d.functionMap[name]; !ok {
		return nil, errors.New(fmt.Sprintf("Method %s not found in package debugKingdom", name))
	}
	switch name {
	case "clear":
		fmt.Print("\033[H\033[2J")
	case "returnStruct":
		structInstance := DebugStruct{
			Field1: 1,
			Field2: "Hello",
			Field3: true,
			Field4: 3.14,
			Field5: []int{1, 2, 3},
			Field6: map[string]int{"a": 1, "b": 2},
		}
		return []eclaType.Type{utils.GoToEclaType(structInstance)}, nil
	}
	return []eclaType.Type{eclaType.Null{}}, nil
}
