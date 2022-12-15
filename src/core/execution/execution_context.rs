use crate::{decoding::DecodingContext, execution::processor::ProcessorFault};

use super::machine::ItaniumMachine;

use colored::*;

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
            println!("Processor Fault! Instruction Index {}\n", self.instruction_index);

            if self.instruction_index != 0 {
                let previous_instruction = &self.decoding_context.executable_instructions.get(self.instruction_index - 1);
            
                if previous_instruction.is_some() {
                    let previous = previous_instruction.unwrap();

                    println!("0x{:08x}: {}", self.decoding_context.instruction_index_to_address[&((self.instruction_index - 1) as u64)], previous.disassembly);
                }
            }

            let written_out = format!("0x{:08x}: {}", self.decoding_context.instruction_index_to_address[&(self.instruction_index as u64)], executable.disassembly);
            
            println!("{}", written_out.red());

            if self.instruction_index + 1 > self.decoding_context.executable_instructions.len() {
                let next_instruction = &self.decoding_context.executable_instructions.get(self.instruction_index + 1);

                if next_instruction.is_some() {
                    let next = next_instruction.unwrap();

                    println!("0x{:08x}: {}", self.decoding_context.instruction_index_to_address[&((self.instruction_index + 1) as u64)], next.disassembly);
                }
            }

            let fault = execution_result.err().unwrap();

            println!("\nFault: {}", fault);

            if fault != ProcessorFault::SoftFault {
                return;
            }
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