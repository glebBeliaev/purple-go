package main

import (
	"fmt"
	"math/rand"
)

func main() {
	numCh := make(chan int)
	squareCh := make(chan int)

	go randomaizer(numCh)
	go square(squareCh, numCh)

	results := saveResults(squareCh)
	printResults(results)
}

func randomaizer(numCh chan int) {
	nums := make([]int, 10)
	for i := 0; i < 10; i++ {
		nums[i] = rand.Intn(101) // от 0 до 100
	}
	for _, n := range nums {
		numCh <- n
	}
	close(numCh)
}

func square(squareCh, numCh chan int) {
	for n := range numCh {
		squareCh <- n * n
	}
	close(squareCh)
}

func saveResults(squareCh chan int) []int {
	var results []int
	for sq := range squareCh {
		results = append(results, sq)
	}
	return results
}

func printResults(results []int) {
	fmt.Println("Все квадраты получены:")
	for i, v := range results {
		fmt.Printf("%d) %d\n", i+1, v)
	}
}
