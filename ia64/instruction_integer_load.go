package ia64

import (
	"Rosalia64/ia64/formats"
	"fmt"
)

func IntegerLoadWithRegister(instructionBits uint64, nextSlot uint64) {
	m := formats.ReadM1(instructionBits, nextSlot)

	fmt.Printf("tabx : %d", m.TableX)
	fmt.Printf("taby : %d", m.TableY)
	fmt.Printf("hint : %d", m.Hint)
	fmt.Printf("r1   : %d", m.R1)
	fmt.Printf("r2   : %d", m.R2)
	fmt.Printf("r3   : %d", m.R3)
	fmt.Printf("qp   : %d", m.QP)
}

func IntegerLoadStore(instructionBits uint64, nextSlot uint64) {
	m := formats.ReadM2(instructionBits, nextSlot)

	fmt.Printf("tabx : %d\n", m.TableX)
	fmt.Printf("taby : %d\n", m.TableY)
	fmt.Printf("hint : %d\n", m.Hint)
	fmt.Printf("r2   : %d\n", m.R2)
	fmt.Printf("r3   : %d\n", m.R3)
	fmt.Printf("qp   : %d\n", m.QP)
}
