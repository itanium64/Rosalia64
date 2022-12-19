use crate::decoding::DecodingContext;

use super::decode_branch_indirect_return::decode_branch_indirect_return;


pub fn decode_part_branch_indirect_misc(context: &mut DecodingContext, slot: u64, next_slot: u64) {
    let tabx = (slot & (0b00000000110000000000000000000000000000000)) >> 31;
    let taby = (slot & (0b00000000001111000000000000000000000000000)) >> 27;

    match tabx {
        2 => {
            match taby {
                1 => {
                    decode_part_branch_indirect_return(context, slot, next_slot)
                }
                _ => {
                    println!("decode_part_branch_indirect_misc: unimplemented tabx and y combination!\ntabx: {}\ntaby: {}", tabx, taby);
                }
            }
        },
        _ => {
            println!("decode_part_branch_indirect_misc: unimplemented tabx and y combination!\ntabx: {}\ntaby: {}", tabx, taby);
        }
    }
}

pub fn decode_part_branch_indirect_return(context: &mut DecodingContext, slot: u64, next_slot: u64) {
    let btype = (slot & (0b00000000000000000000000000000000111000000)) >> 6;

    if btype != 4 {
        println!("decode_part_branch_indirect_return: invalid branch type warning!\nbtype: {}", btype)
    }

    decode_branch_indirect_return(context, slot, next_slot);
}