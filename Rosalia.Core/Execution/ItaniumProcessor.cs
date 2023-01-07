using Rosalia.Core.Execution.Registers;

namespace Rosalia.Core.Execution;

public class ItaniumProcessor {
    public readonly GeneralRegisterCollection   GeneralRegisters;
    public          PredicateRegisterCollection PredicateRegisters;
    public          BranchRegisterCollection    BranchRegisters;
    public          FloatingRegisterCollection  FloatingRegisters;
    public          RegisterStackEngine         RegisterStackEngine;
    public          ulong                       InstructionPointer;

    public ItaniumProcessor() {
        this.GeneralRegisters    = new GeneralRegisterCollection();
        this.PredicateRegisters  = new PredicateRegisterCollection();
        this.BranchRegisters     = new BranchRegisterCollection();
        this.FloatingRegisters   = new FloatingRegisterCollection();
        this.RegisterStackEngine = new RegisterStackEngine();
    }
}
