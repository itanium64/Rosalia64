package decoding

func (decoder *DecoderContext) DecodeIntegerCompareOpcodeC(instructionBits uint64, nextSlot uint64) {
	//tb := (instructionBits & (0b0000100000000000000000000000000000000000000000)) >> 41
	//x2 := (instructionBits & (0b0000011000000000000000000000000000000000000000)) >> 39
	//ta := (instructionBits & (0b0000000100000000000000000000000000000000000000)) >> 38
	//_c := (instructionBits & (0b0000000000000000000000000000100000000000000000)) >> 17
}
