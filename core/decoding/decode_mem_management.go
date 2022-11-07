package decoding

func (decoder *DecoderContext) DecodeSystemMemoryManagment4Plus2bit(instructionBits uint64, nextSlot uint64) {
	tableX := (instructionBits & (0b0000000011000000000000000000000000000000000000)) >> 36
	tableY := (instructionBits & (0b0000000000111100000000000000000000000000000000)) >> 32

	subinstructionTable := [][]func(instructionBits uint64, nextSlot uint64){
		{func(instructionBits, nextSlot uint64) { /* break.m */ }, decoder.DecodeNopMemory},
	}

	subinstructionTable[tableX][tableY](instructionBits, nextSlot)
}

func (decoder *DecoderContext) DecodeSystemMemoryManagment3bit(instructionBits uint64, nextSlot uint64) {
	x3 := (instructionBits & (0b0000011100000000000000000000000000000000000000)) >> 38

	switch x3 {
	case 0:
		decoder.DecodeSystemMemoryManagment4Plus2bit(instructionBits, nextSlot)
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
