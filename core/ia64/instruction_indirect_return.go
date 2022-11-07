package ia64

import (
	"fmt"
	"rosalia64-core/ia64/formats"
)

func BranchIndirectReturn(instructionBits uint64, nextSlot uint64) {
	b4 := formats.ReadB4(instructionBits, nextSlot)

	fmt.Printf("Executing: br b%d\n", b4.B2)
	fmt.Printf("tableX   : %d\n", b4.TableX)
	fmt.Printf("tableY   : %d\n", b4.TableY)
	fmt.Printf("D        : %d\n", b4.D)
	fmt.Printf("WH       : %d\n", b4.WH)
	fmt.Printf("B2       : %d\n", b4.B2)
	fmt.Printf("P        : %d\n", b4.P)
	fmt.Printf("BType    : %d\n", b4.BType)
	fmt.Printf("QP       : %d\n", b4.QP)

	//TODO: do this correctly
	ContinueRunning = false
}
