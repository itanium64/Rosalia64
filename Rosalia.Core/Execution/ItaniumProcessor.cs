using Rosalia.Core.Execution.Registers;

namespace Rosalia.Core.Execution;

public class ItaniumProcessor {
    public readonly GeneralRegisterCollection     GeneralRegisters;
    public readonly PredicateRegisterCollection   PredicateRegisters;
    public readonly BranchRegisterCollection      BranchRegisters;
    public readonly FloatingRegisterCollection    FloatingRegisters;
    public readonly RegisterStackEngine           RegisterStackEngine;
    public readonly ApplicationRegisterCollection ApplicationRegisters;
    public          ulong                         InstructionPointer;

    public ItaniumProcessor() {
        this.GeneralRegisters     = new GeneralRegisterCollection();
        this.PredicateRegisters   = new PredicateRegisterCollection();
        this.BranchRegisters      = new BranchRegisterCollection();
        this.FloatingRegisters    = new FloatingRegisterCollection();
        this.ApplicationRegisters = new ApplicationRegisterCollection(0);
        this.RegisterStackEngine  = new RegisterStackEngine();
    }
}
