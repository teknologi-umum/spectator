#include <stdio.h>
#include <string>

// `calculateTemperature` is a function that accepts 3 arguments as its input:
// `temp` as integer, `from` as string, `to` as string. It returns a float
// as its output.
float calculateTemperature(int temp, std::string from, std::string to)
{
    if (from.compare("Celcius") == 0 && to.compare("Fahrenheit") == 0)
    {
        return (temp * 9 / 5.0f) + 32.0f;
    }
    else if (from.compare("Celcius") == 0 && to.compare("Kelvin") == 0)
    {
        return temp + 273.15f;
    }
    else if (from.compare("Fahrenheit") == 0 && to.compare("Celcius") == 0)
    {
        return (temp - 32) * 5 / 9.0f;
    }
    else if (from.compare("Fahrenheit") == 0 && to.compare("Kelvin") == 0)
    {
        return (temp - 32) * 5 / 9.0f + 273.15f;
    }
    else if (from.compare("Kelvin") == 0 && to.compare("Celcius") == 0)
    {
        return temp - 273.15f;
    }
    else if (from.compare("Kelvin") == 0 && to.compare("Fahrenheit") == 0)
    {
        return (temp - 273.15) * 9 / 5.0f + 32.0f;
    }
    else
    {
        return temp;
    }
}