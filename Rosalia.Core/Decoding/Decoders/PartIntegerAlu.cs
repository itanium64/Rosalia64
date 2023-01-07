namespace Rosalia.Core.Decoding.Decoders;

public static class IntegerAlu {
    // Tags for easier searching
    // integer alu 2+1 2bit+1bit 2-bit+1-bit opcode extensions
    public static void DecodePartIntegerAlu(DecodingContext context, ulong slot, ulong nextSlot) {
        ulong x2a = (slot & (0b00000110000000000000000000000000000000000)) >> 34;
        ulong ve  = (slot & (0b00000001000000000000000000000000000000000)) >> 33;

        switch (x2a) {
            case 2:
                if (ve == 0) {
                    AddsImm14Form.DecodeAddsImm14Form(context, slot, nextSlot);
                } else {
                    Console.WriteLine("DecodePartIntegerAlu: ve = 1 unimplemented.");
                }
                break;
            default:
                Console.WriteLine($"DecodePartIntegerAlu: ve = {ve}; x2a: {x2a} unimplemented");
                break;
        }
    }
}
