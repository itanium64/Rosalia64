package declarations

type Attribute uint64

const (
	ATTRIBUTE_SIGN         Attribute = 1
	ATTRIBUTE_X2A          Attribute = 2
	ATTRIBUTE_VE           Attribute = 3
	ATTRIBUTE_IMMEDIATE    Attribute = 4
	ATTRIBUTE_R1           Attribute = 5
	ATTRIBUTE_R2           Attribute = 6
	ATTRIBUTE_R3           Attribute = 7
	ATTRIBUTE_QP           Attribute = 8
	ATTRIBUTE_D            Attribute = 9
	ATTRIBUTE_WH           Attribute = 10
	ATTRIBUTE_TABX         Attribute = 11
	ATTRIBUTE_TABY         Attribute = 12
	ATTRIBUTE_B2           Attribute = 13
	ATTRIBUTE_P            Attribute = 14
	ATTRIBUTE_BRANCH_TYPE  Attribute = 15
	ATTRIBUTE_M            Attribute = 16
	ATTRIBUTE_HINT         Attribute = 17
	ATTRIBUTE_X            Attribute = 18
	ATTRIBUTE_PR_COMPLETER Attribute = 19
	ATTRIBUTE_PR1          Attribute = 20
	ATTRIBUTE_PR2          Attribute = 21
	ATTRIBUTE_CM4          Attribute = 22
	ATTRIBUTE_COND         Attribute = 23
)

type InstructionAttributeMap map[Attribute]uint64
