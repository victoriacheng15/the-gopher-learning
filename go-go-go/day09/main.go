package main

// Day 09: Printf integer formatting verbs
// %b base2
// %d base10, %+d base10 with sign
// %o base8, %O base8 with with leading 0o
// %x base16, lowercase, %X base16, uppercase, %#x base16 with leading 0x
// %4d pad with spaces to 4 characters, right justified
// %-4d pad with spaces to 4 characters, left justified
// %04d pad with 0s to 4 characters

import "fmt"

func main() {
	i := 10

	fmt.Printf("base2 - %b\n", i)
	fmt.Printf("base10 - %d\n", i)
	fmt.Printf("base10 with a sign - %+d\n", i)
	fmt.Printf("base8 - %o\n", i)
	fmt.Printf("base8 with leading 0o - %O\n", i)
	fmt.Printf("base16 lowercase - %x\n", i)
	fmt.Printf("base16 uppercase - %X\n", i)
	fmt.Printf("base16 with leading 0x - %#x\n", i)
	fmt.Printf("pad with right spaces - %4d\n", i)
	fmt.Printf("pad with left spaces - %-4d\n", i)
	fmt.Printf("pad with right zeros - %04d\n", i)
}
