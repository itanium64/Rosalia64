// See https://aka.ms/new-console-template for more information

using Rosalia.Core.Decoding;
using Rosalia.Core.Execution;
using ExecutionContext=Rosalia.Core.Execution.ExecutionContext;

DecodingContext decodingContext = new DecodingContext(
    new byte[] {
        0x11, 0xF8, 0x24, 0x01, 0x24, 0x25, 0xE0, 0x01, 0x30, 0x00, 0x42, 0x00, 0x00, 0x00, 0x00, 0x20,
    },
    0x1400000000000000
);

decodingContext.NextBundle();

ExecutionContext executionContext = new ExecutionContext(
    decodingContext,
    new ItaniumMachine()
);

executionContext.Run();