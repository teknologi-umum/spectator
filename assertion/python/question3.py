import random as __random

{0}

def main():
    testCases = [
        {
            "got": isSameNumber(100, 212),
            "expected": False
        },
        {
            "got": isSameNumber(25, 25),
            "expected": True
        }
    ]

    for _ in range(4):
        a = __random.randint(0, 9999)
        b = __random.randint(0, 9999)
        expected = a == b
        got = isSameNumber(a, b)
        testCases.append({ "got": got, "expected": expected })

    for _ in range(4):
        a = __random.randint(0, 9999)
        expected = a == a
        got = isSameNumber(a, a)
        testCases.append({ "got": got, "expected": expected })

    for i, test in enumerate(testCases):
        if test["got"] == test["expected"]:
            print(f"# {i+1} PASSING")
        else:
            print(f"# {i+1} FAILED")
            print(f"> EXPECTED {test['expected']}")
            print(f"> GOT {test['got']}")

if __name__ == "__main__":
    main()
