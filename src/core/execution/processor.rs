use std::fmt::Display;

use super::{register_stack_engine::RegisterStackEngine, GeneralRegister, register_types::{PredicateRegister, FloatingRegister}};

#[derive(Debug)]
pub enum ProcessorFault {
    IllegalOperation,
    NATConsumptionFault,
    
}

impl Display for ProcessorFault {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        write!(f, "{:?}", self)
    }
}

pub struct ItaniumProcessor {
    general_registers: [GeneralRegister; 128],
    floating_registers: [FloatingRegister; 128],
    predicate_registers: [PredicateRegister; 64],
    branch_registers: [u64; 8],
    register_stack_engine: RegisterStackEngine,
}

impl ItaniumProcessor {
    pub fn new() -> ItaniumProcessor {
        let mut general_registers: [GeneralRegister; 128] = [GeneralRegister::new(); 128];
        let mut predicate_registers: [PredicateRegister; 64] = [PredicateRegister::new(); 64];
        let mut floating_registers: [FloatingRegister; 128] = [FloatingRegister::new(); 128];

        for i in 0..127 {
            general_registers[i] = GeneralRegister {
                register_id: i as u64,
                value: 0,
                not_a_thing: false
            }
        } 

        for i in 0..63 {
            predicate_registers[i] = PredicateRegister {
                register_id: i as u64,
                value: false,
            }
        } 

        //Has to be set manually to true.
        predicate_registers[0] = PredicateRegister {
            register_id: 0,
            value: true
        };

        for i in 0..127 {
            floating_registers[i] = FloatingRegister {
                register_id: i as u64,
                value: 0.0f64,
            }
        } 

        floating_registers[1] = FloatingRegister {
            register_id: 1,
            value: 1.0f64,
        };

        ItaniumProcessor { 
            general_registers: general_registers, 
            floating_registers: floating_registers, 
            predicate_registers: predicate_registers, 
            branch_registers: [0, 0, 0, 0, 0, 0, 0, 0], 
            register_stack_engine: RegisterStackEngine::new() 
        }
    }


    pub fn retrieve_general_register_mut(&mut self, gr: u64) -> Option<&mut GeneralRegister> {
        if gr < 32 {
            return self.general_registers.get_mut(gr as usize);
        }

        let register = (gr - 32) + self.register_stack_engine.current_frame().register_base;

        return self.general_registers.get_mut(register as usize);
    }

    pub fn retrieve_general_register(&self, gr: u64) -> Option<&GeneralRegister> {
        if gr < 32 {
            return self.general_registers.get(gr as usize);
        }

        let register = (gr - 32) + self.register_stack_engine.current_frame().register_base;

        return self.general_registers.get(register as usize);
    }

    pub fn retrieve_predicate_register_mut(&mut self, pr: u64) -> Option<&mut PredicateRegister> {
        return self.predicate_registers.get_mut(pr as usize);
    }

    pub fn retrieve_predicate_register(&self, pr: u64) -> Option<&PredicateRegister> {
        return self.predicate_registers.get(pr as usize);
    }

    pub fn retrieve_floating_register_mut(&mut self, fr: u64) -> Option<&mut FloatingRegister> {
        return self.floating_registers.get_mut(fr as usize);
    }

    pub fn retrieve_floating_register(&self, fr: u64) -> Option<&FloatingRegister> {
        return self.floating_registers.get(fr as usize);
    }

    pub fn retrieve_branch_register_mut(&mut self, fr: u64) -> Option<&mut u64> {
        return self.branch_registers.get_mut(fr as usize);
    }

    pub fn retrieve_branch_register(&self, fr: u64) -> Option<&u64> {
        return self.branch_registers.get(fr as usize);
    }
}