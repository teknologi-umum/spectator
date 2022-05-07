#include <stdio.h>
#include <stdlib.h>
#include <vector>
#include <time.h>

_REPLACE_ME_WITH_DIRECTIVES_

int findHeaterPower(long power);

_REPLACE_ME_WITH_SOLUTION_

int __workingAnswer(long power) {
    int result = 0;
    while (power != 0) {
        result += power % 10;
        power /= 10;
    }

    return result;
}

typedef struct TestCase {
    int expected;
    int got;
} TestCase;

// creates a random number between min and max
long __randomNumber(int min, long max) {
    return (rand() % (max - min + 1)) + min;
}

int main() {
    srand(time(0));

    std::vector<TestCase> testCases{};
    testCases[0].expected = 19;
    testCases[0].got = findHeaterPower(100212373);

    for (int i = 0; i < 9; i++) {
        int n = __randomNumber(1000000000, 9999999999);
        TestCase testResult{};
        testResult.expected =  __workingAnswer(n);
        testResult.got = findHeaterPower(n);
        testCases.push_back(testResult);
    }

    for (unsigned int i = 0; i < testCases.size(); i++) {
        TestCase test = testCases.at(i);

        if (test.got == test.expected) {
            printf("# %d PASSING\n", i + 1);
        } else {
            printf("# %d FAILED\n", i + 1);
            printf("> EXPECTED %d\n", test.expected);
            printf("> GOT %d\n", test.got);
        }
    }
    return 0;
}
