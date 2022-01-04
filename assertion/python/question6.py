import random as __random

{0}

def main():
    testCases = [
        {
            "got": calculateGrade([73, 67, 38, 33]),
            "expected": [75, 67, 40, 33]
        }
    ]

    def workingAnswer(arr):
        out = []
        for grade in arr:
            if grade >= 38 and grade % 5 != 0 and grade % 5 >= 3:
                out.append(grade + (5 - grade % 5))
                continue

            out.append(grade)

        return out

    for i in range(9):
        inp = []
        arrLength = __random.randint(4, 20)
        for j in range(arrLength):
            inp.append(__random.randint(0, 100))

        expected = workingAnswer(inp)
        got = calculateGrade(inp)
        testCases.append({ "expected": expected, "got": got })

    for i, test in enumerate(testCases):
        if ", ".join(str(x) for x in test["expected"]) ==  ", ".join(str(x) for x in test["got"]):
            print(f"# {i+1} PASSING")
        else:
            print(f"# {i+1} FAILED")
            print("> EXPECTED {}".format(", ".join(str(x) for x in test["expected"])))
            print("> GOT {}".format(", ".join(str(x) for x in test["got"])))

if __name__ == "__main__":
    main()
