namespace Rosalia.Core.Execution.Registers;

public abstract class PredicateRegisterBase {
    public long RegisterId;
    public bool Value;

    public virtual bool RetrieveValue() {
        return this.Value;
    }

    public virtual ProcessorFault WriteValue(bool value) {
        this.Value = value;

        return ProcessorFault.None;
    }
}

public class PredicateRegister : PredicateRegisterBase {
    public PredicateRegister(long registerId) {
        this.RegisterId = registerId;
    }
}

public class PredicateRegisterZero : PredicateRegisterBase {
    public PredicateRegisterZero() {
        this.RegisterId = 0;
    }

    public override bool RetrieveValue() => true;
    public override ProcessorFault WriteValue(bool value) => ProcessorFault.IllegalOperation;
}

public class PredicateRegisterCollection {
    private readonly PredicateRegisterBase[] _registers;

    public PredicateRegisterCollection() {
        this._registers    = new PredicateRegisterBase[64];
        this._registers[0] = new PredicateRegisterZero();

        for (int i = 1; i < 64; i++) {
            this._registers[i] = new PredicateRegister(i);
        }
    }

    public PredicateRegisterBase RetrieveRegister(long register) {
        return this._registers[register];
    }
}
