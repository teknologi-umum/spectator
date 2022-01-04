#include <stdio.h>
#include <stdlib.h>

int isSameNumber(int a, int b);

{0}

struct TestCase {
    int expected;
    int got;
};

// creates a random number between min and max
int __randomNumber(int min, int max) {
    return (rand() % (max - min + 1)) + min;
}

int main() {
    struct TestCase testCases[10];

    testCases[0].expected = 1;
    testCases[0].got = isSameNumber(100, 212);

    testCases[1].expected = 0;
    testCases[1].got = isSameNumber(25, 25);

    for (int i = 2; i < 6; i++) {
        int a = __randomNumber(0, 9999);
        int b = __randomNumber(0, 9999);
        int expected = a == b;
        int got = isSameNumber(a, b);
        testCases[i].got = got;
        testCases[i].expected = expected;
    }

    for (int i = 6; i < 10; i++) {
        int a = __randomNumber(0, 9999);
        int expected = 0;
        int got = isSameNumber(a, a);
        testCases[i].got = got;
        testCases[i].expected = expected;
    }

    for (unsigned int i = 0; i < sizeof(testCases) / sizeof(struct TestCase); i++) {
        struct TestCase test = testCases[i];

        if (test.got == test.expected) {
            printf("# %d PASSING\n", i+1);
        } else {
            printf("# %d FAILED\n", i+1);
            printf("> EXPECTED %d\n", test.expected);
            printf("> GOT %d\n", test.got);
        }
    }
    return 0;
}
