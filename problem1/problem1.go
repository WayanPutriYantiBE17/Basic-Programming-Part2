package main

import "fmt"

func KonversiNilai(nilai int) string {
	// your code here
	huruf :=""
	if nilai >=80 && nilai<= 100 {
		huruf = "A"
	}else if nilai >=65 && nilai<80{
		huruf ="B+"
	}else if nilai>=50 && nilai <65{
		huruf ="B"
	}else if nilai>=35 && nilai<50{
		huruf ="C"
	}else if nilai<35{
		huruf ="D"
	}
	return huruf
}

func main() {
	var nilai int = 80

	fmt.Println(KonversiNilai(nilai))
}
