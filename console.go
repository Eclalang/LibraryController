package LibraryController

import (
	"errors"
	"fmt"
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

func (c *Console) GetVariables() map[string]eclaType.Type {
	args := make(map[string]eclaType.Type)
	// Reset the color
	args["COLOROFF"], _ = eclaType.NewString(console.ColorOff)

	// Regular colors
	args["BLACK"], _ = eclaType.NewString(console.BLACK)
	args["RED"], _ = eclaType.NewString(console.RED)
	args["GREEN"], _ = eclaType.NewString(console.GREEN)
	args["YELLOW"], _ = eclaType.NewString(console.YELLOW)
	args["BLUE"], _ = eclaType.NewString(console.BLUE)
	args["MAGENTA"], _ = eclaType.NewString(console.MAGENTA)
	args["CYAN"], _ = eclaType.NewString(console.CYAN)
	args["WHITE"], _ = eclaType.NewString(console.WHITE)

	// Bold colors
	args["B_BLACK"], _ = eclaType.NewString(console.BBlack)
	args["B_RED"], _ = eclaType.NewString(console.BRed)
	args["B_GREEN"], _ = eclaType.NewString(console.BGreen)
	args["B_YELLOW"], _ = eclaType.NewString(console.BYellow)
	args["B_BLUE"], _ = eclaType.NewString(console.BBlue)
	args["B_MAGENTA"], _ = eclaType.NewString(console.BMagenta)
	args["B_CYAN"], _ = eclaType.NewString(console.BCyan)
	args["B_WHITE"], _ = eclaType.NewString(console.BWhite)

	// Underline colors
	args["U_BLACK"], _ = eclaType.NewString(console.UBlack)
	args["U_RED"], _ = eclaType.NewString(console.URed)
	args["U_GREEN"], _ = eclaType.NewString(console.UGreen)
	args["U_YELLOW"], _ = eclaType.NewString(console.UYellow)
	args["U_BLUE"], _ = eclaType.NewString(console.UBlue)
	args["U_MAGENTA"], _ = eclaType.NewString(console.UMagenta)
	args["U_CYAN"], _ = eclaType.NewString(console.UCyan)
	args["U_WHITE"], _ = eclaType.NewString(console.UWhite)

	// Background colors
	args["ON_BLACK"], _ = eclaType.NewString(console.On_Black)
	args["ON_RED"], _ = eclaType.NewString(console.On_Red)
	args["ON_GREEN"], _ = eclaType.NewString(console.On_Green)
	args["ON_YELLOW"], _ = eclaType.NewString(console.On_Yellow)
	args["ON_BLUE"], _ = eclaType.NewString(console.On_Blue)
	args["ON_MAGENTA"], _ = eclaType.NewString(console.On_Magenta)
	args["ON_CYAN"], _ = eclaType.NewString(console.On_Cyan)
	args["ON_WHITE"], _ = eclaType.NewString(console.On_White)

	// High Intensity colors
	args["I_BLACK"], _ = eclaType.NewString(console.IBlack)
	args["I_RED"], _ = eclaType.NewString(console.IRed)
	args["I_GREEN"], _ = eclaType.NewString(console.IGreen)
	args["I_YELLOW"], _ = eclaType.NewString(console.IYellow)
	args["I_BLUE"], _ = eclaType.NewString(console.IBlue)
	args["I_MAGENTA"], _ = eclaType.NewString(console.IMagenta)
	args["I_CYAN"], _ = eclaType.NewString(console.ICyan)
	args["I_WHITE"], _ = eclaType.NewString(console.IWhite)

	// Bold High Intensity colors
	args["BI_BLACK"], _ = eclaType.NewString(console.BIBlack)
	args["BI_RED"], _ = eclaType.NewString(console.BIRed)
	args["BI_GREEN"], _ = eclaType.NewString(console.BIGreen)
	args["BI_YELLOW"], _ = eclaType.NewString(console.BIYellow)
	args["BI_BLUE"], _ = eclaType.NewString(console.BIBlue)
	args["BI_MAGENTA"], _ = eclaType.NewString(console.BIMagenta)
	args["BI_CYAN"], _ = eclaType.NewString(console.BICyan)
	args["BI_WHITE"], _ = eclaType.NewString(console.BIWhite)

	// High Intensity background colors
	args["ONI_BLACK"], _ = eclaType.NewString(console.On_IBlack)
	args["ONI_RED"], _ = eclaType.NewString(console.On_IRed)
	args["ONI_GREEN"], _ = eclaType.NewString(console.On_IGreen)
	args["ONI_YELLOW"], _ = eclaType.NewString(console.On_IYellow)
	args["ONI_BLUE"], _ = eclaType.NewString(console.On_IBlue)
	args["ONI_MAGENTA"], _ = eclaType.NewString(console.On_IMagenta)
	args["ONI_CYAN"], _ = eclaType.NewString(console.On_ICyan)
	args["ONI_WHITE"], _ = eclaType.NewString(console.On_IWhite)

	return args
}
