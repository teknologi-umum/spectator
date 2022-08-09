#include <stdio.h>
#include <strings.h>

// `mumble` is a function that accepts an argument as string and returns a
// string as its output.
char *mumble(char *input)
{
    int inputLen = strlen(input);
    int triNum = (inputLen * (inputLen + 1)) / 2;
    int resultLen = triNum + inputLen - 1;
    char *result = malloc(resultLen);

    int pos = 0;
    for (int i = 0; i < inputLen; i++)
    {
        for (int j = 0; j <= i; j++)
        {
            char c = input[i];
            result[pos] = j == 0 ? toupper(c) : tolower(c);
            pos++;
        }
        if (pos != resultLen)
        {
            result[pos] = '-';
            pos++;
        }
    }

    return result;
}