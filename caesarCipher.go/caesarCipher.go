package main

import (
	"fmt"
	"strings"
)
//Руна- это символы(их всего 128)!
// Мы используем только 26, тоесть английский алфавит, так как функция шифрования методом цезаря имеет только 26 символов!
func caesar(r rune, shift int) rune {
	s := int(r) + shift
	if s > 'z' {
		return rune(s - 26)
	} else if s < 'a' {
		return rune(s + 26)
	}
	return rune(s)
}

func main() {
	// Обьявляем переменную value!
	value := "Cipher"
	fmt.Println(value)
	value = "lippszmxce"
	//Принимаем rune через Map!
	//Возвращаем руну со сдвигом влево на 4 по символам в значении value!
	result := strings.Map(func(r rune) rune {
		return caesar(r, -4)
	}, value)
	fmt.Println(value, result)
}