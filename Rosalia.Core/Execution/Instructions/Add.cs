using Rosalia.Core.Decoding;

namespace Rosalia.Core.Execution.Instructions;

public class Add {
    public static ProcessorFault AddlImm22Form(ExecutionContext context, Dictionary<InstructionAttribute, ulong> attributes) {
        return ProcessorFault.None;
    }
}
