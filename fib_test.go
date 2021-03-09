package main

import (
	"errors"
	"testing"
)

func TestFibonacci(t *testing.T) {
	const (
		index  = 7
		result = 13
	)
	v, err := fibonacci(index)
	if v != result || err != nil {
		t.Errorf("fibonachi(%d) = %d, %v; want %d, nil", index, err, v, result)
	}
}

func TestAdd(t *testing.T) {
	v, err := add(2, 4)
	if !(err == nil && v == 6) {
		t.Errorf("add(2, 4) = %d, %v; want 6", v, err)
		return
	}

	v, err = add(uintMax-5, 6)
	if !(errors.Is(err, ErrOverflow) && v == 0) {
		t.Errorf("add(%d, 6) = 0, %v; want 0, %v", uintMax-5, err, ErrOverflow)
		return
	}
}
