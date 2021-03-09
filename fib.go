package main

import (
	"errors"
	"fmt"
	"log"
)

// uintMax максимальное значение для реализации uint
const uintMax = ^uint(0)

// ErrOverflow ошибка переполнения числа
var ErrOverflow = errors.New("uint overflow")

func main() {
	v, err := fibonacci(30)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(v)
}

// fibonacci вычисление числа фибоначи для элемента index
func fibonacci(index uint) (uint, error) {
	if index <= 2 {
		return 1, nil
	}
	fib1, err := fibonacci(index - 1)
	if err != nil {
		return 0, err
	}
	fib2, err := fibonacci(index - 2)
	if err != nil {
		return 0, err
	}
	return add(fib1, fib2)
}

// add суммирование двух числе с проверкой на переполнение
func add(left, right uint) (uint, error) {

	if right > 0 {
		if left > uintMax-right {
			return 0, ErrOverflow
		}
	} else {
		if left < uintMax-right {
			return 0, ErrOverflow
		}
	}
	return left + right, nil
}
