import java.util.*;
import java.io.*;
import java.math.*;

_REPLACE_ME_WITH_DIRECTIVES_

public class question3 {

    _REPLACE_ME_WITH_SOLUTION_

    private static boolean workingAnswer(int number1, int number2) {
        return number1 == number2;
    }

    private static List<boolean[]> generateTestCase() {
        var listOfTest = new ArrayList<boolean[]>();

        listOfTest.add(new boolean[] { false, SimilarNumber.isSameNumber(100, 212) });
        listOfTest.add(new boolean[] { true, SimilarNumber.isSameNumber(25, 25) });

        return listOfTest;
    }

    private static List<boolean[]> generateRandomTestCase(int numberOfTC) {
        var listOfTest = new ArrayList<boolean[]>();

        // diffnumber
        for (int i = 0; i < numberOfTC; i++) {

            var number1 = getRandomNumber(0, 9999);
            var number2 = getRandomNumber(0, 9999);

            var expected = workingAnswer(number1, number2);
            var got = SimilarNumber.isSameNumber(number1, number2);

            listOfTest.add(new boolean[] { expected, got });

        }

        // same number
        for (int i = 0; i < numberOfTC; i++) {

            var number1 = getRandomNumber(0, 9999);

            var expected = workingAnswer(number1, number1);
            var got = SimilarNumber.isSameNumber(number1, number1);

            listOfTest.add(new boolean[] { expected, got });

        }
        return listOfTest;
    }

    private static int getRandomNumber(int min, int max) {
        return (int) Math.floor(Math.random() * (max - min + 1) + min);
    }

    public static void main(String[] args) {
        var testCase = generateTestCase();
        var randomTestCase = generateRandomTestCase(4);

        testCase.addAll(randomTestCase);

        var counter = 0;
        for (boolean[] test : testCase) {
            counter++;
            if (test[0] == test[1]) {
                System.out.printf("# %d PASSING\n", counter);
            } else {
                System.out.printf("# %d FAILED\n", counter);
                System.out.printf("> EXPECTED %b\n", test[0]);
                System.out.printf("> GOT %b\n", test[1]);
            }
        }
    }
}