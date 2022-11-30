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

func ExecuteIntegerCompareRegisterForm(attributes declarations.InstructionAttributeMap) {
	__qp := attributes[declarations.ATTRIBUTE_QP]
	__r2 := attributes[declarations.ATTRIBUTE_R2]
	__r3 := attributes[declarations.ATTRIBUTE_R3]
	_pr1 := attributes[declarations.ATTRIBUTE_PR1]
	_pr2 := attributes[declarations.ATTRIBUTE_PR2]
	_prc := attributes[declarations.ATTRIBUTE_PR_COMPLETER]
	cmp4 := attributes[declarations.ATTRIBUTE_CM4]
	cond := attributes[declarations.ATTRIBUTE_COND]

	if *RetrievePredicateRegister(__qp) {
		if _pr1 == _pr2 {
			//illegal operation fault
			return
		}

		r2 := RetrieveGeneralRegister(__r2)
		r3 := RetrieveGeneralRegister(__r3)

		nat := r2.NotAThing || r3.NotAThing

		valR2 := r2.Value
		valR3 := r3.Value

		if cmp4 == 1 {
			valR2 = valR2 ^ ((1 << 31) - 1)
			valR3 = valR3 ^ ((1 << 31) - 1)
		}

		result := false

		switch declarations.ComparisonTypeTbTaC(cond) {
		case declarations.TB_TA_C_LT:
			result = valR2 < valR3
		case declarations.TB_TA_C_LT_NONE:
			result = valR2 < valR3
		case declarations.TB_TA_C_LT_UNC:
			result = valR2 < valR3
		case declarations.TB_TA_C_EQ:
			result = valR2 == valR3
		case declarations.TB_TA_C_NE:
			result = valR2 != valR3
		case declarations.TB_TA_C_GT:
			result = valR2 > valR3
		case declarations.TB_TA_C_LE:
			result = valR2 <= valR3
		case declarations.TB_TA_C_GE:
			result = valR2 >= valR3
		}

		compareRegisterCompleters[declarations.CompareRegisterCompleter(_prc)](_pr1, _pr2, result, nat)
	} else {
		if _prc == uint64(declarations.PR_COMPLETER_UNC) {
			if _pr1 == _pr2 {
				//illegal operation fault
				return
			}

			*RetrievePredicateRegister(_pr1) = false
			*RetrievePredicateRegister(_pr2) = false
		}
	}
}
