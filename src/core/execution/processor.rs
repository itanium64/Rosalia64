use super::register_stack_engine::RegisterStackEngine;

pub enum ProcessorFault {
    IllegalOperation
}

#[derive(Clone, Copy)]
pub struct GeneralRegister {
    register_id: u64,
    value: u64,
    not_a_thing: bool,
}

impl GeneralRegister {
    pub fn write(&mut self, value: u64) -> Result<(), ProcessorFault> {
        if self.register_id == 0 {
            return Err(ProcessorFault::IllegalOperation);
        }

        self.value = value;

        return Ok(());
    }

    pub fn write_nat(&mut self, not_a_thing: bool) -> Result<(), ProcessorFault> {
        if self.register_id == 0 {
            return Err(ProcessorFault::IllegalOperation);
        }

        self.not_a_thing = not_a_thing;

        return Ok(());
    }

    pub fn read(&self) -> u64 {
        return self.value;
    }

    pub fn read_nat(&self) -> bool {
        return self.not_a_thing;
    }
}

#[derive()]
pub struct ItaniumProcessor {
    general_registers: [GeneralRegister; 128],
    floating_registers: [f64; 128],
    predicate_registers: [bool; 64],
    branch_registers: [u64; 8],
    register_stack_engine: RegisterStackEngine,
}

impl ItaniumProcessor {
    pub fn retrieve_general_register(&mut self, gr: u64) -> Option<&mut GeneralRegister> {
        return self.general_registers.get_mut(gr as usize);
    }

    pub fn retrieve_predicate_register(&mut self, pr: u64) -> Option<&mut bool> {
        return self.predicate_registers.get_mut(pr as usize);
    }

    pub fn retrieve_floating_register(&mut self, fr: u64) -> Option<&mut f64> {
        return self.floating_registers.get_mut(fr as usize);
    }

    pub fn retrieve_branch_register(&mut self, fr: u64) -> Option<&mut u64> {
        return self.branch_registers.get_mut(fr as usize);
    }
}