#include <stdio.h>
#include <string>
#include <vector>
#include <numeric>
// `mumble` is a function that accepts an argument as string and returns a
// string as its output.
std::string mumble(std::string input)
{
    // write your code here
    std::vector<std::string> words;

    for (unsigned int i = 0; i < input.size(); i++)
    {
        std::string m = "";
        for (unsigned int j = 0; j <= i; j++)
        {
            char c = input.at(i);
            m.push_back(j == 0 ? std::toupper(c) : std::tolower(c));
        }
        words.push_back(m);
    }

    std::string result = std::accumulate(
        words.begin(), words.end(), std::string(),
        [](const std::string &acc, const std::string &curr) -> std::string
        {
            return acc + (acc.length() > 0 ? "-" : "") + curr;
        });

    return result;
}