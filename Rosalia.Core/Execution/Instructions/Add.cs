using Rosalia.Core.Decoding;
using Rosalia.Core.Execution.Registers;

namespace Rosalia.Core.Execution.Instructions;

public class Add {
    public static ProcessorFault AddImmForm(ExecutionContext context, Dictionary<InstructionAttribute, ulong> attributes) {
        ulong reg1 = attributes[InstructionAttribute.R1];
        ulong reg3 = attributes[InstructionAttribute.R3];
        ulong immd = attributes[InstructionAttribute.Immediate];
        ulong qp   = attributes[InstructionAttribute.QualifyingPredicate];

        ItaniumProcessor processor = context.Machine.Processor;

        if (processor.PredicateRegisters.RetrieveRegister(qp).RetrieveValue()) {
            GeneralRegisterBase r1 = processor.GeneralRegisters.RetrieveRegister(reg1);
            GeneralRegisterBase r3 = processor.GeneralRegisters.RetrieveRegister(reg3);

            ProcessorFault valErr = r1.WriteValue((long) immd + r3.RetrieveValue());
            ProcessorFault natErr = r1.WriteNotAThing(r3.RetrieveNotAThing());

            if (valErr != ProcessorFault.None) {
                return valErr;
            }

            if (natErr != ProcessorFault.None) {
                return natErr;
            }
        }

        return ProcessorFault.None;
    }
}
