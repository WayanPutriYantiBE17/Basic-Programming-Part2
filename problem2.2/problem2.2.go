package main

import (
	"fmt"
)

func FaktorBilanganDesc(n int) string {
	// your code here
    outp := []int{}
    for i := n; i >=1; i-- {
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
	fmt.Println(FaktorBilanganDesc(number))
}
