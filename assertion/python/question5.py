import random as __random

{0}

def main():
    testCases = [
        {
            "expected": "A-Bb-Ccc-Dddd",
            "got": mumble("abcd")
        },
        {
            "expected": "R-Qq-Aaa-Eeee-Zzzzz-Tttttt-Yyyyyyy",
            "got": mumble("RqaEzTy")
        }
    ]

    def workingAnswer(s):
        result = ""
        for i in range(len(s)):
            c = s[i]
            result += c.upper() + c.lower() * i
            if i < len(s) - 1:
                result += "-"
        return result

    characters = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
    for i in range(8):
        chars = ""
        seed = __random.randint(4, 20)
        for j in range(seed):
            chars += __random.choice(characters.split())
        expected = workingAnswer(chars)
        got = mumble(chars)
        testCases.append({ "expected": expected, "got": got })

    for i in range(len(testCases)):
        test = testCases[i]
        if test["expected"] == test["got"]:
            print("# {} PASSING".format(i+1))
        else:
            print("# {} FAILED".format(i+1))
            print("> EXPECTED {}".format(test["expected"]))
            print("> GOT {}".format(test["got"]))

if __name__ == "__main__":
    main()
