use crate::{decoding::{InstructionAttributeMap, InstructionAttribute}, execution::{processor::ProcessorFault, ItaniumMachine}};

pub fn execute_instruction(
    /* Contains Processor and Memory */ machine: &mut ItaniumMachine, 
    /* Contains Function Parameters  */ attributes: &InstructionAttributeMap,
    /*       Instruction Pointer     */ _instruction_index: &mut usize,
    /*    Converts IP to an address  */ _instruction_index_to_address: &HashMap<u64, u64>,
    /*    Converts address to an IP  */ _address_to_instruction_index: &HashMap<u64, u64>,
) -> Result<(), ProcessorFault> {

}