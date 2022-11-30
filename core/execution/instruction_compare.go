package execution

import (
	"rosalia64/core/declarations"
)

//PR Completers
func CompareRegisterCompleterAnd(pr1 uint64, pr2 uint64, result bool, nat bool) {
	if nat || !result {
		*RetrievePredicateRegister(pr1) = false
		*RetrievePredicateRegister(pr2) = false
	}
}

func CompareRegisterCompleterOr(pr1 uint64, pr2 uint64, result bool, nat bool) {
	if !nat && result {
		*RetrievePredicateRegister(pr1) = true
		*RetrievePredicateRegister(pr2) = true
	}
}

func CompareRegisterCompleterOrAndCm(pr1 uint64, pr2 uint64, result bool, nat bool) {
	if !nat && result {
		*RetrievePredicateRegister(pr1) = true
		*RetrievePredicateRegister(pr2) = false
	}
}

func CompareRegisterCompleterUncNone(pr1 uint64, pr2 uint64, result bool, nat bool) {
	if nat {
		*RetrievePredicateRegister(pr1) = false
		*RetrievePredicateRegister(pr2) = false
	} else {
		*RetrievePredicateRegister(pr1) = result
		*RetrievePredicateRegister(pr2) = !result
	}
}

var compareRegisterCompleters map[declarations.CompareRegisterCompleter]func(pr1 uint64, pr2 uint64, result bool, nat bool) = map[declarations.CompareRegisterCompleter]func(pr1 uint64, pr2 uint64, result bool, nat bool){
	declarations.PR_COMPLETER_AND:       CompareRegisterCompleterAnd,
	declarations.PR_COMPLETER_OR:        CompareRegisterCompleterOr,
	declarations.PR_COMPLETER_OR_AND_CM: CompareRegisterCompleterOrAndCm,
	declarations.PR_COMPLETER_NONE:      CompareRegisterCompleterUncNone,
	declarations.PR_COMPLETER_UNC:       CompareRegisterCompleterUncNone,
}

func ExecuteIntegerCompareRegisterFormLT(attributes declarations.InstructionAttributeMap) {
	_qp := attributes[declarations.ATTRIBUTE_QP]
	_r2 := attributes[declarations.ATTRIBUTE_R2]
	_r3 := attributes[declarations.ATTRIBUTE_R3]
	pr1 := attributes[declarations.ATTRIBUTE_PR1]
	pr2 := attributes[declarations.ATTRIBUTE_PR2]
	prc := attributes[declarations.ATTRIBUTE_PR_COMPLETER]
	cm4 := attributes[declarations.ATTRIBUTE_CM4]

	if *RetrievePredicateRegister(_qp) {
		r2 := RetrieveGeneralRegister(_r2)
		r3 := RetrieveGeneralRegister(_r3)

		nat := r2.NotAThing || r3.NotAThing

		valR2 := r2.Value
		valR3 := r3.Value

		if cm4 == 1 {
			valR2 = valR2 ^ ((1 << 31) - 1)
			valR3 = valR3 ^ ((1 << 31) - 1)
		}

		result := valR2 < valR3

		compareRegisterCompleters[declarations.CompareRegisterCompleter(prc)](pr1, pr2, result, nat)
	}
}
