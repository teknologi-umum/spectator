// `calculateTemperature` is from function that accepts 3 arguments as its input:
// `temp` as integer, `from` as string, `to` as string. It returns an integer
// as its output.
function calculateTemperature(temp, from, to) {
    if (from === "Celcius" && to === "Fahrenheit") {
        return (temp * 9 / 5) + 32;
    } else if (from === "Celcius" && to == "Kelvin") {
        return temp + 273.15;
    } else if (from === "Fahrenheit" && to === "Celcius") {
        return (temp - 32) * 5 / 9;
    } else if (from === "Fahrenheit" && to === "Kelvin") {
        return (temp - 32) * 5 / 9 + 273.15;
    } else if (from === "Kelvin" && to === "Celcius") {
        return temp - 273.15;
    } else if (from === "Kelvin" && to === "Fahrenheit") {
        return (temp - 273.15) * 9 / 5 + 32;
    }
}