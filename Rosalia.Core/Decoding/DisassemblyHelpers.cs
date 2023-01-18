namespace Rosalia.Core.Decoding;

public class DisassemblyHelpers {
    public static string FormatQualifyingPredicate(ulong qp) {
        if (qp == 0) {
            return "      ";
        }

        return $"(p{qp})".PadRight(6);
    }
}
