import java.util.*;
import java.io.*;
import java.math.*;

_REPLACE_ME_WITH_DIRECTIVES_

public class question5 {

    _REPLACE_ME_WITH_SOLUTION_
  
    private static String workingAnswer(String input) {
        var inputLen = input.length();

        var triNum = (inputLen * (inputLen + 1)) / 2;
        var resultLen = triNum + inputLen - 1;

        var sb = new StringBuilder();
        sb.setLength(resultLen);

        var pos = 0;
        for (int i = 0; i < inputLen; i++) {
            for (int j = 0; j <= i; j++) {
                var c = input.charAt(i);

                var appended = j == 0 ? Character.toUpperCase(c) : Character.toLowerCase(c);
                sb.setCharAt(pos, appended);
                pos++;
            }
            if (pos != resultLen) {
                sb.setCharAt(pos, '-');
                pos++;
            }
        }
        return sb.toString();
    }

    private static List<Object[]> generateTestCase() {
        var listOfTest = new ArrayList<Object[]>();

        listOfTest.add(new Object[] { "A-Bb-Ccc-Dddd", Mumble.mumble("abcd"), new QuestionAttribute("abcd") });
        listOfTest.add(new Object[] { "R-Qq-Aaa-Eeee-Zzzzz-Tttttt-Yyyyyyy", Mumble.mumble("RqaEzTy"), new QuestionAttribute("RqaEzTy") });

        return listOfTest;
    }

    private static int getRandomNumber(int min, int max) {
        return (int) Math.floor(Math.random() * (max - min + 1) + min);
    }

    private static String genWords(int n) {
        var characters = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz";
        var sb = new StringBuilder();

        sb.setLength(n);
        final var nchar = characters.length();
        for (int i = 0; i < n; i++) {
            var randomChar = characters.charAt(getRandomNumber(0, nchar-1));
            sb.setCharAt(i, randomChar);
        }
        return sb.toString();
    }

    private static List<Object[]> generateRandomTestCase(int numberOfTC) {
        var listOfTest = new ArrayList<Object[]>();

        for (int i = 0; i < numberOfTC; i++) {

            var number = getRandomNumber(4, 20);
            var generated = genWords(number);
            var expected = workingAnswer(generated);
            var got = Mumble.mumble(generated);

            listOfTest.add(new Object[] { expected, got, new QuestionAttribute(generated) });

        }
        return listOfTest;
    }

    public static void main(String[] args) {
        var testCase = generateTestCase();
        var randomTestCase = generateRandomTestCase(8);

        testCase.addAll(randomTestCase);

        var counter = 0;
        for (Object[] test : testCase) {
            counter++;
            if (test[0].equals(test[1])) {
                System.out.printf("# %d PASSING\n", counter);
            } else {
                System.out.printf("# %d FAILED\n", counter);
            }
            QuestionAttribute argument = (QuestionAttribute) test[2];
            System.out.printf("> ARGUMENTS mumble(\"%s\")\n", argument.input);
            System.out.printf("> EXPECTED %s\n", test[0]);
            System.out.printf("> GOT %s\n", test[1]);
        }
    }
}

class QuestionAttribute {
    public String input;

    public QuestionAttribute(String input) {
        this.input = input;
    }
}