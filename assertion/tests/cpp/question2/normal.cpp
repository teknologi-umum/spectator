#include <stdio.h>
#include <string>

// `calculateTemperature` is a function that accepts 3 arguments as its input:
// `temp` as float, `from` as string, `to` as string. It returns a float
// as its output.
float calculateTemperature(float temp, std::string from, std::string to)
{
    if (from.compare("Celcius") == 0 && to.compare("Fahrenheit") == 0)
    {
        return (temp * 9 / 5.0) + 32.0;
    }
    else if (from.compare("Celcius") == 0 && to.compare("Kelvin") == 0)
    {
        return temp + 273.15;
    }
    else if (from.compare("Fahrenheit") == 0 && to.compare("Celcius") == 0)
    {
        return (temp - 32) * 5 / 9.0;
    }
    else if (from.compare("Fahrenheit") == 0 && to.compare("Kelvin") == 0)
    {
        return (temp - 32) * 5 / 9.0 + 273.15;
    }
    else if (from.compare("Kelvin") == 0 && to.compare("Celcius") == 0)
    {
        return temp - 273.15;
    }
    else if (from.compare("Kelvin") == 0 && to.compare("Fahrenheit") == 0)
    {
        return (temp - 273.15) * 9 / 5.0 + 32.0;
    }
    else
    {
        return temp;
    }
}
