namespace Rosalia.Core.Execution.Registers;

public enum ApplicationRegisterDescription {
    RegisterStackConfiguration             = 16,
    BackingStorePointer                    = 17,
    BackingStorePointerStore               = 18,
    RegisterStackEngineNotAThingCollection = 19,
    Ia32FloatingPointControl               = 21,
    Ia32EFlag                              = 24,
    CodeSegmentDescriptor                  = 25,
    CompareAndStoreData                    = 25,
    StackSegmentDescriptor                 = 26,
    Ia32Cflg                               = 27,
    Ia32FloatingPointStatus                = 28,
    Ia32FloatingPointInstruction           = 29,
    Ia32FloatingPointData                  = 30,
    CompareAndExchangeValue                = 32,
    UserNotAThingCollection                = 36,
    FloatingPointStatus                    = 40,
    IntervalTimeCounter                    = 44,
    ResourceUtilizationCounter             = 45,
    PreviousFunctionState                  = 64,
    LoopCount                              = 65,
    EpilogCount                            = 66,
}

public abstract class ApplicationRegisterBase {
    public long RegisterId;
    public ulong Value;

    public virtual ulong RetrieveValue() {
        return this.Value;
    }

    public virtual ProcessorFault WriteValue(ulong value) {
        this.Value = value;

        return ProcessorFault.None;
    }
}

public class BackingStorePointer : ApplicationRegisterBase {
    public BackingStorePointer(ulong value) {
        this.Value      = value;
        this.RegisterId = (long) ApplicationRegisterDescription.BackingStorePointer;
    }
}

public class ApplicationRegister : ApplicationRegisterBase {
    public ApplicationRegister(long registerId) {
        this.RegisterId = registerId;
    }
}

public class ApplicationRegisterCollection {
    private readonly ApplicationRegisterBase[] _registers;

    public ApplicationRegisterCollection(ulong bspValue) {
        this._registers = new ApplicationRegisterBase[128];

        for (int i = 0; i < 128; i++) {
            this._registers[i] = new ApplicationRegister(i);
        }

        this._registers[(int) ApplicationRegisterDescription.BackingStorePointer] = new BackingStorePointer(bspValue);
    }

    public ApplicationRegisterBase RetrieveRegister(ulong register) {
        return this._registers[register];
    }

    public ApplicationRegisterBase RetrieveRegister(ApplicationRegisterDescription desc) {
        return this._registers[(int)desc];
    }
}
