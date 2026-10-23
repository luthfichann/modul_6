package main
import "fmt"

func main() {
	var a, b, i int
	var hasil int = 1

	fmt.Scan(&a, &b)
	for i = 1; i <= b; i++ {
		hasil *= a
	}
	fmt.Println(hasil)
}

// PSEUDOCODE

// program pangkat
// kamus
// 	a, b, i : int
// 	hasil : int <- 1
// algoritma
// 	input(a, b)
// 	for i <- 1 to n do
// 		hasil = hasil * a
// 	endfor
// 	output(hasil)
// endprogram