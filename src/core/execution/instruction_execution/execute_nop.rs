use crate::{execution::{ItaniumMachine, processor::ProcessorFault}, decoding::InstructionAttributeMap};

pub fn execute_nop(machine: &mut ItaniumMachine, attributes: &InstructionAttributeMap) -> Result<(), ProcessorFault> {
    return Ok(())
}