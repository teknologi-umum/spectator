class Heater {
    public static int findHeaterPower(long power) {
        var result = 0;
        while (power != 0) {
            result += power % 10;
            power /= 10;
        }
        return result;
    }
}
