package execution

import (
	"fmt"
	"rosalia64/core/declarations"
	"rosalia64/core/misc"
)

func ExecuteIntegerLoadNoBaseUpdateForm(attributes declarations.InstructionAttributeMap) {
	tableX := attributes[declarations.ATTRIBUTE_TABX]
	tableY := attributes[declarations.ATTRIBUTE_TABY]

	qp := attributes[declarations.ATTRIBUTE_QP]

	r1 := RetrieveGeneralRegister(attributes[declarations.ATTRIBUTE_R1])
	r3 := RetrieveGeneralRegister(attributes[declarations.ATTRIBUTE_R3])

	if r1.Value == r3.Value {
		//Illegal operation fault
		return
	}

	if *RetrievePredicateRegister(qp) {
		//TODO: check_target_register(r1)

		bitLengthTable := []int{
			1, 2, 4, 8,
		}

		countBytes := bitLengthTable[tableX]
		var readBytes []byte
		var value uint64

		var speculative, //Speculative execution is a technique to preload stuff before asked
			//advanced, //adds the load to the ALAT table, also sets the target registers NAT to check for detect deferral of the load
			//checkClear,
			//checkNoClear,
			//acquire, //Ordered read, seems to mean that the bytes are read in order, if false they can appear in any order?
			//bias, //Can ignore, apperantly hints to the implementation to acquire exclusive ownership of the cache line containing addressed data.
			//fill,
			_defer bool

		switch tableY {
		case 0:
		case 1:
			speculative = true
		case 2:
			//advanced = true
		case 3:
			speculative = true
			//advanced = true
		case 4:
			//bias = true
		case 5:
			//acquire = true
		case 6:
			countBytes = 8
			//fill = true
		case 8:
			//checkClear = true
		case 9:
			//checkNoClear = true
		case 10:
			//checkClear = true
			//acquire = true
		default:
			fmt.Printf("ld%d load extension not implemented! decimal %d\n", countBytes, tableY)
		}

		//fmt.Printf("Executing: ld%d r%d = [r%d]\n", countBytes, attributes[declarations.ATTRIBUTE_R1], attributes[declarations.ATTRIBUTE_R3])

		//check := checkClear || checkNoClear

		if !speculative && r3.NotAThing {
			//register_nat_consumption_fault(itype)
			return
		}

		//TODO: figure out what PSR is, and what flags do
		_defer = speculative && (r3.NotAThing /* || PSR.ed */)

		//TODO: maybe look into speculative execution?

		/*
			//Checks a lookup table of preloaded addresses,
			if check && alat_cmp(GENERAL, r1) {
				if checkClear {
					//clear the lookup table of that entry if requested
					alat_invalidate_single_entry(GENERAL, r1)
				}
			} else ::::
		*/

		if !_defer {
			//Translates a virtual address to a physical one
			//paddr := tlb_translate(r3, countBytes, itype, PSR.cpl, &mattr, &defer)
			paddr := r3.Value

			if !_defer {
				//readBytes = mem_read(paddr, size, UM.be, mattr, otype, bias | *ldhint*)
				readBytes = memory[paddr : paddr+uint64(countBytes)]
				value = misc.BytesToUint64(readBytes, countBytes)
			}
		}

		/*if checkClear || advanced {
			//clear the lookup table of that entry if requested
			//alat_invalidate_single_entry(GENERAL, r1)
		}*/

		/*if _defer {
			if speculative {
				//executes a speculative read request
				//r1 = natd_gr_read(paddr, size, UM.be, mattr, otype, bias | *ldhint*)
				//r1.NotAThing = true
			} else {
				//r1.Value = 0
				//r1.NotAThing = false
			}
		} else {*/
		/*	if fill {
				//bitPos := r3 bits 8 to 3
				//r1.Value = readBytes as value
				//r1.NotAThing = RetrieveApplicationRegister(UNAT, bitPos)
		/*	} else {*/

		r1.Value = misc.ZeroExt(value, countBytes*8)
		r1.NotAThing = false

		/*if checkNoClear || advanced /* && ma_is_speculative */ /*{
				//alat_write(GENERAL, r1, paddr, size)
			}
		}*/
	}
}
