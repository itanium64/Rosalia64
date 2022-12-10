use crate::{execution::{ItaniumMachine, processor::ProcessorFault}, decoding::InstructionAttributeMap};

pub fn execute_adds_imm14_form(machine: &mut ItaniumMachine, attributes: &InstructionAttributeMap) -> Result<(), ProcessorFault> {
    return Ok(());
}