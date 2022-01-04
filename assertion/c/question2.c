#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <time.h>

{0}

const char* __CELCIUS = "Celcius";
const char* __FAHRENHEIT = "Fahrenheit";
const char* __KELVIN = "Kelvin";

int __workingAnswer(int n, char* a, char* b) {
     if (strcmp(a, __CELCIUS) == 0 && strcmp(b, __FAHRENHEIT) == 0) {
        return (n * 9 / 5) + 32;
      } else if (strcmp(a, __CELCIUS) == 0 && strcmp(b, __KELVIN) == 0) {
        return n + 273.15;
      } else if (strcmp(a, __FAHRENHEIT) == 0 && strcmp(b, __CELCIUS) == 0) {
        return (n - 32) * 5 / 9;
      } else if (strcmp(a, __FAHRENHEIT) == 0 && strcmp(b, __KELVIN) == 0) {
        return (n - 32) * 5 / 9 + 273.15;
      } else if (strcmp(a, __KELVIN) == 0 && strcmp(b, __CELCIUS) == 0) {
        return n - 273.15;
      } else if (strcmp(a, __KELVIN) == 0 && strcmp(b, __FAHRENHEIT) == 0) {
        return (n - 273.15) * 9 / 5 + 32;
      } else {
        return n;
      }
}

struct TestCase {
    int expected;
    int got;
};

// creates a random number between min and max
int __randomNumber(int min, int max) {
    return (rand() % (max - min + 1)) + min;
}

int main() {
    srand(time(0));

    char temperatures[3][10];
    sprintf(temperatures[0], "%s", "Celcius");
    sprintf(temperatures[1], "%s", "Fahrenheit");
    sprintf(temperatures[2], "%s", "Kelvin");

    struct TestCase testCases[10];
    testCases[0].expected = 212;
    testCases[0].got = calculateTemperature(100, "Celcius", "Fahrenheit");

    testCases[1].expected = 373;
    testCases[1].got = calculateTemperature(212, "Fahrenheit", "Kelvin");

    testCases[2].expected = 273;
    testCases[2].got = calculateTemperature(0, "Celcius", "Kelvin");

    testCases[3].expected = 32;
    testCases[3].got = calculateTemperature(0, "Celcius", "Fahrenheit");

    testCases[4].expected = -459;
    testCases[4].got = calculateTemperature(0, "Kelvin", "Fahrenheit");

    for (unsigned int i = 5; i < sizeof(testCases) / sizeof(struct TestCase); i++) {
        // Generate random test cases
        int randNum = __randomNumber(-500, 500);
        char from[10];
        strncpy(from, temperatures[__randomNumber(0, 2)], 10);
        char to[10];
        strncpy(to, temperatures[__randomNumber(0, 2)], 10);
        int expected = __workingAnswer(randNum, from, to);
        int got = calculateTemperature(randNum, from, to);
        testCases[i].expected = expected;
        testCases[i].got = got;
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
