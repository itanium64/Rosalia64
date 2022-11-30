package decoding

func (decoder *DecoderContext) DecodeIntegerCompareOpcodeC(instructionBits uint64, nextSlot uint64) {
	//tb := (instructionBits & (0b0000100000000000000000000000000000000000000000)) >> 41
	//x2 := (instructionBits & (0b0000011000000000000000000000000000000000000000)) >> 39
	//ta := (instructionBits & (0b0000000100000000000000000000000000000000000000)) >> 38
	//_c := (instructionBits & (0b0000000000000000000000000000100000000000000000)) >> 17

	tb := uint64(1)
	x2 := uint64(3)
	ta := uint64(1)
	_c := uint64(1)

	//doing this instead of a endless if/else tree, way more readable
	//result is exactly the same, we just turn it into an integer
	as5bit := (tb) | (x2 << 1) | (ta << 3) | (_c << 4)

	//if tb ta c together are higher than 2 = .and

	
	
	if x2 == 1 {
		//cmp4
	}
}
