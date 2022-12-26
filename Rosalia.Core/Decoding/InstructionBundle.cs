namespace Rosalia.Core.Decoding;

public struct InstructionBundle {
    public ulong Template;
    public ulong Slot0;
    public ulong Slot1;
    public ulong Slot2;

    public InstructionBundle(ulong lo, ulong hi) {
        ulong templ = (lo & 0b11111);
        ulong slot0 = (lo & 0b000000000001111111111111111111111111111111111111111100000) >> 5;
        ulong slot1 = (lo & 0b111111111110000000000000000000000000000000000000000000000) >> 46 |
                      (hi & 0b000000000000000000000000000111111111111111111111111111111) << 18;
        ulong slot2 = (hi & 0b1111111111111111111111111111111111111111100000000000000000000000) >> 23;

        this.Template = templ;
        this.Slot0    = slot0;
        this.Slot1    = slot1;
        this.Slot2    = slot2;
    }
}
