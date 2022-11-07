package decoding

import (
	"fmt"
	"rosalia64/core/declarations"
	"rosalia64/core/formats"
)

func (decoder *DecoderContext) DecodeAddImmediate22(instructionBits uint64, nextSlot uint64) {
	a5 := formats.ReadA5(instructionBits, nextSlot)

	instructionStruct := declarations.InstructionStruct{
		Attributes: declarations.InstructionAttributeMap{
			declarations.ATTRIBUTE_SIGN:      a5.Sign,
			declarations.ATTRIBUTE_IMMEDIATE: a5.Immediate,
			declarations.ATTRIBUTE_R1:        a5.R1,
			declarations.ATTRIBUTE_R3:        a5.R3,
			declarations.ATTRIBUTE_QP:        a5.QP,
		},
		Disassembly: fmt.Sprintf("addl r%d = %d, r%d", a5.R1, a5.Immediate, a5.R3),
	}

	decoder.ExecutableInstructions = append(decoder.ExecutableInstructions, declarations.AddlImm22)
	decoder.InstructionStructs = append(decoder.InstructionStructs, instructionStruct)
}

func (decoder *DecoderContext) DecodeAddImmediate14(instructionBits uint64, nextSlot uint64) {
	a4 := formats.ReadA4(instructionBits, nextSlot)

	instructionStruct := declarations.InstructionStruct{
		Attributes: declarations.InstructionAttributeMap{
			declarations.ATTRIBUTE_SIGN:      a4.Sign,
			declarations.ATTRIBUTE_X2A:       a4.X2A,
			declarations.ATTRIBUTE_VE:        a4.VE,
			declarations.ATTRIBUTE_IMMEDIATE: a4.Immediate,
			declarations.ATTRIBUTE_R1:        a4.R1,
			declarations.ATTRIBUTE_R3:        a4.R3,
			declarations.ATTRIBUTE_QP:        a4.QP,
		},
		Disassembly: fmt.Sprintf("adds r%d = %d, r%d", a4.R1, a4.Immediate, a4.R3),
	}

	decoder.ExecutableInstructions = append(decoder.ExecutableInstructions, declarations.AddsImm14)
	decoder.InstructionStructs = append(decoder.InstructionStructs, instructionStruct)
}

func (decoder *DecoderContext) DecodeIntegerALU(instructionBits uint64, nextSlot uint64) {
	x2a_ := (instructionBits & (0b0000010000000000000000000000000000000000000000)) >> 39
	ve__ := (instructionBits & (0b0000001000000000000000000000000000000000000000)) >> 38

	switch x2a_ {
	case 2:
		if ve__ == 0 {
			decoder.DecodeAddImmediate14(instructionBits, nextSlot)
		}
	}
}
