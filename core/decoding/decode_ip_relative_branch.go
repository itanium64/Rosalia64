package decoding

import (
	"fmt"
	"rosalia64/core/declarations"
	"rosalia64/core/formats"
	"strings"
)

func (decoder *DecoderContext) DecodeIPRelativeBranch(instructionBits uint64, nextSlot uint64) {
	b1 := formats.ReadB1(instructionBits, nextSlot)

	disassembly := "br."
	trailingDisassembly := ""

	instructionStruct := declarations.InstructionStruct{
		Attributes: declarations.InstructionAttributeMap{
			declarations.ATTRIBUTE_BRANCH_TYPE:  b1.Btype,
			declarations.ATTRIBUTE_IMMEDIATE:    b1.Immediate,
			declarations.ATTRIBUTE_SPECULATION:  b1.Wh,
			declarations.ATTRIBUTE_PREFETCH:     b1.P,
			declarations.ATTRIBUTE_BRANCH_CACHE: b1.D,
		},
		Disassembly: "",
	}

	switch declarations.BranchWhetherHint(b1.Wh) {
	case declarations.BRWH_SPECULATE_TAKEN:
		trailingDisassembly += "sptk."
	case declarations.BRWH_SPECULATE_NOT_TAKEN:
		trailingDisassembly += "spnt."
	case declarations.BRWH_DYNAMIC_TAKEN:
		trailingDisassembly += "dptk."
	case declarations.BRWH_DYNAMIC_NOT_TAKEN:
		trailingDisassembly += "dpnt."
	}

	switch declarations.SequentialPrefetchHint(b1.P) {
	case declarations.PREFETCH_HINT_FEW:
		trailingDisassembly += "few."
	case declarations.PREFETCH_HINT_MANY:
		trailingDisassembly += "many."
	}

	if b1.D == 1 {
		trailingDisassembly += "clr"
	} else {
		trailingDisassembly = strings.TrimSuffix(trailingDisassembly, ".")
	}

	//.cond and none do the same thing
	if b1.Btype == 0 {
		b1.Btype = 2
	}

	switch declarations.IPRelativeBranchType(b1.Btype) {
	case declarations.IPREL_BRANCH_TYPE_COND:
		disassembly += "cond."

		target := decoder.InstructionIndexToAddress[decoder.InstructionIndex] + int64(b1.Immediate)

		instructionStruct.Disassembly = fmt.Sprintf("%s%s loc_%08x", disassembly, trailingDisassembly, target)

		decoder.InstructionStructs = append(decoder.InstructionStructs, instructionStruct)
		decoder.ExecutableInstructions = append(decoder.ExecutableInstructions, declarations.ExecuteIPRelativeBranchCond)
	}
}
