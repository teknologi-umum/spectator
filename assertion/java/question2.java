import java.util.*;
import java.math.BigDecimal;

// _REPLACE_ME_WITH_DIRECTIVES_

public class question2 {

    private static String[] temperatures = { "Celcius", "Fahrenheit", "Kelvin" };

    // _REPLACE_ME_WITH_SOLUTION_
    private static class Temperature {
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

    private static double workingAnswer(double temp, String from, String to) {
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

    private static List<double[]> generateTestCase() {
        var listOfTest = new ArrayList<double[]>();

        listOfTest.add(new double[] { 212, Temperature.calculateTemperature(100, "Celcius", "Fahrenheit") });
        listOfTest.add(new double[] { 373.15, Temperature.calculateTemperature(212, "Fahrenheit", "Kelvin") });
        listOfTest.add(new double[] { 273.15, Temperature.calculateTemperature(0, "Celcius", "Kelvin") });
        listOfTest.add(new double[] { 32, Temperature.calculateTemperature(0, "Celcius", "Fahrenheit") });
        listOfTest.add(new double[] { -459.67, Temperature.calculateTemperature(0, "Kelvin", "Fahrenheit") });

        return listOfTest;
    }

    private static List<double[]> generateRandomTestCase(int numberOfTestCases) {
        var listOfTest = new ArrayList<double[]>();

        for (int i = 0; i < numberOfTestCases; i++) {
            var from = temperatures[getRandomNumber(0, 2)];
            var to = temperatures[getRandomNumber(0, 2)];
            double randomTemperature = getRandomNumber(-500, 500);

            double expected = workingAnswer(randomTemperature, from, to);
            double got = Temperature.calculateTemperature(randomTemperature, from, to);

            listOfTest.add(new double[] { expected, got });
        }

        return listOfTest;
    }

    private static int getRandomNumber(int min, int max) {
        return (int) Math.floor(Math.random() * (max - min + 1) + min);
    }

    public static void main(String[] args) {
        List<double[]> testCase = generateTestCase();
        List<double[]> randomTestCase = generateRandomTestCase(5);

        testCase.addAll(randomTestCase);

        int counter = 0;
        for (double[] test : testCase) {
            BigDecimal expected = BigDecimal.valueOf(test[0]);
            BigDecimal got = BigDecimal.valueOf(test[1]);

            counter++;
            if (expected.compareTo(got) == 0) {
                System.out.printf("# %d PASSING\n", counter);
            } else {
                System.out.printf("# %d FAILED\n", counter);
                System.out.printf("> EXPECTED %f\n", expected);
                System.out.printf("> GOT %f\n", got);
            }
        }
    }
}
