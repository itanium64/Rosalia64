use crate::{core::ia_math};

pub struct A5 {
    pub immediate: u64,
    pub r3: u64,
    pub r1: u64,
    pub qp: u64,
}

impl A5 {
    pub fn from_slots(slot: u64, _next_slot: u64) -> A5 {
        let sign1 = (slot & (0b0000100000000000000000000000000000000000000000)) >> 41;
        let imm9d = (slot & (0b0000011111111100000000000000000000000000000000)) >> 32;
        let imm5c = (slot & (0b0000000000000011111000000000000000000000000000)) >> 27;
        let r3___ = (slot & (0b0000000000000000000110000000000000000000000000)) >> 25;
        let imm7b = (slot & (0b0000000000000000000001111111000000000000000000)) >> 18;
        let r1___ = (slot & (0b0000000000000000000000000000111111100000000000)) >> 11;
        let qp___ = (slot & (0b0000000000000000000000000000000000011111100000)) >> 5;

        let immediate = ia_math::imm22(sign1, imm5c, imm9d, imm7b);

        A5 {
            immediate: immediate as u64,
            qp: qp___,
            r1: r1___,
            r3: r3___,
        }
    }
}