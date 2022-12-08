use std::collections::HashMap;

use crate::core::execution;

use super::{InstructionBundle, SLOT_ORDERS};

pub struct DecodingContext {
    current_address: u64,
    text_section: Vec<u8>,
    text_section_index: usize,
    executable_instructions: Vec<execution::ExecutableInstruction>,
    instruction_index: u64,
    address_to_instruction_index: HashMap<u64, u64>,
    instruction_index_to_address: HashMap<u64, u64>,
}

impl DecodingContext {
    pub fn new(text_section: &Vec<u8>, address_base: u64) -> DecodingContext {
        DecodingContext { 
            current_address: address_base,
            text_section: text_section.clone(),
            text_section_index: 0, 
            executable_instructions: Vec::new(), 
            instruction_index: 0, 
            address_to_instruction_index: HashMap::new(), 
            instruction_index_to_address: HashMap::new()
        }
    }

    pub fn next_bundle(&mut self) {
        let byte_slice = &self.text_section[self.text_section_index..self.text_section_index + 16];
        let byte_bundle = 
            u128::from_le_bytes(
                byte_slice
                    .try_into()
                    .expect("failed to read bundle!")
            );

        if byte_bundle == 0 {
            //I'm like 99% sure the whole bundle being empty means there's nothing there
            return
        }
        
        let instruction_bundle = InstructionBundle::decode(byte_bundle);
        let bundle_pipeline = SLOT_ORDERS[&instruction_bundle.template].clone();

        //We don't need to care about mid-bundle addresses
        //cuz this is used to know where to go to on branch instructions
        //and you cant branch to go inbetween 2 instructions in a bundle
        //you can only go to the beginning of one
        self.address_to_instruction_index.insert(self.current_address, self.instruction_index);
        self.instruction_index_to_address.insert(self.instruction_index, self.current_address);
        
        self.current_address += 16;  //Bundle is 16 bytes
        self.instruction_index += 3; //Bundle has 3 instructions
    }
}