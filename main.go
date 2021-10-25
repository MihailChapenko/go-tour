package main

import (
	"fmt"
	"math"
	"strings"
)

//Sqrt exercise loops and functions
func Sqrt(x float64) float64 {
	z := float64(1)
	for i := 1; i <= 10; i++ {
		fmt.Println(z)
		z -= (z*z - x) / (2 * z)
	}
	return z
}

//WordCount exercise maps
func WordCount(s string) map[string]int {
	result := make(map[string]int)
	words := strings.Fields(s)
	for _, word := range words {
		result[word] += 1
	}
	return result
}

//fibonacci
func fibonacci() func() int {
	total, nextTotal := 0, 1
	return func() int {
		result := total
		total, nextTotal = nextTotal, nextTotal+result
		return result
	}
}

//Pic exercise slices
func Pic(dx, dy int) [][]uint8 {
	var result = make([][]uint8, dy)
	for x := range result {
		result[x] = make([]uint8, dx)
		for y := range result[x] {
			result[x][y] = uint8(x * y)
		}
	}
	return result
}

func main() {
	//loops and functions
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))

	//maps
	//wc.Test(WordCount)

	//fibonacci
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}

	//slices
	//pic.Show(Pic)
}
