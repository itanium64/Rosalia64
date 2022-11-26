package formats

import (
	"rosalia64/core/ia_math"
)

type A4 struct {
	Sign      uint64
	X2A       uint64
	VE        uint64
	Immediate uint64
	R3        uint64
	R1        uint64
	QP        uint64
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

	immediate := ia_math.Imm14(sign_, imm6d, imm7b)

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

	immediate := ia_math.Imm22(sign1, imm5c, imm9d, imm7b)

	return A5{
		Sign:      sign1,
		Immediate: uint64(immediate),
		R1:        r1___,
		R3:        r3___,
		QP:        qp___,
	}
}

type A6 struct {
	Tb uint64
	X2 uint64
	Ta uint64
	P2 uint64
	R3 uint64
	R2 uint64
	C  uint64
	P1 uint64
	Qp uint64
}

func ReadA6(instructionBits uint64, nextSlot uint64) A6 {
	tb := (instructionBits & (0b0000100000000000000000000000000000000000000000)) >> 41
	x2 := (instructionBits & (0b0000011000000000000000000000000000000000000000)) >> 39
	ta := (instructionBits & (0b0000000100000000000000000000000000000000000000)) >> 38
	p2 := (instructionBits & (0b0000000011111100000000000000000000000000000000)) >> 32
	r3 := (instructionBits & (0b0000000000000011111110000000000000000000000000)) >> 25
	r2 := (instructionBits & (0b0000000000000000000001111111000000000000000000)) >> 18
	_c := (instructionBits & (0b0000000000000000000000000000100000000000000000)) >> 17
	p1 := (instructionBits & (0b0000000000000000000000000000011111100000000000)) >> 11
	qp := (instructionBits & (0b0000000000000000000000000000000000011111100000)) >> 5

	return A6{
		Tb: tb,
		X2: x2,
		Ta: ta,
		P2: p2,
		R3: r3,
		R2: r2,
		C:  _c,
		P1: p1,
		Qp: qp,
	}
}
