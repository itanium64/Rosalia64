using Rosalia.Core.Decoding.Decoders;

namespace Rosalia.Core.Decoding;

public delegate void InstructionDecoder(DecodingContext context, ulong slot, ulong nextSlot);

public static class InstructionTables {
    public static readonly Dictionary<ulong, InstructionDecoder> IntegerUnitInstructionTable = new() {
        [9] = AddlImm22Form.DecodeAddlImm22Form
    };

    public static readonly Dictionary<ulong, InstructionDecoder> MemoryUnitInstructionTable = new() {
        [9] = AddlImm22Form.DecodeAddlImm22Form
    };

    public static readonly Dictionary<ulong, InstructionDecoder> FloatUnitInstructionTable = new() {

    };

    public static readonly Dictionary<ulong, InstructionDecoder> BranchUnitInstructionTable = new() {

    };

    public static readonly Dictionary<ulong, InstructionDecoder> ExtendedUnitInstructionTable = new() {

    };

    public static Dictionary<ulong, InstructionDecoder> GetInstructionTable(ExecutionSlotOrders.UnitOrStop unit) {
        return unit switch {
            ExecutionSlotOrders.UnitOrStop.Memory   => MemoryUnitInstructionTable,
            ExecutionSlotOrders.UnitOrStop.Integer  => IntegerUnitInstructionTable,
            ExecutionSlotOrders.UnitOrStop.Float    => FloatUnitInstructionTable,
            ExecutionSlotOrders.UnitOrStop.Branch   => BranchUnitInstructionTable,
            ExecutionSlotOrders.UnitOrStop.Extended => ExtendedUnitInstructionTable,
            ExecutionSlotOrders.UnitOrStop.None     => throw new ArgumentOutOfRangeException(nameof(unit), unit, "Invalid Unit."),
            ExecutionSlotOrders.UnitOrStop.Stop     => throw new ArgumentOutOfRangeException(nameof(unit), unit, "Invalid Unit."),
            ExecutionSlotOrders.UnitOrStop.End      => throw new ArgumentOutOfRangeException(nameof(unit), unit, "Invalid Unit."),
            _                                       => throw new ArgumentOutOfRangeException(nameof(unit), unit, "Invalid Unit.")
        };
    }
}
