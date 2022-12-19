use std::collections::HashMap;

use crate::{decoding::{DecodingContext, instruction_formats::{A5}, InstructionAttribute}, execution::{self, instruction_execution}};

use super::disassembly_helpers::format_qualifying_predicate;

// Tags for easier searching:
// Add Immediate 22 imm22_form addl Imm
pub fn decode_addl_imm22_form(context: &mut DecodingContext, slot: u64, next_slot: u64) {
    let a5 = A5::from_slots(slot, next_slot);

    let disassembly = format!("{} addl r{} = {}, r{}", format_qualifying_predicate(a5.qp), a5.r1, a5.immediate, a5.r3);

    let attributes: HashMap<InstructionAttribute, u64> = HashMap::from([
        (InstructionAttribute::R1, a5.r1),
        (InstructionAttribute::R3, a5.r3),
        (InstructionAttribute::QualifyingPredicate, a5.qp),
        (InstructionAttribute::Immediate, a5.immediate),
    ]);

    let executable_instruction = execution::ExecutableInstruction {
        execution_function: instruction_execution::execute_add_imm_form,
        attributes: attributes,
        disassembly: disassembly
    };

    context.executable_instructions.push(executable_instruction);
}
