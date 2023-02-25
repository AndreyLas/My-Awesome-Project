package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	var operator string
	var result, arabic1, arabic2 int
	var reader = bufio.NewReader(os.Stdin)

	rome := map[string]int{
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

	fmt.Println("Введите знаечние")

	text, _ := reader.ReadString('\n') // ждет ввода данных (строка)

	for i := 0; i < len(text); i++ { // определить знак оператора в строке
		if string(text[i]) == "+" || string(text[i]) == "-" || string(text[i]) == "*" || string(text[i]) == "/" {
			operator = string(text[i])
		}
	}

	splitFunc := func(r rune) bool {
		return strings.ContainsRune("/*-+", r) //отделить строку от операторов

	}

	digit := strings.FieldsFunc(text, splitFunc)

	if len(digit) > 2 { // проверка на колличесво значение в введенной строке
		fmt.Println("Ошибка, формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
		return
	} else if len(digit) < 2 {
		fmt.Println("Ошибка, строка не является математической операцией.")
		return
	}

	digit[0] = strings.TrimSpace(digit[0]) // удалить пробелы для первого числа
	num1, _ := strconv.Atoi(digit[0])      // конвертировать первое значение из строки в число

	digit[1] = strings.TrimSpace(digit[1]) // удалить пробелы для второго числа
	num2, _ := strconv.Atoi(digit[1])      // конвертировать второе значение из строки в число

	if (num1 == 0 && num2 != 0) || (num1 != 0 && num2 == 0) { //проверка систем счисления
		fmt.Println("Ошибка, недопустимое выражение")
		return
	} else if num1 == 0 && num2 == 0 { // определена римская система счисления
		if key, ok := rome[digit[0]]; !ok { // сравнение первого значения с ключом в карте
			fmt.Println("Ошибка, недопустимое первое число")
			return
		} else {
			arabic1 = key
		}

		if key, ok := rome[digit[1]]; !ok { // // сравнение второго значения с ключом в карте
			fmt.Println("Ошибка, недопустимое второе число")
			return
		} else {

			arabic2 = key
		}

		if operator == "+" {
			result = arabic1 + arabic2
		} else if operator == "-" {
			result = arabic1 - arabic2
		} else if operator == "*" {
			result = arabic1 * arabic2
		} else if operator == "/" {
			result = arabic1 / arabic2
		}

		if result <= 0 {
			fmt.Println("Ошибка, недопустимое выражение, в римскоой системе счисления отсутсвуют отрицательные числа и ноль")
			return
		}

		fmt.Println(romanToArabic(result))
		return

	}
	if num1 < 1 || num1 > 10 {
		fmt.Println("Ошибка, недопустимый диапазон")
		return
	} else if num2 < 1 || num2 > 10 {
		fmt.Println("Ошибка, недопустимый диапазон")
		return
	}

	if operator == "+" {
		result = num1 + num2
	} else if operator == "-" {
		result = num1 - num2
	} else if operator == "*" {
		result = num1 * num2
	} else if operator == "/" {
		result = num1 / num2
	}

	fmt.Println(result)
}

func romanToArabic(num int) string {
	var roman = ""
	var numbers = []int{1, 4, 5, 9, 10, 40, 50, 90, 100}
	var romans = []string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C"}
	var index = len(romans) - 1

	for num > 0 {
		for numbers[index] <= num {
			roman += romans[index]
			num -= numbers[index]
		}
		index -= 1
	}

	return roman
}
