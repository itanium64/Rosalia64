package declarations

type InstructionAddressConverter interface {
	GetAddressFromInstructionIndex(address uint64) int64
	GetInstructionIndexFromAddress(address int64) int64
}

var InstructionConverter InstructionAddressConverter
