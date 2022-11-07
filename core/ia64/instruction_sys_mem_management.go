package ia64

func SystemMemoryManagment42(instructionBits uint64, nextSlot uint64) {
	tableX := (instructionBits & (0b0000000011000000000000000000000000000000000000)) >> 36
	tableY := (instructionBits & (0b0000000000111100000000000000000000000000000000)) >> 32

	subinstructionTable := [][]func(instructionBits uint64, nextSlot uint64){
		{func(instructionBits, nextSlot uint64) { /* break.m */ }, NopMemory},
	}

	subinstructionTable[tableX][tableY](instructionBits, nextSlot)
}
