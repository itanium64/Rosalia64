package ia64

import "fmt"

func NopBranch(instructionBits uint64, nextSlot uint64) {
	fmt.Printf("\nExecuting: nop.b\n\n")
}

func NopMemory(instructionBits uint64, nextSlot uint64) {
	fmt.Printf("\nExecuting: nop.m\n\n")
}

func NopInteger(instructionBits uint64, nextSlot uint64) {
	fmt.Printf("\nExecuting: nop.i\n\n")
}
