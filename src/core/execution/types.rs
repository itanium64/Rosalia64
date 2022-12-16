use std::fmt::Display;

use crate::core::decoding;

use super::{processor::{ProcessorFault}, machine::ItaniumMachine};

pub struct ExecutableInstruction {
    pub execution_function: fn(&mut ItaniumMachine, &decoding::InstructionAttributeMap) -> Result<(), ProcessorFault>,
    pub attributes: decoding::InstructionAttributeMap,
    pub disassembly: String
}

impl Display for ExecutableInstruction {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        write!(f, "{}", self.disassembly)
    }
}