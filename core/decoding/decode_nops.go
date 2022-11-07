package decoding

import "rosalia64/core/declarations"

func (decoder *DecoderContext) DecodeNopInteger(instructionBits uint64, nextSlot uint64) {
	instructionStruct := declarations.InstructionStruct{
		Disassembly: "nop.i",
	}

	decoder.ExecutableInstructions = append(decoder.ExecutableInstructions, declarations.NopInteger)
	decoder.InstructionStructs = append(decoder.InstructionStructs, instructionStruct)
}

func (decoder *DecoderContext) DecodeNopBranch(instructionBits uint64, nextSlot uint64) {
	instructionStruct := declarations.InstructionStruct{
		Disassembly: "nop.b",
	}

	decoder.ExecutableInstructions = append(decoder.ExecutableInstructions, declarations.NopMemory)
	decoder.InstructionStructs = append(decoder.InstructionStructs, instructionStruct)
}

func (decoder *DecoderContext) DecodeNopMemory(instructionBits uint64, nextSlot uint64) {
	instructionStruct := declarations.InstructionStruct{
		Disassembly: "nop.m",
	}

	decoder.ExecutableInstructions = append(decoder.ExecutableInstructions, declarations.NopMemory)
	decoder.InstructionStructs = append(decoder.InstructionStructs, instructionStruct)
}
