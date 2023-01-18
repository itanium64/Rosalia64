namespace Rosalia.Core.Decoding;

public static class ExecutionSlotOrders {
    public enum UnitOrStop {
        None,
        Integer,
        Memory,
        Float,
        Branch,
        Extended,
        Stop,
        End,
    }

    public static readonly Dictionary<ulong, UnitOrStop[]> SlotOrders = new() {
         [0x00] = new [] { UnitOrStop.Memory,  UnitOrStop.Integer,  UnitOrStop.Integer,  UnitOrStop.End,     UnitOrStop.None, UnitOrStop.None  },
         [0x01] = new [] { UnitOrStop.Memory,  UnitOrStop.Integer,  UnitOrStop.Integer,  UnitOrStop.Stop,    UnitOrStop.End,  UnitOrStop.None  },
         [0x02] = new [] { UnitOrStop.Memory,  UnitOrStop.Integer,  UnitOrStop.Stop,     UnitOrStop.Integer, UnitOrStop.End,  UnitOrStop.None  },
         [0x03] = new [] { UnitOrStop.Memory,  UnitOrStop.Integer,  UnitOrStop.Stop,     UnitOrStop.Integer, UnitOrStop.Stop, UnitOrStop.End   },
         [0x04] = new [] { UnitOrStop.Memory,  UnitOrStop.Extended, UnitOrStop.Extended, UnitOrStop.End,     UnitOrStop.None, UnitOrStop.None  },
         [0x05] = new [] { UnitOrStop.Memory,  UnitOrStop.Extended, UnitOrStop.Extended, UnitOrStop.Stop,    UnitOrStop.End,  UnitOrStop.None  },
         [0x08] = new [] { UnitOrStop.Memory,  UnitOrStop.Memory,   UnitOrStop.Integer,  UnitOrStop.End,     UnitOrStop.None, UnitOrStop.None  },
         [0x09] = new [] { UnitOrStop.Memory,  UnitOrStop.Memory,   UnitOrStop.Integer,  UnitOrStop.Stop,    UnitOrStop.End,  UnitOrStop.None  },
         [0x0A] = new [] { UnitOrStop.Memory,  UnitOrStop.Stop,     UnitOrStop.Memory,   UnitOrStop.Integer, UnitOrStop.End,  UnitOrStop.None  },
         [0x0B] = new [] { UnitOrStop.Memory,  UnitOrStop.Stop,     UnitOrStop.Memory,   UnitOrStop.Integer, UnitOrStop.Stop, UnitOrStop.End   },
         [0x0C] = new [] { UnitOrStop.Memory,  UnitOrStop.Float,    UnitOrStop.Integer,  UnitOrStop.End,     UnitOrStop.None, UnitOrStop.None  },
         [0x0D] = new [] { UnitOrStop.Memory,  UnitOrStop.Float,    UnitOrStop.Integer,  UnitOrStop.End,     UnitOrStop.None, UnitOrStop.None  },
         [0x0E] = new [] { UnitOrStop.Memory,  UnitOrStop.Memory,   UnitOrStop.Float,    UnitOrStop.End,     UnitOrStop.None, UnitOrStop.None  },
         [0x0F] = new [] { UnitOrStop.Memory,  UnitOrStop.Memory,   UnitOrStop.Float,    UnitOrStop.Stop,    UnitOrStop.End,  UnitOrStop.None  },
         [0x10] = new [] { UnitOrStop.Memory,  UnitOrStop.Integer,  UnitOrStop.Branch,   UnitOrStop.End,     UnitOrStop.None, UnitOrStop.None  },
         [0x11] = new [] { UnitOrStop.Memory,  UnitOrStop.Integer,  UnitOrStop.Branch,   UnitOrStop.Stop,    UnitOrStop.End,  UnitOrStop.None  },
         [0x12] = new [] { UnitOrStop.Memory,  UnitOrStop.Branch,   UnitOrStop.Branch,   UnitOrStop.End,     UnitOrStop.None, UnitOrStop.None  },
         [0x13] = new [] { UnitOrStop.Memory,  UnitOrStop.Branch,   UnitOrStop.Branch,   UnitOrStop.Stop,    UnitOrStop.End,  UnitOrStop.None  },
         [0x16] = new [] { UnitOrStop.Branch,  UnitOrStop.Branch,   UnitOrStop.Branch,   UnitOrStop.End,     UnitOrStop.None, UnitOrStop.None  },
         [0x17] = new [] { UnitOrStop.Branch,  UnitOrStop.Branch,   UnitOrStop.Branch,   UnitOrStop.Stop,    UnitOrStop.End,  UnitOrStop.None  },
         [0x18] = new [] { UnitOrStop.Memory,  UnitOrStop.Memory,   UnitOrStop.Branch,   UnitOrStop.End,     UnitOrStop.None, UnitOrStop.None  },
         [0x19] = new [] { UnitOrStop.Memory,  UnitOrStop.Memory,   UnitOrStop.Branch,   UnitOrStop.Stop,    UnitOrStop.End,  UnitOrStop.None  },
         [0x1C] = new [] { UnitOrStop.Memory,  UnitOrStop.Float,    UnitOrStop.Branch,   UnitOrStop.End,     UnitOrStop.None, UnitOrStop.None  },
         [0x1D] = new [] { UnitOrStop.Memory,  UnitOrStop.Float,    UnitOrStop.Branch,   UnitOrStop.Stop,    UnitOrStop.End,  UnitOrStop.None  }
    };
}
