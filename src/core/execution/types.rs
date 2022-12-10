use crate::core::decoding;

use super::{processor::{ProcessorFault}, machine::ItaniumMachine};

pub struct ExecutableInstruction {
    pub execution_function: fn(&mut ItaniumMachine, &decoding::InstructionAttributeMap) -> Result<(), ProcessorFault>,
    pub attributes: decoding::InstructionAttributeMap,
    pub disassembly: String
}
