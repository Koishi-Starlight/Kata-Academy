package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getResult(num1, num2 string, operator string) int {
	var value int
	convNum1, err := strconv.Atoi(num1)
	errorHandler(err)
	convNum2, err := strconv.Atoi(num2)
	errorHandler(err)
	switch operator {
	case "+":
		value = convNum1 + convNum2
	case "-":
		value = convNum1 - convNum2
	case "*":
		value = convNum1 * convNum2
	case "/":
		value = convNum1 / convNum2
	default:
		panic("Unknown operator")
	}
	return value
}

func errorHandler(err error) {
	if err != nil {
		panic(err)
	}
}

var roman = map[string]int{
	"I":    1,
	"II":   2,
	"III":  3,
	"IV":   4,
	"V":    5,
	"VI":   6,
	"VII":  7,
	"VIII": 8,
	"IX":   9,
	"X":    10,
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(`Enter mathematical operations in following format: "1 + 1".
Available numbers 0 - 10, I - X
Available operations: +, -, *, /
Example: 1 + 2
Example 2: I + VI
Example 3: 8 - 2
PS: Roman numbers don't support negative.'`)
	text, _ := reader.ReadString('\n')         // Awaiting input, string type
	text = strings.TrimSpace(text)             // Removing \n from input
	operationSlice := strings.Split(text, " ") //Making a slice, removes spaces as separator
	if len(operationSlice) != 3 {              //Checks if you entered correct amount of arguments. 2 numbers and 1 operator
		panic("Incorrect amount of arguments.")
	}
	num1 := operationSlice[0]
	num2 := operationSlice[2]
	operator := operationSlice[1]
	fmt.Println(num1, num2, operator)
	fmt.Println(roman[num1], roman[num2])
	//fmt.Printf("Your result is %d", getResult(num1, num2, operator))
}
