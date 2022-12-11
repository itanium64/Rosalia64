use phf::phf_map;

use super::{DecodingContext, instruction_decoding, UnitOrStop};

pub static INSTRUCTION_TABLE_INTEGER: phf::Map<u32, fn(&mut DecodingContext, slot: u64, next_slot: u64)> = phf_map! {
    8u32 => instruction_decoding::decode_part_integer_alu,
    9u32 => instruction_decoding::decode_addl_imm22_form
};

pub static INSTRUCTION_TABLE_MEMORY: phf::Map<u32, fn(&mut DecodingContext, slot: u64, next_slot: u64)> = phf_map! {
    4u32 => instruction_decoding::decode_part_int_load_store,
    9u32 => instruction_decoding::decode_addl_imm22_form,
};

pub static INSTRUCTION_TABLE_FLOAT: phf::Map<u32, fn(&mut DecodingContext, slot: u64, next_slot: u64)> = phf_map! {
    
};

pub static INSTRUCTION_TABLE_BRANCH: phf::Map<u32, fn(&mut DecodingContext, slot: u64, next_slot: u64)> = phf_map! {
    2u32 => instruction_decoding::decode_nop_branch,
};

pub static INSTRUCTION_TABLE_EXTENDED: phf::Map<u32, fn(&mut DecodingContext, slot: u64, next_slot: u64)> = phf_map! {
    
};

pub fn get_unit_instruction_table(unit: &UnitOrStop) -> &'static phf::Map<u32, fn(&mut DecodingContext, slot: u64, next_slot: u64)> {
    match unit {
        UnitOrStop::Integer => &INSTRUCTION_TABLE_INTEGER,
        UnitOrStop::Memory => &INSTRUCTION_TABLE_MEMORY,
        UnitOrStop::Float => &INSTRUCTION_TABLE_FLOAT,
        UnitOrStop::Branch => &INSTRUCTION_TABLE_BRANCH,
        UnitOrStop::Extended => &INSTRUCTION_TABLE_EXTENDED,
        _ => panic!("Invalid unit!")
    }
}