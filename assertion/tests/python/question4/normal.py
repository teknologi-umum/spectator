# `findHeaterPower` is function that accepts an argument as long and returns
# an integer as its output.
def findHeaterPower(Input):
    output = 0

    for i in range(len(str(Input))):
        output += int(str(Input)[i])

    return output