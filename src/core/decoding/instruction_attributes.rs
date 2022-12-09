use std::collections::HashMap;

#[derive(Debug, PartialEq, Eq, PartialOrd, Ord, Hash)]
pub enum InstructionAttribute {
    R1,
    R2,
    R3,
    Immediate,
    QualifyingPredicate,
}

pub type InstructionAttributeMap = HashMap<InstructionAttribute, u64>;