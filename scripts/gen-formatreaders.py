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

def generateVariableDecoders(names, lengths, tab):
    if len(names) != len(lengths):
        print("There have to be an equal amount of names and lengths!")
        quit()

    zeroString = ""

    if skipOpcode:
        zeroString = "0000"

    requiredTotalBits = 46

    longestNameLength = 0

    stringTab = ""

    if tab:
        stringTab = "\t"

    for i in range(len(names)):
        length = len(names[i]) 

        if length > longestNameLength:
            longestNameLength = length

    for i in range(len(names)):
        underscores = (longestNameLength - len(names[i])) * '_'
        variableString = f"{stringTab}{underscores}{names[i]} := (instructionBits & (0b"
        bitString = zeroString

        length = int(lengths[i])

        bitString += '1' * length
        zeroString += '0' * length

        remainingZeroes = 46 - len(bitString)
        bitString += '0' * remainingZeroes

        variableString += bitString

        variableString += f")) >> {remainingZeroes}"

        print(variableString)

# Generation Starts Here

immediateCreated = False

print(f"type {formatName.upper()} struct" + " {")

for name in names:
    if name.lower().startswith("imm"):
        if immediateCreated == False:
            print("\tImmediate uint64")
            immediateCreated = True
        
        continue
    
    print(f"\t{name.capitalize()} uint64")

print("}\n")

print(f"func Read{formatName.upper()}(instructionBits uint64, nextSlot uint64) {formatName.upper()} " + "{")

generateVariableDecoders(names, lengths, True)

print("}")