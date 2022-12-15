use crate::{decoding::{InstructionAttributeMap, InstructionAttribute}, execution::{processor::ProcessorFault, machine::ItaniumMachine}};

pub fn execute_add_imm_form(machine: &mut ItaniumMachine, attributes: &InstructionAttributeMap) -> Result<(), ProcessorFault> {
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

    return Err(ProcessorFault::IllegalOperation);
    return Ok(())
}