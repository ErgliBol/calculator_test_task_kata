package main

import (
	"bufio"   // пакет чтения ввода
	"fmt"     // пакет форматированного вывода
	"os"      // пакет работы с операционной системой
	"strconv" // пакет конвертации строк в числа и наоборот
	"strings" // пакет работы со строками
)

// карта преобразования римских чисел в арабские
var roman = map[string]int{
	"C": 100, "XC": 90, "L": 50, "XL": 40, "X": 10, "IX": 9,
	"VIII": 8, "VII": 7, "VI": 6, "V": 5, "IV": 4, "III": 3, "II": 2, "I": 1,
}

// массивы обратного преобразования арабских чисел в римские
var convIntToRoman = []int{100, 90, 50, 40, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
var romanSymbols = []string{"C", "XC", "L", "XL", "X", "IX", "VIII", "VII", "VI", "V", "IV", "III", "II", "I"}

// каждый оператор ассоциирован с соответствующей функцией
var operators = map[string]func(int, int) int{
	"+": func(a, b int) int { return a + b },
	"-": func(a, b int) int { return a - b },
	"/": func(a, b int) int { return a / b },
	"*": func(a, b int) int { return a * b },
}

// константы хранения сообщений об ошибках
const (
	LOW    = "Ошибка: строка не является математической операцией."
	HIGH   = "Ошибка: формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *)."
	SCALE  = "Ошибка: используются одновременно разные системы счисления."
	DIV    = "Ошибка: в римской системе нет отрицательных чисел."
	ZERO   = "Ошибка: в римской системе нет числа 0."
	RANGE  = "Ошибка: калькулятор работает только с числами от 1 до 10 включительно."
	NEGNUM = "Ошибка: калькулятор не работает с отрицательными числами."
)

// основная функция, принимает строку и выполняет соответствующую операцию
func base(s string) {
	var operator string           // переменная хранения оператора
	var stringsFound int          // счётчик строк, найденных в выражении
	numbers := make([]int, 0)     // слайс хранения арабских чисел
	romans := make([]string, 0)   // слайс хранения римских чисел
	romansToInt := make([]int, 0) // слайс хранения римских чисел в виде арабских

	// определяем оператор и разделяем строку на две части
	for idx := range operators {
		if strings.Contains(s, idx) {
			operator = idx
			data := strings.Split(s, operator)
			if len(data) != 2 {
				panic(HIGH) // если количество операндов не равно двум, выдаём ошибку
			}
			for _, elem := range data {
				num, err := strconv.Atoi(elem) // попытка преобразовать элемент в число
				if err != nil {
					stringsFound++ // если преобразование не удалось, значит это римское число
					romans = append(romans, elem)
				} else {
					numbers = append(numbers, num) // добавляем арабское число в соответствующий слайс
				}
			}
			break
		}
	}

	// обрабатка различных случаев на основе найденных строк и оператора
	switch {
	case operator == "":
		panic(LOW) // если оператор не найден, выдаём ошибку
	case stringsFound == 1:
		panic(SCALE) // если найдены как римские, так и арабские числа, выдаём ошибку
	case stringsFound == 0:
		// если оба числа арабские
		if len(numbers) != 2 || numbers[0] < 1 || numbers[0] > 10 || numbers[1] < 1 || numbers[1] > 10 {
			panic(RANGE) // если числа не в диапазоне от 1 до 10, выдаём ошибку
		}
		result := operators[operator](numbers[0], numbers[1]) // вычисляем результат операции
		if result < 0 {
			panic(NEGNUM) // если результат отрицательный, выдаём ошибку
		}
		fmt.Println(result) // выводим результат
	case stringsFound == 2:
		// если оба числа римские
		for _, elem := range romans {
			if val, ok := roman[elem]; ok && val > 0 && val < 11 {
				romansToInt = append(romansToInt, val) // преобразуем римские числа в арабские
			} else {
				panic(RANGE) // если римское число не в диапазоне от 1 до 10, выдаём ошибку
			}
		}
		result := operators[operator](romansToInt[0], romansToInt[1]) // вычисляем результат операции
		if result < 1 {
			panic(DIV) // если результат меньше 1, выдаём ошибку
		}
		fmt.Println(intToRoman(result)) // выводим результат в виде римского числа
	}
}

// Ффнкция для преобразования арабского числа в римское
func intToRoman(num int) string {
	var romanNum strings.Builder // используем strings.Builder для построения строки
	for i := 0; num > 0; i++ {
		for num >= convIntToRoman[i] {
			num -= convIntToRoman[i]              // вычитаем значение римского числа
			romanNum.WriteString(romanSymbols[i]) // добавляем соответствующий символ к результату
		}
	}
	return romanNum.String() // возвращаем строку с римским числом
}

// основная функция программы
func main() {
	fmt.Println("Welcome to kata-calculator")
	fmt.Println("Допустимые операции: +, -, *, /")
	fmt.Println("Работает с числами от 1 до 10 включительно (арабские и римские числа)")
	reader := bufio.NewReader(os.Stdin) // создаём объект для чтения ввода
	for {
		console, _ := reader.ReadString('\n') // читаем строку из ввода
		if strings.Count(console, "\n") > 1 {
			panic(HIGH) // если введено больше одной строки, выдаём ошибку
		}
		s := strings.ReplaceAll(console, " ", "")   // удаляем пробелы из строки
		base(strings.ToUpper(strings.TrimSpace(s))) // преобразуем строку в верхний регистр и убираем пробелы по краям, передаём в функцию base
	}
}
