class Temperature {
    // test
    // test
    public static double calculateTemperature(double temp, String from, String to) {
        if (isC(from) && isF(to))
            return (temp * 9 / 5) + 32;
        if (isC(from) && isK(to))
            return (temp + 273.15);
        if (isF(from) && isC(to))
            return (temp - 32) * 5 / 9;
        if (isF(from) && isK(to))
            return ((temp - 32) * 5 / 9 + 273.15);
        if (isK(from) && isC(to))
            return (temp - 273.15);
        if (isK(from) && isF(to))
            return ((temp - 273.15) * 9 / 5 + 32);
        return 0;
    }

    private static boolean isC(String tempName) {
        return tempName.equalsIgnoreCase("Celcius");
    }

    private static boolean isF(String tempName) {
        return tempName.equalsIgnoreCase("Fahrenheit");
    }

    private static boolean isK(String tempName) {
        return tempName.equalsIgnoreCase("Kelvin");
    }
}
