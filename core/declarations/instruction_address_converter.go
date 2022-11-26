package declarations

type InstructionAddressConverter interface {
	GetAddressFromInstructionIndex(address uint64) int64
	GetInstructionIndexFromAddress(address int64) uint64
}

var AddressToInstructionIndex InstructionAddressConverter
