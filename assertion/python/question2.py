import random as __random
from decimal import Decimal

_REPLACE_ME_WITH_DIRECTIVES_

_REPLACE_ME_WITH_SOLUTION_

def main():
    testCases = [
        {
            "got": calculateTemperature(100, "Celcius", "Fahrenheit"),
            "expected": 212
        },
        {
            "got": calculateTemperature(212, "Fahrenheit", "Kelvin"),
            "expected": 373.15
        },
        {
            "got": calculateTemperature(0, "Celcius", "Kelvin"),
            "expected": 273.15
        },
        {
            "got": calculateTemperature(0, "Celcius", "Fahrenheit"),
            "expected": 32
        },
        {
            "got": calculateTemperature(0, "Kelvin", "Fahrenheit"),
            "expected": -459.67
        }
    ]

    def workingAnswer(n, a, b):
        if a == "Celcius" and b == "Fahrenheit":
            return (n * 9 / 5) + 32
        elif a == "Celcius" and b == "Kelvin":
            return n + 273.15
        elif a == "Fahrenheit" and b == "Celcius":
            return (n - 32) * 5 / 9
        elif a == "Fahrenheit" and b == "Kelvin":
            return (n - 32) * 5 / 9 + 273.15
        elif a == "Kelvin" and b == "Celcius":
            return n - 273.15
        elif a == "Kelvin" and b == "Fahrenheit":
            return (n - 273.15) * 9 / 5 + 32

        return n

    temperatures = ["Celcius", "Fahrenheit", "Kelvin"]
    for _ in range(5):
        fromTemperature = __random.choice(temperatures)
        toTemperature = __random.choice(temperatures)
        n = __random.randint(-500, 500)
        expected = workingAnswer(n, fromTemperature, toTemperature)
        got = calculateTemperature(n, fromTemperature, toTemperature)
        testCases.append({ "expected": expected, "got": got })

    for i, test in enumerate(testCases):
        got = test["got"]
        expected = test["expected"]

        if type(got) == None:
            print(f"# {i+1} FAILED")
            print(f"> EXPECTED {expected}")
            print(f"> GOT {got}")
            continue

        if round(float(got), 2) == round(float(expected), 2):
            print(f'# {i+1} PASSING')
        else:
            print(f"# {i+1} FAILED")
            print(f"> EXPECTED { round(float(test['expected']), 2) }")
            print(f"> GOT { round(float(test['got']), 2) }")

if __name__ == "__main__":
    main()
