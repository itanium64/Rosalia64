use phf::phf_map;

use crate::core::execution;

#[derive(Clone)]
pub enum InstructionOrStop {
    None,
    Integer,
    Memory,
    Float,
    Branch,
    Extended,
    Stop,
    End,
}

type BundlePipeline = [InstructionOrStop; 6];

pub static SLOT_ORDERS: phf::Map<u32, BundlePipeline> = phf_map! {
    0x00u32 => [InstructionOrStop::Memory,  InstructionOrStop::Integer,  InstructionOrStop::Integer,  InstructionOrStop::End,     InstructionOrStop::None, InstructionOrStop::None],
    0x01u32 => [InstructionOrStop::Memory,  InstructionOrStop::Integer,  InstructionOrStop::Integer,  InstructionOrStop::Stop,    InstructionOrStop::End,  InstructionOrStop::None],
    0x02u32 => [InstructionOrStop::Memory,  InstructionOrStop::Integer,  InstructionOrStop::Stop,     InstructionOrStop::Integer, InstructionOrStop::End,  InstructionOrStop::None],
    0x03u32 => [InstructionOrStop::Memory,  InstructionOrStop::Integer,  InstructionOrStop::Stop,     InstructionOrStop::Integer, InstructionOrStop::Stop, InstructionOrStop::End],
    0x04u32 => [InstructionOrStop::Memory,  InstructionOrStop::Extended, InstructionOrStop::Extended, InstructionOrStop::End,     InstructionOrStop::None, InstructionOrStop::None],
    0x05u32 => [InstructionOrStop::Memory,  InstructionOrStop::Extended, InstructionOrStop::Extended, InstructionOrStop::Stop,    InstructionOrStop::End,  InstructionOrStop::None],
    0x08u32 => [InstructionOrStop::Memory,  InstructionOrStop::Memory,   InstructionOrStop::Integer,  InstructionOrStop::End,     InstructionOrStop::None, InstructionOrStop::None],
    0x09u32 => [InstructionOrStop::Memory,  InstructionOrStop::Memory,   InstructionOrStop::Integer,  InstructionOrStop::Stop,    InstructionOrStop::End,  InstructionOrStop::None],
    0x0Au32 => [InstructionOrStop::Memory,  InstructionOrStop::Stop,     InstructionOrStop::Memory,   InstructionOrStop::Integer, InstructionOrStop::End,  InstructionOrStop::None],
    0x0Bu32 => [InstructionOrStop::Memory,  InstructionOrStop::Stop,     InstructionOrStop::Memory,   InstructionOrStop::Integer, InstructionOrStop::Stop, InstructionOrStop::End],
    0x0Cu32 => [InstructionOrStop::Memory,  InstructionOrStop::Float,    InstructionOrStop::Integer,  InstructionOrStop::End,     InstructionOrStop::None, InstructionOrStop::None],
    0x0Du32 => [InstructionOrStop::Memory,  InstructionOrStop::Float,    InstructionOrStop::Integer,  InstructionOrStop::End,     InstructionOrStop::None, InstructionOrStop::None],
    0x0Eu32 => [InstructionOrStop::Memory,  InstructionOrStop::Memory,   InstructionOrStop::Float,    InstructionOrStop::End,     InstructionOrStop::None, InstructionOrStop::None],
    0x0Fu32 => [InstructionOrStop::Memory,  InstructionOrStop::Memory,   InstructionOrStop::Float,    InstructionOrStop::Stop,    InstructionOrStop::End,  InstructionOrStop::None],
    0x10u32 => [InstructionOrStop::Memory,  InstructionOrStop::Integer,  InstructionOrStop::Branch,   InstructionOrStop::End,     InstructionOrStop::None, InstructionOrStop::None],
    0x11u32 => [InstructionOrStop::Memory,  InstructionOrStop::Integer,  InstructionOrStop::Branch,   InstructionOrStop::Stop,    InstructionOrStop::End,  InstructionOrStop::None],
    0x12u32 => [InstructionOrStop::Memory,  InstructionOrStop::Branch,   InstructionOrStop::Branch,   InstructionOrStop::End,     InstructionOrStop::None, InstructionOrStop::None],
    0x13u32 => [InstructionOrStop::Memory,  InstructionOrStop::Branch,   InstructionOrStop::Branch,   InstructionOrStop::Stop,    InstructionOrStop::End,  InstructionOrStop::None],
    0x16u32 => [InstructionOrStop::Branch,  InstructionOrStop::Branch,   InstructionOrStop::Branch,   InstructionOrStop::End,     InstructionOrStop::None, InstructionOrStop::None],
    0x17u32 => [InstructionOrStop::Branch,  InstructionOrStop::Branch,   InstructionOrStop::Branch,   InstructionOrStop::Stop,    InstructionOrStop::End,  InstructionOrStop::None],
    0x1Cu32 => [InstructionOrStop::Memory,  InstructionOrStop::Float,    InstructionOrStop::Branch,   InstructionOrStop::End,     InstructionOrStop::None, InstructionOrStop::None],
    0x1Du32 => [InstructionOrStop::Memory,  InstructionOrStop::Float,    InstructionOrStop::Branch,   InstructionOrStop::Stop,    InstructionOrStop::End,  InstructionOrStop::None],
};