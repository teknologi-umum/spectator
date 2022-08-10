#include <stdio.h>
#include <strings.h>

// `calculateTemperature` is a function that accepts 3 arguments as its input:
// `temp` as integer, `from` as string, `to` as string. It returns a float
// as its output.
float calculateTemperature(int temp, char *from, char *to)
{
    if (strcmp(from, "Celcius") == 0 && strcmp(to, "Fahrenheit") == 0)
        return (temp * 9 / 5.0f) + 32.0f;
    if (strcmp(from, "Celcius") == 0 && strcmp(to, "Kelvin") == 0)
        return temp + 273.15f;
    if (strcmp(from, "Fahrenheit") == 0 && strcmp(to, "Celcius") == 0)
        return (temp - 32) * 5 / 9.0f;
    if (strcmp(from, "Fahrenheit") == 0 && strcmp(to, "Kelvin") == 0)
        return (temp - 32) * 5 / 9.0f + 273.15f;
    if (strcmp(from, "Kelvin") == 0 && strcmp(to, "Celcius") == 0)
        return temp - 273.15f;
    if (strcmp(from, "Kelvin") == 0 && strcmp(to, "Fahrenheit") == 0)
        return (temp - 273.15) * 9 / 5.0f + 32.0f;

    return temp;
}
