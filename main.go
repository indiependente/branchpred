package main

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
	"time"

	"github.com/indiependente/branchpred/perf"
	"github.com/indiependente/branchpred/processor"
)

// ELMNTS is the number of elements in the slice.
const ELMNTS = 32768

// ITRTNS is the number of iterations needed to perform the test.
const ITRTNS = 100000

func main() {
	done := make(chan struct{})
	data := genData(ELMNTS)
	sort.Ints(data)
	var goroutines = 2

	go runAndPrint("ConditionalSum", processor.ConditionalSum, data, ITRTNS, done)
	go runAndPrint("ConditionalSumNoBranching", processor.ConditionalSumNoBranching, data, ITRTNS, done)

	for goroutines > 0 {
		select {
		case <-done:
			goroutines--
		case <-time.NewTimer(100 * time.Second).C:
			fmt.Fprintf(os.Stderr, "Timed out after 100s")
		}
	}

}

func genData(n int) []int {
	rand.Seed(time.Now().UnixNano())
	data := make([]int, n)
	for i := 0; i < n; i++ {
		data[i] = rand.Intn(256)
	}
	return data
}

func runAndPrint(name string, f func([]int, int) int, data []int, iterations int, done chan struct{}) {
	defer func() { done <- struct{}{} }()

	sum, d := perf.TimeSumFunc(f, data, iterations)
	fmt.Printf("%s = %d.\tElapsed time = %.2fs\n", name, sum, d.Seconds())
}
