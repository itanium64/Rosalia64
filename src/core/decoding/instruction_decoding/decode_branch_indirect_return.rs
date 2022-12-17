use std::{collections::HashMap};

use crate::{decoding::{DecodingContext, instruction_formats::B4, InstructionAttribute}, execution::{self, instruction_execution}};

use super::disassembly_helpers::format_qualifying_predicate;

pub fn decode_branch_indirect_return(context: &mut DecodingContext, slot: u64, next_slot: u64) {
    let b4 = B4::from_slots(slot, next_slot);

    let branch_whether_completer_table = [
        ".sptk",
        ".spnt",
        ".dptk",
        ".dpnt"
    ];

    let sequential_prefetch_hint_completer = [
        ".few",
        ".many"
    ];

    let branch_cache_deallocation_hint_completer = [
        "",
        ".clr"
    ];

    let disassembly = format!("{} br.ret{}{}{} b{}", format_qualifying_predicate(b4.qp), branch_whether_completer_table[b4.wh as usize], sequential_prefetch_hint_completer[b4.p as usize], branch_cache_deallocation_hint_completer[b4.b2 as usize], b4.b2);

    let attributes: HashMap<InstructionAttribute, u64> = HashMap::from([
        (InstructionAttribute::B2, b4.b2),
        (InstructionAttribute::BranchType, b4.btype),
        (InstructionAttribute::BranchWhetherHint, b4.wh),
        (InstructionAttribute::SequentialPrefetchHint, b4.p),
        (InstructionAttribute::BranchDeallocateHint, b4.d),
        (InstructionAttribute::QualifyingPredicate, b4.qp),
    ]); 

    let executable_instruction = execution::ExecutableInstruction {
        execution_function: instruction_execution::execute_indirect_branch,
        disassembly: disassembly,
        attributes: attributes,
    };

    context.executable_instructions.push(executable_instruction);
}