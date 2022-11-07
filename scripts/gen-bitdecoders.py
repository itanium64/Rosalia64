import sys

namesArg = ""
lengthsArg = ""
skipOpcode = True

for arg in sys.argv:
    arg = arg.lower()

    if arg.startswith("-dontskip"):
        skipOpcode = False

    if arg.startswith("-names"):
        namesArg = arg.split(":")[1]
    
    if arg.startswith("-lengths"):
        lengthsArg = arg.split(":")[1]

names = namesArg.split(",")
lengths = lengthsArg.split(",")

if len(names) != len(lengths):
    print("There have to be an equal amount of names and lengths!")
    quit()

zeroString = ""

if skipOpcode:
    zeroString = "0000"

requiredTotalBits = 46

for i in range(len(names)):
    variableString = f"{names[i]} := (instructionBits & (0b"
    bitString = zeroString

    length = int(lengths[i])

    bitString += '1' * length
    zeroString += '0' * length

    remainingZeroes = 46 - len(bitString)
    bitString += '0' * remainingZeroes

    variableString += bitString

    variableString += f")) >> {remainingZeroes}"
    
    print(variableString)