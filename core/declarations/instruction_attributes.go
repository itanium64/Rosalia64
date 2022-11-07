package declarations

type Attribute uint64

const (
	ATTRIBUTE_SIGN        Attribute = 0
	ATTRIBUTE_X2A         Attribute = 1
	ATTRIBUTE_VE          Attribute = 2
	ATTRIBUTE_IMMEDIATE   Attribute = 3
	ATTRIBUTE_R1          Attribute = 4
	ATTRIBUTE_R2          Attribute = 5
	ATTRIBUTE_R3          Attribute = 6
	ATTRIBUTE_QP          Attribute = 7
	ATTRIBUTE_D           Attribute = 8
	ATTRIBUTE_WH          Attribute = 9
	ATTRIBUTE_TABX        Attribute = 10
	ATTRIBUTE_TABY        Attribute = 11
	ATTRIBUTE_B2          Attribute = 12
	ATTRIBUTE_P           Attribute = 13
	ATTRIBUTE_BRANCH_TYPE Attribute = 14
	ATTRIBUTE_M           Attribute = 15
	ATTRIBUTE_HINT        Attribute = 16
	ATTRIBUTE_X           Attribute = 17
)

type InstructionAttributeMap map[Attribute]uint64
