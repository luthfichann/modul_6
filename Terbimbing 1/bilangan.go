package main
import "fmt"

func main() {
	var a, b, j int

	fmt.Scan(&a)
	fmt.Scan(&b)
	for j = a; j <= b; j++ {
		fmt.Print(j, " ")
	} 
}

// PSEUDOCODE

// program bilangan
// kamus
// 	a, b, j : int
// algoritma
// 	input(a)
// 	input(b)
// 	for j <- a to j <= b do
// 		output(j, " ")
// 	endfor
// endprogram

