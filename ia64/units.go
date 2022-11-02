package ia64

type Unit uint8

type UnitInstructionTable map[uint64]func(instructionBits uint64, nextSlot uint64)

const (
	Invalid Unit = 0
	I_Unit  Unit = 1
	M_Unit  Unit = 2
	F_Unit  Unit = 3
	B_Unit  Unit = 4
)

type UnitOrder struct {
	Slot0 Unit
	Slot1 Unit
	Slot2 Unit
}

var I_UnitInstructionTable UnitInstructionTable = UnitInstructionTable{
	8: IntegerALU,
}

var M_UnitInstructionTable UnitInstructionTable = UnitInstructionTable{
	9: AddImmediate22,
}

var B_UnitInstructionTable UnitInstructionTable = UnitInstructionTable{
	2: NopBranch,
}

func GetInstructionTable(unit Unit) UnitInstructionTable {
	switch unit {
	case I_Unit:
		return I_UnitInstructionTable
	case M_Unit:
		return M_UnitInstructionTable
	case B_Unit:
		return B_UnitInstructionTable
	}

	return nil
}

var UnitTable map[uint64]UnitOrder = map[uint64]UnitOrder{
	0x11: {
		Slot0: M_Unit,
		Slot1: I_Unit,
		Slot2: B_Unit,
	},
}
