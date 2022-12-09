use std::fmt::Display;

use phf::phf_map;

#[derive(Clone, PartialEq, Eq)]
pub enum UnitOrStop {
    None,
    Integer,
    Memory,
    Float,
    Branch,
    Extended,
    Stop,
    End,
}

impl Display for UnitOrStop {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        match self {
            UnitOrStop::Integer => write!(f, "Integer"),
            UnitOrStop::Memory => write!(f, "Memory"),
            UnitOrStop::Float => write!(f, "Float"),
            UnitOrStop::Branch => write!(f, "Branch"),
            UnitOrStop::Extended => write!(f, "Extended"),
            _ => panic!("Invalid unit!")
        }
    }
}

type BundlePipeline = [UnitOrStop; 6];

pub static SLOT_ORDERS: phf::Map<u32, BundlePipeline> = phf_map! {
    0x00u32 => [UnitOrStop::Memory,  UnitOrStop::Integer,  UnitOrStop::Integer,  UnitOrStop::End,     UnitOrStop::None, UnitOrStop::None],
    0x01u32 => [UnitOrStop::Memory,  UnitOrStop::Integer,  UnitOrStop::Integer,  UnitOrStop::Stop,    UnitOrStop::End,  UnitOrStop::None],
    0x02u32 => [UnitOrStop::Memory,  UnitOrStop::Integer,  UnitOrStop::Stop,     UnitOrStop::Integer, UnitOrStop::End,  UnitOrStop::None],
    0x03u32 => [UnitOrStop::Memory,  UnitOrStop::Integer,  UnitOrStop::Stop,     UnitOrStop::Integer, UnitOrStop::Stop, UnitOrStop::End],
    0x04u32 => [UnitOrStop::Memory,  UnitOrStop::Extended, UnitOrStop::Extended, UnitOrStop::End,     UnitOrStop::None, UnitOrStop::None],
    0x05u32 => [UnitOrStop::Memory,  UnitOrStop::Extended, UnitOrStop::Extended, UnitOrStop::Stop,    UnitOrStop::End,  UnitOrStop::None],
    0x08u32 => [UnitOrStop::Memory,  UnitOrStop::Memory,   UnitOrStop::Integer,  UnitOrStop::End,     UnitOrStop::None, UnitOrStop::None],
    0x09u32 => [UnitOrStop::Memory,  UnitOrStop::Memory,   UnitOrStop::Integer,  UnitOrStop::Stop,    UnitOrStop::End,  UnitOrStop::None],
    0x0Au32 => [UnitOrStop::Memory,  UnitOrStop::Stop,     UnitOrStop::Memory,   UnitOrStop::Integer, UnitOrStop::End,  UnitOrStop::None],
    0x0Bu32 => [UnitOrStop::Memory,  UnitOrStop::Stop,     UnitOrStop::Memory,   UnitOrStop::Integer, UnitOrStop::Stop, UnitOrStop::End],
    0x0Cu32 => [UnitOrStop::Memory,  UnitOrStop::Float,    UnitOrStop::Integer,  UnitOrStop::End,     UnitOrStop::None, UnitOrStop::None],
    0x0Du32 => [UnitOrStop::Memory,  UnitOrStop::Float,    UnitOrStop::Integer,  UnitOrStop::End,     UnitOrStop::None, UnitOrStop::None],
    0x0Eu32 => [UnitOrStop::Memory,  UnitOrStop::Memory,   UnitOrStop::Float,    UnitOrStop::End,     UnitOrStop::None, UnitOrStop::None],
    0x0Fu32 => [UnitOrStop::Memory,  UnitOrStop::Memory,   UnitOrStop::Float,    UnitOrStop::Stop,    UnitOrStop::End,  UnitOrStop::None],
    0x10u32 => [UnitOrStop::Memory,  UnitOrStop::Integer,  UnitOrStop::Branch,   UnitOrStop::End,     UnitOrStop::None, UnitOrStop::None],
    0x11u32 => [UnitOrStop::Memory,  UnitOrStop::Integer,  UnitOrStop::Branch,   UnitOrStop::Stop,    UnitOrStop::End,  UnitOrStop::None],
    0x12u32 => [UnitOrStop::Memory,  UnitOrStop::Branch,   UnitOrStop::Branch,   UnitOrStop::End,     UnitOrStop::None, UnitOrStop::None],
    0x13u32 => [UnitOrStop::Memory,  UnitOrStop::Branch,   UnitOrStop::Branch,   UnitOrStop::Stop,    UnitOrStop::End,  UnitOrStop::None],
    0x16u32 => [UnitOrStop::Branch,  UnitOrStop::Branch,   UnitOrStop::Branch,   UnitOrStop::End,     UnitOrStop::None, UnitOrStop::None],
    0x17u32 => [UnitOrStop::Branch,  UnitOrStop::Branch,   UnitOrStop::Branch,   UnitOrStop::Stop,    UnitOrStop::End,  UnitOrStop::None],
    0x1Cu32 => [UnitOrStop::Memory,  UnitOrStop::Float,    UnitOrStop::Branch,   UnitOrStop::End,     UnitOrStop::None, UnitOrStop::None],
    0x1Du32 => [UnitOrStop::Memory,  UnitOrStop::Float,    UnitOrStop::Branch,   UnitOrStop::Stop,    UnitOrStop::End,  UnitOrStop::None],
};