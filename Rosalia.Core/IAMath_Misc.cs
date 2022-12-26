namespace Rosalia.Core;

public partial class IAMath {
    public static long IntPow(long x, long n)
    {
        if (n == 0) {
            return 1;
        }

        if (n == 1) {
            return x;
        }

        long y = IntPow(x, 2);

        if ((n % 2) == 0) {
            return y * y;
        }

        return x * y * y;
    }

    public static long ZeroExt(long val, long pos) {
        long and = IntPow(2, pos) - 1;

        return val & and;
    }
}
