#include <stdio.h>
#include <stdlib.h>
#include <string>
#include <vector>

std::vector<int> calculateGrade(std::vector<int> input);

{0}

std::vector<int> __workingAnswer(std::vector<int> input) {
    std::vector<int> result;

    for (auto& grade : input) {
        if (grade >= 38 && grade % 5 != 0 && grade % 5 >= 3) {
            result.insert(result.end(), grade + (5 - grade % 5));
        } else {
            result.insert(result.end(), grade);
        }
    }

    return result;
}

typedef struct TestCase {
    std::vector<int> expected;
    std::vector<int> got;
} TestCase;

// creates a random number between min and max
long __randomNumber(int min, long max) {
    return (rand() % (max - min + 1)) + min;
}

std::vector<int> genVector(int n) {
    std::vector<int> result;
    for (int i = 0; i < n; i++) {
        result.insert(result.end(), __randomNumber(0, 100));
    }
    return result;
}

int main() {
    std::vector<TestCase> testCases{
        {.expected = {75, 67, 40, 33},
         .got = calculateGrade({75, 67, 40, 33})}};

    for (int i = 0; i < 8; i++) {
        int len = __randomNumber(4, 20);
        std::vector<int> input = genVector(len);
        std::vector<int> expected = __workingAnswer(input);
        std::vector<int> got = calculateGrade(input);
        testCases.insert(testCases.end(),
                         (TestCase){ .expected = expected, .got = got });
    }

    for (unsigned int i = 0; i < testCases.size(); i++) {
        TestCase test = testCases.at(i);

        if (test.got == test.expected) {
            printf("# %d PASSING\n", i + 1);
        } else {
            std::string expected(test.expected.begin(), test.expected.end());
            std::string got(test.got.begin(), test.got.end());
            printf("# %d FAILED\n", i + 1);
            printf("> EXPECTED %s\n", expected.c_str());
            printf("> GOT %s\n", got.c_str());
        }
    }
    return 0;
}
