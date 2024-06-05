package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Карты для преобразования римских чисел в арабские и наоборот
var romanToArabic = map[string]int{
	"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
}

var arabicToRoman = map[int]string{
	1: "I", 2: "II", 3: "III", 4: "IV", 5: "V", 6: "VI", 7: "VII", 8: "VIII", 9: "IX", 10: "X",
}

// Главная функция - точка входа в программу
func main() {
	defer func() { // Декларация отложенной функции для обработки паники
		if r := recover(); r != nil {
			fmt.Println("Произошла ошибка:", r)
		}
	}()

	scanner := bufio.NewScanner(os.Stdin) // Создаем новый сканер для чтения ввода с консоли
	fmt.Println("Введите выражение (например, 3+4 или VI*II):")
	for scanner.Scan() { // Запускаем бесконечный цикл для чтения ввода
		input := scanner.Text()           // Считываем введенную строку
		result := calculate(input)        // Вызываем функцию calculate для вычисления результата
		fmt.Println("Результат:", result) // Выводим результат
		fmt.Println("Введите выражение:") // Приглашаем пользователя ввести новое выражение
	}
}

// Функция для вычисления результата выражения
func calculate(input string) string {
	input = strings.TrimSpace(input) // Удаляем пробелы в начале и конце строки
	// Создаем регулярное выражение для парсинга входной строки
	re := regexp.MustCompile(`^\s*(\d+|I{1,3}|IV|VI{0,3}|IX|X)\s*([+\-*/])\s*(\d+|I{1,3}|IV|VI{0,3}|IX|X)\s*$`)
	matches := re.FindStringSubmatch(input)  // Ищем совпадения в строке
	if matches == nil || len(matches) != 4 { // Проверяем корректность ввода
		panic("неверный формат ввода")
	}

	a, op, b := matches[1], matches[2], matches[3] // Извлекаем числа и оператор из совпадений

	if isRoman(a) && isRoman(b) { // Если оба числа римские
		return calculateRoman(a, op, b) // Вызываем функцию для обработки римских чисел
	} else if isArabic(a) && isArabic(b) { // Если оба числа арабские
		return calculateArabic(a, op, b) // Вызываем функцию для обработки арабских чисел
	} else {
		panic("смешивание арабских и римских чисел") // Вызываем панику, если числа смешанные
	}
}

// Функция для проверки, является ли строка римским числом
func isRoman(s string) bool {
	_, ok := romanToArabic[s] // Проверяем наличие числа в карте римских чисел
	return ok
}

// Функция для проверки, является ли строка арабским числом
func isArabic(s string) bool {
	_, err := strconv.Atoi(s) // Пытаемся преобразовать строку в целое число
	return err == nil         // Возвращаем true, если преобразование успешно, иначе false
}

// Функция для обработки арифметических операций с римскими числами
func calculateRoman(a, op, b string) string {
	ai := romanToArabic[a] // Преобразуем римское число a в арабское
	bi := romanToArabic[b] // Преобразуем римское число b в арабское

	result := performOperation(ai, bi, op) // Выполняем арифметическую операцию

	if result < 1 { // Проверяем, чтобы результат был больше 1
		panic("результат меньше единицы")
	}

	return toRoman(result) // Преобразуем результат обратно в римское число
}

// Функция для обработки арифметических операций с арабскими числами
func calculateArabic(a, op, b string) string {
	ai, _ := strconv.Atoi(a) // Преобразуем строку a в целое число
	bi, _ := strconv.Atoi(b) // Преобразуем строку b в целое число

	if ai < 1 || ai > 10 || bi < 1 || bi > 10 { // Проверяем, чтобы числа были в диапазоне от 1 до 10
		panic("числа должны быть в диапазоне от 1 до 10")
	}

	result := performOperation(ai, bi, op) // Выполняем арифметическую операцию

	return strconv.Itoa(result) // Преобразуем результат обратно в строку
}

// Функция для выполнения арифметической операции над двумя числами
func performOperation(a, b int, op string) int {
	switch op {
	case "+": // Операция сложения
		return a + b
	case "-": // Операция вычитания
		return a - b
	case "*": // Операция умножения
		return a * b
	case "/": // Операция деления
		if b == 0 {
			panic("деление на ноль") // Вызываем панику при делении на ноль
		}
		return a / b
	default:
		panic("неверная операция") // Вызываем панику для неизвестных операций
	}
}

// Функция для преобразования арабского числа в римское
func toRoman(num int) string {
	if num < 1 { // Проверка на допустимость числа
		return ""
	}

	// Списки значений и соответствующих римских символов
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	var result strings.Builder // Используем strings.Builder для построения строки

	for i, value := range values { // Перебираем все значения
		for num >= value { // Пока num больше или равно текущему значению
			result.WriteString(symbols[i]) // Добавляем соответствующий римский символ к результату
			num -= value                   // Вычитаем значение из num
		}
	}

	return result.String() // Возвращаем результат в виде строки
}
