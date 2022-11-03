package ia64

const (
	MaxGeneralRegisterCount   = 128
	MaxPredicateRegisterCount = 64
)

type Register struct {
	Value     uint64
	NotAThing bool
}

type IAProcessorState struct {
	GeneralRegisters   [MaxGeneralRegisterCount]Register
	PredicateRegisters [MaxPredicateRegisterCount]bool
}

var Processor IAProcessorState = IAProcessorState{}

func RetrieveGeneralRegister(r uint64) *Register {
	return &Processor.GeneralRegisters[r]
}

func RetrievePredicateRegister(pr uint64) bool {
	//PR0 is hardwired to return true
	if pr == 1 {
		return true
	}

	return Processor.PredicateRegisters[pr]
}

func SetPredicateRegister(qp uint64, value bool) {
	//We don't have to worry about PR0 because in the Retrieve it always return 1 is PR0 is retrieved.
	Processor.PredicateRegisters[qp] = value
}
