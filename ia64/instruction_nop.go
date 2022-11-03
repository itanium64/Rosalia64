package ia64

import "fmt"

func NopBranch(instructionBits uint64, nextSlot uint64) {
	fmt.Printf("\nExecuting: nop.b\n")
}
