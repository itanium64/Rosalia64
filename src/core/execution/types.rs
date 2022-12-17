use std::{fmt::Display, collections::HashMap};

use crate::core::decoding;

use super::{processor::{ProcessorFault}, ItaniumMachine};

pub struct ExecutableInstruction {
    pub execution_function: fn(&mut ItaniumMachine, &decoding::InstructionAttributeMap, &mut usize, &HashMap<u64, u64>, &HashMap<u64, u64>) -> Result<(), ProcessorFault>,
    pub attributes: decoding::InstructionAttributeMap,
    pub disassembly: String
}

impl Display for ExecutableInstruction {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        write!(f, "{}", self.disassembly)
    }
}