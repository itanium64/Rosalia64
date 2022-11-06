package ia64

func IntegerALU(instructionBits uint64, nextSlot uint64) {
=	x2a_ := (instructionBits & (0b0000010000000000000000000000000000000000000000)) >> 39
	ve__ := (instructionBits & (0b0000001000000000000000000000000000000000000000)) >> 38

	switch x2a_ {
	case 2:
		if ve__ == 0 {
			AddImmediate14(instructionBits, nextSlot)
		}
	}
}

func IntegerLoadStoreSemaphoreFR(instructionBits uint64, nextSlot uint64) {
	m := (instructionBits & (0b0000100000000000000000000000000000000000000000)) >> 41
	x := (instructionBits & (0b0000000000000100000000000000000000000000000000)) >> 32

	if m == 1 {
		if x == 1 {
			//unused
			return
		}

		IntegerLoadWithRegister(instructionBits, nextSlot)
	} else {
		if x == 1 {
			//Semaphore/get FR
			return
		}

		IntegerLoadStore(instructionBits, nextSlot)
	}
}
