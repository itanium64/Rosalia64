package declarations

type CompareRegisterCompleter uint64

const (
	PR_COMPLETER_NONE      CompareRegisterCompleter = 1
	PR_COMPLETER_UNC       CompareRegisterCompleter = 2
	PR_COMPLETER_OR        CompareRegisterCompleter = 3
	PR_COMPLETER_AND       CompareRegisterCompleter = 4
	PR_COMPLETER_OR_AND_CM CompareRegisterCompleter = 5
)