package ia64

var M_UnitInstructionTable UnitInstructionTable = UnitInstructionTable{
	9: AddImmediate22,
}

func AddImmediate22(instructionBits uint64) {
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
