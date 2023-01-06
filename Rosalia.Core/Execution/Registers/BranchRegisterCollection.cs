namespace Rosalia.Core.Execution.Registers;

public class BranchRegister {
    public long  RegisterId;
    public ulong Value;

    public BranchRegister(long registerId) {
        this.RegisterId = registerId;
    }

    public ulong RetrieveValue() {
        return this.Value;
    }

    public ProcessorFault WriteValue(ulong value) {
        this.Value = value;

        return ProcessorFault.None;
    }
}

public class BranchRegisterCollection {
    private readonly BranchRegister[] _registers;

    public BranchRegisterCollection() {
        this._registers = new BranchRegister[8];

        for (int i = 0; i < 8; i++) {
            this._registers[i] = new BranchRegister(i);
        }
    }

    public BranchRegister RetrieveRegister(long register) {
        return this._registers[register];
    }
}
