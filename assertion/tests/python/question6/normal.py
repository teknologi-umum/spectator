# `calculateGrade` is a function that accepts a list of integer as its input and
# returns a list of integer as its output
def calculateGrade(Input):
    out = []
    for grade in Input:
        if grade >= 38 and grade % 5 != 0 and grade % 5 >= 3:
            out.append(grade + (5 - grade % 5))
            continue

        out.append(grade)

    return out