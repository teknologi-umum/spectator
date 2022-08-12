#include <stdio.h>
#include <stdlib.h>
#include <vector>
#include <time.h>
#include <string>

_REPLACE_ME_WITH_DIRECTIVES_

bool isSameNumber(int a, int b);

_REPLACE_ME_WITH_SOLUTION_

typedef struct TestCase
{
    int expected;
    int got;
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

    std::vector<TestCase> testCases(2);
    testCases[0].expected = false;
    testCases[0].got = isSameNumber(100, 212);
    testCases[0].arguments = "isSameNumber(100, 212)";
    testCases[1].expected = true;
    testCases[1].got = isSameNumber(25, 25);
    testCases[1].arguments = "isSameNumber(25, 25)";

    for (int i = 0; i < 4; i++)
    {
        int a = __randomNumber(0, 9999);
        int b = __randomNumber(0, 9999);

        TestCase testResult{};
        testResult.expected = a == b;
        testResult.got = isSameNumber(a, b);
        testResult.arguments = "isSameNumber(" + std::to_string(a) + ", " + std::to_string(b) + ")";
        testCases.push_back(testResult);
    }

    for (int i = 0; i < 4; i++)
    {
        int a = __randomNumber(0, 9999);
        TestCase testResult{};
        testResult.expected = true;
        testResult.got = isSameNumber(a, a);
        testResult.arguments = "isSameNumber(" + std::to_string(a) + ", " + std::to_string(a) + ")";
        testCases.push_back(testResult);
    }

    for (unsigned int i = 0; i < testCases.size(); i++)
    {
        TestCase test = testCases.at(i);

        if (test.got == test.expected)
        {
            printf("# %d PASSING\n", i + 1);
        }
        else
        {
            printf("# %d FAILED\n", i + 1);
        }

        printf("> ARGUMENTS %s\n", test.arguments.c_str());
        printf("> EXPECTED %d\n", test.expected);
        printf("> GOT %d\n", test.got);
    }
    return 0;
}
