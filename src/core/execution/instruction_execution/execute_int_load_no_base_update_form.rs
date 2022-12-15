use crate::{execution::{ItaniumMachine, processor::ProcessorFault}, decoding::{InstructionAttributeMap, InstructionAttribute}};

pub fn execute_int_load_no_base_update_form(machine: &mut ItaniumMachine, attributes: &InstructionAttributeMap) -> Result<(), ProcessorFault> {
    let ___r1 = attributes[&InstructionAttribute::R1];
    let ___r3 = attributes[&InstructionAttribute::R3];
    let _hint = attributes[&InstructionAttribute::Hint];
    let tab_x = attributes[&InstructionAttribute::TableX];
    let tab_y = attributes[&InstructionAttribute::TableY];
    let ___qp = attributes[&InstructionAttribute::QualifyingPredicate];

    let mut speculative: bool = false;
    let mut advanced: bool = false;
    let mut bias: bool = false;
    let mut acquire: bool = false;
    let mut fill: bool = false;
    let mut check_clear: bool = false;
    let mut check_no_clear: bool = false;
    let mut no_clear: bool = false;

    let mut count_bytes;

    let r1 = machine.processor.retrieve_general_register(___r1).unwrap();
    let r3 = machine.processor.retrieve_general_register(___r3).unwrap();

    match tab_x {
        0 => {},
        1 => {
            speculative = true;
        },
        2 => {
            advanced = true;
        },
        3 => {
            speculative = true;
            advanced = true;
        },
        4 => {
            bias = true;
        },
        5 => {
            acquire = true;
        },
        6 => {
            count_bytes = 8;
            fill = true
        },
        8 => {
            check_clear = true;
        },
        9 => {
            check_no_clear = true
        },
        10 => {
            check_clear = true;
            acquire = true;
        }
        _ => {
            println!("execute_int_load_no_base_update_form: unimplemented tab_x value!\ntab_x: {}", tab_x);
        }
    }
    
    if !speculative && r3.not_a_thing {
        return Err(ProcessorFault::NATConsumptionFault);
    }

    let defer_exception = speculative && (r3.not_a_thing /* || PSR.ed (Processor Status Register; Exception Deferral) */);
    
    return Ok(())
}