use std::collections::HashMap;

use crate::{core::execution, decoding::{get_unit_instruction_table}};

use super::{InstructionBundle, SLOT_ORDERS, UnitOrStop};

pub struct DecodingContext<'a> {
    current_address: u64,
    text_section: &'a [u8],
    text_section_index: usize,
    text_section_size: usize,
    pub executable_instructions: Vec<execution::ExecutableInstruction>,
    instruction_index: u64,
    pub address_to_instruction_index: HashMap<u64, u64>,
    pub instruction_index_to_address: HashMap<u64, u64>,
}

impl DecodingContext<'_> {
    pub fn new<'byteref>(text_section: &'byteref [u8], text_section_size: usize, address_base: u64) -> DecodingContext {
        DecodingContext { 
            current_address: address_base,
            text_section: text_section.clone(),
            text_section_index: 0, 
            text_section_size: text_section_size,
            executable_instructions: Vec::new(), 
            instruction_index: 0, 
            address_to_instruction_index: HashMap::new(), 
            instruction_index_to_address: HashMap::new()
        }
    }

    pub fn decode_all(&mut self) {
        while self.instruction_index * 16 != self.text_section_size as u64 {
            self.next_bundle()
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
        self.text_section_index += 16;
        self.instruction_index += 3; //Bundle has 3 instructions

        let mut unit_slot0: Option<UnitOrStop> = None;
        let mut unit_slot1: Option<UnitOrStop> = None;
        let mut unit_slot2: Option<UnitOrStop> = None;

        let mut pipeline_index: usize = 0;

        while unit_slot0.is_none() || unit_slot1.is_none() || unit_slot2.is_none() {
            let current_item = bundle_pipeline[pipeline_index].clone();

            pipeline_index += 1;

            if current_item == UnitOrStop::None || current_item == UnitOrStop::Stop {
                continue;
            } else if current_item == UnitOrStop::End {
                break;
            }

            if unit_slot0.is_none() {
                unit_slot0 = Some(current_item);
                continue;
            }

            if unit_slot1.is_none() {
                unit_slot1 = Some(current_item);
                continue;
            }

            if unit_slot2.is_none() {
                unit_slot2 = Some(current_item);
                continue;
            }          
        }

        self.decode_instruction_slot(instruction_bundle.slot0, instruction_bundle.slot1, unit_slot0.unwrap());
        self.decode_instruction_slot(instruction_bundle.slot1, instruction_bundle.slot2, unit_slot1.unwrap());
        self.decode_instruction_slot(instruction_bundle.slot2, 0b0000000000000000000000, unit_slot2.unwrap());
    }

    fn decode_instruction_slot(&mut self, slot: u64, next_slot: u64, unit: UnitOrStop) {
        let mask = 0b1111 << 37;
        let major_opcode = ((slot & mask) >> 37) as u32;

        let unit_table = get_unit_instruction_table(&unit);
        let retrieved = unit_table.get(&major_opcode);

        match retrieved {
            Some(decoder) => decoder(self, slot, next_slot),
            None => {
                println!("decode_instruction_slot: Major Opcode {} unimplemented for {} unit", major_opcode, unit)
            }
        }
    }
}