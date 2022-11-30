package decoding

import (
	"rosalia64/core/declarations"
	"rosalia64/core/formats"
)

type DisassemblyWithInstructionTable struct {
	Disassembly string
	Instruction declarations.ExecutableInstruction
}

var opCodeCInstructionTable map[uint64]DisassemblyWithInstructionTable = map[uint64]DisassemblyWithInstructionTable{
	0: {Instruction: declarations.IntegerCompareRegisterFormLT, Disassembly: ".lt"},
	1: {Instruction: declarations.IntegerCompareRegisterFormLT, Disassembly: ".lt"},
	2: {Instruction: declarations.IntegerCompareRegisterFormEQ, Disassembly: ".eq"},
	3: {Instruction: declarations.IntegerCompareRegisterFormNE, Disassembly: ".ne"},
	4: {Instruction: declarations.IntegerCompareRegisterFormGT, Disassembly: ".gt"},
	5: {Instruction: declarations.IntegerCompareRegisterFormLE, Disassembly: ".le"},
	6: {Instruction: declarations.IntegerCompareRegisterFormGE, Disassembly: ".ge"},
	7: {Instruction: declarations.IntegerCompareRegisterFormLT, Disassembly: ".lt"},
}

func (decoder *DecoderContext) DecodeIntegerCompareOpcodeC(instructionBits uint64, nextSlot uint64) {
	a6 := formats.ReadA6(instructionBits, nextSlot)

	tbTaC := (a6.Tb) | (a6.Ta << 1) | (a6.C << 2)

	prRegCompleter := declarations.PR_COMPLETER_NONE

	disassembly := "cmp"

	//0 = full 64bit compare, 1 = 32bit compare
	cm4 := 0

	if a6.X2 == 1 {
		//cmp4
		cm4 = 1

		disassembly += "4"
	}

	instruction := opCodeCInstructionTable[tbTaC]

	disassembly += instruction.Disassembly

	switch tbTaC {
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
			declarations.ATTRIBUTE_QP:           a6.Qp,
			declarations.ATTRIBUTE_R2:           a6.R2,
			declarations.ATTRIBUTE_R3:           a6.R3,
			declarations.ATTRIBUTE_PR1:          a6.P1,
			declarations.ATTRIBUTE_PR2:          a6.P2,
			declarations.ATTRIBUTE_PR_COMPLETER: uint64(prRegCompleter),
			declarations.ATTRIBUTE_CM4:          uint64(cm4),
		},
		Disassembly: disassembly,
	}

	decoder.ExecutableInstructions = append(decoder.ExecutableInstructions, instruction.Instruction)
	decoder.InstructionStructs = append(decoder.InstructionStructs, instructionStruct)
}
