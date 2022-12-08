use crate::core::decoding;

pub struct ExecutableInstruction {
    execution_function: fn(decoding::InstructionAttributeMap),
    attributes: decoding::InstructionAttributeMap,
    disassembly: String
}
