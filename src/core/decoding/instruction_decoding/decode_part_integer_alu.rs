use crate::decoding::DecodingContext;

use super::decode_adds_imm14_form;

// Tags for easier searching
// integer alu 2+1 2bit+1bit 2-bit+1-bit opcode extensions 
pub fn decode_part_integer_alu(context: &mut DecodingContext, slot: u64, next_slot: u64) {
    let x2a = (slot & (0b00000100000000000000000000000000000000000)) >> 34;
	let _ve = (slot & (0b00000010000000000000000000000000000000000)) >> 33;

    match x2a {
        2 => {
            if _ve == 0 {
                decode_adds_imm14_form(context, slot, next_slot);
            } else {
                println!("decode_part_integer_alu: Unimplemented _ve warning!\nx2a: {}\nve: {}", x2a, _ve)
            }
        },
        _ => {
            println!("decode_part_integer_alu: Unimplemented x2a warning!\nx2a: {}\nve: {}", x2a, _ve)
        }
    }
}