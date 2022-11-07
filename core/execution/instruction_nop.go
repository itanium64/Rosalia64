package execution

import (
	"fmt"
	"rosalia64/core/declarations"
)

func ExecuteNopBranch(attributes declarations.InstructionAttributeMap) {
	fmt.Printf("\nExecuting: nop.b\n\n")
}

func ExecuteNopMemory(attributes declarations.InstructionAttributeMap) {
	fmt.Printf("\nExecuting: nop.m\n\n")
}

func ExecuteNopInteger(attributes declarations.InstructionAttributeMap) {
	fmt.Printf("\nExecuting: nop.i\n\n")
}
