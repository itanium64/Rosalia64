package decoding

import (
	"fmt"
	"rosalia64/core/declarations"

	"lukechampine.com/uint128"
)

type DecoderContext struct {
	ExecutableInstructions []declarations.ExecutableInstruction
	InstructionStructs     []declarations.InstructionStruct
}

func (decoder *DecoderContext) NextBundle(bundle [16]byte) {
	asUint128 := uint128.FromBytes(bundle[:])

	if asUint128.Hi == 0 && asUint128.Lo == 0 {
		return
	}

	template := asUint128.Lo & 0b11111

	unitOrder := UnitTable[template]

	slot0 := (asUint128.Lo & 0b000000000001111111111111111111111111111111111111111100000)
	slot1 := (asUint128.Lo&0b111111111110000000000000000000000000000000000000000000000)>>41 |
		(asUint128.Hi&0b000000000000000000000000000111111111111111111111111111111)<<23
	slot2 := (asUint128.Hi & 0b1111111111111111111111111111111111111111100000000000000000000000) >> 18

	//fmt.Printf("\n\n\nNEW BUNDLE: template (decimal): %d\n\n\n", template)
	//
	//fmt.Printf("high : %064b\n", asUint128.Hi)
	//fmt.Printf("low  :                                                                 %064b\n     :\n", asUint128.Lo)
	//fmt.Printf("whole: %064b%064b\n", asUint128.Hi, asUint128.Lo)
	//fmt.Printf("slot0:                                                                 %064b\n", slot0)
	//fmt.Printf("slot1:                        %064b\n", slot1)
	//fmt.Printf("slot2: %064b\n", slot2<<18)

	decoder.decodeInstructionSlot(slot0, slot1, unitOrder.Slot0)
	decoder.decodeInstructionSlot(slot1, slot2, unitOrder.Slot1)
	decoder.decodeInstructionSlot(slot2, 0b000, unitOrder.Slot2)
}

func (decoder *DecoderContext) decodeInstructionSlot(slot uint64, nextSlot uint64, unit Unit) {
	majorOpcode := slot & (0b1111 << 42) >> 42

	table := GetInstructionTable(unit)

	instruction, exists := table[majorOpcode]

	if !exists {
		fmt.Printf("\nUNIMPLEMENTED!!!: \n")
		fmt.Printf("unit : %s\n", UnitToString(unit))
		fmt.Printf("major: %d\n", majorOpcode)

		return
	}

	instruction(slot, nextSlot)
}
