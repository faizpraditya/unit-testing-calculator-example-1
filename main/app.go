package main

import (
	services "calculator/service"
	"flag"
	"fmt"
)

func main() {
	num1 := flag.Int("num1", 0, "first number")
	num2 := flag.Int("num2", 0, "second number")
	opr := flag.String("opr", "add", "Calculator operator")
	flag.Parse()

	switch *opr {
	case "add":
		myCalc := services.Calculator{
			*num1, *num2,
		}
		fmt.Println(myCalc.Add())
	case "sub":
		myCalc := services.Calculator{
			*num1, *num2,
		}
		fmt.Println(myCalc.Sub())
	}
}
