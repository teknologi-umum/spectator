#include <stdio.h>
#include <stdlib.h>
#include <ctype.h>
#include <string.h>
#include <time.h>

_REPLACE_ME_WITH_DIRECTIVES_

char* mumble(char* input);

_REPLACE_ME_WITH_SOLUTION_

char* __workingAnswer(const char* input) {
    int inputLen = strlen(input);
    // `triNum` is the final length of the result. Since we want `abcd` to be
    // `a-bb-ccc-dddd`, it's basically 1 + 2 + 3 + 4 <- this notation is called
    // triangular number. We must add `inputLen - 1` because we need to also
    // count the total of the hyphens used as a separator.
    // i hate c string, they need some whacky ass malloc ritual or some shit
    // ...that, or i'm just dumb, most likely the latter
    int triNum = (inputLen * (inputLen + 1)) / 2;
    int resultLen = triNum + inputLen - 1;
    char* result = malloc(resultLen);

    int pos = 0;
    for (int i = 0; i < inputLen; i++) {
        for (int j = 0; j <= i; j++) {
            char c = input[i];
            result[pos] = j == 0 ? toupper(c) : tolower(c);
            pos++;
        }
        if (pos != resultLen) {
            result[pos] = '-';
            pos++;
        }
    }

    return result;
}

typedef struct TestCase {
    char* expected;
    char* got;
} TestCase;

// creates a random number between min and max
long __randomNumber(int min, long max) {
    return (rand() % (max - min + 1)) + min;
}

const char* characters = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz";
char* __genWords(int n) {
    char* result = malloc(n + 1); // 1 for null terminating character

    const int nchar = strlen(characters);
    for (int i = 0; i < n; i++) {
        char randomChar = characters[__randomNumber(0, nchar)];
        result[i] = randomChar;
    }

    return result;
}

int main() {
    srand(time(0));

    TestCase testCases[10] = {
        {.expected = "A-Bb-Ccc-Dddd",
         .got = mumble("abcd")},
        {.expected = "R-Qq-Aaa-Eeee-Zzzzz-Tttttt-Yyyyyyy",
         .got = mumble("RqaEzTy")}};

    for (int i = 2; i < 10; i++) {
        int n = __randomNumber(4, 20);
        char* word = __genWords(n);
        char* expected = __workingAnswer(word);
        char* got = mumble(word);
        testCases[i].expected = expected;
        testCases[i].got = got;
        free(word);
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
    }
    return 0;
}
