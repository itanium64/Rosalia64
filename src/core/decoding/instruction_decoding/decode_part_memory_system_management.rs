use crate::decoding::DecodingContext;

use super::decode_nop_memory;

//Tags:
//Opcode 0 System Memory 3-bit 3bit Opcode Extensions
pub fn decode_part_memory_system_management(context: &mut DecodingContext, slot: u64, next_slot: u64) {
    let x3 = (slot & (0b00000111000000000000000000000000000000000)) >> 33;

    match x3 {
        0 => {
            decode_part_memory_system_management_4bit_2bit(context, slot, next_slot);
        }
        _ => {
            println!("decode_part_memory_system_management: unimplemented x3 warning\nx3: {}", x3)
        }
    }
}

//Tags
//Opcode 0 System Memory 4+2 4-bit+2-bit 4bit+2bit extensions 42
pub fn decode_part_memory_system_management_4bit_2bit(context: &mut DecodingContext, slot: u64, next_slot: u64) {
    let tabx = (slot & (0b00000000110000000000000000000000000000000)) >> 31;
    let taby = (slot & (0b00000000001111000000000000000000000000000)) >> 27;

    match taby {
        1 => {
            match tabx {
                0 => {
                    decode_nop_memory(context, slot, next_slot);
                }
                _ => {
                    println!("decode_part_memory_system_management_4bit_2bit: unimplemented tabx & taby combination\ntabx: {}\ntaby: {}", tabx, taby)
                }
            }
        }
        _ => {
            println!("decode_part_memory_system_management_4bit_2bit: unimplemented tabx & taby combination\ntabx: {}\ntaby: {}", tabx, taby)
        }
    }
}