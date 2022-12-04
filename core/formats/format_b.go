package formats

import "rosalia64/core/ia_math"

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

type B1 struct {
	Sign      uint64
	D         uint64
	Wh        uint64
	Immediate uint64
	P         uint64
	Btype     uint64
	Qp        uint64
}

func ReadB1(instructionBits uint64, nextSlot uint64) B1 {
	__sign := (instructionBits & (0b0000100000000000000000000000000000000000000000)) >> 41
	_____d := (instructionBits & (0b0000010000000000000000000000000000000000000000)) >> 40
	____wh := (instructionBits & (0b0000001100000000000000000000000000000000000000)) >> 38
	imm20b := (instructionBits & (0b0000000011111111111111111111000000000000000000)) >> 18
	_____p := (instructionBits & (0b0000000000000000000000000000100000000000000000)) >> 17
	_btype := (instructionBits & (0b0000000000000000000000000000000011100000000000)) >> 11
	____qp := (instructionBits & (0b0000000000000000000000000000000000011111100000)) >> 5

	immediate := ia_math.SignExt(__sign<<20|imm20b, 21) << 4

	return B1{
		Sign:      __sign,
		D:         _____d,
		Wh:        ____wh,
		Immediate: uint64(immediate),
		P:         _____p,
		Btype:     _btype,
		Qp:        ____qp,
	}
}
