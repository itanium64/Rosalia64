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

	switch as5bit {
	case 0b0_0000:
		//cmp.lt
	case 0b0_0001:
		//cmp.lt.unc
	case 0b0_0010:
		//cmp.eq.and
	case 0b0_0011:
		//cmp.ne.eq
	case 0b0_0100:
		//cmp.gt.and
	case 0b_0101:
		//cmp.le.and
	case 0b0_0110:
		//cmp.ge.and
	case 0b0_0111:
		//cmp.lt.and

	// x2 = 1

	case 0b0_1000:
		//cmp4.lt
	case 0b0_1001:
		//cmp4.lt.unc
	case 0b0_1010:
		//cmp4.eq.and
	case 0b0_1011:
		//cmp4.ne.eq
	case 0b0_1100:
		//cmp4.gt.and
	case 0b0_1101:
		//cmp4.le.and
	case 0b0_1110:
		//cmp4.ge.and
	case 0b0_1111:
		//cmp4.lt.and
	}
}
