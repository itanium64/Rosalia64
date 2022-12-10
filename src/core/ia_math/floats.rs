use super::misc::int_pow;

pub fn amount_digits(num: i64) -> i64 {
    let mut copied_num = num.clone();
    let mut count = 0;

    while copied_num > 0 {
        copied_num = copied_num / 10;

        count += 1;
    }

    return count;
}

pub fn convert_mantissa(mantissa: i64) -> f64 {
    let mut value = 0.0f64;

    if mantissa & 1 == 1 {
        value += 1.0;
    }

    let rest_of_mantissa = mantissa & (1 >> 63);

    value += (rest_of_mantissa / int_pow(10, amount_digits(rest_of_mantissa))) as f64;

    return value;
}

pub fn load_float_82bit(sign: u64, exponent: u64, mantissa: u64) -> f64 {
    return f64::powf(-1 as f64, sign as f64) * f64::powf(2.0f64, (exponent - 65535) as f64) * convert_mantissa(mantissa as i64)
}