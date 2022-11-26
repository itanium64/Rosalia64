package execution

import (
	"rosalia64/core/declarations"
)

func ExecuteAddlImm22(attributes declarations.InstructionAttributeMap) {
	reg1 := attributes[declarations.ATTRIBUTE_R1]
	reg3 := attributes[declarations.ATTRIBUTE_R3]
	immd := attributes[declarations.ATTRIBUTE_IMMEDIATE]
	__qp := attributes[declarations.ATTRIBUTE_QP]

	if *RetrievePredicateRegister(__qp) {
		//todo: check_target_register(r1)

		r1 := RetrieveGeneralRegister(reg1)
		r3 := RetrieveGeneralRegister(reg3)

		r1.Value = int64(immd) + r3.Value
		r1.NotAThing = r3.NotAThing
	}
}

func ExecuteAddsImm14(attributes declarations.InstructionAttributeMap) {
	reg1 := attributes[declarations.ATTRIBUTE_R1]
	reg3 := attributes[declarations.ATTRIBUTE_R3]
	immd := attributes[declarations.ATTRIBUTE_IMMEDIATE]
	__qp := attributes[declarations.ATTRIBUTE_QP]

	if *RetrievePredicateRegister(__qp) {
		//todo: check_target_register(r1)

		r1 := RetrieveGeneralRegister(reg1)
		r3 := RetrieveGeneralRegister(reg3)

		r1.Value = int64(immd) + r3.Value
		r1.NotAThing = r3.NotAThing
	}
}
