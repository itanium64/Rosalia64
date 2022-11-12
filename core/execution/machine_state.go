package execution

const (
	MaxGeneralRegisterCount   = 128
	MaxPredicateRegisterCount = 64
)

const (
	SizeKB         = 1024
	SizeMB         = 1024 * 1024
	StackSizeBytes = SizeMB
)

type RegisterID uint64

const (
	RegisterEAX RegisterID = 8
	RegisterSP  RegisterID = 12
)

type Register struct {
	RegisterID RegisterID
	Value      uint64
	NotAThing  bool
}

type IAProcessorState struct {
	GeneralRegisters    [MaxGeneralRegisterCount]Register
	PredicateRegisters  [MaxPredicateRegisterCount]bool
	RegisterStackEngine RegisterStackEngine
}

var processor IAProcessorState = IAProcessorState{}
var memory []byte

var ContinueRunning bool

var CurrentExecutionContext ExecutionContext

func RetrieveGeneralRegister(r uint64) *Register {
	return &processor.GeneralRegisters[r]
}

func RetrievePredicateRegister(pr uint64) bool {
	//PR0 is hardwired to return true
	if pr == 0 {
		return true
	}

	return processor.PredicateRegisters[pr]
}

func SetPredicateRegister(qp uint64, value bool) {
	//We don't have to worry about PR0 because in the Retrieve it always return 1 is PR0 is retrieved.
	processor.PredicateRegisters[qp] = value
}

func InitializeMachine(ram uint64) {
	ramSize := ram * 1024

	memory = make([]byte, ramSize)
	ContinueRunning = true

	for i := 0; i != MaxGeneralRegisterCount; i++ {
		processor.GeneralRegisters[i] = Register{
			RegisterID: RegisterID(i),
		}
	}

	processor.GeneralRegisters[12].Value = ramSize - StackSizeBytes
}
