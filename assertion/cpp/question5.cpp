#include <stdio.h>
#include <stdlib.h>
#include <string>
#include <vector>
#include <numeric>
#include <time.h>

std::string mumble(std::string input);

{0}

std::string __workingAnswer(std::string input) {
    std::vector<std::string> words;

    for (unsigned int i = 0; i < input.size(); i++) {
        std::string m = "";
        for (unsigned int j = 0; j <= i; j++) {
            char c = input.at(i);
            m.insert(m.end(), j == 0 ? std::toupper(c) : std::tolower(c));
        }
        words.insert(words.end(), m);
    }

    std::string result = std::accumulate(
        words.begin(), words.end(), std::string(),
        [](const std::string& acc, const std::string& curr) -> std::string {
              return acc + (acc.length() > 0 ? "-" : "") + curr;
        });

    return result;
}

typedef struct TestCase {
    std::string expected;
    std::string got;
} TestCase;

// creates a random number between min and max
long __randomNumber(int min, long max) {
    return (rand() % (max - min + 1)) + min;
}

const std::string characters = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz";
std::string __genWords(int n) {
    std::string result;

    const int nchar = characters.size();
    for (int i = 0; i < n; i++) {
        char randomChar = characters[__randomNumber(0, nchar)];
        result.append(&characters[n]);
    }

    return result;
}

int main() {
    srand(time(0));

    std::vector<TestCase> testCases{
        {.expected = "A-Bb-Ccc-Dddd",
         .got = mumble("abcd")},
        {.expected = "R-Qq-Aaa-Eeee-Zzzzz-Tttttt-Yyyyyyy",
         .got = mumble("RqaEzTy")}};

    for (int i = 0; i < 8; i++) {
        int n = __randomNumber(4, 20);
        std::string word = __genWords(n);
        std::string expected = __workingAnswer(word);
        std::string got = mumble(word);
        testCases.insert(testCases.end(),
                         (TestCase){ .expected = expected, .got = got });
    }

    for (unsigned int i = 0; i < testCases.size(); i++) {
        TestCase test = testCases.at(i);

        if (test.got == test.expected) {
            printf("# %d PASSING\n", i + 1);
        } else {
            printf("# %d FAILED\n", i + 1);
            printf("> EXPECTED %s\n", test.expected.c_str());
            printf("> GOT %s\n", test.got.c_str());
        }
    }
    return 0;
}

