from math import exp
import random as __random

_REPLACE_ME_WITH_DIRECTIVES_

_REPLACE_ME_WITH_SOLUTION_

def main():
    testCases = [
        {
            "arguments": "findHeaterPower(100212373)",
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
        arguments = f"findHeaterPower({n})"
        testCases.append({ "got": got, "expected": expected, "arguments": arguments })

    for i, test in enumerate(testCases):
        if type(test["got"]) != int and type(test["got"]) != float and type(test["got"]) != str:
            print(f"# {i+1} FAILED")
            print(f"> ARGUMENTS {test['arguments']}")
            print(f"> EXPECTED {test['expected']}")
            print(f"> GOT {test['got']}")
            continue

        if type(test["got"]) == str and test["got"].isdigit() == False:
            print(f"# {i+1} FAILED")
            print(f"> ARGUMENTS {test['arguments']}")
            print(f"> EXPECTED {test['expected']}")
            print(f"> GOT {test['got']}")
            continue

        if int(test["got"]) == int(test["expected"]):
            print(f"# {i+1} PASSING")
            print(f"> ARGUMENTS {test['arguments']}")
            print(f"> EXPECTED {test['expected']}")
            print(f"> GOT {test['got']}")
        else:
            print(f"# {i+1} FAILED")
            print(f"> ARGUMENTS {test['arguments']}")
            print(f"> EXPECTED {test['expected']}")
            print(f"> GOT {test['got']}")

if __name__ == "__main__":
    main()
