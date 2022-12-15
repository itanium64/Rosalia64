use crate::{execution::{ItaniumMachine, processor::ProcessorFault}, decoding::{InstructionAttributeMap, InstructionAttribute}};

pub fn execute_int_store_no_base_update_form(machine: &mut ItaniumMachine, attributes: &InstructionAttributeMap) -> Result<(), ProcessorFault> {
    let processor = &mut machine.processor;

    let reg2 = attributes[&InstructionAttribute::R2];
    let reg3 = attributes[&InstructionAttribute::R3];
    let tabx = attributes[&InstructionAttribute::TableX];
    let __qp = attributes[&InstructionAttribute::QualifyingPredicate];

    let bit_length_table: [u64; 4] = [1, 2, 4, 8];

    if processor.retrieve_predicate_register(__qp).unwrap().read() {
        let count_bytes = bit_length_table[tabx as usize];

        let r2 = processor.retrieve_general_register(reg2).unwrap();
        let r3 = processor.retrieve_general_register(reg3).unwrap();

        if r2.not_a_thing | r3.not_a_thing {
            return Err(ProcessorFault::NATConsumptionFault);
        }

        let address = r3.read();

        let bytes = address.to_le_bytes();

        let mut execution_result:Result<(), ProcessorFault> = Ok(());

        if address + count_bytes >= machine.memory_size as u64 {
            println!("execute_int_load_no_base_update_form: write caused overflow!");

            execution_result = Err(ProcessorFault::SoftFault);
        }

        for x in address..address + count_bytes {
            let index = (x - address) % machine.memory_size as u64;

            machine.memory[x as usize] = bytes[index as usize]
        }

        return execution_result;
    }
    
    return Ok(())
}