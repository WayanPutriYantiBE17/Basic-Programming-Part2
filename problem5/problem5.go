package main

import (
	"fmt"
	"math"
)

func Pangkat(base, pangkat int) int {
	// your code here
	hasil:=math.Pow(float64(base),float64(pangkat))

	result := int(hasil)
	return result

}

func main() {
	fmt.Println(Pangkat(2, 3))  // 8
	fmt.Println(Pangkat(5, 3))  // 125
	fmt.Println(Pangkat(10, 2)) // 100
	fmt.Println(Pangkat(2, 5))  // 32
	fmt.Println(Pangkat(7, 3))  // 343
}
