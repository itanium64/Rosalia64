package ia64

const (
	MaxGeneralRegisterCount = 128
)

type Register struct {
	Value     uint64
	NotAThing bool
}

type IAProcessorState struct {
	GeneralRegisters [MaxGeneralRegisterCount]Register
}

var Processor IAProcessorState = IAProcessorState{}
