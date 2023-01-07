using Rosalia.Core.Decoding;

namespace Rosalia.Core.Execution;

public delegate ProcessorFault ExecutionFunction(ExecutionContext context, Dictionary<InstructionAttribute, ulong> attributes);

public struct ExecutableInstruction {
    public ExecutionFunction                       ExecutionFunction;
    public Dictionary<InstructionAttribute, ulong> Attributes;
    public string                                  Disassembly;

    public override string ToString() => this.Disassembly;
}
