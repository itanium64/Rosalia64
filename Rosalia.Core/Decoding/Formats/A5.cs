// ReSharper disable InconsistentNaming
namespace Rosalia.Core.Decoding.Formats;

public struct A5 {
    public ulong S;
    public ulong Immediate;
    public ulong R3;
    public ulong R1;
    public ulong Qp;

    public A5(ulong slot, ulong nextSlot) {
        ulong ____s = (slot & (0b00001000000000000000000000000000000000000)) >> 36;
        ulong imm9d = (slot & (0b00000111111111000000000000000000000000000)) >> 27;
        ulong imm5c = (slot & (0b00000000000000111110000000000000000000000)) >> 22;
        ulong ___r3 = (slot & (0b00000000000000000001100000000000000000000)) >> 20;
        ulong imm7b = (slot & (0b00000000000000000000011111110000000000000)) >> 13;
        ulong ___r1 = (slot & (0b00000000000000000000000000001111111000000)) >> 6;
        ulong ___qp = (slot & (0b00000000000000000000000000000000000111111)) >> 0;

        return new A5 {
            S         = ____s,
            Immediate = (ulong)immediate,
            R3        = ___r3,
            R1        = ___r1,
            Qp        = ___qp,
        };
    }
}
