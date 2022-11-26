package execution

import (
	"rosalia64/core/declarations"
)

func ExecuteBranchIndirectReturn(attributes declarations.InstructionAttributeMap) {
	b2___ := attributes[declarations.ATTRIBUTE_B2]
	//tabx_ := attributes[declarations.ATTRIBUTE_TABX]
	//taby_ := attributes[declarations.ATTRIBUTE_TABY]
	//____d := attributes[declarations.ATTRIBUTE_D]
	//___wh := attributes[declarations.ATTRIBUTE_WH]
	//____p := attributes[declarations.ATTRIBUTE_P]
	//btype := attributes[declarations.ATTRIBUTE_BRANCH_TYPE]
	___qp := attributes[declarations.ATTRIBUTE_QP]

	if *RetrievePredicateRegister(___qp) {
		//TODO: register stack frames

		addrToJump := *RetrieveBranchRegister(b2___)

		//TODO: see if this is correct
		if addrToJump == 0 {
			ContinueRunning = false
			return
		}

		addrToII := declarations.InstructionConverter.GetInstructionIndexFromAddress(addrToJump)

		CurrentExecutionContext.InstructionIndex = addrToII
	}
}
