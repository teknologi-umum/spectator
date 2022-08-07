#include <stdio.h>

// `calculateGrade` is a function that accepts 2 arguments as its input: `size` as integer
// and `grade` as an array of integer. It returns an array of integer as its
// output
int *calculateGrade(int size, int *grade)
{
    for (int i = 0; i < size; i++)
    {
        int g = grade[i];
        if (g >= 38 && g % 5 != 0 && g % 5 >= 3)
        {
            grade[i] = g + (5 - g % 5);
        }
        else
        {
            grade[i] = g;
        }
    }

    return grade;
}