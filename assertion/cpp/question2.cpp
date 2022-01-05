#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <string>
#include <time.h>
#include <vector>

int calculateTemperature(int temp, std::string from, std::string to);

{0}

const std::string __CELCIUS = "Celcius";
const std::string __FAHRENHEIT = "Fahrenheit";
const std::string __KELVIN = "Kelvin";

bool isC(std::string unit) { return unit.compare(__CELCIUS) == 0; }
bool isF(std::string unit) { return unit.compare(__FAHRENHEIT) == 0; }
bool isK(std::string unit) { return unit.compare(__KELVIN) == 0; }

int __workingAnswer(int n, std::string a, std::string b) {
    if (isC(a) && isF(b)) return (n * 9 / 5) + 32;
    if (isC(a) && isK(b)) return n + 273.15;
    if (isF(a) && isC(b)) return (n - 32) * 5 / 9;
    if (isF(a) && isK(b)) return (n - 32) * 5 / 9 + 273.15;
    if (isK(a) && isC(b)) return n - 273.15;
    if (isK(a) && isF(b)) return (n - 273.15) * 9 / 5 + 32;
    return n;
}

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

    std::vector<std::string> temperatures{"Celcius", "Fahrenheit", "Kelvin"};

    std::vector<TestCase> testCases{
        {.expected = 212,
         .got = calculateTemperature(100, "Celcius", "Fahrenheit")},
        {.expected = 373,
         .got = calculateTemperature(212, "Fahrenheit", "Kelvin")},
        {.expected = 273,
         .got = calculateTemperature(0, "Celcius", "Kelvin")},
        {.expected = 32,
         .got = calculateTemperature(0, "Celcius", "Fahrenheit")},
        {.expected = -459,
         .got = calculateTemperature(0, "Kelvin", "Fahrenheit")}};

    for (unsigned int i = 0; i < 5; i++) {
        // Generate random test cases
        int randNum = __randomNumber(-500, 500);
        std::string from = temperatures.at(__randomNumber(0, 2));
        std::string to = temperatures.at(__randomNumber(0, 2));
        int expected = __workingAnswer(randNum, from, to);
        int got = calculateTemperature(randNum, from, to);

        testCases.insert(testCases.end(),
                         (TestCase){.expected = expected, .got = got});
    }

    for (unsigned int i = 0; i < testCases.size(); i++) {
        TestCase test = testCases.at(i);

        if (test.got == test.expected) {
            printf("# %d PASSING\n", i + 1);
        } else {
            printf("# %d FAILED\n", i + 1);
            printf("> EXPECTED %d\n", test.expected);
            printf("> GOT %d\n", test.got);
        }
    }

    return 0;
}
