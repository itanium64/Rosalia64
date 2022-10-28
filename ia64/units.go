package ia64

type Unit uint8

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

var UnitTable map[uint64]UnitOrder = map[uint64]UnitOrder{
	0x11: {
		Slot0: M_Unit,
		Slot1: I_Unit,
		Slot2: B_Unit,
	},
}
