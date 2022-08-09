#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <time.h>

_REPLACE_ME_WITH_DIRECTIVES_

int *calculateGrade(int size, int *input);

_REPLACE_ME_WITH_SOLUTION_

int *__workingAnswer(int size, int *input)
{
    for (int i = 0; i < size; i++)
    {
        int grade = input[i];
        if (grade >= 38 && grade % 5 != 0 && grade % 5 >= 3)
        {
            input[i] = grade + (5 - grade % 5);
        }
        else
        {
            input[i] = grade;
        }
    }

    return input;
}

typedef struct TestCase
{
    char *expected;
    char *got;
    char *arguments;
} TestCase;

// creates a random number between min and max
int __randomNumber(int min, int max)
{
    return (rand() % (max - min + 1)) + min;
}

int *genArray(int len)
{
    int *result = malloc(len * sizeof(int));
    for (int i = 0; i < len; i++)
    {
        result[i] = __randomNumber(10, 100);
    }
    return result;
}

// this is poorman's implementation
char *arrayToString(int len, int *arr)
{
    int index = 0;
    char *result = malloc(255 * sizeof(char));

    index += sprintf(result, "%c", '[');
    for (int i = 0; i < len - 1; i++)
    {
        index += sprintf(result + index, "%d, ", arr[i]);
    }
    sprintf(result + index, "%d%c", arr[len - 1], ']');

    return result;
}

int main()
{
    srand(time(0));

    TestCase testCases[10];

    int input[] = {73, 67, 38, 33};
    int answer[] = {75, 67, 40, 33};

    testCases[0].expected = arrayToString(4, answer);
    testCases[0].got = arrayToString(4, calculateGrade(4, input));
    testCases[0].arguments = "calculateGrade(4, [73, 67, 38, 33])";

    for (int i = 1; i < 10; i++)
    {
        TestCase *test = testCases + i;

        int len = __randomNumber(4, 20);

        int *input = genArray(len);
        char *args = arrayToString(len, input);
        int *input2 = malloc(len * sizeof(int));
        memcpy(input2, input, len * sizeof(int));

        int *expected = __workingAnswer(len, input);
        int *got = calculateGrade(len, input2);
        test->expected = arrayToString(len, expected);
        test->got = arrayToString(len, got);
        test->arguments = malloc(sizeof(char[250]));
        sprintf(test->arguments, "calculateGrade(%d, %s)", len, args);

        free(input);
        free(input2);
        free(args);
    }

    for (unsigned int i = 0; i < sizeof(testCases) / sizeof(TestCase); i++)
    {
        TestCase *test = testCases + i;

        if (strcmp(test->got, test->expected) == 0)
        {
            printf("# %d PASSING\n", i + 1);
        }
        else
        {
            printf("# %d FAILED\n", i + 1);
        }

        printf("> ARGUMENTS %s\n", test->arguments);
        printf("> EXPECTED %s\n", test->expected);
        printf("> GOT %s\n", test->got);

        free(test->expected);
        free(test->got);
        if (i >= 1)
        {
            free(test->arguments);
        }
    }
    return 0;
}
