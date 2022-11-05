package formats

type M1_2_4 struct {
	M      uint64
	TableY uint64
	TableX uint64
	Hint   uint64
	X      uint64
	R3     uint64
	R2     uint64
	R1     uint64
	QP     uint64
}

func ReadM1(instructionBits uint64, nextSlot uint64) M1_2_4 {
	m___ := (instructionBits & (0b0000100000000000000000000000000000000000000000)) >> 41
	tabx := (instructionBits & (0b0000011110000000000000000000000000000000000000)) >> 37
	taby := (instructionBits & (0b0000000001100000000000000000000000000000000000)) >> 35
	hint := (instructionBits & (0b0000000000011000000000000000000000000000000000)) >> 33
	x___ := (instructionBits & (0b0000000000000100000000000000000000000000000000)) >> 32
	r3__ := (instructionBits & (0b0000000000000011111110000000000000000000000000)) >> 25
	r2__ := (instructionBits & (0b0000000000000000000001111111000000000000000000)) >> 18
	r1__ := (instructionBits & (0b0000000000000000000000000000111111100000000000)) >> 11
	qp__ := (instructionBits & (0b0000000000000000000000000000000000011111100000)) >> 5

	return M1_2_4{
		M:      m___,
		TableX: tabx,
		TableY: taby,
		Hint:   hint,
		X:      x___,
		R3:     r3__,
		R2:     r2__,
		R1:     r1__,
		QP:     qp__,
	}
}

func ReadM2(instructionBits uint64, nextSlot uint64) M1_2_4 {
	return ReadM1(instructionBits, nextSlot)
}

func ReadM4(instructionBits uint64, nextSlot uint64) M1_2_4 {
	return ReadM1(instructionBits, nextSlot)
}
