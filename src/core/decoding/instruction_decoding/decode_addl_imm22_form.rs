use crate::decoding::{DecodingContext, instruction_formats::{self, A5}};

impl DecodingContext<'_> {
    // Tags for easier searching:
    // Add Immediate 22 imm22_form addl Imm
    pub fn decode_addl_imm22_form(&mut self, slot: u64, next_slot: u64) {
        let a5 = A5::from_slots(slot, next_slot);

        println!("{}", a5.immediate)
    }
}