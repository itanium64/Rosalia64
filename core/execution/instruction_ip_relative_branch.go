package execution

import "rosalia64/core/declarations"

func ExecuteIPRelativeBranchCond(attributes declarations.InstructionAttributeMap) {
	___qp := attributes[declarations.ATTRIBUTE_QP]

	if *RetrievePredicateRegister(___qp) {

	}
}
