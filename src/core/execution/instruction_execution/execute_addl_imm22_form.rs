use crate::{decoding::{InstructionAttributeMap, InstructionAttribute}, execution::{ItaniumProcessor, processor::ProcessorFault}};

pub fn execute_addl_imm22_form(processor: &mut ItaniumProcessor, attributes: InstructionAttributeMap) -> Result<(), ProcessorFault> {
    let reg1 = attributes[&InstructionAttribute::R1];
    let reg3 = attributes[&InstructionAttribute::R3];
    let immd = attributes[&InstructionAttribute::Immediate];
    let __qp = attributes[&InstructionAttribute::QualifyingPredicate];

    if *processor.retrieve_predicate_register(__qp).unwrap() {
        let mut r1 = *processor.retrieve_general_register(reg1).unwrap();
        let r3 = *processor.retrieve_general_register(reg3).unwrap();

        let val_err = r1.write(immd + r3.read());
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