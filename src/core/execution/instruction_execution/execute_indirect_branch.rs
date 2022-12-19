use std::collections::HashMap;

use crate::{execution::{processor::ProcessorFault, ItaniumMachine}, decoding::{InstructionAttributeMap, InstructionAttribute}};

pub fn execute_indirect_branch(
    /* Contains Processor and Memory */ machine: &mut ItaniumMachine, 
    /* Contains Function Parameters  */ attributes: &InstructionAttributeMap,
    /*       Instruction Pointer     */ instruction_index: &mut usize,
    /*    Converts IP to an address  */ instruction_index_to_address: &HashMap<u64, u64>,
    /*    Converts address to an IP  */ address_to_instruction_index: &HashMap<u64, u64>,
) -> Result<(), ProcessorFault> {
    let processor = &mut machine.processor;

    let b2 = attributes[&InstructionAttribute::B2];
    let qp = attributes[&InstructionAttribute::QualifyingPredicate];

    if processor.retrieve_predicate_register(qp).unwrap().read() {
        let jump_address = processor.retrieve_branch_register(b2).unwrap();

        if jump_address == &0 {
            machine.continue_running = false;
            return Ok(())
        }
    }
    
    return Ok(())
}