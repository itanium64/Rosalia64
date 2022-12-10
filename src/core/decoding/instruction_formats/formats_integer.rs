use crate::{core::ia_math};

pub struct A5 {
    pub immediate: u64,
    pub r3: u64,
    pub r1: u64,
    pub qp: u64,
}

impl A5 {
    pub fn from_slots(slot: u64, _next_slot: u64) -> A5 {
        let sign1 = (slot & (0b00001000000000000000000000000000000000000)) >> 36;
        let imm9d = (slot & (0b00000111111111000000000000000000000000000)) >> 27;
        let imm5c = (slot & (0b00000000000000111110000000000000000000000)) >> 22;
        let r3___ = (slot & (0b00000000000000000001100000000000000000000)) >> 20;
        let imm7b = (slot & (0b00000000000000000000011111110000000000000)) >> 13;
        let r1___ = (slot & (0b00000000000000000000000000001111111000000)) >> 06;
        let qp___ = (slot & (0b00000000000000000000000000000000000111111)) >> 00;

        let immediate = ia_math::imm22(sign1, imm5c, imm9d, imm7b);

        A5 {
            immediate: immediate as u64,
            qp: qp___,
            r1: r1___,
            r3: r3___,
        }
    }
}
pub struct A4 {
    pub sign: u64,
    pub x2a: u64,
    pub ve: u64,
    pub immediate: u64,
    pub r3: u64,
    pub r1: u64,
    pub qp: u64,
}

impl A4 {
    pub fn from_slots(slot: u64, _next_slot: u64) -> A4 {
        let sign_ = (slot & (0b00001000000000000000000000000000000000000)) >> 47;
	    let x2a__ = (slot & (0b00000110000000000000000000000000000000000)) >> 34;
	    let ve___ = (slot & (0b00000001000000000000000000000000000000000)) >> 33;
	    let imm6d = (slot & (0b00000000111111000000000000000000000000000)) >> 27;
	    let r3___ = (slot & (0b00000000000000111111100000000000000000000)) >> 20;
	    let imm7b = (slot & (0b00000000000000000000011111110000000000000)) >> 13;
	    let r1___ = (slot & (0b00000000000000000000000000001111111000000)) >> 06;
	    let qp___ = (slot & (0b00000000000000000000000000000000000111111)) >> 00;

        let immediate = ia_math::imm14(sign_, imm6d, imm7b);

        A4 { 
            sign: sign_, 
            x2a: x2a__, 
            ve: ve___, 
            immediate: immediate as u64, 
            r3: r3___, 
            r1: r1___, 
            qp: qp___ 
        }
    }
}