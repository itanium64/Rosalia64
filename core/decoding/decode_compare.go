package decoding

import (
	"fmt"
	"rosalia64/core/declarations"
	"rosalia64/core/formats"
)

type DisassemblyWithInstructionTable struct {
	Disassembly string
	Instruction declarations.ExecutableInstruction
}

var opCodeCdismTable map[declarations.ComparisonTypeTbTaC]string = map[declarations.ComparisonTypeTbTaC]string{
	declarations.TB_TA_C_LT_NONE: ".lt",
	declarations.TB_TA_C_LT_UNC:  ".lt",
	declarations.TB_TA_C_EQ:      ".eq",
	declarations.TB_TA_C_NE:      ".ne",
	declarations.TB_TA_C_GT:      ".gt",
	declarations.TB_TA_C_LE:      ".le",
	declarations.TB_TA_C_GE:      ".ge",
	declarations.TB_TA_C_LT:      ".lt",
}

func (decoder *DecoderContext) DecodeIntegerCompareOpcodeCRegisterForm(instructionBits uint64, nextSlot uint64) {

}

func (decoder *DecoderContext) DecodeIntegerCompareOpcodeCImmediate(instructionBits uint64, nextSlot uint64) {
	a8 := formats.ReadA8(instructionBits, nextSlot)

	TaC := (a8.Ta) | (a8.C << 1)

	prRegCompleter := declarations.PR_COMPLETER_NONE

	disassembly := "cmp"

	//0 = full 64bit compare, 1 = 32bit compare
	cm4 := 0

	if a8.X2 == 3 {
		//cmp4
		cm4 = 1

		disassembly += "4"
	}

	comparisonDisassembly := opCodeCdismTable[declarations.ComparisonTypeTbTaC(TaC)]

	disassembly += comparisonDisassembly

	switch TaC {
	case 0:
		prRegCompleter = declarations.PR_COMPLETER_NONE
	case 1:
		prRegCompleter = declarations.PR_COMPLETER_UNC
		disassembly += ".unc"
	default:
		//if tb ta c together are higher than 2 = .and
		prRegCompleter = declarations.PR_COMPLETER_AND
		disassembly += ".and"
	}

	instructionStruct := declarations.InstructionStruct{
		Attributes: declarations.InstructionAttributeMap{
			declarations.ATTRIBUTE_QP:           a8.Qp,
			declarations.ATTRIBUTE_IMMEDIATE:    a8.Immediate,
			declarations.ATTRIBUTE_R3:           a8.R3,
			declarations.ATTRIBUTE_PR1:          a8.P1,
			declarations.ATTRIBUTE_PR2:          a8.P2,
			declarations.ATTRIBUTE_PR_COMPLETER: uint64(prRegCompleter),
			declarations.ATTRIBUTE_CM4:          uint64(cm4),
		},
		Disassembly: fmt.Sprintf("%s p%d, p%d = %d, r%d", disassembly, a8.P1, a8.P2, a8.Immediate, a8.R3),
	}

	decoder.ExecutableInstructions = append(decoder.ExecutableInstructions, declarations.IntegerCompareRegisterForm)
	decoder.InstructionStructs = append(decoder.InstructionStructs, instructionStruct)
}

func (decoder *DecoderContext) DecodeIntegerCompareOpcodeC(instructionBits uint64, nextSlot uint64) {
	x2 := (instructionBits & (0b0000011000000000000000000000000000000000000000)) >> 39

	if x2 <= 1 {
		decoder.DecodeIntegerCompareOpcodeCRegisterForm(instructionBits, nextSlot)
	} else {
		decoder.DecodeIntegerCompareOpcodeCImmediate(instructionBits, nextSlot)
	}
}
