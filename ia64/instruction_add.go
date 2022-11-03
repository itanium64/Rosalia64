package ia64

import "fmt"

func AddImmediate22(instructionBits uint64, nextSlot uint64) {
	sign1 := (instructionBits & (0b0000100000000000000000000000000000000000000000)) >> 40
	imm9d := (instructionBits & (0b0000011111111100000000000000000000000000000000)) >> 32
	imm5c := (instructionBits & (0b0000000000000011111000000000000000000000000000)) >> 27
	r3___ := (instructionBits & (0b0000000000000000000110000000000000000000000000)) >> 25
	imm7b := (instructionBits & (0b0000000000000000000001111111000000000000000000)) >> 18
	r1___ := (instructionBits & (0b0000000000000000000000000000111111100000000000)) >> 11
	qp___ := (instructionBits & (0b0000000000000000000000000000000000011111100000)) >> 5

	immediate := Imm22(sign1, imm5c, imm9d, imm7b)

	fmt.Printf("\nExecuting: addl r1 = imm22, r3\n")
	fmt.Printf("sign     : %d\n", sign1)
	fmt.Printf("r1       : %d\n", r1___)
	fmt.Printf("r3       : %d\n", r3___)
	fmt.Printf("imm14    : %d\n", immediate)
	fmt.Printf("qp       : %d\n", qp___)

	if RetrievePredicateRegister(qp___) {
		r1 := RetrieveGeneralRegister(r1___)
		r3 := RetrieveGeneralRegister(r3___)

		r1.Value = uint64(immediate) + r3.Value
		r1.NotAThing = r3.NotAThing
	}
}

func AddImmediate14(instructionBits uint64, nextSlot uint64) {
	sign_ := (instructionBits & (0b0000100000000000000000000000000000000000000000)) >> 40
	imm6d := (instructionBits & (0b0000000011111100000000000000000000000000000000)) >> 32
	r3___ := (instructionBits & (0b0000000000000011111110000000000000000000000000)) >> 25
	imm7b := (instructionBits & (0b0000000000000000000001111111000000000000000000)) >> 18
	r1___ := (instructionBits & (0b0000000000000000000000000000111111100000000000)) >> 11
	qp___ := (instructionBits & (0b0000000000000000000000000000000000011111100000)) >> 5

	immediate := Imm14(sign_, imm6d, imm7b)

	fmt.Printf("\nExecuting: adds r1 = imm14, r3\n")
	fmt.Printf("sign     : %d\n", sign_)
	fmt.Printf("r1       : %d\n", r1___)
	fmt.Printf("r3       : %d\n", r3___)
	fmt.Printf("imm14    : %d\n", immediate)
	fmt.Printf("qp       : %d\n", qp___)

	if RetrievePredicateRegister(qp___) {
		r1 := RetrieveGeneralRegister(r1___)
		r3 := RetrieveGeneralRegister(r3___)

		r1.Value = uint64(immediate) + r3.Value
		r1.NotAThing = r3.NotAThing
	}
}
