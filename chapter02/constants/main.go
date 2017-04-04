package main

import "fmt"

// declaring a constant
const p string = "const string"
// multi declaration of constants
const (
	cPi = 3.14
	cLanguage = "en-GB"
	cA = iota  // 0 iota small amount every time called it increments within the confines of the const(parenthesis)
	cB = iota  // 1
	cC = iota  // 2
)

func main() {
	const q = 4
	fmt.Println(q, "\n", p)
	fmt.Println(cPi, "\n", cLanguage)
	fmt.Println(cA,cB,cC)
	bitwise()
}

func bitwise() {
	const( _ = iota  // 0
	 KB = 1 << (iota * 10) // 1 << (1 * 10) shifts the bit to left 10 places 2^10
	 MB = 1 << (iota * 10))  // 1 << (2 * 10) shifts the bit left 20 places 2^20

	fmt.Printf("KB - %b - %d\n", KB, KB)
	fmt.Printf("MB - %b - %d\n", MB, MB)

}