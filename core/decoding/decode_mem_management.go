package decoding

import "fmt"

func (decoder *DecoderContext) DecodeSystemMemoryManagment4Plus2bit(instructionBits uint64, nextSlot uint64) {
	tableX := (instructionBits & (0b0000000011000000000000000000000000000000000000)) >> 36
	tableY := (instructionBits & (0b0000000000111100000000000000000000000000000000)) >> 32

	subinstructionTable := [][]func(instructionBits uint64, nextSlot uint64){
		{func(instructionBits, nextSlot uint64) { /* break.m */ }, decoder.DecodeNopMemory, nil, nil},
		{nil, nil, nil, nil},
		{nil, nil, nil, nil},
		{nil, nil, nil, nil},
		{nil, nil, nil, nil},
		{nil, nil, nil, nil},
		{nil, nil, nil, nil},
		{nil, nil, nil, nil},
		{nil, nil, nil, nil},
		{nil, nil, nil, nil},
		{nil, nil, nil, nil},
		{nil, nil, nil, nil},
		{nil, nil, nil, nil},
		{nil, nil, nil, nil},
		{nil, nil, nil, nil},
	}

	tableResult := subinstructionTable[tableY][tableX]

	if tableResult != nil {
		tableResult(instructionBits, nextSlot)
	} else {
		fmt.Printf("UNIMPLEMENTED!!!: DecodeSystemMemoryManagment4Plus2bit\nTableX: %d\nTableY %d\n\n", tableX, tableY)
	}
}

func (decoder *DecoderContext) DecodeSystemMemoryManagment3bit(instructionBits uint64, nextSlot uint64) {
	x3 := (instructionBits & (0b0000011100000000000000000000000000000000000000)) >> 38

	switch x3 {
	case 0:
		decoder.DecodeSystemMemoryManagment4Plus2bit(instructionBits, nextSlot)
	case 4:
		fmt.Printf("UNIMPLEMENTED!!!: DecodeSystemMemoryManagment3bit\nx3: %d\n\n", x3)
		//chk.a.nc : int
	case 5:
		fmt.Printf("UNIMPLEMENTED!!!: DecodeSystemMemoryManagment3bit\nx3: %d\n\n", x3)
		//chk.a.clr : int
	case 6:
		fmt.Printf("UNIMPLEMENTED!!!: DecodeSystemMemoryManagment3bit\nx3: %d\n\n", x3)
		//chk.a.nc : fp
	case 7:
		fmt.Printf("UNIMPLEMENTED!!!: DecodeSystemMemoryManagment3bit\nx3: %d\n\n", x3)
		//chk.a.clr : fp
	}
}
