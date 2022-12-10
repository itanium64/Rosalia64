use super::ItaniumProcessor;

pub struct ItaniumMachine {
    pub processor: ItaniumProcessor,
    pub memory: Vec<u8>
} 

impl ItaniumMachine {
    pub fn new(memory_size_bytes: usize) -> ItaniumMachine {
        let processor = ItaniumProcessor::new();
        let memory = vec![0u8; memory_size_bytes];

        ItaniumMachine { 
            processor: processor, 
            memory: memory 
        }
    }
}