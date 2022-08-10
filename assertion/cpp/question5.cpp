#include <stdio.h>
#include <stdlib.h>
#include <string>
#include <vector>
#include <numeric>
#include <time.h>

_REPLACE_ME_WITH_DIRECTIVES_

std::string mumble(std::string input);

_REPLACE_ME_WITH_SOLUTION_

std::string __workingAnswer(std::string input)
{
    std::vector<std::string> words;

    for (unsigned int i = 0; i < input.size(); i++)
    {
        std::string m = "";
        for (unsigned int j = 0; j <= i; j++)
        {
            char c = input.at(i);
            m.push_back(j == 0 ? std::toupper(c) : std::tolower(c));
        }
        words.push_back(m);
    }

    std::string result = std::accumulate(
        words.begin(), words.end(), std::string(),
        [](const std::string &acc, const std::string &curr) -> std::string
        {
            return acc + (acc.length() > 0 ? "-" : "") + curr;
        });

    return result;
}

typedef struct TestCase
{
    std::string expected;
    std::string got;
    std::string arguments;
} TestCase;

// creates a random number between min and max
int __randomNumber(int min, int max)
{
    return (rand() % (max - min + 1)) + min;
}

const std::string characters = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz";
std::string __genWords(int n)
{
    std::string result;

    const int nchar = characters.size() - 1;
    for (int i = 0; i < n; i++)
    {
        const char randomChar = characters.at(__randomNumber(0, nchar));
        result.append(1, randomChar);
    }

    return result;
}

int main()
{
    srand(time(0));

    std::vector<TestCase> testCases(2);
    testCases[0].expected = "A-Bb-Ccc-Dddd";
    testCases[0].got = mumble("abcd");
    testCases[0].arguments = "mumble(\"abcd\")";
    testCases[1].expected = "R-Qq-Aaa-Eeee-Zzzzz-Tttttt-Yyyyyyy";
    testCases[1].got = mumble("RqaEzTy");
    testCases[1].arguments = "mumble(\"RqaEzTy\")";

    for (int i = 0; i < 8; i++)
    {
        int n = __randomNumber(4, 20);
        std::string word = __genWords(n);

        TestCase testResult{};
        testResult.expected = __workingAnswer(word);
        testResult.got = mumble(word);
        testResult.arguments = "mumble(\"" + word + "\")";
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
        printf("> EXPECTED %s\n", test.expected.c_str());
        printf("> GOT %s\n", test.got.c_str());
    }
    return 0;
}
