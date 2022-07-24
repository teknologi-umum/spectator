# `calculateTemperature` is a function that accepts 3 arguments as its input:
# `temp` as integer, `from` as string, `to` as string. It returns a float as
# its output.
def calculateTemperature(temp, From, To):
    celcius = "Celcius"
    fahrenheit = "Fahrenheit"
    kelvin = "Kelvin"

    if From == celcius and To == fahrenheit:
        return str((temp * 9 / 5) + 32)
    elif From == celcius and To == kelvin:
        return str(temp + 273.15)
    elif From == fahrenheit and To == celcius:
        return str((temp - 32) * 5 / 9)
    elif From == fahrenheit and To == kelvin:
        return str((temp - 32) * 5 / 9 + 273.15)
    elif From == kelvin and To == celcius:
        return str(temp - 273.15)
    elif From == kelvin and To == fahrenheit:
        return str((temp - 273.15) * 9 / 5 + 32)
    else:
        return str(temp)