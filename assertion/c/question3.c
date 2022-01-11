#include <stdio.h>
#include <stdlib.h>
#include <time.h>

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
    srand(time(0));

    TestCase testCases[10] = {
        {.expected = 0,
         .got = isSameNumber(100, 212)},
        {.expected = 1,
         .got = isSameNumber(25, 25)}};;

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
        int expected = 1;
        int got = isSameNumber(a, a);
        testCases[i].got = got;
        testCases[i].expected = expected;
    }

    // `sizeof` returns the size of the memory used, not the length of the
    // array so we need to divide it by the size of the struct
    for (unsigned int i = 0; i < sizeof(testCases) / sizeof(TestCase); i++) {
        TestCase test = testCases[i];

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
