// WARNING!!!!!!!!!!!
// this thing does not work

#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <time.h>

int* calculateGrade(int size, int* input);

{0}

int* calculateGrade(int size, int* input) {
    int* result = malloc(size * sizeof(int));

    for (int i = 0; i < size; i++) {
        int grade = input[i];
        if (grade >= 38 && grade % 5 != 0 && grade % 5 >= 3) {
            result[i] = grade + (5 - grade % 5);
        } else {
            result[i] = grade;
        }
    }

    return result;
}

int* __workingAnswer(int size, int* input) {
    int* result = malloc(size * sizeof(int));

    for (int i = 0; i < size; i++) {
        int grade = input[i];
        if (grade >= 38 && grade % 5 != 0 && grade % 5 >= 3) {
            result[i] = grade + (5 - grade % 5);
        } else {
            result[i] = grade;
        }
    }

    return result;
}

typedef struct TestCase {
    int* expected;
    int* got;
} TestCase;

// creates a random number between min and max
long __randomNumber(int min, long max) {
    return (rand() % (max - min + 1)) + min;
}

int* genArray(int len) {
    int* result = malloc(len * sizeof(int));
    for (int i = 0; i < len; i++) {
        result[i] = __randomNumber(0, 100);
    }
    return result;
}

// this is poorman's implementation
char* arrayToString(int* arr) {
    char* result = malloc(sizeof(arr));

    for (int i = 0, len = sizeof(arr); i < len; i++) {
        result[i] = arr[i];
    }

    return result;
}

int main() {
    srand(time(0));
    TestCase testCases[10];

    int answer[] = {75, 67, 40, 33};
    testCases[0].expected = answer;
    testCases[0].got = calculateGrade(4, answer);

    for (int i = 1; i < 10; i++) {
        int len = __randomNumber(4, 20);
        int* input = genArray(len);
        int* expected = __workingAnswer(len, input);
        int* got = calculateGrade(len, input);
        testCases[i].expected = expected;
        testCases[i].got = got;
    }

    for (unsigned int i = 0, len = sizeof(testCases) / sizeof(TestCase); i < len; i++) {
        TestCase test = testCases[i];

        if (memcmp(test.got, test.expected, sizeof(&test.got))) {
            printf("# %d PASSING\n", i + 1);
        } else {
            // std::string expected(test.expected.begin(), test.expected.end());
            // std::string got(test.got.begin(), test.got.end());
            printf("# %d FAILED\n", i + 1);
            printf("> EXPECTED %s\n", arrayToString(test.expected));
            printf("> GOT %s\n", arrayToString(test.got));
        }
    }
    return 0;
}
