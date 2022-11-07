package decoding

import (
	"fmt"
	"rosalia64/core/declarations"
	"rosalia64/core/formats"
)

func (decoder *DecoderContext) DecodeBranchIndirectReturn(instructionBits uint64, nextSlot uint64) {
	b4 := formats.ReadB4(instructionBits, nextSlot)

	instructionStruct := declarations.InstructionStruct{
		Attributes: declarations.InstructionAttributeMap{
			declarations.ATTRIBUTE_D:           b4.D,
			declarations.ATTRIBUTE_WH:          b4.WH,
			declarations.ATTRIBUTE_TABX:        b4.TableX,
			declarations.ATTRIBUTE_TABY:        b4.TableY,
			declarations.ATTRIBUTE_B2:          b4.B2,
			declarations.ATTRIBUTE_P:           b4.P,
			declarations.ATTRIBUTE_BRANCH_TYPE: b4.BType,
			declarations.ATTRIBUTE_QP:          b4.QP,
		},
		Disassembly: fmt.Sprintf("br b%d", b4.B2),
	}

	decoder.ExecutableInstructions = append(decoder.ExecutableInstructions, declarations.BranchIndirectReturn)
	decoder.InstructionStructs = append(decoder.InstructionStructs, instructionStruct)
}

func (decoder *DecoderContext) BranchIndirectMiscellaneous(instructionBits uint64, nextSlot uint64) {
	tableX := (instructionBits & (0b0000000011000000000000000000000000000000000000)) >> 36
	tableY := (instructionBits & (0b0000000000111100000000000000000000000000000000)) >> 32

	subinstructionTable := [][]func(instructionBits, nextSlot uint64){
		{nil, nil, nil, nil},
		{nil, nil, decoder.DecodeBranchIndirectReturn, nil},
		//this table still goes on
	}

	subinstructionTable[tableY][tableX](instructionBits, nextSlot)
}
