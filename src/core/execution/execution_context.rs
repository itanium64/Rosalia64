use crate::{decoding::DecodingContext, execution::processor::ProcessorFault};

use super::machine::ItaniumMachine;

use colored::*;

pub struct ExecutionContext<'a, 'b> {
    pub decoding_context: DecodingContext<'a>,
    pub machine: &'b mut ItaniumMachine,
    pub instruction_index: usize,
    pub paused: bool
}

impl ExecutionContext<'_, '_> {
    pub fn new<'a, 'b>(decoding_context: DecodingContext<'a>, machine: &'b mut ItaniumMachine) -> ExecutionContext<'a, 'b> {
        ExecutionContext { 
            decoding_context: decoding_context, 
            machine: machine, 
            instruction_index: 0, 
            paused: false 
        }
    }

    pub fn step(&mut self) {
        let executable = &self.decoding_context.executable_instructions[self.instruction_index];
        
        self.instruction_index += 1;

        let execution_result = (executable.execution_function)(&mut self.machine, &executable.attributes, &mut self.instruction_index, &self.decoding_context.instruction_index_to_address, &self.decoding_context.address_to_instruction_index);
    }

    pub fn pause(&mut self) {
        self.paused = true
    }
    
    pub fn run(&mut self) {
        while self.machine.continue_running && !self.paused {
            self.step()
        }
    }
}