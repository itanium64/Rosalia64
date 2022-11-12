package execution

import "rosalia64/core/structures"

type RegisterStackWindow struct {
	RegisterBase uint64
	SizeOfFrame  uint64

	CountInputRegisters  uint64
	CountLocalRegisters  uint64
	CountOutputRegisters uint64
}

type RegisterStackEngine struct {
	StackWindows structures.Stack[RegisterStackWindow]
}

func (stack *RegisterStackEngine) NewFrame(inputRegisters uint64) {
	currentFrame := stack.CurrentFrame()

	stack.StackWindows.Push(RegisterStackWindow{
		//Subtracting the `inputRegisters` here,
		//Because if there are 2 output registers for example,
		//The beginning 2 registers `r32` and `r33`, have to refrence
		//The same registers to work as input parameters to functions.
		RegisterBase:        (currentFrame.RegisterBase + currentFrame.SizeOfFrame) - inputRegisters,
		CountInputRegisters: inputRegisters,
		SizeOfFrame:         inputRegisters,
	})
}

func (stack *RegisterStackEngine) Allocate(localRegisters uint64, outputRegisters uint64) {
	//TODO: figure out if re-allocation clears registers
	//      and what happens in situaions where for example there are
	//      2 locals and 4 outputs, and it turns to 4 locals and 2 outputs
	//      will the 4 locals have the old values of the outputs? are they cleared?

	currentFrame := stack.CurrentFrame()

	currentFrame.CountLocalRegisters = localRegisters
	currentFrame.CountOutputRegisters = outputRegisters

	currentFrame.SizeOfFrame = currentFrame.CountInputRegisters + currentFrame.CountLocalRegisters + currentFrame.CountOutputRegisters
}

func (stack *RegisterStackEngine) CurrentFrame() *RegisterStackWindow {
	return stack.StackWindows.Top()
}
