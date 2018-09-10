package perf

import (
	"time"
)

// TimeSumFunc runs the function f on the provided arguments and returns its output and the elapsed time.
func TimeSumFunc(f func([]int, int) int, data []int, iterations int) (sum int, d time.Duration) {
	start := time.Now()
	sum = f(data, iterations)
	d = time.Since(start)
	return sum, d
}
