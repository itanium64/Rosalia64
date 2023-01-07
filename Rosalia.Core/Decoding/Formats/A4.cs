// ReSharper disable InconsistentNaming
namespace Rosalia.Core.Decoding.Formats;

public struct A4 {
    public ulong S;
    public ulong X2a;
    public ulong Ve;
    public ulong Immediate;
    public ulong R3;
    public ulong R1;
    public ulong Qp;

    public A4(ulong slot, ulong nextSlot) {
        ulong sign_ = (slot & (0b00001000000000000000000000000000000000000)) >> 47;
        ulong x2a__ = (slot & (0b00000110000000000000000000000000000000000)) >> 34;
        ulong ve___ = (slot & (0b00000001000000000000000000000000000000000)) >> 33;
        ulong imm6d = (slot & (0b00000000111111000000000000000000000000000)) >> 27;
        ulong r3___ = (slot & (0b00000000000000111111100000000000000000000)) >> 20;
        ulong imm7b = (slot & (0b00000000000000000000011111110000000000000)) >> 13;
        ulong r1___ = (slot & (0b00000000000000000000000000001111111000000)) >> 06;
        ulong qp___ = (slot & (0b00000000000000000000000000000000000111111)) >> 00;

        long immediate = IAMath.Imm14(sign_, imm6d, imm7b);

        this.S         = sign_;
        this.X2a       = x2a__;
        this.Ve        = ve___;
        this.Immediate = (ulong) immediate;
        this.R3        = r3___;
        this.R1        = r1___;
        this.Qp        = qp___;
    }
}
