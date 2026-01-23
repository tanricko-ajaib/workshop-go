package fib

import (
	"fmt"
	"testing"
)

func Test_Fib(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{0, 0}, {1, 1}, {2, 1}, {3, 2}, {4, 3},
		{5, 5}, {6, 8}, {7, 13}, {8, 21}, {9, 34},
		{10, 55}, {15, 610}, {20, 6765}, {25, 75025},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.n), func(t *testing.T) {
			if got := Fib(tt.n); got != tt.want {
				t.Errorf("fib(%d) = %d; want %d", tt.n, got, tt.want)
			}
		})
	}
}
