using Rosalia.Core.Decoding.Formats;
using Rosalia.Core.Execution;
using Rosalia.Core.Execution.Instructions;

namespace Rosalia.Core.Decoding.Decoders;

public static class AddsImm14Form {
    public static void DecodeAddsImm14Form(DecodingContext context, ulong slot, ulong nextSlot) {
        A4 a4 = new A4(slot, nextSlot);

        string disassembly = $"{DisassemblyHelpers.FormatQualifyingPredicate(a4.Qp)} adds r{a4.R1} = {a4.Immediate}, r{a4.R3}";

        Dictionary<InstructionAttribute, ulong> attributes = new() {
            [InstructionAttribute.R1] = a4.R1,
            [InstructionAttribute.R3] = a4.R3,
            [InstructionAttribute.Immediate] = a4.Immediate,
            [InstructionAttribute.QualifyingPredicate] = a4.Qp,
        };

        ExecutableInstruction instruction = new ExecutableInstruction {
            ExecutionFunction = Add.AddImmForm,
            Attributes        = attributes,
            Disassembly       = disassembly,
        };

        context.ExecutableInstructions.Add(instruction);
    }
}
