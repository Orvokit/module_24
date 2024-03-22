//Напишите функцию, которая на вход принимает массив предложений (длинных строк) и массив символов типа rune,
//а возвращает 2D-массив, где на позиции [i][j] стоит индекс вхождения символа j из chars
//в последнее слово в предложении i (строку надо разбить на слова и взять последнее).
//То есть сигнатура следующая:
//func parseTest(sentences []string, chars []rune)

package main

import (
	"fmt"
	"strings"
)

func parseTest(sentences []string, charts []rune) [3][3]int {
	var s []string
	var a [3][3]int

	for i := 0; i < len(sentences); i++ {
		s = strings.Split(sentences[i], " ")
		fmt.Println("В каком предложении сейчас ищем:", s)
		size := len(s)
		fmt.Println("В каком слове сейчас ищем:", s[size-1])
		for j := 0; j < len(charts); j++ {
			fmt.Println("Какой символ ищем:", string(charts[j]))
			a[i][j] = strings.Index(s[size-1], string(charts[j]))
			if a[i][j] != -1 {
				a[i][j] = a[i][j] / 2
			}
			fmt.Println("Полученный индекс:", a[i][j])
		}
	}
	return a
}

func main() {
	sentences := []string{"Как дышит грудь свежо и ёмко — слова не выразят ничьи!",
		"Как по оврагам в полдень громко на пену прядают ручьи!",
		"В эфире песнь дрожит и тает, на глыбе зеленеет рожь — и голос нежный напевает: «Еще весну переживешь!»"}
	charts := []rune{'и', 'у', 'ш'}

	fmt.Println("Исходный текст:")
	for _, v := range sentences {
		fmt.Println(v)
	}
	fmt.Println(parseTest(sentences, charts))
}
