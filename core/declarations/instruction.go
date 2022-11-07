package declarations

type InstructionStruct struct {
	Attributes  InstructionAttributeMap
	Disassembly string
}

type ExecutableInstruction func(attributes InstructionAttributeMap)
