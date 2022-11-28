package decoding

type Unit uint8

type UnitInstructionTable map[uint64]func(instructionBits uint64, nextSlot uint64)

const (
	Invalid Unit = 0
	I_Unit  Unit = 1
	M_Unit  Unit = 2
	F_Unit  Unit = 3
	B_Unit  Unit = 4
	L_Unit  Unit = 6
	X_Unit  Unit = 7
)

type UnitOrder struct {
	Slot0 Unit
	Slot1 Unit
	Slot2 Unit
}

var I_UnitInstructionTable UnitInstructionTable = UnitInstructionTable{}
var M_UnitInstructionTable UnitInstructionTable = UnitInstructionTable{}
var B_UnitInstructionTable UnitInstructionTable = UnitInstructionTable{}
var F_UnitInstructionTable UnitInstructionTable = UnitInstructionTable{}
var L_UnitInstructionTable UnitInstructionTable = UnitInstructionTable{}
var X_UnitInstructionTable UnitInstructionTable = UnitInstructionTable{}

func GetInstructionTable(unit Unit) UnitInstructionTable {
	switch unit {
	case I_Unit:
		return I_UnitInstructionTable
	case M_Unit:
		return M_UnitInstructionTable
	case F_Unit:
		return F_UnitInstructionTable
	case B_Unit:
		return B_UnitInstructionTable
	case L_Unit:
		return L_UnitInstructionTable
	case X_Unit:
		return X_UnitInstructionTable
	}

	return nil
}

func UnitToString(unit Unit) string {
	switch unit {
	case I_Unit:
		return "I Unit"
	case M_Unit:
		return "M Unit"
	case F_Unit:
		return "F Unit"
	case B_Unit:
		return "B Unit"
	case L_Unit:
		return "L Unit"
	case X_Unit:
		return "X Unit"
	}

	return ""
}

var UnitTable map[uint64]UnitOrder = map[uint64]UnitOrder{
	0x00: {
		Slot0: M_Unit,
		Slot1: I_Unit,
		Slot2: I_Unit,
	},
	0x01: {
		Slot0: M_Unit,
		Slot1: I_Unit,
		Slot2: I_Unit,
		//Stop
	},
	0x02: {
		Slot0: M_Unit,
		Slot1: I_Unit,
		//Stop
		Slot2: I_Unit,
	},
	0x03: {
		Slot0: M_Unit,
		Slot1: I_Unit,
		//Stop
		Slot2: I_Unit,
		//Stop
	},
	0x04: {
		Slot0: M_Unit,
		Slot1: L_Unit,
		Slot2: X_Unit,
	},
	0x05: {
		Slot0: M_Unit,
		Slot1: L_Unit,
		Slot2: X_Unit,
		//Stop
	},

	//0x06 and 0x07 don't seem to have any units

	0x08: {
		Slot0: M_Unit,
		Slot1: M_Unit,
		Slot2: I_Unit,
	},
	0x09: {
		Slot0: M_Unit,
		Slot1: M_Unit,
		Slot2: I_Unit,
		//Stop
	},
	0x0A: {
		Slot0: M_Unit,
		//Stop
		Slot1: M_Unit,
		Slot2: I_Unit,
	},
	0x0B: {
		Slot0: M_Unit,
		//Stop
		Slot1: M_Unit,
		Slot2: I_Unit,
		//Stop
	},
	0x0C: {
		Slot0: M_Unit,
		Slot1: F_Unit,
		Slot2: I_Unit,
	},
	0x0D: {
		Slot0: M_Unit,
		Slot1: F_Unit,
		Slot2: I_Unit,
		//Stop
	},
	0x0E: {
		Slot0: M_Unit,
		Slot1: M_Unit,
		Slot2: F_Unit,
	},
	0x0F: {
		Slot0: M_Unit,
		Slot1: M_Unit,
		Slot2: F_Unit,
		//Stop
	},
	0x10: {
		Slot0: M_Unit,
		Slot1: I_Unit,
		Slot2: B_Unit,
	},
	0x11: {
		Slot0: M_Unit,
		Slot1: I_Unit,
		Slot2: B_Unit,
		//Stop
	},
	0x12: {
		Slot0: M_Unit,
		Slot1: B_Unit,
		Slot2: B_Unit,
	},
	0x13: {
		Slot0: M_Unit,
		Slot1: B_Unit,
		Slot2: B_Unit,
		//Stop
	},

	//0x14 and 0x15 don't seem to have any units

	0x16: {
		Slot0: B_Unit,
		Slot1: B_Unit,
		Slot2: B_Unit,
	},
	0x17: {
		Slot0: B_Unit,
		Slot1: B_Unit,
		Slot2: B_Unit,
		//Stop
	},
	0x18: {
		Slot0: M_Unit,
		Slot1: M_Unit,
		Slot2: B_Unit,
	},
	0x19: {
		Slot0: M_Unit,
		Slot1: M_Unit,
		Slot2: B_Unit,
		//Stop
	},

	//0x1A and 0x1B don't seem to have any units

	0x1C: {
		Slot0: M_Unit,
		Slot1: F_Unit,
		Slot2: B_Unit,
	},
	0x1D: {
		Slot0: M_Unit,
		Slot1: F_Unit,
		Slot2: B_Unit,
		//Stop
	},

	//0x1E and 0x1F don't seem to have any units
}

var DecodingContext *DecoderContext

func InitializeDecoderAndTables() {
	DecodingContext = &DecoderContext{
		AddressToInstructionIndex: make(map[uint64]int64),
		InstructionIndexToAddress: make(map[int64]int64),
	}

	B_UnitInstructionTable[0] = DecodingContext.BranchIndirectMiscellaneous
	B_UnitInstructionTable[2] = DecodingContext.DecodeNopBranch

	I_UnitInstructionTable[0] = DecodingContext.DecodeIntegerMisc3bit
	I_UnitInstructionTable[8] = DecodingContext.DecodeIntegerALU

	M_UnitInstructionTable[0] = DecodingContext.DecodeSystemMemoryManagment3bit
	M_UnitInstructionTable[4] = DecodingContext.DecodeIntegerLoadStoreSemaphoreFR1bit
	M_UnitInstructionTable[8] = DecodingContext.DecodeIntegerALU
	M_UnitInstructionTable[9] = DecodingContext.DecodeAddImmediate22
	M_UnitInstructionTable[12] = DecodingContext.DecodeIntegerCompareOpcodeC
}
