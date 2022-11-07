package ia64

import (
	"encoding/binary"
	"fmt"
	"rosalia64/core/ia64/formats"
)

func IntegerStoreRegister(m formats.M1_2_4) {
	r2 := RetrieveGeneralRegister(m.R2)
	r3 := RetrieveGeneralRegister(m.R3)

	if RetrievePredicateRegister(m.QP) {
		bitLengthTable := []uint64{
			1, 2, 4, 8,
		}

		countBytes := bitLengthTable[m.TableX]
		regAsBytes := make([]byte, 8)

		fmt.Printf("Executing: st%d [r%d] = r%d\n", countBytes, m.R3, m.R2)

		//regAsBytes
		if r3.NotAThing || r2.NotAThing {
			//register_nat_consumption_fault(WRITE)
			return
		}

		//Translates virtual address to physical
		//paddr = tlb_translate(GR[r3], size, WRITE, PSR.cpl, &mattr, &tmp_unused);
		paddr := r3.Value

		//mem_write(GR[r2], paddr, size, UM.be, mattr, otype, sthint);

		binary.LittleEndian.PutUint64(regAsBytes, r2.Value)

		for i := paddr; i != paddr+countBytes; i++ {
			memory[i] = regAsBytes[i-paddr]
		}
	}
}
