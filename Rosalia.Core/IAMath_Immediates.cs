// ReSharper disable InconsistentNaming
namespace Rosalia.Core;

public partial class IAMath {
    public static long SignExt(ulong i, ulong n) {
        return (long)(((i) << (int)(64 - (n))) >> (int)(64 - (n)));
    }

    public static long Imm22(ulong sign, ulong imm5c, ulong imm9d, ulong imm7b) {
        return SignExt(sign << 21 | imm5c << 16 | imm9d << 7 | imm7b, 22);
    }

    public static long Imm14(ulong sign, ulong imm6d, ulong imm7b) {
        return SignExt(sign << 13 | imm6d << 7 | imm7b, 14);
    }
}
