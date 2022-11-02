package ia64

import "fmt"

var I_UnitInstructionTable UnitInstructionTable = UnitInstructionTable{
	8: IntegerALU,
}

func IntegerALU(instructionBits uint64) {
	x2a_ := (instructionBits & (0b0000010000000000000000000000000000000000000000)) >> 39
	ve__ := (instructionBits & (0b0000001000000000000000000000000000000000000000)) >> 38

	fmt.Printf("x2a : %d", x2a_)
	fmt.Printf("ve  : %d", ve__)

	switch x2a_ {
	case 2:
		if ve__ == 0 {
			AddImmediate14(instructionBits)
		}
	}
}

func AddImmediate14(instructionBits uint64) {
	sign_ := (instructionBits & (0b0000100000000000000000000000000000000000000000)) >> 40
	imm6d := (instructionBits & (0b0000000011111100000000000000000000000000000000)) >> 32
	r3___ := (instructionBits & (0b0000000000000011111110000000000000000000000000)) >> 25
	imm7b := (instructionBits & (0b0000000000000000000001111111000000000000000000)) >> 18
	r1___ := (instructionBits & (0b0000000000000000000000000000111111100000000000)) >> 11
	//qp___ := (instructionBits & (0b0000000000000000000000000000000000011111110000)) >> 4

	//TODO: if( PR[qp] )

	immediate := Imm14(sign_, imm6d, imm7b)

	Processor.GeneralRegisters[r1___].Value = uint64(immediate) + Processor.GeneralRegisters[r3___].Value
	Processor.GeneralRegisters[r1___].NotAThing = Processor.GeneralRegisters[r3___].NotAThing

	//endif
}
