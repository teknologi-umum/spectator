#include <stdio.h>
#include <stdlib.h>
#include <string>
#include <vector>
#include <time.h>
#include <sstream>

_REPLACE_ME_WITH_DIRECTIVES_

std::vector<int> calculateGrade(std::vector<int> input);

_REPLACE_ME_WITH_SOLUTION_

std::vector<int> __workingAnswer(std::vector<int> input)
{
    std::vector<int> result;

    for (auto &grade : input)
    {
        if (grade >= 38 && grade % 5 != 0 && grade % 5 >= 3)
        {
            result.push_back(grade + (5 - grade % 5));
        }
        else
        {
            result.push_back(grade);
        }
    }

    return result;
}

typedef struct TestCase
{
    std::vector<int> expected;
    std::vector<int> got;
    std::string arguments;
} TestCase;

// creates a random number between min and max
int __randomNumber(int min, int max)
{
    return (rand() % (max - min + 1)) + min;
}

std::vector<int> genVector(int n)
{
    std::vector<int> result;
    for (int i = 0; i < n; i++)
    {
        result.push_back(__randomNumber(0, 100));
    }
    return result;
}

int main()
{
    srand(time(0));

    std::vector<TestCase> testCases(1);
    testCases[0].expected = {75, 67, 40, 33};
    testCases[0].got = calculateGrade({75, 67, 40, 33});
    testCases[0].arguments = "calculateGrade({75, 67, 40, 33})";

    for (int i = 0; i < 8; i++)
    {
        int len = __randomNumber(4, 20);
        std::vector<int> input = genVector(len);

        TestCase testResult{};
        testResult.expected = __workingAnswer(input);
        testResult.got = calculateGrade(input);

        std::stringstream ss;
        for (size_t i = 0; i < input.size(); ++i)
        {
            if (i != 0)
            {
                ss << ", ";
            }

            ss << input[i];
        }

        testResult.arguments = "calculateGrade({" + ss.str() + "})";
        testCases.push_back(testResult);
    }

    for (unsigned int i = 0; i < testCases.size(); i++)
    {
        TestCase test = testCases.at(i);

        std::stringstream expectedStream;
        for (size_t i = 0; i < test.expected.size(); ++i)
        {
            if (i != 0)
            {
                expectedStream << ", ";
            }

            expectedStream << test.expected[i];
        };

        std::stringstream gotStream;
        for (size_t i = 0; i < test.got.size(); ++i)
        {
            if (i != 0)
            {
                gotStream << ", ";
            }

            gotStream << test.got[i];
        }

        if (gotStream.str() == expectedStream.str())
        {
            printf("# %d PASSING\n", i + 1);
        }
        else
        {
            printf("# %d FAILED\n", i + 1);
        }

        printf("> ARGUMENTS %s\n", test.arguments.c_str());
        printf("> EXPECTED %s\n", expectedStream.str().c_str());
        printf("> GOT %s\n", gotStream.str().c_str());
    }
    return 0;
}
