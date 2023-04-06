package main

import "testing"

func TestSum(t *testing.T) {
	total := SumTwoValues(10, 10)

	if total != 20 {
		t.Errorf("Sum result is invalid. Result %d. Expected %d", total, 20)
	}
}
