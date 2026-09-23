package main

import "fmt"
import "flag"
import "strconv"
import "os"

type App struct {
	num1 int
	num2 int
	mode string
	operator string
	float1 float64
	float2 float64
	resultInt int
	resultFloat float64
}

func getFlags(app *App) error {

	flag.Parse()
	f := flag.Args()

	if len(f) != 3 {
		return fmt.Errorf("Program received %v arguments. Needs exactly 3.", len(f))
	}

	app.mode = "integer"
	combined := f[0] + f[2]

	for i := 0; i < len(combined); i++ {
		if combined[i] == '.' {
			app.mode = "float"
		}
	}
	
	var err1, err2 error
	switch app.mode {
	case "integer":
		app.num1, err1 = strconv.Atoi(f[0])
		app.num2, err2 = strconv.Atoi(f[2])
	case "float":
		app.float1, err1 = strconv.ParseFloat(f[0], 64)
		app.float2, err2 = strconv.ParseFloat(f[2], 64)
	}
	if err1 != nil || err2 != nil {
		return fmt.Errorf("Error while parsing numbers with mode %v, number 1: %v, number 2: %v", app.mode, err1, err2)
	}


	switch f[1] {
	case "+", "-", "*", "/":
		app.operator = f[1]
	default:
		return fmt.Errorf("Error while parsing the operator: not a valid operator")
	}

	return nil
}

func calcResult(app *App) {
	switch app.mode {
	case "integer":
		calcInteger(app)
		break
	case "float":
		calcFloat(app)
	}
}

func calcInteger(app *App) {
	switch app.operator {
	case "+":
		app.resultInt = app.num1 + app.num2
	case "-":
		app.resultInt = app.num1 - app.num2
	case "*":
		app.resultInt = app.num1 * app.num2
	case "/":
		app.mode = "float"
		if app.num2 == 0 {
			fmt.Println("Can not divide by 0")
			os.Exit(1)
		}
		app.resultFloat = float64(app.num1) / float64(app.num2)
	}


}

func calcFloat(app *App) {
	switch app.operator {
	case "+":
		app.resultFloat = app.float1 + app.float2
	case "-":
		app.resultFloat = app.float1 - app.float2
	case "*":
		app.resultFloat = app.float1 * app.float2
	case "/":
		app.mode = "float"
		if app.float2 == 0 {
			fmt.Println("Can not divide by 0")
			os.Exit(1)
		}
		app.resultFloat = app.float1 / app.float2
	}

}


func printResult(app *App) {

	switch app.mode {
	case "integer":
		fmt.Println(app.resultInt)
	case "float":
		fmt.Println(app.resultFloat)
	default:
		fmt.Println("something went wrong")
	}

}


func main() {

	var app = App{}

	err := getFlags(&app)
	if err != nil {
		fmt.Printf("error during flag parsing: %v\n", err)
		os.Exit(1)
	}

	calcResult(&app)

	printResult(&app)

}
