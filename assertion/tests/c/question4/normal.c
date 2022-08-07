#include <stdio.h>

// `findHeaterPower` is function that accepts an argument as long and
// returns an integer as its output.
int findHeaterPower(long input)
{
    int result = 0;
    while (input != 0)
    {
        result += input % 10;
        input /= 10;
    }

    return result;
}