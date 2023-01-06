namespace Rosalia.Core.Execution.Registers;

public abstract class FloatingRegisterBase {
    public long RegisterId;
    public double Value;

    public virtual double RetrieveValue() {
        return this.Value;
    }

    public virtual ProcessorFault WriteValue(double value) {
        this.Value = value;

        return ProcessorFault.None;
    }
}

public class FloatingRegister : FloatingRegisterBase {
    public FloatingRegister(long registerId) {
        this.RegisterId = registerId;
    }
}

public class FloatingRegisterZero : FloatingRegisterBase {
    public FloatingRegisterZero() {
        this.RegisterId = 0;
    }

    public override double RetrieveValue() => 0.0;
    public override ProcessorFault WriteValue(double value) => ProcessorFault.IllegalOperation;
}

public class FloatingRegisterOne : FloatingRegisterBase {
    public FloatingRegisterOne() {
        this.RegisterId = 1;
    }

    public override double RetrieveValue() => 1.0;
    public override ProcessorFault WriteValue(double value) => ProcessorFault.IllegalOperation;
}

public class FloatingRegisterCollection {
    private readonly FloatingRegisterBase[] _registers;

    public FloatingRegisterCollection() {
        this._registers    = new FloatingRegisterBase[128];
        this._registers[0] = new FloatingRegisterZero();
        this._registers[1] = new FloatingRegisterOne();

        for (int i = 2; i < 128; i++) {
            this._registers[i] = new FloatingRegister(i);
        }
    }

    public FloatingRegisterBase RetrieveRegister(long register) {
        return this._registers[register];
    }
}
