package decoding

import (
	"fmt"
	"rosalia64/core/declarations"
	"rosalia64/core/formats"
)

func (decoder *DecoderContext) DecodeIntegerStore(instructionBits uint64, nextSlot uint64) {
	m := formats.ReadM1(instructionBits, nextSlot)

	bitLengthTable := []uint64{
		1, 2, 4, 8,
	}

	instructionStruct := declarations.InstructionStruct{
		Attributes: declarations.InstructionAttributeMap{
			declarations.ATTRIBUTE_TABX: m.TableX,
			declarations.ATTRIBUTE_R2:   m.R2,
			declarations.ATTRIBUTE_R3:   m.R3,
			declarations.ATTRIBUTE_QP:   m.QP,
		},
		Disassembly: fmt.Sprintf("st%d [r%d] = r%d", bitLengthTable[m.TableX], m.R3, m.R2),
	}

	decoder.ExecutableInstructions = append(decoder.ExecutableInstructions, declarations.IntegerStoreRegister)
	decoder.InstructionStructs = append(decoder.InstructionStructs, instructionStruct)
}

func (decoder *DecoderContext) DecodeIntegerLoad(instructionBits uint64, nextSlot uint64) {
	m := formats.ReadM1(instructionBits, nextSlot)

	bitLengthTable := []uint64{
		1, 2, 4, 8,
	}

	instructionStruct := declarations.InstructionStruct{
		Attributes: declarations.InstructionAttributeMap{
			declarations.ATTRIBUTE_TABX: m.TableX,
			declarations.ATTRIBUTE_TABY: m.TableY,
			declarations.ATTRIBUTE_R1:   m.R1,
			declarations.ATTRIBUTE_R3:   m.R3,
			declarations.ATTRIBUTE_QP:   m.QP,
		},
		Disassembly: fmt.Sprintf("ld%d r%d = [r%d]", bitLengthTable[m.TableX], m.R1, m.R3),
	}

	decoder.ExecutableInstructions = append(decoder.ExecutableInstructions, declarations.IntegerLoadNoBaseUpdateForm)
	decoder.InstructionStructs = append(decoder.InstructionStructs, instructionStruct)
}

func (decoder *DecoderContext) DecodeIntegerLoadStoreExtensions(instructionBits uint64, nextSlot uint64) {
	taby := (instructionBits & (0b0000011110000000000000000000000000000000000000)) >> 37

	if taby >= 12 {
		decoder.DecodeIntegerStore(instructionBits, nextSlot)
	} else {
		decoder.DecodeIntegerLoad(instructionBits, nextSlot)
	}
}

func (decoder *DecoderContext) DecodeIntegerLoadStoreSemaphoreFR1bit(instructionBits uint64, nextSlot uint64) {
	m := (instructionBits & (0b0000100000000000000000000000000000000000000000)) >> 41
	x := (instructionBits & (0b0000000000000100000000000000000000000000000000)) >> 32

	if m == 1 {
		if x == 1 {
			//unused
			return
		}

		//Integer Load +Reg Opcode Extensions
		fmt.Printf("UNIMPLEMENTED!!! DecodeIntegerLoadStoreSemaphoreFR1bit\nm: 1\nx: 0\n\n")
	} else {
		if x == 1 {
			//Semaphore/get FR
			fmt.Printf("UNIMPLEMENTED!!! DecodeIntegerLoadStoreSemaphoreFR1bit\nm: 0\nx: 0\n\n")
			return
		}

		decoder.DecodeIntegerLoadStoreExtensions(instructionBits, nextSlot)
	}
}
