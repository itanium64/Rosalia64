use std::env;

use exe::{VecPE, PE, ImageFileMachine};
use rosalia_core::{decoding::DecodingContext, execution::{ItaniumMachine, ExecutionContext}};

extern crate rosalia_core;

fn main() {
    let args: Vec<String> = env::args().collect();

    let image = VecPE::from_disk_file(&args[1]).unwrap();
    let opt_header = image.get_nt_headers_64().unwrap();

    //make sure it's a IA64 binary
    if opt_header.file_header.machine != ImageFileMachine::IA64 as u16 {
        println!("Executable has to be a IA64 Binary!");
        return;
    } 

    let entrypoint = 
        image
            .get_section_by_name( String::from(".text") )
            .expect("No code section found in executable!");

    let text_data = 
        entrypoint
            .read(&image)
            .expect("Failed to read code!");

    let mut decoding_context = 
        DecodingContext::new(
            &text_data[0..entrypoint.virtual_size as usize], 
            entrypoint.virtual_size as usize,
            opt_header.optional_header.image_base + entrypoint.virtual_address.0 as u64
        );

    decoding_context.decode_all();

    let mut machine = ItaniumMachine::new(512 * 1024 * 1024);
    let mut execution_context = ExecutionContext::new(decoding_context, &mut machine);

    execution_context.run();

    //Rust journey comes to an end as there is no way for me to make registers work.
}
