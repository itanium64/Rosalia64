using Rosalia.Core.Execution;

namespace Rosalia.Core.Decoding;

public class DecodingContext : IDisposable {
    private MemoryStream _textStream;
    private BinaryReader _textReader;

    private ulong _instructionIndex;
    private ulong _currentAddress;

    private readonly Dictionary<ulong, ulong> _addressToInstructionIndex;
    private readonly Dictionary<ulong, ulong> _instructionIndexToAddress;

    public readonly List<ExecutableInstruction> ExecutableInstructions;

    public DecodingContext(byte[] textSection, ulong addressBase) {
        this._textStream = new MemoryStream(textSection);
        this._textReader = new BinaryReader(this._textStream);

        this._addressToInstructionIndex = new Dictionary<ulong, ulong>();
        this._instructionIndexToAddress = new Dictionary<ulong, ulong>();

        this._currentAddress = addressBase;

        this.ExecutableInstructions = new List<ExecutableInstruction>();
    }

    public void NextBundle() {
        ulong lo = this._textReader.ReadUInt64();
        ulong hi = this._textReader.ReadUInt64();

        InstructionBundle bundle = new InstructionBundle(lo, hi);

        ExecutionSlotOrders.UnitOrStop[] unitOrder = ExecutionSlotOrders.SlotOrders[bundle.Template];

        this._addressToInstructionIndex.Add(this._currentAddress, this._instructionIndex);
        this._instructionIndexToAddress.Add(this._instructionIndex, this._currentAddress);

        this._currentAddress   += 16;
        this._instructionIndex += 3;

        if (lo == 0 && hi == 0) {
            return;
        }

        ExecutionSlotOrders.UnitOrStop unitSlot0 = ExecutionSlotOrders.UnitOrStop.None;
        ExecutionSlotOrders.UnitOrStop unitSlot1 = ExecutionSlotOrders.UnitOrStop.None;
        ExecutionSlotOrders.UnitOrStop unitSlot2 = ExecutionSlotOrders.UnitOrStop.None;

        int slotOrderIndex = 0;

        while (unitSlot0 == ExecutionSlotOrders.UnitOrStop.None ||
               unitSlot1 == ExecutionSlotOrders.UnitOrStop.None ||
               unitSlot2 == ExecutionSlotOrders.UnitOrStop.None
        ) {
            ExecutionSlotOrders.UnitOrStop currentItem = unitOrder[slotOrderIndex];

            slotOrderIndex++;

            if(currentItem == ExecutionSlotOrders.UnitOrStop.None || currentItem == ExecutionSlotOrders.UnitOrStop.Stop) {
                continue;
            } if (currentItem == ExecutionSlotOrders.UnitOrStop.End) {
                break;
            }

            if (unitSlot0 == ExecutionSlotOrders.UnitOrStop.None) {
                unitSlot0 = currentItem;
            }

            if (unitSlot1 == ExecutionSlotOrders.UnitOrStop.None) {
                unitSlot1 = currentItem;
            }

            if (unitSlot2 == ExecutionSlotOrders.UnitOrStop.None) {
                unitSlot2 = currentItem;
            }
        }

        this.DecodeInstructionSlot(bundle.Slot0, bundle.Slot1, unitSlot0);
        this.DecodeInstructionSlot(bundle.Slot1, bundle.Slot2, unitSlot1);
        this.DecodeInstructionSlot(bundle.Slot2, 0b0000000000, unitSlot2);
    }

    private void DecodeInstructionSlot(ulong slot, ulong nextSlot, ExecutionSlotOrders.UnitOrStop unit) {
        ulong mask = (ulong)0b1111 << 37;
        ulong majorOpcode = ((slot & mask) >> 37);

        InstructionDecoder decoder =
            InstructionTables
                .GetInstructionTable(unit)!
                .GetValueOrDefault(majorOpcode, null);

        if (decoder == null) {
            Console.WriteLine($"Major Opcode {majorOpcode} not implemented for {unit} unit.");
            return;
        }

        decoder(this, slot, nextSlot);
    }

    public ulong ConvertInstructionPointer(ulong ip) {
        return this._addressToInstructionIndex.GetValueOrDefault(ip, 0ul);
    }

    public void Dispose() {
        this._textStream.Dispose();
        this._textReader.Dispose();

        GC.SuppressFinalize(this);
    }
}
