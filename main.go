package main

import (
	"bufio"   //ввод/вывод буферизованных данных
	"fmt"     //форматированный ввод/вывод
	"os"      //доступ к операционной системе (например, чтения ввода)
	"strconv" //преобразование строк в другие типы данных
	"strings" //работа со строками
)

var roman = map[string]int{ // определение мапы для преобразования римских чисел в арабские
	"C": 100, "XC": 90, "L": 50, "XL": 40, "X": 10, "IX": 9,
	"VIII": 8, "VII": 7, "VI": 6, "V": 5, "IV": 4, "III": 3, "II": 2, "I": 1,
}

var convIntToRoman = []int{100, 90, 50, 40, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1} // Список арабских чисел для обратного преобразования
var romanSymbols = []string{"C", "XC", "L", "XL", "X", "IX", "VIII", "VII", "VI", "V", "IV", "III", "II", "I"}

var operators = map[string]func(int, int) int{ // определение мапы операторов и их функций
	"+": func(a, b int) int { return a + b },
	"-": func(a, b int) int { return a - b },
	"/": func(a, b int) int { return a / b },
	"*": func(a, b int) int { return a * b },
}

const ( // константы для сообщений об ошибках
	LOW    = "Ошибка: строка не является математической операцией."
	HIGH   = "Ошибка: формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *)."
	SCALE  = "Ошибка: используются одновременно разные системы счисления."
	DIV    = "Ошибка: в римской системе нет отрицательных чисел."
	ZERO   = "Ошибка: в римской системе нет числа 0."
	RANGE  = "Ошибка: калькулятор работает только с числами от 1 до 10 включительно."
	NEGNUM = "Ошибка: калькулятор не работает с отрицательными числами."
)

func base(s string) { //основная функция, выполняющая всю логику (определяет оператор, обрабатывает ошибки,выводит результат)
	var operator string
	var stringsFound int
	numbers := make([]int, 0)
	romans := make([]string, 0)
	romansToInt := make([]int, 0)

	for idx := range operators { // поиск оператора в строке и разделение строки на два операнда
		if strings.Contains(s, idx) {
			operator = idx
			data := strings.Split(s, operator)
			if len(data) != 2 {
				panic(HIGH)
			}
			for _, elem := range data {
				num, err := strconv.Atoi(elem)
				if err != nil {
					stringsFound++
					romans = append(romans, elem)
				} else {
					numbers = append(numbers, num)
				}
			}
			break
		}
	}

	switch { // обработка ошибок и выполнение операции
	case operator == "":
		panic(LOW)
	case stringsFound == 1:
		panic(SCALE)
	case stringsFound == 0:
		if len(numbers) != 2 || numbers[0] < 1 || numbers[0] > 10 || numbers[1] < 1 || numbers[1] > 10 {
			panic(RANGE)
		}
		result := operators[operator](numbers[0], numbers[1])
		if result < 0 {
			panic(NEGNUM)
		}
		fmt.Println(result)
	case stringsFound == 2:
		for _, elem := range romans {
			if val, ok := roman[elem]; ok && val > 0 && val < 11 {
				romansToInt = append(romansToInt, val)
			} else {
				panic(RANGE)
			}
		}
		result := operators[operator](romansToInt[0], romansToInt[1])
		if result < 0 {
			panic(DIV)
		}
		fmt.Println(intToRoman(result))
	}
}

func intToRoman(num int) string { // функция для преобразования арабских чисел в римские
	var romanNum strings.Builder
	for i := 0; num > 0; i++ {
		for num >= convIntToRoman[i] {
			num -= convIntToRoman[i]
			romanNum.WriteString(romanSymbols[i])
		}
	}
	return romanNum.String()
}

func main() { //за работу)
	fmt.Println("kataКалькулятор")
	fmt.Println("Допустимые операции: +, -, *, /")
	fmt.Println("Работает с числами от 1 до 10 включительно (арабские и римские числа)")
	reader := bufio.NewReader(os.Stdin)
	for {
		console, _ := reader.ReadString('\n')
		if strings.Count(console, "\n") > 1 {
			panic(HIGH)
		}
		s := strings.ReplaceAll(console, " ", "")
		base(strings.ToUpper(strings.TrimSpace(s)))
	}
}
