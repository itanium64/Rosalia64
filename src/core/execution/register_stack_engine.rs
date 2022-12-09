pub struct StackWindow {
    register_base: u64,
    size_of_frame: u64,

    count_input_registers: u64,
    count_local_registers: u64,
    count_output_registers: u64,
    count_rotating_registers: u64,
}

pub struct RegisterStackEngine {
    stack_windows: Vec<StackWindow>
}

impl RegisterStackEngine {
    pub fn new() -> RegisterStackEngine {
        RegisterStackEngine { stack_windows: Vec::new() }
    }

    pub fn current_frame(&self) -> &mut StackWindow {
        return &mut self.stack_windows[self.stack_windows.len() - 1];
    }

    //Creates a new frame after a function call
    pub fn new_frame(&mut self, input_registers: u64) {
        let current_frame = self.current_frame();

        self.stack_windows.push(StackWindow { 
            register_base: (current_frame.register_base + current_frame.size_of_frame) - input_registers, 
            size_of_frame: input_registers, 
            count_input_registers: input_registers, 
            count_local_registers: 0, 
            count_output_registers: 0, 
            count_rotating_registers: 0 
        })
    }

    //Allocates more registers if necessary
    pub fn allocate(&mut self, local_registers: u64, output_registers: u64) {
        let current_frame = self.current_frame();

        current_frame.count_local_registers = local_registers;
        current_frame.count_output_registers = output_registers;

        current_frame.size_of_frame = current_frame.count_input_registers + current_frame.count_local_registers + current_frame.count_output_registers
    }
}