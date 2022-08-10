#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <string>
#include <time.h>
#include <vector>
#include <cmath>
#include <cfloat>

_REPLACE_ME_WITH_DIRECTIVES_

float calculateTemperature(float temp, std::string from, std::string to);

_REPLACE_ME_WITH_SOLUTION_

const std::string __CELCIUS = "Celcius";
const std::string __FAHRENHEIT = "Fahrenheit";
const std::string __KELVIN = "Kelvin";

bool isC(std::string unit) { return unit.compare(__CELCIUS) == 0; }
bool isF(std::string unit) { return unit.compare(__FAHRENHEIT) == 0; }
bool isK(std::string unit) { return unit.compare(__KELVIN) == 0; }

float __workingAnswer(int n, std::string a, std::string b)
{
    if (isC(a) && isF(b))
        return (n * 9 / 5.0f) + 32.0f;
    if (isC(a) && isK(b))
        return n + 273.15f;
    if (isF(a) && isC(b))
        return (n - 32) * 5 / 9.0f;
    if (isF(a) && isK(b))
        return (n - 32) * 5 / 9.0f + 273.15f;
    if (isK(a) && isC(b))
        return n - 273.15f;
    if (isK(a) && isF(b))
        return (n - 273.15f) * 9 / 5.0f + 32.0f;
    return n;
}

typedef struct TestCase
{
    float expected;
    float got;
    std::string arguments;
} TestCase;

// creates a random number between min and max
int __randomNumber(int min, int max)
{
    return (rand() % (max - min + 1)) + min;
}

int main()
{
    srand(time(0));

    std::vector<std::string> temperatures{"Celcius", "Fahrenheit", "Kelvin"};

    std::vector<TestCase> testCases(5);

    testCases[0].expected = 212.0;
    testCases[0].got = calculateTemperature(100, "Celcius", "Fahrenheit");
    testCases[0].arguments = "calculateTemperature(100, \"Celcius\", \"Fahrenheit\")";
    testCases[1].expected = 373.15;
    testCases[1].got = calculateTemperature(212, "Fahrenheit", "Kelvin");
    testCases[1].arguments = "calculateTemperature(212, \"Fahrenheit\", \"Kelvin\")";
    testCases[2].expected = 273.15;
    testCases[2].got = calculateTemperature(0, "Celcius", "Kelvin");
    testCases[2].arguments = "calculateTemperature(0, \"Celcius\", \"Kelvin\")";
    testCases[3].expected = 32.0;
    testCases[3].got = calculateTemperature(0, "Celcius", "Fahrenheit");
    testCases[3].arguments = "calculateTemperature(0, \"Celcius\", \"Fahrenheit\")";
    testCases[4].expected = -459.67;
    testCases[4].got = calculateTemperature(0, "Kelvin", "Fahrenheit");
    testCases[4].arguments = "calculateTemperature(0, \"Kelvin\", \"Fahrenheit\")";

    for (unsigned int i = 0; i < 5; i++)
    {
        // Generate random test cases
        int randNum = __randomNumber(-500, 500);
        std::string from = temperatures.at(__randomNumber(0, 2));
        std::string to = temperatures.at(__randomNumber(0, 2));

        TestCase testResult{};
        testResult.expected = __workingAnswer(randNum, from, to);
        testResult.got = calculateTemperature(randNum, from, to);
        testResult.arguments = "calculateTemperature(" + std::to_string(randNum) + ", \"" + from + "\", \"" + to + "\")";
        testCases.push_back(testResult);
    }

    for (unsigned int i = 0; i < testCases.size(); i++)
    {
        TestCase test = testCases.at(i);

        if (fabs(test.got - test.expected) < FLT_EPSILON)
        {
            printf("# %d PASSING\n", i + 1);
        }
        else
        {
            printf("# %d FAILED\n", i + 1);
        }
        printf("> ARGUMENTS %s\n", test.arguments.c_str());
        printf("> EXPECTED %.2f\n", test.expected);
        printf("> GOT %.2f\n", test.got);
    }

    return 0;
}
