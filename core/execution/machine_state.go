package execution

const (
	MaxGeneralRegisterCount       = 128
	MaxFloatingPointRegisterCount = 128
	MaxPredicateRegisterCount     = 64
	MaxBranchRegisterCount        = 8
)

const (
	SizeKB         = 1024
	SizeMB         = 1024 * 1024
	StackSizeBytes = SizeMB
)

type RegisterID uint64

const (
	RegisterSP RegisterID = 12
)

type GeneralRegister struct {
	RegisterID RegisterID
	Value      int64
	NotAThing  bool
}

type FloatingRegister struct {
	RegisterID RegisterID
	Value      float64
}

type IAProcessorState struct {
	GeneralRegisters    [MaxGeneralRegisterCount]GeneralRegister
	PredicateRegisters  [MaxPredicateRegisterCount]bool
	FloatingRegisters   [MaxFloatingPointRegisterCount]FloatingRegister
	BranchRegisters     [MaxBranchRegisterCount]int64
	RegisterStackEngine RegisterStackEngine
}

var processor IAProcessorState = IAProcessorState{}
var memory []byte

var ContinueRunning bool

var CurrentExecutionContext ExecutionContext

func RetrieveGeneralRegister(r uint64) *GeneralRegister {
	if r == 0 {
		return &GeneralRegister{
			Value:      0,
			NotAThing:  false,
			RegisterID: RegisterID(r),
		}
	}

	if r < 32 {
		return &processor.GeneralRegisters[r]
	} else {
		register := (r - 32) + processor.RegisterStackEngine.CurrentFrame().RegisterBase

		return &processor.GeneralRegisters[register]
	}
}

func RetrievePredicateRegister(pr uint64) *bool {
	trueValue := true

	//PR0 is hardwired to return true
	if pr == 0 {
		return &trueValue
	}

	return &processor.PredicateRegisters[pr]
}

func RetrieveBranchRegister(br uint64) *int64 {
	return &processor.BranchRegisters[br]
}

func RetrieveFloatingPointRegister(fr uint64) *FloatingRegister {
	switch fr {
	case 0:
		return &FloatingRegister{
			Value:      0,
			RegisterID: RegisterID(fr),
		}
	case 1:
		return &FloatingRegister{
			Value:      1,
			RegisterID: RegisterID(fr),
		}
	default:
		return &processor.FloatingRegisters[fr]
	}
}

func InitializeMachine(ram int64) {
	ramSize := ram * 1024

	memory = make([]byte, ramSize)
	ContinueRunning = true

	for i := 0; i != MaxGeneralRegisterCount; i++ {
		processor.GeneralRegisters[i] = GeneralRegister{
			RegisterID: RegisterID(i),
		}
	}

	processor.GeneralRegisters[RegisterSP].Value = ramSize - StackSizeBytes

	processor.FloatingRegisters[0].Value = 0
	processor.FloatingRegisters[1].Value = 1
}
