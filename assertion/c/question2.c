#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <time.h>
#include <math.h>
#include <float.h>

_REPLACE_ME_WITH_DIRECTIVES_

float calculateTemperature(int temp, char *from, char *to);

_REPLACE_ME_WITH_SOLUTION_

const char *__CELCIUS = "Celcius";
const char *__FAHRENHEIT = "Fahrenheit";
const char *__KELVIN = "Kelvin";

int isC(const char *unit) { return strcmp(unit, __CELCIUS) == 0; }
int isF(const char *unit) { return strcmp(unit, __FAHRENHEIT) == 0; }
int isK(const char *unit) { return strcmp(unit, __KELVIN) == 0; }

float __workingAnswer(int n, const char *a, const char *b)
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
    char *arguments;
} TestCase;

// creates a random number between min and max
int __randomNumber(int min, int max)
{
    return (rand() % (max - min + 1)) + min;
}

int main()
{
    srand(time(0));

    char temperatures[3][12] = {"Celcius", "Fahrenheit", "Kelvin"};

    TestCase testCases[10] = {
        {.expected = 212,
         .got = calculateTemperature(100, "Celcius", "Fahrenheit"),
         .arguments = "calculateTemperature(100, \"Celcius\", \"Fahrenheit\")"},
        {.expected = 373.15,
         .got = calculateTemperature(212, "Fahrenheit", "Kelvin"),
         .arguments = "calculateTemperature(212, \"Fahrenheit\", \"Kelvin\")"},
        {.expected = 273.15,
         .got = calculateTemperature(0, "Celcius", "Kelvin"),
         .arguments = "calculateTemperature(0, \"Celcius\", \"Kelvin\")"},
        {.expected = 32,
         .got = calculateTemperature(0, "Celcius", "Fahrenheit"),
         .arguments = "calculateTemperature(0, \"Celcius\", \"Fahrenheit\")"},
        {.expected = -459.67,
         .got = calculateTemperature(0, "Kelvin", "Fahrenheit"),
         .arguments = "calculateTemperature(0, \"Kelvin\", \"Fahrenheit\")"}};

    // `sizeof` returns the size of the memory used, not the length of the
    // array so we need to divide it by the size of the struct
    for (unsigned int i = 5; i < sizeof(testCases) / sizeof(TestCase); i++)
    {
        TestCase *test = &testCases[i];

        // Generate random test cases
        int randNum = __randomNumber(-500, 500);

        char *from = temperatures[__randomNumber(0, 2)];
        char *to = temperatures[__randomNumber(0, 2)];
        float expected = __workingAnswer(randNum, from, to);
        float got = calculateTemperature(randNum, from, to);

        test->expected = expected;
        test->got = got;
        test->arguments = malloc(sizeof(char) * 100);
        sprintf(test->arguments, "calculateTemperature(%d, \"%s\", \"%s\")", randNum, from, to);
    }

    for (unsigned int i = 0; i < sizeof(testCases) / sizeof(TestCase); i++)
    {
        TestCase *test = testCases + i;

        if (fabs(test->got - test->expected) < FLT_EPSILON)
        {
            printf("# %d PASSING\n", i + 1);
        }
        else
        {
            printf("# %d FAILED\n", i + 1);
        }

        printf("> ARGUMENTS %s\n", test->arguments);
        printf("> EXPECTED %.2f\n", test->expected);
        printf("> GOT %.2f\n", test->got);

        if (i >= 5)
        {
            free(test->arguments);
        }
    }

    return 0;
}
