//Напишите функцию, которая принимает массив чисел, а возвращает два массива: один из чётных чисел, второй из нечётных.

package main

import (
	"fmt"
	"math/rand"
	"time"
)

const size = 10

func create() [size]int {
	var arr [size]int
	rand.NewSource(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		arr[i] = rand.Intn(100)
	}
	return arr
}

func findParity(arr [size]int) ([]int, []int) {
	even := make([]int, 0, 0) //четные
	odd := make([]int, 0, 0)  //нечетные

	for i := 0; i < size; i++ {
		if arr[i]%2 == 0 {
			even = append(even, arr[i])
		} else {
			odd = append(odd, arr[i])
		}
	}
	return even, odd
}

func main() {
	arr := create()
	fmt.Println("Сгенерированный массив", arr)
	even, odd := findParity(arr)
	fmt.Println("Массив четных чисел", even)
	fmt.Println("Массив нечетных чисел", odd)
}
