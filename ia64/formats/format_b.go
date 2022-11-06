package formats

type B4 struct {
	D      uint64
	WH     uint64
	TableX uint64
	TableY uint64
	B2     uint64
	P      uint64
	BType  uint64
	QP     uint64
}

func ReadB4(instructionBits uint64, nextSlot uint64) B4 {
	d_____ := (instructionBits & (0b0000010000000000000000000000000000000000000000)) >> 40
	wh____ := (instructionBits & (0b0000001100000000000000000000000000000000000000)) >> 38
	tableX := (instructionBits & (0b0000000011000000000000000000000000000000000000)) >> 36
	tableY := (instructionBits & (0b0000000000111100000000000000000000000000000000)) >> 32
	b2____ := (instructionBits & (0b0000000000000000000000000111000000000000000000)) >> 18
	p_____ := (instructionBits & (0b0000000000000000000000000000100000000000000000)) >> 17
	btype_ := (instructionBits & (0b0000000000000000000000000000000011100000000000)) >> 11
	qp____ := (instructionBits & (0b0000000000000000000000000000000000011111100000)) >> 5

	return B4{
		D:      d_____,
		WH:     wh____,
		TableX: tableX,
		TableY: tableY,
		B2:     b2____,
		P:      p_____,
		BType:  btype_,
		QP:     qp____,
	}
}
