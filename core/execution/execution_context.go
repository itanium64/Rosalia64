package execution

import "rosalia64/core/declarations"

type ExecutionContext struct {
	InstructionIndex       int64
	Paused                 bool
	ExecutableInstructions []declarations.ExecutableInstruction
	InstructionStructs     []declarations.InstructionStruct
}

func NewExecutionContext(instructions []declarations.ExecutableInstruction, instructionStructs []declarations.InstructionStruct, instructionIndex int64) {
	CurrentExecutionContext = ExecutionContext{
		InstructionIndex:       instructionIndex,
		Paused:                 false,
		ExecutableInstructions: instructions,
		InstructionStructs:     instructionStructs,
	}
}

func (context *ExecutionContext) Step() {
	executable := context.ExecutableInstructions[context.InstructionIndex]
	attributes := context.InstructionStructs[context.InstructionIndex]

	executable(attributes.Attributes)

	context.InstructionIndex++
}

func (context *ExecutionContext) Run() {
	for ContinueRunning && !context.Paused {
		context.Step()
	}
}

func (context *ExecutionContext) Pause() {
	context.Paused = true
}
