package main

import "fmt"

func main()  {
	var n float64
	var e float64 = 0
	fmt.Scan(&n)
	for i:=0.0; i<n; i++{
		fact := 1.0
		for j:=1.0; j<=i; j++{
			fact *= j
		}
		e += 1.0 /fact
	}

	fmt.Print(e)
}