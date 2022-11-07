package decoding

import "rosalia64/core/declarations"

type DecoderContext struct {
	ExecutableInstructions []declarations.ExecutableInstruction
	InstructionStructs     []declarations.InstructionStruct
}

var DecodingContext *DecoderContext

func InitializeDecoderAndTables() {
	DecodingContext = &DecoderContext{}

	B_UnitInstructionTable[0] = DecodingContext.BranchIndirectMiscellaneous
	B_UnitInstructionTable[2] = DecodingContext.DecodeNopBranch

	I_UnitInstructionTable[0] = DecodingContext.DecodeIntegerMisc3bit
	I_UnitInstructionTable[8] = DecodingContext.DecodeIntegerALU

	M_UnitInstructionTable[0] = DecodingContext.DecodeSystemMemoryManagment3bit
	M_UnitInstructionTable[4] = DecodingContext.DecodeIntegerLoadStoreSemaphoreFR1bit
	M_UnitInstructionTable[8] = DecodingContext.DecodeIntegerALU
	M_UnitInstructionTable[9] = DecodingContext.DecodeAddImmediate22
}
