package execution

import (
	"encoding/binary"
	"rosalia64/core/declarations"
)

func ExecuteIntegerStoreRegister(attributes declarations.InstructionAttributeMap) {
	tx := attributes[declarations.ATTRIBUTE_TABX]
	qp := attributes[declarations.ATTRIBUTE_QP]

	r2 := RetrieveGeneralRegister(attributes[declarations.ATTRIBUTE_R2])
	r3 := RetrieveGeneralRegister(attributes[declarations.ATTRIBUTE_R3])

	if *RetrievePredicateRegister(qp) {
		bitLengthTable := []int64{
			1, 2, 4, 8,
		}

		countBytes := bitLengthTable[tx]
		regAsBytes := make([]byte, 8)

		//regAsBytes
		if r3.NotAThing || r2.NotAThing {
			//register_nat_consumption_fault(WRITE)
			return
		}

		//Translates virtual address to physical
		//paddr = tlb_translate(GR[r3], size, WRITE, PSR.cpl, &mattr, &tmp_unused);
		paddr := r3.Value

		//mem_write(GR[r2], paddr, size, UM.be, mattr, otype, sthint);

		binary.LittleEndian.PutUint64(regAsBytes, uint64(r2.Value))

		for i := paddr; i != paddr+countBytes; i++ {
			memory[i] = regAsBytes[i-paddr]
		}
	}
}
