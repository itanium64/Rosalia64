import sys

namesArg = ""
lengthsArg = ""
formatName = ""
skipOpcode = True

for arg in sys.argv:
    arg = arg.lower()

    if arg.startswith("-dontskip"):
        skipOpcode = False

    if arg.startswith("-names"):
        namesArg = arg.split(":")[1]
    
    if arg.startswith("-lengths"):
        lengthsArg = arg.split(":")[1]

    if arg.startswith("-format"):
        formatName = arg.split(":")[1]

names = namesArg.split(",")
lengths = lengthsArg.split(",")

underscoredNameDict = {}

def generateVariableDecoders(names, lengths, tab):
    if len(names) != len(lengths):
        print("There have to be an equal amount of names and lengths!")
        quit()

    zeroString = ""

    if skipOpcode:
        zeroString = "0000"

    requiredTotalBits = 41

    longestNameLength = 0

    stringTab = ""

    if tab:
        stringTab = "        "

    for i in range(len(names)):
        length = len(names[i])

        if length > longestNameLength:
            longestNameLength = length

    for i in range(len(names)):

        underscores = (longestNameLength - len(names[i])) * '_'
        variableString = f"{stringTab}ulong {underscores}{names[i]} = (slot & (0b"
        bitString = zeroString

        underscoredNameDict[names[i]] = f"{underscores}{names[i]}"

        length = int(lengths[i])

        bitString += '1' * length
        zeroString += '0' * length

        remainingZeroes = 41 - len(bitString)
        bitString += '0' * remainingZeroes

        variableString += bitString

        variableString += f")) >> {remainingZeroes};"

        if names[i] != "_":
            print(variableString)

# Generation Starts Here

immediateCreated = False

print("// ReSharper disable InconsistentNaming")
print("namespace Rosalia.Core.Decoding.Decoders;\n")

print(f"public struct {formatName.upper()}" + " {")

for name in names:
    if name.lower().startswith("imm"):
        if immediateCreated == False:
            print("    public ulong Immediate;")
            immediateCreated = True

        continue

    if name != "_":
        print(f"    public ulong {name.capitalize()};")

print(f"\n    public {formatName.upper()}(ulong slot, ulong nextSlot) " + "{")
generateVariableDecoders(names, lengths, True)

print("")

for name in names:
    if name.lower().startswith("imm"):
        if immediateCreated == True:
            print(f"        this.Immediate = (ulong)immediate;")
            immediateCreated = False

        continue
    if name != "_":
        print(f"        this.{name.capitalize()} = {underscoredNameDict[name]};")
print("    }")

print("}")