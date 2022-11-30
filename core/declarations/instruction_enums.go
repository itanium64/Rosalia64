package declarations

type CompareRegisterCompleter uint64
type ComparisonTypeTbTaC uint64

const (
	PR_COMPLETER_NONE      CompareRegisterCompleter = 1
	PR_COMPLETER_UNC       CompareRegisterCompleter = 2
	PR_COMPLETER_OR        CompareRegisterCompleter = 3
	PR_COMPLETER_AND       CompareRegisterCompleter = 4
	PR_COMPLETER_OR_AND_CM CompareRegisterCompleter = 5
)

const (
	TB_TA_C_LT_NONE ComparisonTypeTbTaC = 0
	TB_TA_C_LT_UNC  ComparisonTypeTbTaC = 1
	TB_TA_C_EQ      ComparisonTypeTbTaC = 2
	TB_TA_C_NE      ComparisonTypeTbTaC = 3
	TB_TA_C_GT      ComparisonTypeTbTaC = 4
	TB_TA_C_LE      ComparisonTypeTbTaC = 5
	TB_TA_C_GE      ComparisonTypeTbTaC = 6
	TB_TA_C_LT      ComparisonTypeTbTaC = 7
)
