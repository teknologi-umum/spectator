#include <stdio.h>
#include <stdlib.h>
#include <vector>

int isSameNumber(int a, int b);

{0}

typedef struct TestCase {
    int expected;
    int got;
} TestCase;

// creates a random number between min and max
int __randomNumber(int min, int max) {
    return (rand() % (max - min + 1)) + min;
}

int main() {
    std::vector<TestCase> testCases{
        {.expected = false,
         .got = isSameNumber(100, 212)},
        {.expected = true,
         .got = isSameNumber(25, 25)}};

    for (int i = 0; i < 4; i++) {
        int a = __randomNumber(0, 9999);
        int b = __randomNumber(0, 9999);
        int expected = a == b;
        int got = isSameNumber(a, b);
        testCases.insert(testCases.end(),
                         (TestCase){ .expected = expected, .got = got });
    }

    for (int i = 0; i < 4; i++) {
        int a = __randomNumber(0, 9999);
        int expected = true;
        int got = isSameNumber(a, a);
        testCases.insert(testCases.end(),
                         (TestCase){ .expected = expected, .got = got });
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
