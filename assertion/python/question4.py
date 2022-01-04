from math import exp
import random as __random

{0}

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

    for i in range(len(testCases)):
        test = testCases[i]
        if test["got"] == test["expected"]:
            print("# {} PASSING".format(i+1))
        else:
            print("# {} FAILED".format(i+1))
            print("> EXPECTED {}".format(test["expected"]))
            print("> GOT {}".format(test["got"]))

if __name__ == "__main__":
    main()
