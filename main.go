package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

var (
	//Мапы арабских и римских цифр для конвертации полученной строки в int
	arabicNumerals = map[string]int{"-10": -10, "-9": -9, "-8": -8, "-7": -7, "-6": -6, "-5": -5, "-4": -4, "-3": -3, "-2": -2, "-1": -1, "1": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9, "10": 10}
	romanNumerals  = map[string]int{"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10}
	//Структура для записи ответа римскими цифрами
	reverseRomanNumerals = []struct {
		Value  int
		Symbol string
	}{
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}
)

// Непосредственное выполнение математической операциии
func calculation(operand1, operand2 int, operation string) (result int) {
	switch {
	case operation == "+":
		result = operand1 + operand2
	case operation == "-":
		result = operand1 - operand2
	case operation == "*":
		result = operand1 * operand2
	case operation == "/":
		result = operand1 / operand2
	default:
		fmt.Println(errors.New("error: invalid mathematical expression"))
		os.Exit(1)
	}
	return result
}

// Конвертация строк с арабскими числами, полученных при вводе, в int для передачи в функцию calculation
func arabicCalculation(stringOperand1, stringOperand2, operation string) (result int) {
	operand1 := arabicNumerals[stringOperand1]
	operand2 := arabicNumerals[stringOperand2]
	if operand1 == 0 || operand2 == 0 {
		fmt.Println(errors.New("error: number is out of range or input format is invalid"))
		os.Exit(1)
	}
	result = calculation(operand1, operand2, operation)
	return result
}

// Конвертация чисел int в римские числа
func intToRoman(num int) string {
	result := ""
	for _, numeral := range reverseRomanNumerals {
		for num >= numeral.Value {
			result += numeral.Symbol
			num -= numeral.Value
		}
	}
	return result
}

// Конвертация строк с римскими числами, полученных при вводе, в int для передачи в функцию calculation
func romanCalculation(romanOperand1, romanOperand2, operation string) string {
	operand1 := romanNumerals[romanOperand1]
	operand2 := romanNumerals[romanOperand2]
	if operand1 == 0 || operand2 == 0 {
		fmt.Println(errors.New("error: number is out of range or input format is invalid"))
		os.Exit(1)
	}
	arabicResult := calculation(operand1, operand2, operation)
	if arabicResult <= 0 {
		fmt.Println(errors.New("error: result is out of range. roman numerals cannot be less than one"))
		os.Exit(1)
	}
	return intToRoman(arabicResult)
}

func main() {
	var operand1, operand2, operation string
	reader := bufio.NewReader(os.Stdin) //Считывание строки из терминала вместе с пробелами
	fmt.Println("Type an mathematical expression:")
	expression, _ := reader.ReadString('\n')
	expression = strings.TrimSpace(expression) //Удаление пробелов и другие символов из начала и конца строки
	//Считывание операндов и символа математической операции из полученной строки
	_, err := fmt.Sscanf(expression, "%s %s %s\n", &operand1, &operation, &operand2)
	if err != nil {
		fmt.Println(errors.New("error: invalid input format (example: 2 * 2)"))
		return
	}
	//Анализ имеющегося математического выражения
	if arabicNumerals[operand1] != 0 && arabicNumerals[operand2] != 0 { //Проверка строк на написание арабскими цифрами
		result := arabicCalculation(operand1, operand2, operation)
		fmt.Println("RESULT:", result) //Вывод результата математического выражения с арабскими числами
	} else {
		operand1 = strings.ToUpper(operand1) //Приведение строк к стандартному представлению римских чисел заглавными буквами
		operand2 = strings.ToUpper(operand2)
		if romanNumerals[operand1] != 0 && romanNumerals[operand2] != 0 { //Проверка строк на написание римскими цифрами
			result := romanCalculation(operand1, operand2, operation)
			fmt.Println("RESULT:", result) //Вывод результата математического выражения с римскими числами
		} else {
			fmt.Println(errors.New("error: invalid operand. number is out of range or belongs to another numeral system"))
			os.Exit(1)
		}
	}
}
