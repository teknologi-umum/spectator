#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <time.h>

_REPLACE_ME_WITH_DIRECTIVES_

int* calculateGrade(int size, int* input);

_REPLACE_ME_WITH_SOLUTION_

int* __workingAnswer(int size, int* input) {
    for (int i = 0; i < size; i++) {
        int grade = input[i];
        if (grade >= 38 && grade % 5 != 0 && grade % 5 >= 3) {
            input[i] = grade + (5 - grade % 5);
        } else {
            input[i] = grade;
        }
    }

    return input;
}

typedef struct TestCase {
    char* expected;
    char* got;
} TestCase;

// creates a random number between min and max
long __randomNumber(int min, long max) {
    return (rand() % (max - min + 1)) + min;
}

int* genArray(int len) {
    int* result = malloc(len * sizeof(int));
    for (int i = 0; i < len; i++) {
        result[i] = __randomNumber(10, 100);
    }
    return result;
}

// this is poorman's implementation
char* arrayToString(int len, int* arr) {
    char* result = malloc(255 * sizeof(char));

    sprintf(&result[0], "%c", '[');
    for (int i = 0; i < len - 1; i++) {
        sprintf(&result[strlen(result)], "%d, ", arr[i]);
    }
    sprintf(&result[strlen(result)], "%d%c", arr[len], ']');

    return result;
}

int main() {
    srand(time(0));

    TestCase testCases[10];

    int answer[] = {75, 67, 40, 33};

    int* input = malloc(4 * sizeof(int));
    input[0] = 73;
    input[1] = 67;
    input[2] = 38;
    input[3] = 33;

    testCases[0].expected = arrayToString(4, answer);
    testCases[0].got = arrayToString(4, calculateGrade(4, input));

    for (int i = 1; i < 10; i++) {
        int len = __randomNumber(4, 20);

        int* input = genArray(len);
        // if I don't add +1 here it'll lead to UB, idk why but this thing works
        int* input2 = malloc(len * sizeof(int) + 1);
        memcpy(input2, input, len * sizeof(int) + 1);

        int* expected = __workingAnswer(len, input);
        int* got = calculateGrade(len, input2);

        testCases[i].expected = arrayToString(len, expected);
        testCases[i].got = arrayToString(len, got);
    }

    for (unsigned int i = 0, len = sizeof(testCases) / sizeof(TestCase); i < len; i++) {
        TestCase test = testCases[i];

        if (strcmp(test.got, test.expected) == 0) {
            printf("# %d PASSING\n", i + 1);
        } else {
            printf("# %d FAILED\n", i + 1);
            printf("> EXPECTED %s\n", test.expected);
            printf("> GOT %s\n", test.got);
        }

        free(test.got);
        free(test.expected);
    }
    return 0;
}
