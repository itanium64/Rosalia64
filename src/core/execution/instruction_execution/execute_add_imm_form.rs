use std::collections::HashMap;

use crate::{decoding::{InstructionAttributeMap, InstructionAttribute}, execution::{processor::ProcessorFault, ItaniumMachine}};

pub fn execute_add_imm_form(
    /* Contains Processor and Memory */ machine: &mut ItaniumMachine, 
    /* Contains Function Parameters  */ attributes: &InstructionAttributeMap,
    /*       Instruction Pointer     */ _instruction_index: &mut usize,
    /*    Converts IP to an address  */ _instruction_index_to_address: &HashMap<u64, u64>,
    /*    Converts address to an IP  */ _address_to_instruction_index: &HashMap<u64, u64>,
) -> Result<(), ProcessorFault> {
    let processor = &mut machine.processor;

    let reg1 = attributes[&InstructionAttribute::R1];
    let reg3 = attributes[&InstructionAttribute::R3];
    let immd = attributes[&InstructionAttribute::Immediate];
    let __qp = attributes[&InstructionAttribute::QualifyingPredicate];

    if processor.retrieve_predicate_register(__qp).unwrap().read() {
        let mut r1 = *processor.retrieve_general_register_mut(reg1).unwrap();
        let r3 = *processor.retrieve_general_register(reg3).unwrap();

        let r3val = r3.read() as i64;
        let imval = immd as i64;

        let val_err = r1.write( (imval + r3val) as u64);
        let nat_err = r1.write_nat(r3.read_nat());

        if val_err.is_err() {
            return val_err;
        }

        if nat_err.is_err() {
            return nat_err;
        }
    }

    return Ok(())
}