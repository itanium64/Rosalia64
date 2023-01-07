namespace Rosalia.Core.Execution;

public class ItaniumMachine {
    public readonly ItaniumProcessor Processor;
    public          bool             ContinueRunning;

    public ItaniumMachine() {
        this.Processor = new ItaniumProcessor();
    }
}
