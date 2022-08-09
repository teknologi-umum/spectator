import java.util.*;
import java.io.*;
import java.math.*;

_REPLACE_ME_WITH_DIRECTIVES_

public class question3 {
    _REPLACE_ME_WITH_SOLUTION_

    private static boolean workingAnswer(int number1, int number2) {
        return number1 == number2;
    }

    private static List<Object[]> generateTestCase() {
        var listOfTest = new ArrayList<Object[]>();
        int number1;
        int number2;

        number1 = 100;
        number2 = 212;
        listOfTest.add(new Object[] { false, SimilarNumber.isSameNumber(number1, number2), new QuestionAttribute(number1, number2) });

        number1 = 25;
        number2 = 25;
        listOfTest.add(new Object[] { true, SimilarNumber.isSameNumber(number1, number2), new QuestionAttribute(number1, number2) });

        return listOfTest;
    }

    private static List<Object[]> generateRandomTestCase(int numberOfTC) {
        var listOfTest = new ArrayList<Object[]>();

        // diffnumber
        for (int i = 0; i < numberOfTC; i++) {

            var number1 = getRandomNumber(0, 9999);
            var number2 = getRandomNumber(0, 9999);

            var expected = workingAnswer(number1, number2);
            var got = SimilarNumber.isSameNumber(number1, number2);

            listOfTest.add(new Object[] { expected, got, new QuestionAttribute(number1, number2)});
        }

        // same number
        for (int i = 0; i < numberOfTC; i++) {

            var number1 = getRandomNumber(0, 9999);

            var expected = workingAnswer(number1, number1);
            var got = SimilarNumber.isSameNumber(number1, number1);

            listOfTest.add(new Object[] { expected, got, new QuestionAttribute(number1, number1) });

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
        for (Object[] test : testCase) {
            counter++;
            if (test[0] == test[1]) {
                System.out.printf("# %d PASSING\n", counter);
            } else {
                System.out.printf("# %d FAILED\n", counter);
            }
            QuestionAttribute argument = (QuestionAttribute) test[2];
            System.out.printf("> ARGUMENTS isSameNumber(%d, %d)\n", argument.number1, argument.number2);
            System.out.printf("> EXPECTED %b\n", test[0]);
            System.out.printf("> GOT %b\n", test[1]);
        }
    }
}

class QuestionAttribute {
    public int number1;
    public int number2;

    public QuestionAttribute(int number1, int number2) {
        this.number1 = number1;
        this.number2 = number2;
    }
}