use crate::decoding::DecodingContext;

use super::machine::ItaniumMachine;

pub struct ExecutionContext<'a, 'b> {
    decoding_context: DecodingContext<'a>,
    machine: &'b mut ItaniumMachine,
    instruction_index: usize,
    paused: bool
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
        let execution_result = (executable.execution_function)(&mut self.machine, &executable.attributes);

        if execution_result.is_err() {
            println!("Processor Fault! Instruction Index {}", self.instruction_index);
            println!(": {}", executable.disassembly);
            println!("Fault: {}", execution_result.err().unwrap())
        }
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