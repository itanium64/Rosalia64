package execution

import (
	"fmt"
	"rosalia64/core/declarations"
)

func ExecuteAddlImm22(attributes declarations.InstructionAttributeMap) {
	sign := attributes[declarations.ATTRIBUTE_SIGN]
	reg1 := attributes[declarations.ATTRIBUTE_R1]
	reg3 := attributes[declarations.ATTRIBUTE_R3]
	immd := attributes[declarations.ATTRIBUTE_IMMEDIATE]
	__qp := attributes[declarations.ATTRIBUTE_QP]

	fmt.Printf("\nExecuting: addl r1 = imm22, r3\n")
	fmt.Printf("sign     : %d\n", sign)
	fmt.Printf("r1       : %d\n", reg1)
	fmt.Printf("r3       : %d\n", reg3)
	fmt.Printf("imm22    : %d\n", immd)
	fmt.Printf("qp       : %d\n", __qp)

	if RetrievePredicateRegister(__qp) {
		//todo: check_target_register(r1)

		r1 := RetrieveGeneralRegister(reg1)
		r3 := RetrieveGeneralRegister(reg3)

		r1.Value = uint64(immd) + r3.Value
		r1.NotAThing = r3.NotAThing
	}
}

func ExecuteAddsImm14(attributes declarations.InstructionAttributeMap) {
	sign := attributes[declarations.ATTRIBUTE_SIGN]
	reg1 := attributes[declarations.ATTRIBUTE_R1]
	reg3 := attributes[declarations.ATTRIBUTE_R3]
	immd := attributes[declarations.ATTRIBUTE_IMMEDIATE]
	__qp := attributes[declarations.ATTRIBUTE_QP]

	fmt.Printf("\nExecuting: adds r1 = imm14, r3\n")
	fmt.Printf("sign     : %d\n", sign)
	fmt.Printf("r1       : %d\n", reg1)
	fmt.Printf("r3       : %d\n", reg3)
	fmt.Printf("imm14    : %d\n", immd)
	fmt.Printf("qp       : %d\n", __qp)

	if RetrievePredicateRegister(__qp) {
		//todo: check_target_register(r1)

		r1 := RetrieveGeneralRegister(reg1)
		r3 := RetrieveGeneralRegister(reg3)

		r1.Value = uint64(immd) + r3.Value
		r1.NotAThing = r3.NotAThing
	}
}
