package ia64

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
