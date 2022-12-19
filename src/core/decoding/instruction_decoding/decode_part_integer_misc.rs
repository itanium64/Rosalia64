use crate::decoding::DecodingContext;

use super::decode_nop_integer;

// Tags:
// Misc I-Unit 3-bit 3bit Opcode Extensions x3 Miscellaneous
pub fn decode_part_integer_misc(context: &mut DecodingContext, slot: u64, next_slot: u64) {
    let x3 = (slot & (0b00000111000000000000000000000000000000000)) >> 33;

    match x3 {
        0 => {
            decode_part_integer_misc_6bit(context, slot, next_slot);
        },
        _ => {
            println!("decode_part_integer_misc: Unimplemented x3 warning!\nx3: {}", x3)
        }
    }
}

// Tags:
// Misc I-Unit 6-bit 6bit Opcode Extensions
pub fn decode_part_integer_misc_6bit(context: &mut DecodingContext, slot: u64, next_slot: u64) {
    let taby = (slot & (0b00000000000011000000000000000000000000000)) >> 27;
    let tabx = (slot & (0b00000000111100000000000000000000000000000)) >> 29;

    match tabx {
        0 => {
            match taby {
                0 => {
                    //break.i
                }
                1 => {
                    //it can be hint.i too but those are the same format so i dont really mind
                    decode_nop_integer(context, slot, next_slot)
                },
                0xA => {
                    //mov.i to ar - imm8
                }
                _ => {
                    println!("decode_part_integer_misc_6bit: invalid tabx & taby combination!\ntabx: {}\ntaby: {}", tabx, taby);
                }
            }
        },  
        _ => {
            println!("decode_part_integer_misc_6bit: Unimplemented tabx warning!\ntabx: {}", tabx)
        }
    }
}