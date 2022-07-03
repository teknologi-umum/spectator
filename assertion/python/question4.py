from math import exp
import random as __random

_REPLACE_ME_WITH_DIRECTIVES_

_REPLACE_ME_WITH_SOLUTION_

def main():
    testCases = [
        {
            "got": findHeaterPower(100212373),
            "expected": 19
        }
    ]

    def workingAnswer(n):
        s = str(n)
        sum = 0
        for i in range(len(s)):
            sum += int(s[i])
        return sum

    for _ in range(9):
        n = __random.randint(1000000000, 9999999999)
        expected = workingAnswer(n)
        got = findHeaterPower(n)
        testCases.append({ "got": got, "expected": expected })

    for i, test in enumerate(testCases):
        if type(test["got"]) != int and type(test["got"]) != float:
            print(f"# {i+1} FAILED")
            print(f"> EXPECTED {test['expected']}")
            print(f"> GOT {test['got']}")
            continue

        if test["got"] == test["expected"]:
            print(f"# {i+1} PASSING")
        else:
            print(f"# {i+1} FAILED")
            print(f"> EXPECTED {test['expected']}")
            print(f"> GOT {test['got']}")

if __name__ == "__main__":
    main()
