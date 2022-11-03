package ia64

import "fmt"

func IntegerLoadWithRegister(instructionBits uint64, nextSlot uint64) {
	tab_ := (instructionBits & (0b0000011110000000000000000000000000000000000000)) >> 37
	tab2 := (instructionBits & (0b0000000001100000000000000000000000000000000000)) >> 35
	hint := (instructionBits & (0b0000000000011000000000000000000000000000000000)) >> 33
	r3__ := (instructionBits & (0b0000000000000011111110000000000000000000000000)) >> 25
	r2__ := (instructionBits & (0b0000000000000000000001111111000000000000000000)) >> 18
	r1__ := (instructionBits & (0b0000000000000000000000000000111111100000000000)) >> 11
	qp__ := (instructionBits & (0b0000000000000000000000000000000000011111100000)) >> 5

	fmt.Printf("tab  : %d", tab_)
	fmt.Printf("tab2 : %d", tab2)
	fmt.Printf("hint : %d", hint)
	fmt.Printf("r1   : %d", r1__)
	fmt.Printf("r2   : %d", r2__)
	fmt.Printf("r3   : %d", r3__)
	fmt.Printf("qp   : %d", qp__)
}

func IntegerLoadStore(instructionBits uint64, nextSlot uint64) {
	tab_ := (instructionBits & (0b0000011110000000000000000000000000000000000000)) >> 37
	//                            0100011001000000111100011111000000000000000000
	tab2 := (instructionBits & (0b0000000001100000000000000000000000000000000000)) >> 35
	hint := (instructionBits & (0b0000000000011000000000000000000000000000000000)) >> 33
	r3__ := (instructionBits & (0b0000000000000011111110000000000000000000000000)) >> 25
	r2__ := (instructionBits & (0b0000000000000000000001111111000000000000000000)) >> 18
	qp__ := (instructionBits & (0b0000000000000000000000000000000000011111100000)) >> 5

	fmt.Printf("tab  : %d\n", tab_)
	fmt.Printf("tab2 : %d\n", tab2)
	fmt.Printf("hint : %d\n", hint)
	fmt.Printf("r2   : %d\n", r2__)
	fmt.Printf("r3   : %d\n", r3__)
	fmt.Printf("qp   : %d\n", qp__)
}
