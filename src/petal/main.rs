use std::env;

use exe::{VecPE, ImageOptionalHeader64, PE, ImageFileMachine};

extern crate rosalia_core;

fn main() {
    let args: Vec<String> = env::args().collect();

    let image = VecPE::from_disk_file(&args[1]).unwrap();
    let opt_header = image.get_nt_headers_64().unwrap();

    println!("Machine: {}; is IA64?: {}", opt_header.file_header.machine, opt_header.file_header.machine == ImageFileMachine::IA64 as u16)
}
