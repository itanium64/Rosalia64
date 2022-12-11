use std::collections::HashMap;

#[derive(Debug, PartialEq, Eq, PartialOrd, Ord, Hash)]
pub enum InstructionAttribute {
    R1,
    R2,
    R3,
    Immediate,
    QualifyingPredicate,
    TableX,
    TableY,
    Hint,
}

pub type InstructionAttributeMap = HashMap<InstructionAttribute, u64>;

#[derive(Debug, PartialEq, Eq, PartialOrd, Ord, Hash)]
pub enum LocalityHintCompleter {
    None = 0,
    NonTemporal1 = 1,
    NonTemporal2 = 2,
    NonTemporalAll = 3
}