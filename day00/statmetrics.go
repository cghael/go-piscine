package main

import (
	"fmt"
	"math"
	"sort"
)

func ftCountMean(numbers []int) {
	var sum float32
	for _, v := range numbers {
		sum += float32(v)
	}
	var mean = sum / float32(len(numbers))
	fmt.Printf("Mean: %.2f\n", mean)
}

func ftCountMedian(numbers []int) {
	switch len(numbers) % 2 {
	case 0:
		fmt.Printf("Median: %.2f\n", float32(numbers[len(numbers)/2]+
			numbers[len(numbers)/2-1])/2)
	case 1:
		fmt.Println("Median:", numbers[len(numbers)/2])
	}
}

func ftCountMode(numbers []int) {
	var frq int
	var mode int

	m := make(map[int]int)
	for _, value := range numbers {
		m[value]++
	}
	for key, value := range m {
		if frq < value {
			frq = value
			mode = key
		}
	}
	fmt.Printf("Mode: %v\n", mode)
}

func ftCountSD(numbers []int) {
	var sum float64
	var mean float64
	var sumSqr float64

	for _, v := range numbers {
		sum += float64(v)
	}
	mean = sum / float64(len(numbers))
	for _, v := range numbers {
		sumSqr += math.Pow(float64(v)-mean, 2)
	}
	fmt.Printf("SD: %.2f\n", math.Sqrt(sumSqr/float64(len(numbers))))
}

func ftCountMetrics(args []string, numbers []int) {
	sort.Ints(numbers)
	for _, v := range args {
		switch v {
		case "-mn":
			ftCountMean(numbers)
		case "-md":
			ftCountMedian(numbers)
		case "-mo":
			ftCountMode(numbers)
		case "-sd":
			ftCountSD(numbers)
		}
	}
}
