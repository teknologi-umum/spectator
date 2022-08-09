import java.util.*;
import java.io.*;
import java.math.*;

_REPLACE_ME_WITH_DIRECTIVES_

public class question4 {

    _REPLACE_ME_WITH_SOLUTION_

    private static int workingAnswer(long power) {
        var result = 0;
        while (power != 0) {
            result += power % 10;
            power /= 10;
        }
        return result;
    }

    private static List<Object[]> generateTestCase() {
        var listOfTest = new ArrayList<Object[]>();

        long power = 100212373;
        listOfTest.add(new Object[] { 19, Heater.findHeaterPower(power), new QuestionAttribute(power) });

        return listOfTest;
    }

    private static List<Object[]> generateRandomTestCase(int numberOfTC) {
        var listOfTest = new ArrayList<Object[]>();

        for (var i = 0; i < numberOfTC; i++) {

            var number = getRandomNumber(1000000000, 9999999999L);

            var expected = workingAnswer(number);
            var got = Heater.findHeaterPower(number);

            listOfTest.add(new Object[] { expected, got, new QuestionAttribute(number) });

        }
        return listOfTest;
    }

    private static long getRandomNumber(int min, long max) {
        return (long) Math.floor(Math.random() * (max - min + 1) + min);
    }

    public static void main(String[] args) {
        var testCase = generateTestCase();
        var randomTestCase = generateRandomTestCase(9);

        testCase.addAll(randomTestCase);

        var counter = 0;
        for (Object[] test : testCase) {
            counter++;
            if (test[0] == test[1]) {
                System.out.printf("# %d PASSING\n", counter);
            } else {
                System.out.printf("# %d FAILED\n", counter);
            }
            QuestionAttribute argument = (QuestionAttribute) test[2];
            System.out.printf("> ARGUMENTS findHeaterPower(%d)\n", argument.power);
            System.out.printf("> EXPECTED %d\n", test[0]);
            System.out.printf("> GOT %d\n", test[1]);
        }
    }
}

class QuestionAttribute {
    public long power;

    public QuestionAttribute(long power) {
        this.power = power;
    }
}
