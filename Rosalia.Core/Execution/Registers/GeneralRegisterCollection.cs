namespace Rosalia.Core.Execution.Registers;

public abstract class GeneralRegisterBase {
    public long RegisterId;
    public long Value;
    public bool NotAThing;

    public virtual long RetrieveValue() {
        return this.Value;
    }

    public virtual bool RetrieveNotAThing() {
        return this.NotAThing;
    }

    public virtual ProcessorFault WriteValue(long value) {
        this.Value = value;

        return ProcessorFault.None;
    }

    public virtual ProcessorFault WriteNotAThing(bool nat) {
        this.NotAThing = nat;

        return ProcessorFault.None;
    }
}

public class GeneralRegister : GeneralRegisterBase {
    public GeneralRegister(long registerId) {
        this.RegisterId = registerId;
    }
}

public class GeneralRegisterZero : GeneralRegisterBase {
    public GeneralRegisterZero() {
        this.RegisterId = 0;
    }

    public override long RetrieveValue() => 0;
    public override bool RetrieveNotAThing() => false;
    public override ProcessorFault WriteValue(long value) => ProcessorFault.IllegalOperation;
    public override ProcessorFault WriteNotAThing(bool nat) => ProcessorFault.IllegalOperation;
}

public class GeneralRegisterCollection {
    private readonly GeneralRegisterBase[] _registers;

    public GeneralRegisterCollection() {
        this._registers    = new GeneralRegisterBase[128];
        this._registers[0] = new GeneralRegisterZero();

        for (int i = 1; i < 128; i++) {
            this._registers[i] = new GeneralRegister(i);
        }
    }

    public GeneralRegisterBase RetrieveRegister(long register) {
        return this._registers[register];
    }
}
