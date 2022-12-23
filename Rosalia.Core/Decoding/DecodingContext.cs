namespace Rosalia.Core.Decoding;

public class DecodingContext : IDisposable {
    private MemoryStream _textStream;
    private BinaryReader _textReader;

    private ulong                    _instructionIndex;
    private Dictionary<ulong, ulong> _addressToInstructionIndex;
    private Dictionary<ulong, ulong> _instructionIndexToAddress;

    public DecodingContext(byte[] textSection) {
        this._textStream = new MemoryStream(textSection);
        this._textReader = new BinaryReader(this._textStream);

        this._addressToInstructionIndex = new Dictionary<ulong, ulong>();
        this._instructionIndexToAddress = new Dictionary<ulong, ulong>();
    }

    public void NextBundle() {
        InstructionBundle bundle = new InstructionBundle(
            lo: this._textReader.ReadUInt64(),
            hi: this._textReader.ReadUInt64()
        );
    }

    public void Dispose() {
        this._textStream.Dispose();
        this._textReader.Dispose();
    }
}
