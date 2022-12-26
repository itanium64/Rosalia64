namespace Rosalia.Core;

public partial class IAMath {
    public static int AmountDigits(ulong num) {
        int count = 0;

        while (num > 0) {
            num = num / 10;
            count++;
        }

        return count;
    }

    public static double ConvertMantissa(ulong mantissa) {
        double value = 0.0;

        if ((mantissa & 1) == 1) {
            value += 1.0;
        }

        ulong restOfMantissa = mantissa & 9223372036854775807;

        value += (restOfMantissa / (ulong)IntPow(10, AmountDigits(restOfMantissa)));

        return value;
    }

    public static double LoadFloat82bit(ulong sign, ulong exponent, ulong mantissa) {
        if (exponent == 0) {
            return Math.Pow(-1, sign) * Math.Pow(2, -16382) * ConvertMantissa(mantissa);
        } else {
            return Math.Pow(-1, sign) * Math.Pow(2.0, (exponent - 65535)) * ConvertMantissa(mantissa);
        }
    }
}
