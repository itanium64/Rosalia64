using Rosalia.Core.Decoding;

namespace Rosalia.Core.Execution;

public delegate ProcessorFault ExecutionFunction(ExecutionContext context, Dictionary<InstructionAttribute, ulong> attributes);

public struct ExecutableInstruction {
    public ExecutionFunction                       ExecutionFunction;
    public Dictionary<InstructionAttribute, ulong> Attributes;
    public string                                  Disassembly;

    public override string ToString() => this.Disassembly;
}

public class ExecutionContext {
    public readonly DecodingContext DecodingContext;
    public readonly ItaniumMachine  Machine;

    private bool            _paused;

    public ExecutionContext(DecodingContext decodingContext, ItaniumMachine machine) {
        this.DecodingContext = decodingContext;
        this.Machine         = machine;

        this._paused          = false;
    }

    public void Step() {
        ulong instructionIndex = this.DecodingContext.ConvertInstructionPointer(this.Machine.Processor.InstructionPointer);
        ExecutableInstruction instruction = this.DecodingContext.ExecutableInstructions[(int)instructionIndex];

        ProcessorFault fault = instruction.ExecutionFunction(this, instruction.Attributes);

        if (fault != ProcessorFault.None) {
            Console.WriteLine($"Fault thrown: {fault}; On: {instruction.Disassembly}");
        }
    }

    public void Pause() {
        this._paused = true;
    }

    public void Run() {
        while (this.Machine.ContinueRunning && !this._paused) {
            this.Step();
        }

        Console.WriteLine($"IA64 Return Code: {this.Machine.Processor.GeneralRegisters.RetrieveRegister(8).RetrieveValue()}");
    }
}
