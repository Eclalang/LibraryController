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
	variables := make(map[string]eclaType.Type)
	// Reset the color
	variables["COLOROFF"], _ = eclaType.NewString(console.ColorOff)

	// Regular colors
	variables["BLACK"], _ = eclaType.NewString(console.BLACK)
	variables["RED"], _ = eclaType.NewString(console.RED)
	variables["GREEN"], _ = eclaType.NewString(console.GREEN)
	variables["YELLOW"], _ = eclaType.NewString(console.YELLOW)
	variables["BLUE"], _ = eclaType.NewString(console.BLUE)
	variables["MAGENTA"], _ = eclaType.NewString(console.MAGENTA)
	variables["CYAN"], _ = eclaType.NewString(console.CYAN)
	variables["WHITE"], _ = eclaType.NewString(console.WHITE)

	// Bold colors
	variables["B_BLACK"], _ = eclaType.NewString(console.BBlack)
	variables["B_RED"], _ = eclaType.NewString(console.BRed)
	variables["B_GREEN"], _ = eclaType.NewString(console.BGreen)
	variables["B_YELLOW"], _ = eclaType.NewString(console.BYellow)
	variables["B_BLUE"], _ = eclaType.NewString(console.BBlue)
	variables["B_MAGENTA"], _ = eclaType.NewString(console.BMagenta)
	variables["B_CYAN"], _ = eclaType.NewString(console.BCyan)
	variables["B_WHITE"], _ = eclaType.NewString(console.BWhite)

	// Underline colors
	variables["U_BLACK"], _ = eclaType.NewString(console.UBlack)
	variables["U_RED"], _ = eclaType.NewString(console.URed)
	variables["U_GREEN"], _ = eclaType.NewString(console.UGreen)
	variables["U_YELLOW"], _ = eclaType.NewString(console.UYellow)
	variables["U_BLUE"], _ = eclaType.NewString(console.UBlue)
	variables["U_MAGENTA"], _ = eclaType.NewString(console.UMagenta)
	variables["U_CYAN"], _ = eclaType.NewString(console.UCyan)
	variables["U_WHITE"], _ = eclaType.NewString(console.UWhite)

	// Background colors
	variables["ON_BLACK"], _ = eclaType.NewString(console.On_Black)
	variables["ON_RED"], _ = eclaType.NewString(console.On_Red)
	variables["ON_GREEN"], _ = eclaType.NewString(console.On_Green)
	variables["ON_YELLOW"], _ = eclaType.NewString(console.On_Yellow)
	variables["ON_BLUE"], _ = eclaType.NewString(console.On_Blue)
	variables["ON_MAGENTA"], _ = eclaType.NewString(console.On_Magenta)
	variables["ON_CYAN"], _ = eclaType.NewString(console.On_Cyan)
	variables["ON_WHITE"], _ = eclaType.NewString(console.On_White)

	// High Intensity colors
	variables["I_BLACK"], _ = eclaType.NewString(console.IBlack)
	variables["I_RED"], _ = eclaType.NewString(console.IRed)
	variables["I_GREEN"], _ = eclaType.NewString(console.IGreen)
	variables["I_YELLOW"], _ = eclaType.NewString(console.IYellow)
	variables["I_BLUE"], _ = eclaType.NewString(console.IBlue)
	variables["I_MAGENTA"], _ = eclaType.NewString(console.IMagenta)
	variables["I_CYAN"], _ = eclaType.NewString(console.ICyan)
	variables["I_WHITE"], _ = eclaType.NewString(console.IWhite)

	// Bold High Intensity colors
	variables["BI_BLACK"], _ = eclaType.NewString(console.BIBlack)
	variables["BI_RED"], _ = eclaType.NewString(console.BIRed)
	variables["BI_GREEN"], _ = eclaType.NewString(console.BIGreen)
	variables["BI_YELLOW"], _ = eclaType.NewString(console.BIYellow)
	variables["BI_BLUE"], _ = eclaType.NewString(console.BIBlue)
	variables["BI_MAGENTA"], _ = eclaType.NewString(console.BIMagenta)
	variables["BI_CYAN"], _ = eclaType.NewString(console.BICyan)
	variables["BI_WHITE"], _ = eclaType.NewString(console.BIWhite)

	// High Intensity background colors
	variables["ONI_BLACK"], _ = eclaType.NewString(console.On_IBlack)
	variables["ONI_RED"], _ = eclaType.NewString(console.On_IRed)
	variables["ONI_GREEN"], _ = eclaType.NewString(console.On_IGreen)
	variables["ONI_YELLOW"], _ = eclaType.NewString(console.On_IYellow)
	variables["ONI_BLUE"], _ = eclaType.NewString(console.On_IBlue)
	variables["ONI_MAGENTA"], _ = eclaType.NewString(console.On_IMagenta)
	variables["ONI_CYAN"], _ = eclaType.NewString(console.On_ICyan)
	variables["ONI_WHITE"], _ = eclaType.NewString(console.On_IWhite)

	return variables
}
