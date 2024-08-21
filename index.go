package main

import (
    "fmt"
    "sort"
)

// A function to calculate the sum of squares of numbers in a slice
func sumOfSquares(numbers []int) int {
    sum := 0
    for _, number := range numbers {
        sum += number * number
    }
    return sum
}

// A function to calculate the mean of numbers in a slice
func mean(numbers []int) float64 {
    sum := 0
    for _, number := range numbers {
        sum += number
    }
    return float64(sum) / float64(len(numbers))
}

// A function to calculate the median of numbers in a slice
func median(numbers []int) float64 {
    sortedNumbers := make([]int, len(numbers))
    copy(sortedNumbers, numbers)
    sort.Ints(sortedNumbers)

    n := len(sortedNumbers)
    if n%2 == 0 {
        return float64(sortedNumbers[n/2-1]+sortedNumbers[n/2]) / 2
    } else {
        return float64(sortedNumbers[n/2])
    }
}

func main() {
    numbers := []int{1, 2, 3, 4, 5}
    fmt.Printf("Sum of squares: %d\n", sumOfSquares(numbers))
    fmt.Printf("Mean: %.2f\n", mean(numbers))
    fmt.Printf("Median: %.2f\n", median(numbers))
}
