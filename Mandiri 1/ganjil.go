package main
import "fmt"

func main() {
	var a, i, ganjil int

	fmt.Scan(&a)
	for i = 1; i <= a; i++ {
		ganjil = 2 * i - 1
		fmt.Print(ganjil, " ")
	}
}