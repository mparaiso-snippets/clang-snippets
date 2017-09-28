/**
 * Addison.Wesley.Algorithms.4th.Edition.Mar.2011.ISBN.032157351X.pdf
 * page 128
 */
package main

import (
	"io"
	"os"
	"strconv"
	"strings"
)

type Any = interface{}

type Stack []Any

func (stack *Stack) Push(value Any) {
	*stack = append([]Any{value}, *stack...)
}
func (stack *Stack) Pop() (result Any) {
	if len(*stack) == 0 {
		return nil
	}
	result, *stack = (*stack)[0], (*stack)[1:]
	return result
}

func (stack *Stack) Length() int {
	return len((*stack))
}

const (
	expression = "(45 + (5*6))"
	numbers    = "0123456789"
)

func main() {
	operandStack := new(Stack)
	operatorStack := new(Stack)
	reader := strings.NewReader(expression)

	for { // while there are runes in the string
		c, _, err := reader.ReadRune()
		print("character", string(c), "\n")
		if err != nil {
			if err == io.EOF { /* if end of string break the loop */
				break

			} else { /* else exit program with error */
				print("Error reading expression\n", err.String())
				os.Exit(1)
			}
		}

		if strings.IndexRune(numbers, c) > -1 { // if rune is a digit
			if number, err := strconv.ParseFloat(string(c), 64); err != nil {
				print("Error converting ", string(c), " to a number\n")
				os.Exit(1)
			} else {
				operandStack.Push(number) // push to operand stack
			}
		} else if c == '*' || c == '/' || c == '-' || c == '+' { // if run is an operator
			operatorStack.Push(c) // push to operator stack
		} else if c == ')' {
			second := operandStack.Pop().(float64)
			first := operandStack.Pop().(float64)
			operator := operatorStack.Pop().(rune)
			switch operator {
			case '*':
				operandStack.Push(first * second)
			case '/':
				operandStack.Push(first / second)
			case '+':
				operandStack.Push(first + second)
			case '-':
				operandStack.Push(first - second)
			}
		} else if c != ' ' {
			print("Error illegal character ", c, "\n")
		}
	}
	print("operandStack.Length() : ", operandStack.Length(), "\n", "operatorStack.Length() : ", operatorStack.Length(), "\n")
	if operandStack.Length() > 1 || operatorStack.Length() > 0 {
		print("Malformed mathematical expression : " + expression + "\n")
		os.Exit(1)
	} else {
		print("result for : "+expression+" = ", strconv.FormatFloat(operandStack.Pop().(float64), 'f', -1, 64))
		os.Exit(0)
	}
}
