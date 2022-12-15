use std::collections::HashMap;

use crate::{decoding::{DecodingContext, InstructionAttribute}, execution::{instruction_execution, self}};

use super::disassembly_helpers::format_qualifying_predicate;

fn decode_nop(context: &mut DecodingContext, disassembly: String, imm20a: u64, qp: u64, i: u64) {
    let marker_disassembly: String;

    let immediate = i << 20 | imm20a;
    
    match immediate {
        0 => marker_disassembly = String::from(" "),
        _ => marker_disassembly = format!(" // marker: {}", immediate)
    }

    let attributes: HashMap<InstructionAttribute, u64> = HashMap::from([
        (InstructionAttribute::Immediate, immediate),
        (InstructionAttribute::QualifyingPredicate, qp),
    ]); 

    let executable_instruction = execution::ExecutableInstruction {
        execution_function: instruction_execution::execute_nop,
        attributes: attributes,
        disassembly: format!("{}{}", disassembly, marker_disassembly)
    };

    context.executable_instructions.push(executable_instruction);
}

fn decode_nop_markers(slot: u64) -> (u64, u64, u64) {
    let _____i = (slot & (0b00001000000000000000000000000000000000000)) >> 36;
    let imm20a = (slot & (0b00000000000000011111111111111111111000000)) >> 6;
    let ____qp = (slot & (0b00000000000000000000000000000000000111111)) >> 0;

    return (_____i, imm20a, ____qp);
}

pub fn decode_nop_integer(context: &mut DecodingContext, slot: u64, _next_slot: u64) {
    let read = decode_nop_markers(slot);

    let disassembly = format!("{} nop.i", format_qualifying_predicate(read.2));
    
    decode_nop(context, disassembly, read.1, read.2, read.0);
}

pub fn decode_nop_branch(context: &mut DecodingContext, slot: u64, _next_slot: u64) {
    let read = decode_nop_markers(slot);

    let disassembly = format!("{} nop.b", format_qualifying_predicate(read.2));
    
    decode_nop(context, disassembly, read.1, read.2, read.0);
}

pub fn decode_nop_memory(context: &mut DecodingContext, slot: u64, _next_slot: u64) {
    let read = decode_nop_markers(slot);

    let disassembly = format!("{} nop.m", format_qualifying_predicate(read.2));
    
    decode_nop(context, disassembly, read.1, read.2, read.0);
}

pub fn decode_nop_float(context: &mut DecodingContext, slot: u64, _next_slot: u64) {
    let read = decode_nop_markers(slot);

    let disassembly = format!("{} nop.f", format_qualifying_predicate(read.2));
    
    decode_nop(context, disassembly, read.1, read.2, read.0);
}
