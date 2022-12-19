use super::processor::ProcessorFault;

#[derive(Clone, Copy, Debug)]
pub struct GeneralRegister {
    pub(crate) register_id: u64,
    pub(crate) value: u64,
    pub(crate) not_a_thing: bool,
}

impl GeneralRegister {
    pub fn new() -> GeneralRegister {
        GeneralRegister { register_id: 0, value: 0, not_a_thing: false }
    }

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

#[derive(Clone, Copy, Debug)]
pub struct PredicateRegister {
    pub(crate) register_id: u64,
    pub(crate) value: bool
}

impl PredicateRegister {
    pub fn new() -> PredicateRegister {
        PredicateRegister { register_id: 0, value: false }
    }

    pub fn write(&mut self, value: bool) -> Result<(), ProcessorFault> {
        if self.register_id == 0 {
            return Err(ProcessorFault::IllegalOperation);
        }

        self.value = value;

        return Ok(());
    }

    pub fn read(&self) -> bool {
        return self.value;
    }
}

#[derive(Clone, Copy, Debug)]
pub struct FloatingRegister {
    pub(crate) register_id: u64,
    pub(crate) value: f64
}

impl FloatingRegister {
    pub fn new() -> FloatingRegister {
        FloatingRegister { register_id: 0, value: 0.0f64 }
    }

    pub fn write(&mut self, value: f64) -> Result<(), ProcessorFault> {
        if self.register_id == 0 || self.register_id == 1 {
            return Err(ProcessorFault::IllegalOperation);
        }

        self.value = value;

        return Ok(());
    }

    pub fn read(&self) -> f64 {
        return self.value;
    }
}
