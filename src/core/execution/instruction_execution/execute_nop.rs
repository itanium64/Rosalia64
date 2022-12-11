use crate::{execution::{ItaniumMachine, processor::ProcessorFault}, decoding::InstructionAttributeMap};

pub fn execute_nop(_machine: &mut ItaniumMachine, _attributes: &InstructionAttributeMap) -> Result<(), ProcessorFault> {
    return Ok(())
}