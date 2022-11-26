package ia_math

import "math"

func AmountDigits(num int64) int64 {
	count := int64(0)

	for num > 0 {
		num = num / 10
		count++
	}

	return count
}

func ConvertMantissa(mantissa int64) float64 {
	value := 0.0

	if mantissa&1 == 1 {
		value += 1
	}

	restOfMantissa := mantissa & (1 >> 63)

	value += float64(restOfMantissa / IntPow(10, AmountDigits(restOfMantissa)))

	return value
}

func LoadFloat82bit(sign uint64, exponent uint64, mantissa uint64) float64 {
	return math.Pow(-1, float64(sign)) * math.Pow(2, float64(exponent-65535)) * ConvertMantissa(int64(mantissa))
}
