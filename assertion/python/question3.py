import random as __random

_REPLACE_ME_WITH_DIRECTIVES_

_REPLACE_ME_WITH_SOLUTION_

def main():
    testCases = [
        {
            "arguments": "isSameNumber(100, 212)",
            "got": isSameNumber(100, 212),
            "expected": False
        },
        {
            "arguments": "isSameNumber(25, 25)",
            "got": isSameNumber(25, 25),
            "expected": True
        }
    ]

    for _ in range(4):
        a = __random.randint(0, 9999)
        b = __random.randint(0, 9999)
        expected = a == b
        got = isSameNumber(a, b)
        arguments = f"isSameNumber({a}, {b})"
        testCases.append({ "got": got, "expected": expected, "arguments": arguments })

    for _ in range(4):
        a = __random.randint(0, 9999)
        expected = a == a
        got = isSameNumber(a, a)
        arguments = f"isSameNumber({a}, {a})"
        testCases.append({ "got": got, "expected": expected, "arguments": arguments })

    for i, test in enumerate(testCases):
        if test["got"] == test["expected"]:
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
