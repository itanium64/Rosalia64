use crate::core::decoding;

pub struct ExecutableInstruction {
    pub execution_function: fn(decoding::InstructionAttributeMap),
    pub attributes: decoding::InstructionAttributeMap,
    pub disassembly: String
}
