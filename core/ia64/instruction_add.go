package ia64

import (
	"fmt"
	"rosalia64-core/ia64/formats"
)

func AddImmediate22(instructionBits uint64, nextSlot uint64) {
	a5 := formats.ReadA5(instructionBits, nextSlot)

	fmt.Printf("\nExecuting: addl r1 = imm22, r3\n")
	fmt.Printf("sign     : %d\n", a5.Sign)
	fmt.Printf("r1       : %d\n", a5.R1)
	fmt.Printf("r3       : %d\n", a5.R3)
	fmt.Printf("imm22    : %d\n", a5.Immediate)
	fmt.Printf("qp       : %d\n", a5.QP)

	if RetrievePredicateRegister(a5.QP) {
		//todo: check_target_register(r1)

		r1 := RetrieveGeneralRegister(a5.R1)
		r3 := RetrieveGeneralRegister(a5.R3)

		r1.Value = uint64(a5.Immediate) + r3.Value
		r1.NotAThing = r3.NotAThing
	}
}

func AddImmediate14(instructionBits uint64, nextSlot uint64) {
	a4 := formats.ReadA4(instructionBits, nextSlot)

	fmt.Printf("\nExecuting: adds r1 = imm14, r3\n")
	fmt.Printf("sign     : %d\n", a4.Sign)
	fmt.Printf("r1       : %d\n", a4.R1)
	fmt.Printf("r3       : %d\n", a4.R3)
	fmt.Printf("imm14    : %d\n", a4.Immediate)
	fmt.Printf("qp       : %d\n", a4.QP)

	if RetrievePredicateRegister(a4.QP) {
		//todo: check_target_register(r1)

		r1 := RetrieveGeneralRegister(a4.R1)
		r3 := RetrieveGeneralRegister(a4.R3)

		r1.Value = uint64(a4.Immediate) + r3.Value
		r1.NotAThing = r3.NotAThing
	}
}
