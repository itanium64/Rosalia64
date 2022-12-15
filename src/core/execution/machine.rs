use super::ItaniumProcessor;

pub struct ItaniumMachine {
    pub processor: ItaniumProcessor,
    pub memory: Vec<u8>,
    pub memory_size: usize,
    pub continue_running: bool
} 

impl ItaniumMachine {
    pub fn new(memory_size_bytes: usize) -> ItaniumMachine {
        let processor = ItaniumProcessor::new();
        let memory = vec![0u8; memory_size_bytes];

        ItaniumMachine { 
            processor: processor, 
            memory: memory,
            memory_size: memory_size_bytes,
            continue_running: true,
        }
    }
}