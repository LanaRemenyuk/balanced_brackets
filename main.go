package main

import (
	"fmt"
)

func main() {
	fmt.Println(isBalanced("{(/)}()"))
}

func isBalanced(s string) string {
	// сразу проверяем, если в строке нечётное кол-во символов,то скобки несбалансированы
	if len(s)%2 != 0 {
		return "Скобки несбалансированы"
	}

	// создаём стек для хранения открывающих скобок
	stack := []rune{}
	// бежим по строке
	for _, bracket := range s {
		le := len(stack) - 1
		// если скобка открывающая, добавляем её в стек
		if bracket == '{' || bracket == '[' || bracket == '(' {
			stack = append(stack, bracket)
			continue
			/* проверяем не пуст ли стек и соответствуют ли
			закрывающая скобка открывающей с верхушки стека
			*/
		} else if len(stack) > 0 {
			topElem := stack[le]
			if topElem == '(' && bracket == ')' || topElem == '[' && bracket == ']' ||
				topElem == '{' && bracket == '}' {
				// снимаем верхний элемент
				stack = stack[:le]
			} else {
				// если закрывающая скобка не матчится с верхней не стеке
				return "Скобки несбалансированы"
			}
		} else {
			// если ввели не скобки
			return "Скобки несбалансированы"
		}
	}
	// проверяем пуст ли стек
	if len(stack) == 0 {
		return "Скобки сбалансированы"
	} else {
		return "Скобки несбалансированы"
	}
}
