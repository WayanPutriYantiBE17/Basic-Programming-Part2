package main

import (
	"fmt"
)

func FullPrima(n int) bool {
	if n < 2 {
		return false
	}

	// Cek apakah bilangan n merupakan bilangan prima
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			return false
		}
	}
	if n>9{
	// Memisahkan digit puluhan dan satuan dari bilangan n
	satuan := n % 10
	puluhan := n / 10

	// Cek apakah bilangan puluhan dan satuan juga bilangan prima
	if puluhan < 2 || !FullPrima(puluhan) {
		return false
	}
	if satuan < 2 || !FullPrima(satuan) {
		return false
	}
	}

	// Jika semua kondisi terpenuhi, bilangan n merupakan full prima
	return true
}


func main() {
	fmt.Println(FullPrima(2))  // true
	fmt.Println(FullPrima(3))  // true
	fmt.Println(FullPrima(11)) // false
	fmt.Println(FullPrima(13)) // false
	fmt.Println(FullPrima(23)) // true
	fmt.Println(FullPrima(29)) // false
	fmt.Println(FullPrima(37)) // true
	fmt.Println(FullPrima(41)) // false
	fmt.Println(FullPrima(43)) // false
	fmt.Println(FullPrima(53)) // true
}
