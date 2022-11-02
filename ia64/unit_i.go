package ia64

func IntegerALU(instructionBits uint64, nextSlot uint64) {
	x2a_ := (instructionBits & (0b0000010000000000000000000000000000000000000000)) >> 39
	ve__ := (instructionBits & (0b0000001000000000000000000000000000000000000000)) >> 38

	switch x2a_ {
	case 2:
		if ve__ == 0 {
			AddImmediate14(instructionBits, nextSlot)
		}
	}
}
