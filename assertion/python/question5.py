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
        result = []
        for index,letter in enumerate(s):
            result.append( (letter*(index+1)).capitalize() )
        return "-".join(result)

    characters = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
    for i in range(8):
        seed = __random.randint(4, 20)
        chars = __random.choices(characters, k=seed)
        expected = workingAnswer(chars)
        got = mumble(chars)
        testCases.append({ "expected": expected, "got": got })

    for i, test in enumerate(testCases):
        if test["expected"] == test["got"]:
            print(f"# {i+1} PASSING")
        else:
            print(f"# {i+1} FAILED")
            print(f"> EXPECTED {test['expected']}")
            print(f"> GOT {test['got']}")

if __name__ == "__main__":
    main()
