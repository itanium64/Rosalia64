package decoding

func (decoder *DecoderContext) DecodeIntegerMisc3bit(instructionBits uint64, nextSlot uint64) {
	x3 := (instructionBits & (0b0000011100000000000000000000000000000000000000)) >> 38

	switch x3 {
	case 0:
		decoder.DecodeIntegerMisc6bitExt(instructionBits, nextSlot)
	}

}

func (decoder *DecoderContext) DecodeIntegerMisc6bitExt(instructionBits uint64, nextSlot uint64) {
	tableX := (instructionBits & (0b0000000011000000000000000000000000000000000000)) >> 36
	tableY := (instructionBits & (0b0000000000111100000000000000000000000000000000)) >> 32

	subinstructionTable := [][]func(instructionBits, nextSlot uint64){
		{nil, nil, nil, nil},
		{decoder.DecodeNopInteger, nil, nil, nil},
		//this table still goes on
	}

	subinstructionTable[tableY][tableX](instructionBits, nextSlot)
}
