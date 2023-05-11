package main

import (
	"fmt"
)

func FaktorBilangan(n int) string {
    outp := []int{}
    for i := 1; i <= n; i++ {
        if n%i == 0 {
            outp = append(outp, i)
        }
    }
    output := ""
    for _, factor := range outp {
        output += fmt.Sprintf("%d\n", factor)
    }
    return output
}

func main() {
	var number int
	fmt.Scanf("%d", &number)
	fmt.Println(FaktorBilangan(number))
}
