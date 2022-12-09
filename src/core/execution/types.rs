use crate::core::decoding;

use super::processor::{ItaniumProcessor, ProcessorFault};

pub struct ExecutableInstruction {
    pub execution_function: fn(&mut ItaniumProcessor, decoding::InstructionAttributeMap) -> Result<(), ProcessorFault>,
    pub attributes: decoding::InstructionAttributeMap,
    pub disassembly: String
}
