package formats

type A4 struct {
	Sign      uint64
	X2A       uint64
	VE        uint64
	Immediate uint64
	R3        uint64
	R1        uint64
	QP        uint64
}

type A5 struct {
	Sign      uint64
	Immediate uint64
	R3        uint64
	R1        uint64
	QP        uint64
}

func ReadA5(instructionBits uint64, nextSlot uint64) A5 {
	sign1 := (instructionBits & (0b0000100000000000000000000000000000000000000000)) >> 41
	imm9d := (instructionBits & (0b0000011111111100000000000000000000000000000000)) >> 32
	imm5c := (instructionBits & (0b0000000000000011111000000000000000000000000000)) >> 27
	r3___ := (instructionBits & (0b0000000000000000000110000000000000000000000000)) >> 25
	imm7b := (instructionBits & (0b0000000000000000000001111111000000000000000000)) >> 18
	r1___ := (instructionBits & (0b0000000000000000000000000000111111100000000000)) >> 11
	qp___ := (instructionBits & (0b0000000000000000000000000000000000011111100000)) >> 5

	immediate := Imm22(sign1, imm5c, imm9d, imm7b)

	return A5{
		Sign:      sign1,
		Immediate: uint64(immediate),
		R1:        r1___,
		R3:        r3___,
		QP:        qp___,
	}
}

func ReadA4(instructionBits uint64, nextSlot uint64) A4 {
	sign_ := (instructionBits & (0b0000100000000000000000000000000000000000000000)) >> 41
	x2a__ := (instructionBits & (0b0000011000000000000000000000000000000000000000)) >> 39
	ve___ := (instructionBits & (0b0000000100000000000000000000000000000000000000)) >> 38
	imm6d := (instructionBits & (0b0000000011111100000000000000000000000000000000)) >> 32
	r3___ := (instructionBits & (0b0000000000000011111110000000000000000000000000)) >> 25
	imm7b := (instructionBits & (0b0000000000000000000001111111000000000000000000)) >> 18
	r1___ := (instructionBits & (0b0000000000000000000000000000111111100000000000)) >> 11
	qp___ := (instructionBits & (0b0000000000000000000000000000000000011111100000)) >> 5

	immediate := Imm14(sign_, imm6d, imm7b)

	return A4{
		Sign:      sign_,
		X2A:       x2a__,
		VE:        ve___,
		Immediate: uint64(immediate),
		R1:        r1___,
		R3:        r3___,
		QP:        qp___,
	}
}
