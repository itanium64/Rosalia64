use std::collections::HashMap;

use crate::{decoding::{DecodingContext, instruction_formats::M1_2_4, InstructionAttribute}, execution::{self, instruction_execution}};

use super::disassembly_helpers::format_qualifying_predicate;

//Tags for easier searching:
// Integer Load Store Semaphore Get FR 1-bit 1bit Opcode Extensions
// 4-27
pub fn decode_part_int_load_store(context: &mut DecodingContext, slot: u64, next_slot: u64) {
    let m = (slot & (0b00001000000000000000000000000000000000000)) >> 36;
    let x = (slot & (0b00000000000001000000000000000000000000000)) >> 27;

    let combined = m << 1 | x;

    match combined {
        0 => {
            decode_part_int_load_store_extensions(context, slot, next_slot);
        },
        _ => {
            println!("decode_part_int_load_store: Unimplemented m and x warning!\nm: {}\nx: {}\ncombined: {}", m, x, combined)
        }
    }
}

//Tags for easier searching:
// Integer Load Store Opcode Extensions
// 4-29
pub fn decode_part_int_load_store_extensions(context: &mut DecodingContext, slot: u64, next_slot: u64) {
    let table_y = (slot & (0b00000111100000000000000000000000000000000)) >> 32;

    if table_y >= 12 {
        decode_integer_store(context, slot, next_slot);
    } else {
        decode_integer_load(context, slot, next_slot);
    }
}

//no_base_update_form
pub fn decode_integer_load(context: &mut DecodingContext, slot: u64, next_slot: u64) {
    let m = M1_2_4::from_slots(slot, next_slot);

    let m = M1_2_4::from_slots(slot, next_slot);

    let bit_length_table: [u64; 4] = [1, 2, 4, 8];

    let hint_table = [
        "",
        ".s",
        ".a",
        ".sa",
        ".bias",
        ".acq",
        ".fill",
        "",
        ".c.clr",
        ".c.nc",
        ".c.clr.acq",
    ];

    let locality_hint_table = [
        "",
        ".nt1",
        ".nt2",
        ".nta"
    ];

    let disassembly = format!("{} ld{}{}{} r{} = [r{}]", format_qualifying_predicate(m.qp), bit_length_table[m.tab_x as usize], hint_table[m.tab_y as usize], locality_hint_table[m.hint as usize], m.r1, m.r3);

    let attributes: HashMap<InstructionAttribute, u64> = HashMap::from([
        (InstructionAttribute::R1, m.r1),
        (InstructionAttribute::R3, m.r3),
        (InstructionAttribute::Hint, m.hint),
        (InstructionAttribute::TableX, m.tab_x),
        (InstructionAttribute::TableY, m.tab_y),
        (InstructionAttribute::QualifyingPredicate, m.qp),
    ]); 

    let executable_instruction = execution::ExecutableInstruction {
        execution_function: instruction_execution::execute_int_load_no_base_update_form,
        attributes: attributes,
        disassembly: disassembly
    };

    context.executable_instructions.push(executable_instruction);
}


pub fn decode_integer_store(context: &mut DecodingContext, slot: u64, next_slot: u64) {

}