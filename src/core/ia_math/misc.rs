pub fn int_pow(x: i64, n: i64) -> i64 {
    if n == 0 {
        return 1
    }

    if n == 1 {
        return x
    }

    let y = int_pow(x, n / 2);

    if n % 2 == 0 {
        return y * y;
    }

    return x * y * y;
}

pub fn zero_ext(val: i64, pos: i64) -> i64 {
    let and = int_pow(2, pos) - 1;

    return val & and;
}

