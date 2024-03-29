import random as __random

_REPLACE_ME_WITH_DIRECTIVES_

_REPLACE_ME_WITH_SOLUTION_

def main():
    testCases = [
        {
            "arguments": "mumble(\"abcd\")",
            "expected": "A-Bb-Ccc-Dddd",
            "got": mumble("abcd")
        },
        {
            "arguments": "mumble(\"RqaEzTy\")",
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
        arguments = f"mumble({chars})"
        testCases.append({ "expected": expected, "got": got, "arguments": arguments })

    for i, test in enumerate(testCases):
        if type(test["got"]) != str:
            print(f"# {i+1} FAILED")
            print(f"> ARGUMENTS {test['arguments']}")
            print(f"> EXPECTED {test['expected']}")
            print(f"> GOT \"{test['got']}\"")
            continue

        if test["expected"] == test["got"]:
            print(f"# {i+1} PASSING")
            print(f"> ARGUMENTS {test['arguments']}")
            print(f"> EXPECTED \"{test['expected']}\"")
            print(f"> GOT \"{test['got']}\"")
        else:
            print(f"# {i+1} FAILED")
            print(f"> ARGUMENTS {test['arguments']}")
            print(f"> EXPECTED \"{test['expected']}\"")
            print(f"> GOT \"{test['got']}\"")

if __name__ == "__main__":
    main()
