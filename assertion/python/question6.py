import random as __random

_REPLACE_ME_WITH_DIRECTIVES_

_REPLACE_ME_WITH_SOLUTION_

def main():
    testCases = [
        {
            "arguments": "calculateGrade([73, 67, 38, 33])",
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
        arrLength = __random.randint(4, 20)
        inp = __random.sample(range(0,100), arrLength)
        
        expected = workingAnswer(inp)
        got = calculateGrade(inp)
        arguments = f"calculateGrade([{', '.join(map(str, inp))}])"
        testCases.append({ "expected": expected, "got": got, "arguments": arguments })

    for i, test in enumerate(testCases):
        if ", ".join( map(str,test["expected"]) ) ==  ", ".join( map(str,test["got"]) ):
            print(f"# {i+1} PASSING")
            print(f"> ARGUMENTS {test['arguments']}")
            print(f"> EXPECTED [{ ', '.join(map(str,test['expected'])) }]" )
            print(f"> GOT [{ ', '.join(map(str,test['got'])) }]" )
        else:
            print(f"# {i+1} FAILED")
            print(f"> ARGUMENTS {test['arguments']}")
            print(f"> EXPECTED [{ ', '.join(map(str,test['expected'])) }]" )
            print(f"> GOT [{ ', '.join(map(str,test['got'])) }]" )

if __name__ == "__main__":
    main()
