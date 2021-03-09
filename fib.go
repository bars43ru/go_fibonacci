package main

import (
	"errors"
	"fmt"
	"log"
	"math/big"
)

// uintMax максимальное значение для реализации uint
const uintMax = ^uint(0)

// ErrOverflow ошибка переполнения числа
var ErrOverflow = errors.New("uint overflow")
var ErrIndex = errors.New("the index starts at 1")

func main() {
	v, err := fibonacciRecursive(5)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("for 5 recursive algorithm", v)

	v, err = fibonacciLoop(70)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("for 70 loop algorithm", v)

	v2, err := fibonacciCalc(10000)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("for 10000 calc. algorithm", v2)
}

// fibonacciRecursive вычисление числа фибоначи для элемента index через рекурсию, где первые начальные значения в ряде 1, 1
func fibonacciRecursive(index uint) (uint, error) {
	if index == 0 {
		return 0, ErrIndex
	}
	if index <= 2 {
		return 1, nil
	}
	fib1, err := fibonacciRecursive(index - 1)
	if err != nil {
		return 0, err
	}
	fib2, err := fibonacciRecursive(index - 2)
	if err != nil {
		return 0, err
	}
	return add(fib1, fib2)
}

// fibonacciLoop вычисление числа фибоначи для элемента index через цикл, где первые начальные значения в ряде 1, 1
func fibonacciLoop(index uint) (uint, error) {

	if index == 0 {
		return 0, ErrIndex
	}

	var v1, v2 uint = 1, 1

	var i uint
	for i = 3; i <= index; i++ {
		v, err := add(v1, v2)
		if err != nil {
			return 0, err
		}
		v1, v2 = v2, v
	}

	return v2, nil
}

// fibonacciCalc вычисление числа фибоначи для элемента index
func fibonacciCalc(index uint) (*big.Int, error) {

	v, _ := fibonacci(index)
	return v, nil
}

func fibonacci(index uint) (*big.Int, *big.Int) {

	if index == 0 {
		return big.NewInt(0), big.NewInt(1)
	}

	a, b := fibonacci(index / 2)
	c := big.NewInt(0).Mul(a, big.NewInt(0).Sub(big.NewInt(0).Mul(b, big.NewInt(2)), a))
	d := big.NewInt(0).Add(big.NewInt(0).Mul(a, a), big.NewInt(0).Mul(b, b))

	if index%2 == 0 {
		return c, d
	} else {
		return d, big.NewInt(0).Add(c, d)
	}
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
