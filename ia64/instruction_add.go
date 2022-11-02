package ia64

func AddImmediate22(instructionBits uint64, nextSlot uint64) {
	sign1 := (instructionBits & (0b0000100000000000000000000000000000000000000000)) >> 40
	imm9d := (instructionBits & (0b0000011111111100000000000000000000000000000000)) >> 32
	imm5c := (instructionBits & (0b0000000000000011111000000000000000000000000000)) >> 27
	reg3_ := (instructionBits & (0b0000000000000000000110000000000000000000000000)) >> 25
	imm7b := (instructionBits & (0b0000000000000000000001111111100000000000000000)) >> 18
	reg1_ := (instructionBits & (0b0000000000000000000000000000011111110000000000)) >> 10
	//qp___ := (instructionBits & (0b0000000000000000000000000000000000001111110000)) >> 4

	immediate := Imm22(sign1, imm5c, imm9d, imm7b)

	//TODO: if( PR[qp] )

	Processor.GeneralRegisters[reg1_].Value = uint64(immediate) + Processor.GeneralRegisters[reg3_].Value
	Processor.GeneralRegisters[reg1_].NotAThing = Processor.GeneralRegisters[reg3_].NotAThing

	//endif
}

func AddImmediate14(instructionBits uint64, nextSlot uint64) {
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
