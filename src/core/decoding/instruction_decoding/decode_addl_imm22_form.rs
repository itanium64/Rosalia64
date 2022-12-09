use crate::{decoding::{DecodingContext, instruction_formats::{A5}}, execution::{self, ExecutableInstruction}};

use super::disassembly_helpers::format_qualifying_predicate;

impl DecodingContext<'_> {
    // Tags for easier searching:
    // Add Immediate 22 imm22_form addl Imm
    pub fn decode_addl_imm22_form(&mut self, slot: u64, next_slot: u64) {
        let a5 = A5::from_slots(slot, next_slot);

        let disassembly = format!("{} addl r{} = {}, r{}", format_qualifying_predicate(a5.qp), a5.r1, a5.immediate, a5.r3);

        //self.executable_instructions.push(executable_instruction);

        println!("{}", a5.immediate)
    }
}