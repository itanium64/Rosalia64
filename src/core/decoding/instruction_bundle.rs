pub struct InstructionBundle {
    pub template: u32,
    pub slot0: u64,
    pub slot1: u64,
    pub slot2: u64,
}

impl InstructionBundle {
    pub fn decode(input: u128) -> InstructionBundle {
        let mut instruction_bundle: InstructionBundle = InstructionBundle { template: 0, slot0: 0, slot1: 0, slot2: 0 };

        // Yes, I did just do this
        instruction_bundle.template = u32::try_from(input & 0b00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000011111 >> 0).expect("Failed to convert template to a u32");
        instruction_bundle.slot0    = u64::try_from(input & 0b00000000000000000000000000000000000000000000000000000000000000000000000000000000001111111111111111111111111111111111111111100000 >> 5).expect("Failed to convert slot 0 to a u64");
        instruction_bundle.slot1    = u64::try_from(input & 0b00000000000000000000000000000000000000000111111111111111111111111111111111111111110000000000000000000000000000000000000000000000 >> 46).expect("Failed to convert slot 1 to a u64");
        instruction_bundle.slot2    = u64::try_from(input & 0b11111111111111111111111111111111111111111000000000000000000000000000000000000000000000000000000000000000000000000000000000000000 >> 87).expect("Failed to convert slot 2 to a u64");

        instruction_bundle
    }
}