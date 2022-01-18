import java.util.*;
import java.io.*;
import java.math.*;

public class question2 {

    private static String[] temperatures = { "Celcius", "Fahrenheit", "Kelvin" };

    _REPLACE_ME_

    private static boolean isC(String tempName) {
        return tempName.equalsIgnoreCase("Celcius");
    }

    private static boolean isF(String tempName) {
        return tempName.equalsIgnoreCase("Fahrenheit");
    }

    private static boolean isK(String tempName) {
        return tempName.equalsIgnoreCase("Kelvin");
    }

    private static int workingAnswer(int temp, String from, String to) {
        if (isC(from) && isF(to))
            return (temp * 9 / 5) + 32;
        if (isC(from) && isK(to))
            return (int) (temp + 273.15);
        if (isF(from) && isC(to))
            return (temp - 32) * 5 / 9;
        if (isF(from) && isK(to))
            return (int) ((temp - 32) * 5 / 9 + 273.15);
        if (isK(from) && isC(to))
            return (int) (temp - 273.15);
        if (isK(from) && isF(to))
            return (int) ((temp - 273.15) * 9 / 5 + 32);
        return 0;
    }

    private static List<int[]> generateTestCase() {
        List<int[]> listOfTest = new ArrayList<int[]>();

        listOfTest.add(new int[] { 212, Temperature.calculateTemperature(100, "Celcius", "Fahrenheit") });
        listOfTest.add(new int[] { 373, Temperature.calculateTemperature(212, "Fahrenheit", "Kelvin") });
        listOfTest.add(new int[] { 273, Temperature.calculateTemperature(0, "Celcius", "Kelvin") });
        listOfTest.add(new int[] { 32, Temperature.calculateTemperature(0, "Celcius", "Fahrenheit") });
        listOfTest.add(new int[] { -459, Temperature.calculateTemperature(0, "Kelvin", "Fahrenheit") });

        return listOfTest;
    }

    public static List<int[]> generateRandomTestCase(int numberOfTC) {
        List<int[]> listOfTest = new ArrayList<int[]>();

        for (int i = 0; i < numberOfTC; i++) {
            String from = temperatures[getRandomNumber(0, 2)];
            String to = temperatures[getRandomNumber(0, 2)];
            int randTemp = getRandomNumber(-500, 500);

            int expected = workingAnswer(randTemp, from, to);
            int got = Temperature.calculateTemperature(randTemp, from, to);

            listOfTest.add(new int[] { expected, got });

        }
        return listOfTest;
    }

    private static int getRandomNumber(int min, int max) {
        return (int) Math.floor(Math.random() * (max - min + 1) + min);
    }

    public static void main(String[] args) {
        List<int[]> testCase = generateTestCase();
        List<int[]> randomTestCase = generateRandomTestCase(5);

        testCase.addAll(randomTestCase);

        int counter = 0;
        for (int[] test : testCase) {
            counter++;
            if (test[0] == test[1]) {
                System.out.printf("# %d PASSING\n", counter);
            } else {
                System.out.printf("# %d FAILED\n", counter);
                System.out.printf("> EXPECTED %d\n", test[0]);
                System.out.printf("> GOT %d\n", test[1]);
            }
        }
    }
}
