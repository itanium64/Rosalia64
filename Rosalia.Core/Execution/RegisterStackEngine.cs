using System.Collections.Concurrent;

namespace Rosalia.Core.Execution;

public class StackWindow {
    public ulong RegisterBase;
    public ulong SizeOfFrame;

    public ulong CountInputRegisters;
    public ulong CountLocalRegisters;
    public ulong CountOutputRegisters;
    public ulong CountRotatingRegisters;
}

public class RegisterStackEngine {
    public ConcurrentStack<StackWindow> _stackWindows;

    public RegisterStackEngine() {
        this._stackWindows = new ConcurrentStack<StackWindow>();
    }

    public StackWindow CurrentFrame() {
        this._stackWindows.TryPeek(out StackWindow window);

        return window;
    }

    public void NewFrame(ulong inputRegisters) {
        StackWindow currentFrame = this.CurrentFrame();

        this._stackWindows.Push(new StackWindow {
            RegisterBase = (currentFrame.RegisterBase + currentFrame.SizeOfFrame) - inputRegisters,
            SizeOfFrame = inputRegisters,
            CountInputRegisters = inputRegisters,
            CountLocalRegisters = 0,
            CountOutputRegisters = 0,
            CountRotatingRegisters = 0
        });
    }

    public void Allocate(ulong localRegisters, ulong outputRegisters, ulong rotatingRegisters) {
        StackWindow currentFrame = this.CurrentFrame();

        currentFrame.CountLocalRegisters = localRegisters;
        currentFrame.CountOutputRegisters = outputRegisters;
        currentFrame.CountRotatingRegisters = rotatingRegisters;

        currentFrame.SizeOfFrame = currentFrame.CountInputRegisters + currentFrame.CountOutputRegisters + currentFrame.CountLocalRegisters + currentFrame.CountRotatingRegisters;
    }
}
