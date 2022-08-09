import java.util.*;
import java.math.BigDecimal;

_REPLACE_ME_WITH_DIRECTIVES_

public class question2 {

    private static String[] temperatures = { "Celcius", "Fahrenheit", "Kelvin" };

    _REPLACE_ME_WITH_SOLUTION_

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

    private static List<Object[]> generateTestCase() {
        var listOfTest = new ArrayList<Object[]>();

        listOfTest.add(new Object[] { 212.0, Temperature.calculateTemperature(100, "Celcius", "Fahrenheit"),
                new QuestionAttribute(100, "Celcius", "Fahrenheit") });
        listOfTest.add(new Object[] { 373.15, Temperature.calculateTemperature(212, "Fahrenheit", "Kelvin"),
                new QuestionAttribute(212, "Fahrenheit", "Kelvin") });
        listOfTest.add(new Object[] { 273.15, Temperature.calculateTemperature(0, "Celcius", "Kelvin"),
                new QuestionAttribute(0, "Celcius", "Kelvin") });
        listOfTest.add(new Object[] { 32.0, Temperature.calculateTemperature(0, "Celcius", "Fahrenheit"),
                new QuestionAttribute(0, "Celcius", "Fahrenheit") });
        listOfTest.add(new Object[] { 212.0, Temperature.calculateTemperature(373.15, "Kelvin", "Fahrenheit"),
                new QuestionAttribute(373.15, "Kelvin", "Fahrenheit") });

        return listOfTest;
    }

    private static List<Object[]> generateRandomTestCase(int numberOfTestCases) {
        var listOfTest = new ArrayList<Object[]>();

        for (int i = 0; i < numberOfTestCases; i++) {
            var from = temperatures[getRandomNumber(0, 2)];
            var to = temperatures[getRandomNumber(0, 2)];
            int randomTemperature = getRandomNumber(-500, 500);

            double expected = workingAnswer(randomTemperature, from, to);
            double got = Temperature.calculateTemperature(randomTemperature, from, to);

            listOfTest.add(new Object[] { expected, got, new QuestionAttribute(randomTemperature, from, to) });
        }

        return listOfTest;
    }

    private static int getRandomNumber(int min, int max) {
        return (int) Math.floor(Math.random() * (max - min + 1) + min);
    }

    public static void main(String[] args) {
        List<Object[]> testCase = generateTestCase();
        List<Object[]> randomTestCase = generateRandomTestCase(5);

        testCase.addAll(randomTestCase);

        int counter = 0;
        for (Object[] test : testCase) {
            BigDecimal expected = BigDecimal.valueOf((Double) test[0]);
            BigDecimal got = BigDecimal.valueOf((Double) test[1]);

            counter++;

            if (expected.compareTo(got) == 0) {
                System.out.printf("# %d PASSING\n", counter);
            } else {
                System.out.printf("# %d FAILED\n", counter);
            }
            QuestionAttribute argument = (QuestionAttribute) test[2];
            System.out.printf("> ARGUMENTS calculateTemperatures(%f, \"%s\", \"%s\")\n", argument.temp, argument.from,
                    argument.to);
            System.out.printf("> EXPECTED %f\n", expected);
            System.out.printf("> GOT %f\n", got);
        }
    }
}

class QuestionAttribute {
    public double temp;
    public String from;
    public String to;

    public QuestionAttribute(double temp, String from, String to) {
        this.temp = temp;
        this.from = from;
        this.to = to;
    }
}