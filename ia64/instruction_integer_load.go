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

func IntegerLoad(m formats.M1_2_4) {
	r1 := RetrieveGeneralRegister(m.R1)
	r3 := RetrieveGeneralRegister(m.R3)

	if r1.Value == r3.Value {
		//Illegal operation fault
		return
	}

	if RetrievePredicateRegister(m.QP) {
		//TODO: check_target_register(r1)

		bitLengthTable := []int{
			1, 2, 4, 8,
		}

		countBytes := bitLengthTable[m.TableX]
		var readBytes []byte

		var speculative, advanced, checkClear, checkNoClear, acquire, bias, fill, _defer bool

		switch m.TableY {
		case 1:
			speculative = true
		case 2:
			advanced = true
		case 3:
			speculative = true
			advanced = true
		case 4:
			bias = true
		case 5:
			acquire = true
		case 6:
			countBytes = 8
			fill = true
		case 8:
			checkClear = true
		case 9:
			checkNoClear = true
		case 10:
			checkClear = true
			acquire = true
		default:
			fmt.Printf("ld%d load extension not implemented! decimal %d", m.TableY)
		}

		//check := checkClear || checkNoClear

		if !speculative && r3.NotAThing {
			//register_nat_consumption_fault(itype)
		}

		//TODO: figure out what PSR is, and what flags do
		_defer = speculative && (r3.NotAThing /* || PSR.ed */)

		//TODO: maybe look into speculative execution?

		/*
			if check && alat_cmp(GENERAL, r1) {
				if checkClear {
					alat_invalidate_single_entry(GENERAL, r1)
				}
			} else ::::
		*/

		if !_defer {
			//paddr := tlb_translate(r3, countBytes, itype, PSR.cpl, &mattr, &defer)

			if !_defer {
				
			}
		}
	}
}

func IntegerLoadStore(instructionBits uint64, nextSlot uint64) {
	m := formats.ReadM1(instructionBits, nextSlot)

	if m.TableY >= 12 {
		//store
	} else {
		IntegerLoad(m)
	}

	fmt.Printf("tabx : %d\n", m.TableX)
	fmt.Printf("taby : %d\n", m.TableY)
	fmt.Printf("hint : %d\n", m.Hint)
	fmt.Printf("r2   : %d\n", m.R2)
	fmt.Printf("r3   : %d\n", m.R3)
	fmt.Printf("qp   : %d\n", m.QP)
}
