use std::collections::HashMap;

use crate::{decoding::{DecodingContext, instruction_formats::A4, InstructionAttribute}, execution::{instruction_execution, self}};

use super::disassembly_helpers::format_qualifying_predicate;

pub fn decode_adds_imm14_form(context: &mut DecodingContext, slot: u64, next_slot: u64) {
    let a4 = A4::from_slots(slot, next_slot);

    let disassembly = format!("{} adds r{} = {}, r{}", format_qualifying_predicate(a4.qp), a4.r1, a4.immediate, a4.r3);
    
    let attributes: HashMap<InstructionAttribute, u64> = HashMap::from([
        (InstructionAttribute::R1, a4.r1),
        (InstructionAttribute::R3, a4.r3),
        (InstructionAttribute::QualifyingPredicate, a4.qp),
        (InstructionAttribute::Immediate, a4.immediate),
    ]); 

    let executable_instruction = execution::ExecutableInstruction {
        execution_function: instruction_execution::execute_adds_imm14_form,
        attributes: attributes,
        disassembly: disassembly
    };

    context.executable_instructions.push(executable_instruction);
}