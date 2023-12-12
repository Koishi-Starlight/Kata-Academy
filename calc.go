package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getIntResult(num1, num2 int, operator string) int { // Basic operation function
	var value int
	switch operator {
	case "+":
		value = num1 + num2
	case "-":
		value = num1 - num2
	case "*":
		value = num1 * num2
	case "/":
		value = num1 / num2
	default:
		panic("Illegal operator")
	}
	return value
}
func intToRoman(num int) string { //Arabic num to Roman num converter, using subtraction method.
	var roman string

	var notations []string = []string{"C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	var value []int = []int{100, 90, 50, 40, 10, 9, 5, 4, 1}

	for i := 0; num > 0; i++ {
		for num >= value[i] {
			roman += notations[i]
			num -= value[i]
		}
	}
	return roman
}
func GetRomanResults(num1, num2, operator string) string { //Using roman map, passing value for calculation, passing result to converter
	result := getIntResult(availableRoman[num1], availableRoman[num2], operator)
	if result < 0 {
		panic("The operation resulted with negative number")
	}
	return intToRoman(result)
}
var availableRoman = map[string]int{ // Map used for checking if entered number is available for use and convertion.
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
var availableInt = map[string]bool{ // Same as roman, except convertion
	"1":  true,
	"2":  true,
	"3":  true,
	"4":  true,
	"5":  true,
	"6":  true,
	"7":  true,
	"8":  true,
	"9":  true,
	"10": true,
}

func main() {
	reader := bufio.NewReader(os.Stdin) //Invoking a read from the console, based on OS standart input
	fmt.Println(`Enter mathematical operation in following format: "a + b, a - b, a * b, a / b".
Available numbers (1,2,3,4,5… 10), (I,II,III,IV,V… X)
Available operations: +, -, *, /.
Valid operations are following;
Example 1): 1 + 2
Example 2): I + VI
Example 3): 8 - 2
Invalid operations;
Example 1): I - X
Example 2): 1 + V
Example 3): 
Note: Roman numbers don't have negatives, I - V would result in error.'`)
	text, _ := reader.ReadString('\n')         //Awaiting input, string type
	text = strings.TrimSpace(text)             //Removing \n from input ((Tbh idk how it works and why it doesn't remove spaces??))
	operationSlice := strings.Split(text, " ") //Making a slice, uses "spaces" as separator
	if len(operationSlice) != 3 {              //Checks if you entered correct amount of arguments. 2 numbers and 1 operator
		panic("Incorrect amount of arguments.")
	}

	num1 := operationSlice[0] //Assigning numbers and operators after slicing the text.
	num2 := operationSlice[2]
	operator := operationSlice[1]

	_, hasRoman1 := availableRoman[num1] //Checks availability of entered numbers in maps, after slice.
	_, hasRoman2 := availableRoman[num2]
	_, hasInt1 := availableInt[num1]
	_, hasInt2 := availableInt[num2]

	if hasRoman1 && hasRoman2 { // Roman calculation part. (Operator is checked in Switch)
		fmt.Println("Result: ", GetRomanResults(num1, num2, operator))
	} else if hasInt1 && hasInt2 {
		convInt1, _ := strconv.Atoi(num1) // Int calculation part. I didn't use convertion then " if num > 0 && num < 11 " as checker,
		convInt2, _ := strconv.Atoi(num2) // cuz only 10 numbers allowed, so I used a map instead.
		fmt.Println("Result: ", getIntResult(convInt1, convInt2, operator))

	} else {
		panic("Invalid operation provided, be sure to follow the instructions.")
	}
}
