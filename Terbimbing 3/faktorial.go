package main
import "fmt"

func main() {
	var a, i int
	var hasil int = 1
	
	fmt.Scan(&a)
	for i = 1; i <= a; a-=1 {
		hasil *= a
	}
	fmt.Println(hasil)
}

// PSEUDOCODE

// program faktorial
// kamus
// 	a, i : int
// 	hasil : int <- 1
// algoritma
// 	input(a, b)
// 	for i <- 1 to n do
// 		hasil = hasil * a
// 	endfor
// 	output(hasil)
// endprogram