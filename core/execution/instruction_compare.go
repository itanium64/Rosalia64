package execution

import (
	"rosalia64/core/declarations"
)

//PR Completers
func CompareRegisterCompleterAnd(pr1 uint64, pr2 uint64, result bool, nat bool) {
	if nat || !result {
		*RetrievePredicateRegister(pr1) = 0
		*RetrievePredicateRegister(pr2) = 0
	}
}

func CompareRegisterCompleterOr(pr1 uint64, pr2 uint64, result bool, nat bool) {
	if !nat && result {
		*RetrievePredicateRegister(pr1) = 1
		*RetrievePredicateRegister(pr2) = 1
	}
}

func CompareRegisterCompleterOrAndCm(pr1 uint64, pr2 uint64, result bool, nat bool) {
	if !nat && result {
		*RetrievePredicateRegister(pr1) = 1
		*RetrievePredicateRegister(pr2) = 0
	}
}

func CompareRegisterCompleterUncNone(pr1 uint64, pr2 uint64, result bool, nat bool) {
	if nat {
		*RetrievePredicateRegister(pr1) = 0
		*RetrievePredicateRegister(pr2) = 0
	} else {
		*RetrievePredicateRegister(pr1) = result
		*RetrievePredicateRegister(pr2) = !result
	}
}

var compareRegisterCompleters map[uint64]func(result bool, nat bool) = map[uint64]func(pr1 uint64, pr2 uint64, result bool, nat bool){
	declarations.PR_COMPLETER_AND:       CompareRegisterCompleterAnd,
	declarations.PR_COMPLETER_OR:        CompareRegisterCompleterOr,
	declarations.PR_COMPLETER_OR_AND_CM: CompareRegisterCompleterOrAndCm,
	declarations.PR_COMPLETER_NONE:      CompareRegisterCompleterUncNone,
	declarations.PR_COMPLETER_UNC:       CompareRegisterCompleterUncNone
}

func ExecuteCompareRegisterFormLT(attributes declarations.InstructionAttributeMap) {
	if *RetrievePredicateRegister(attributes.QP) {
		r2 := RetrieveGeneralRegister(attributes.R2)
		r3 := RetrieveGeneralRegister(attributes.R3)

		nat := r2.NotAThing || r3.NotAThing

		result := r2.Value < r3.Value

		compareRegisterCompleters[attributes.ATTRIBUTE_PR_COMPLETER](attributes.PR1, attributes.PR2, result, nat)
	}
}