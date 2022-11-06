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

func SystemMemoryManagment(instructionBits uint64, nextSlot uint64) {
	x3 := (instructionBits & (0b0000011100000000000000000000000000000000000000)) >> 38

	switch x3 {
	case 0:
		SystemMemoryManagment42(instructionBits, nextSlot)
	case 4:
		//chk.a.nc : int
	case 5:
		//chk.a.clr : int
	case 6:
		//chk.a.nc : fp
	case 7:
		//chk.a.clr : fp
	}
}

func BranchIndirectMiscellaneous(instructionBits uint64, nextSlot uint64) {
	tableX := (instructionBits & (0b0000000011000000000000000000000000000000000000)) >> 36
	tableY := (instructionBits & (0b0000000000111100000000000000000000000000000000)) >> 32

	subinstructionTable := [][]func(instructionBits, nextSlot uint64){
		{nil, nil, nil, nil},
		{nil, nil, BranchIndirectReturn, nil},
		//this table still goes on
	}

	subinstructionTable[tableY][tableX](instructionBits, nextSlot)
}

func IntegerMisc(instructionBits uint64, nextSlot uint64) {
	x3 := (instructionBits & (0b0000011100000000000000000000000000000000000000)) >> 38

	switch x3 {
	case 0:
		IntegerMisc6bitExt(instructionBits, nextSlot)
	}

}

func IntegerMisc6bitExt(instructionBits uint64, nextSlot uint64) {
	tableX := (instructionBits & (0b0000000011000000000000000000000000000000000000)) >> 36
	tableY := (instructionBits & (0b0000000000111100000000000000000000000000000000)) >> 32

	subinstructionTable := [][]func(instructionBits, nextSlot uint64){
		{nil, nil, nil, nil},
		{NopInteger, nil, nil, nil},
		//this table still goes on
	}

	subinstructionTable[tableY][tableX](instructionBits, nextSlot)
}
