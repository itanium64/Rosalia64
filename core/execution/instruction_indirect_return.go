package execution

import (
	"fmt"
	"rosalia64/core/declarations"
)

func ExecuteBranchIndirectReturn(attributes declarations.InstructionAttributeMap) {
	b2___ := attributes[declarations.ATTRIBUTE_B2]
	tabx_ := attributes[declarations.ATTRIBUTE_TABX]
	taby_ := attributes[declarations.ATTRIBUTE_TABY]
	____d := attributes[declarations.ATTRIBUTE_D]
	___wh := attributes[declarations.ATTRIBUTE_WH]
	____p := attributes[declarations.ATTRIBUTE_P]
	btype := attributes[declarations.ATTRIBUTE_BRANCH_TYPE]
	___qp := attributes[declarations.ATTRIBUTE_QP]

	fmt.Printf("Executing: br b%d\n", b2___)
	fmt.Printf("tableX   : %d\n", tabx_)
	fmt.Printf("tableY   : %d\n", taby_)
	fmt.Printf("D        : %d\n", ____d)
	fmt.Printf("WH       : %d\n", ___wh)
	fmt.Printf("B2       : %d\n", b2___)
	fmt.Printf("P        : %d\n", ____p)
	fmt.Printf("BType    : %d\n", btype)
	fmt.Printf("QP       : %d\n", ___qp)

	if *RetrievePredicateRegister(___qp) {
		//TODO: register stack frames
		//TODO: see if this is correct
		if *RetrieveBranchRegister(b2___) == 0 {
			ContinueRunning = false
		}
	}
}
