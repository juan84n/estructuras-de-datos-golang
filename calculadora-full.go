package main

import (
	"fmt"
	"strconv"
	"strings"
)

func test() {
	var expresion string
	var num1 int
	flag := false
	var operator string
	//var operators []string
	//var numbers []int
	//var result int

	fmt.Print("Expresión arítmetica =>")
	fmt.Scanln(&expresion)

	valores := strings.Split(expresion, "")
	//fmt.Println("el array es ", valores)
	for _, valor := range valores {
		num, error := strconv.Atoi(valor)
		//fmt.Println("el numero y el error son", num, error)
		if error != nil {
			operator = valor
			flag = true
			//operators = append(operators, valor)
		} else {
			if flag {
				returnOperation(num1, num, operator)
			}
			num1 = num
			//numbers = append(numbers, num)
		}
	}
	/*if len(operators) >= len(numbers) {
		fmt.Println("Hay un problema con la operación")
	} else {
		fmt.Println("los slices son", numbers, operators)
		for i:= 0; i < len(operators); i++ {
			if i == 0 {
				returnOperation(numbers[0], numbers[1], operators[i])
			} else {
				returnOperation(numbers[i + 1], numbers[i + 2], operators[i])
			}
		}
	}*/

}

func returnOperation(num1, num2 int, operator string) int {
	switch operator {
	case "+":
		return num1 + num2
	case "-":
		return num1 - num2
	case "*":
		return num1 * num2
	case "/":
		return num1 / num2
	default:
		return 0

	}
}
