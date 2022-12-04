package declarations

type CompareRegisterCompleter uint64
type ComparisonTypeTbTaC uint64
type IPRelativeBranchType uint64
type BranchWhetherHint uint64
type IndirectCallWhetherHint uint64
type BranchCacheHint uint64
type SequentialPrefetchHint uint64

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

const (
	IPREL_BRANCH_TYPE_NONE  IPRelativeBranchType = 1
	IPREL_BRANCH_TYPE_COND  IPRelativeBranchType = 2
	IPREL_BRANCH_TYPE_CALL  IPRelativeBranchType = 3
	IPREL_BRANCH_TYPE_RET   IPRelativeBranchType = 4
	IPREL_BRANCH_TYPE_IA    IPRelativeBranchType = 5
	IPREL_BRANCH_TYPE_CLOOP IPRelativeBranchType = 6
	IPREL_BRANCH_TYPE_CTOP  IPRelativeBranchType = 7
	IPREL_BRANCH_TYPE_CEXIT IPRelativeBranchType = 8
	IPREL_BRANCH_TYPE_WTOP  IPRelativeBranchType = 9
	IPREL_BRANCH_TYPE_WEXIT IPRelativeBranchType = 10
)

const (
	BRWH_SPECULATE_TAKEN     BranchWhetherHint = 0
	BRWH_SPECULATE_NOT_TAKEN BranchWhetherHint = 1
	BRWH_DYNAMIC_TAKEN       BranchWhetherHint = 2
	BRWH_DYNAMIC_NOT_TAKEN   BranchWhetherHint = 3
)

const (
	ICWH_SPECULATE_TAKEN     IndirectCallWhetherHint = 1
	ICWH_SPECULATE_NOT_TAKEN IndirectCallWhetherHint = 3
	ICWH_DYNAMIC_TAKEN       IndirectCallWhetherHint = 5
	ICWH_DYNAMIC_NOT_TAKEN   IndirectCallWhetherHint = 7
)

const (
	BRANCH_CACHE_CLEAR BranchCacheHint = 1
)

const (
	PREFETCH_HINT_FEW  SequentialPrefetchHint = 0
	PREFETCH_HINT_MANY SequentialPrefetchHint = 1
)
