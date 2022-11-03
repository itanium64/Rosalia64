# Rosalia64

Rosalia64 is a Itanium (IA64) Emulator built to run Itanium .exe Files. 
It also serves as a learning tool for me to figure out how a Architecture I find interesting works
aswell as satisfy my curiosity about Itanium.

## Development

Currently Rosalia is in heavy development and pretty much nothing is final,
current plan is for Rosalia to run Windows IA64 PE Binaries but this may change,
as Rosalia really only cares that It gets IA64 instructions that it can execute.

The Emulator is written in Go as it's quite a fast language and it's one I'm pretty
comfortable with, it also provides a simple syntax that can be easily adapted to other
languages if you so choose.

## Sources

1. [Intel® IA-64 Architecture Software Developer's manual](http://refspecs.linux-foundation.org/IA64-softdevman-vol3.pdf)

Covers how Bundles are formatted, how singular instructions are formatted, how Opcodes relate to what operation is executed,
for each instruction it also writes down pseudocode what the instruction actually does on the CPU side which is really helpful.

2. [IA-64 Application Instruction Set Architecture Guide](https://redirect.cs.umbc.edu/portal/help/architecture/aig.pdf)

The 1. Document refrences stuff from this document a fair few times, like for the explanation for what a 'stop' is.
It also covers alot about the format and instructions but I personally prefer 1. for this.
