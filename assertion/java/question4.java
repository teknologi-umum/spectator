import java.util.*;
import java.io.*;
import java.math.*;

public class question4 {

    _REPLACE_ME_

    private static int workingAnswer(long power) {
        int result = 0;
        while (power != 0) {
            result += power % 10;
            power /= 10;
        }
        return result;
    }

    private static List<int[]> generateTestCase() {
        List<int[]> listOfTest = new ArrayList<int[]>();

        listOfTest.add(new int[] { 19, Heater.findHeaterPower(100212373) });

        return listOfTest;
    }

    public static List<int[]> generateRandomTestCase(int numberOfTC) {
        List<int[]> listOfTest = new ArrayList<int[]>();

        for (int i = 0; i < numberOfTC; i++) {

            long number = getRandomNumber(1000000000, 9999999999L);

            int expected = workingAnswer(number);
            int got = Heater.findHeaterPower(number);

            listOfTest.add(new int[] { expected, got });

        }
        return listOfTest;
    }

    private static long getRandomNumber(int min, long max) {
        return (long) Math.floor(Math.random() * (max - min + 1) + min);
    }

    public static void main(String[] args) {
        List<int[]> testCase = generateTestCase();
        List<int[]> randomTestCase = generateRandomTestCase(9);

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
