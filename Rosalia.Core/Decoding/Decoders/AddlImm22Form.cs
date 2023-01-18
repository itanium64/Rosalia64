using Rosalia.Core.Decoding.Formats;
using Rosalia.Core.Execution;
using Rosalia.Core.Execution.Instructions;

namespace Rosalia.Core.Decoding.Decoders;

public class AddlImm22Form {
    // Tags for easier searching:
    // Add Immediate 22 imm22_form addl Imm
    public static void DecodeAddlImm22Form(DecodingContext context, ulong slot, ulong nextSlot) {
        A5 a5 = new A5(slot, nextSlot);

        string disassembly = $"{DisassemblyHelpers.FormatQualifyingPredicate(a5.Qp)} addl r{a5.R1} = {a5.Immediate}, {a5.R3}";

        Dictionary<InstructionAttribute, ulong> attributes = new() {
            [InstructionAttribute.R1]                  = a5.R1,
            [InstructionAttribute.R3]                  = a5.R3,
            [InstructionAttribute.QualifyingPredicate] = a5.Qp,
            [InstructionAttribute.Immediate]           = a5.Immediate
        };

        ExecutableInstruction instruction = new ExecutableInstruction {
            ExecutionFunction = Add.AddImmForm,
            Attributes        = attributes,
            Disassembly       = disassembly
        };

        context.ExecutableInstructions.Add(instruction);
    }
}
