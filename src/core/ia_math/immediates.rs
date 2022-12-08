pub fn sign_ext(i: u64, n: u64) -> i64 {
    (((i) << (64 - (n))) >> (64 - (n))) as i64
}

pub fn imm22(sign: u64, imm5c: u64, imm9d: u64, imm7b: u64) -> i64 {
    sign_ext(sign << 21 | imm5c << 16 | imm9d << 7 | imm7b, 22)
}

pub fn imm14(sign: u64, imm6d: u64, imm7b: u64) -> i64 {
    sign_ext(sign << 13 | imm6d << 7 | imm7b, 14)
}