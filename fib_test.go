package main

import "testing"

func TestFibonacci(t *testing.T) {
	const (
		index = 7
		result = 13
	)
	v := fibonacci(index)
	if v != result {
		t.Errorf("fibonachi(%d) = %d; want %d", index, v, result)
	}
}
