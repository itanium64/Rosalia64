use std::collections::HashMap;

use crate::{execution::{processor::ProcessorFault, ItaniumMachine}, decoding::{InstructionAttributeMap, InstructionAttribute}};

pub fn execute_int_load_no_base_update_form(
    /* Contains Processor and Memory */ machine: &mut ItaniumMachine, 
    /* Contains Function Parameters  */ attributes: &InstructionAttributeMap,
    /*       Instruction Pointer     */ _instruction_index: &mut usize,
    /*    Converts IP to an address  */ _instruction_index_to_address: &HashMap<u64, u64>,
    /*    Converts address to an IP  */ _address_to_instruction_index: &HashMap<u64, u64>,
) -> Result<(), ProcessorFault> {
    let processor = &mut machine.processor;
    
    let ___r1 = attributes[&InstructionAttribute::R1];
    let ___r3 = attributes[&InstructionAttribute::R3];
    let _hint = attributes[&InstructionAttribute::Hint];
    let tab_x = attributes[&InstructionAttribute::TableX];
    let tab_y = attributes[&InstructionAttribute::TableY];
    let ___qp = attributes[&InstructionAttribute::QualifyingPredicate];

    let byte_table: [usize; 4] = [1, 2, 4, 8];

    let mut speculative: bool = false;
    let mut advanced: bool = false;
    let mut bias: bool = false;
    let mut acquire: bool = false;
    let mut fill: bool = false;
    let mut check: bool = false;
    let mut check_clear: bool = false;
    let mut check_no_clear: bool = false;
    let mut no_clear: bool = false;
    let mut translate_address: bool = false;
    let mut read_memory: bool = true;

    let mut count_bytes = byte_table[tab_x as usize];

    let mut r1 = *processor.retrieve_general_register_mut(___r1).unwrap();
    let r3 = *processor.retrieve_general_register(___r3).unwrap();

    match tab_x {
        0 => {},
        1 => {
            speculative = true;
        },
        2 => {
            advanced = true;
        },
        3 => {
            speculative = true;
            advanced = true;
        },
        4 => {
            bias = true;
        },
        5 => {
            acquire = true;
        },
        6 => {
            count_bytes = 8;
            fill = true
        },
        8 => {
            check_clear = true;
            check = true;
        },
        9 => {
            check_no_clear = true;
            check = true;
        },
        10 => {
            check_clear = true;
            acquire = true;
        }
        _ => {
            println!("execute_int_load_no_base_update_form: unimplemented tab_x value!\ntab_x: {}", tab_x);
        }
    }
    
    if !speculative && r3.not_a_thing {
        return Err(ProcessorFault::NATConsumptionFault);
    }

    let defer_exception = false; //speculative && (r3.not_a_thing /* || PSR.ed (Processor Status Register; Exception Deferral) */);
    
    //if(check && false /* alat_cmp(GENERAL, r1) */) {
    //    translate_address = alat_translate_address_on_hit(ldtype, GENERAL, r1)
    //}

    let mut read_value: u64 = 0;

    if !translate_address {
        if check_clear || advanced {
            //alat_invalidate_single_entry(GENERAL, r1);
        }
    } else {
        let read_address = r3.read() as usize;

        if !defer_exception && read_memory {
            let read_bytes = &machine.memory[read_address..read_address + count_bytes];

            read_value = u64::from_le_bytes(read_bytes.try_into().unwrap());
        }

        if check_clear || advanced {
            //alat_invalidate_single_entry(GENERAL, r1);
        }

        if defer_exception {
            if speculative {
                //r1 = natd_gr_read(address, size, UM.be, mattr, otype, bias | ldhint)
                //r1.nat = 1
            } else  {
                let val_result = r1.write(0);
                let nat_result = r1.write_nat(false);

                if val_result.is_err() {
                    return val_result;
                }

                if nat_result.is_err() {
                    return nat_result;
                }
            }
        } else {
            let val_result = r1.write(read_value);
            let nat_result = r1.write_nat(false);

            if val_result.is_err() {
                return val_result;
            }

            if nat_result.is_err() {
                return nat_result;
            }
        }

        if (check_no_clear || advanced) && false /* ma_is_speculative(mattr) */ {
            //alat_write(ldtype, GENERAL, r1, paddr, size);
        }
    }

    return Ok(())
}