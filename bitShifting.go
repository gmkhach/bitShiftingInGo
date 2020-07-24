package main

import "fmt"

func main() {
	x := 2
	fmt.Printf("Decimal: %d\tBinary: %b\n", x, x)

	// In binary 2 will be 0000 0010
	// After shifting all of the bits by one to the left we get 0000 0100, which is decimal 4
	y := x << 1
	fmt.Printf("Decimal: %d\tBinary: %b\n", y, y)

	// The binary value of the following sequence changes by shifting over to the left 10 times
	b := 1
	kb := 1024
	mb := 1024 * kb
	gb := 1024 * mb

	fmt.Printf("kilobyte:\t%d\t\t%b\n", b, b)
	fmt.Printf("kilobyte:\t%d\t\t%b\n", kb, kb)
	fmt.Printf("megabyte:\t%d\t\t%b\n", mb, mb)
	fmt.Printf("gigabyte:\t%d\t%b\n", gb, gb)

	// We can acomplish the same thing using iota
	const (
		// byte is a reserved word, so we will use bite
		bite     = 1 << (iota * 10)
		kilobyte = 1 << (iota * 10)
		megabyte = 1 << (iota * 10)
		gigabyte = 1 << (iota * 10)
	)

	fmt.Printf("kilobyte:\t%d\t\t%b\n", bite, bite)
	fmt.Printf("kilobyte:\t%d\t\t%b\n", kilobyte, kilobyte)
	fmt.Printf("megabyte:\t%d\t\t%b\n", megabyte, megabyte)
	fmt.Printf("gigabyte:\t%d\t%b\n", gigabyte, gigabyte)

}
