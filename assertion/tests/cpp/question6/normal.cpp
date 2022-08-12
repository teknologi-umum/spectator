#include <stdio.h>
#include <vector>

// `calculateGrade` is a function that accepts a vector of integer as its input
// and returns a vector of integer as its output
std::vector<int> calculateGrade(std::vector<int> grade)
{
    // write your code here
    std::vector<int> result;

    for (auto &g : grade)
    {
        if (g >= 38 && g % 5 != 0 && g % 5 >= 3)
        {
            result.push_back(g + (5 - g % 5));
        }
        else
        {
            result.push_back(g);
        }
    }

    return result;
}