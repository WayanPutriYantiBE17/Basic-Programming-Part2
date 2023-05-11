package main

import (
	"fmt"
	"math"
)

func PrimeNumber(number int) bool {
	// your code here

	if number<2{
		return false
	}
	sqrtN :=int(math.Sqrt(float64(number)))
	for i:=2;i<=sqrtN;i++{
		if number%i== 0{
		return false
}
}
return true
}


func main() {
	fmt.Println(PrimeNumber(3))
	fmt.Println(PrimeNumber(7))
	fmt.Println(PrimeNumber(10))
}
