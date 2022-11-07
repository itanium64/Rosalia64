package decoding

import (
	"fmt"
	"rosalia64/core/declarations"
	"rosalia64/core/misc"
)

func (decoder *DecoderContext) DecodeAddImmediate22(instructionBits uint64, nextSlot uint64) {
	sign1 := (instructionBits & (0b0000100000000000000000000000000000000000000000)) >> 41
	imm9d := (instructionBits & (0b0000011111111100000000000000000000000000000000)) >> 32
	imm5c := (instructionBits & (0b0000000000000011111000000000000000000000000000)) >> 27
	r3___ := (instructionBits & (0b0000000000000000000110000000000000000000000000)) >> 25
	imm7b := (instructionBits & (0b0000000000000000000001111111000000000000000000)) >> 18
	r1___ := (instructionBits & (0b0000000000000000000000000000111111100000000000)) >> 11
	qp___ := (instructionBits & (0b0000000000000000000000000000000000011111100000)) >> 5

	immediate := misc.Imm22(sign1, imm5c, imm9d, imm7b)

	instructionStruct := declarations.InstructionStruct{
		Attributes: declarations.InstructionAttributeMap{
			declarations.ATTRIBUTE_SIGN:      sign1,
			declarations.ATTRIBUTE_IMMEDIATE: uint64(immediate),
			declarations.ATTRIBUTE_R1:        r1___,
			declarations.ATTRIBUTE_R3:        r3___,
			declarations.ATTRIBUTE_QP:        qp___,
		},
		Disassembly: fmt.Sprintf("addl r%d = %d, r%d", r1___, immediate, r3___),
	}

	decoder.ExecutableInstructions = append(decoder.ExecutableInstructions, declarations.AddlImm22)
	decoder.InstructionStructs = append(decoder.InstructionStructs, instructionStruct)
}

func (decoder *DecoderContext) DecodeAddImmediate14(instructionBits uint64, nextSlot uint64) {
	sign_ := (instructionBits & (0b0000100000000000000000000000000000000000000000)) >> 41
	x2a__ := (instructionBits & (0b0000011000000000000000000000000000000000000000)) >> 39
	ve___ := (instructionBits & (0b0000000100000000000000000000000000000000000000)) >> 38
	imm6d := (instructionBits & (0b0000000011111100000000000000000000000000000000)) >> 32
	r3___ := (instructionBits & (0b0000000000000011111110000000000000000000000000)) >> 25
	imm7b := (instructionBits & (0b0000000000000000000001111111000000000000000000)) >> 18
	r1___ := (instructionBits & (0b0000000000000000000000000000111111100000000000)) >> 11
	qp___ := (instructionBits & (0b0000000000000000000000000000000000011111100000)) >> 5

	immediate := misc.Imm14(sign_, imm6d, imm7b)

	instructionStruct := declarations.InstructionStruct{
		Attributes: declarations.InstructionAttributeMap{
			declarations.ATTRIBUTE_SIGN:      sign_,
			declarations.ATTRIBUTE_X2A:       x2a__,
			declarations.ATTRIBUTE_VE:        ve___,
			declarations.ATTRIBUTE_IMMEDIATE: uint64(immediate),
			declarations.ATTRIBUTE_R1:        r1___,
			declarations.ATTRIBUTE_R3:        r3___,
			declarations.ATTRIBUTE_QP:        qp___,
		},
		Disassembly: fmt.Sprintf("adds r%d = %d, r%d", r1___, immediate, r3___),
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
