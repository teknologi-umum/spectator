import java.util.*;
import java.io.*;
import java.math.*;
_REPLACE_ME_WITH_DIRECTIVES_

public class question5 {

    _REPLACE_ME_WITH_SOLUTION_
  
    private static String workingAnswer(String input) {
        int inputLen = input.length();

        int triNum = (inputLen * (inputLen + 1)) / 2;
        int resultLen = triNum + inputLen - 1;

        StringBuilder sb = new StringBuilder();
        sb.setLength(resultLen);

        int pos = 0;
        for (int i = 0; i < inputLen; i++) {
            for (int j = 0; j <= i; j++) {
                char c = input.charAt(i);

                char appended = j == 0 ? Character.toUpperCase(c) : Character.toLowerCase(c);
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

    private static List<String[]> generateTestCase() {
        List<String[]> listOfTest = new ArrayList<String[]>();

        listOfTest.add(new String[] { "A-Bb-Ccc-Dddd", Mumble.mumble("abcd") });
        listOfTest.add(new String[] { "R-Qq-Aaa-Eeee-Zzzzz-Tttttt-Yyyyyyy", Mumble.mumble("RqaEzTy") });

        return listOfTest;
    }

    private static int getRandomNumber(int min, int max) {
        return (int) Math.floor(Math.random() * (max - min + 1) + min);
    }

    private static String genWords(int n) {
        String characters = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz";
        StringBuilder sb = new StringBuilder();

        sb.setLength(n);
        final int nchar = characters.length();
        for (int i = 0; i < n; i++) {
            char randomChar = characters.charAt(getRandomNumber(0, nchar-1));
            sb.setCharAt(i, randomChar);
        }
        return sb.toString();
    }

    private static List<String[]> generateRandomTestCase(int numberOfTC) {
        List<String[]> listOfTest = new ArrayList<String[]>();

        for (int i = 0; i < numberOfTC; i++) {

            int number = getRandomNumber(4, 20);
            String generated = genWords(number);
            String expected = workingAnswer(generated);
            String got = Mumble.mumble(generated);

            listOfTest.add(new String[] { expected, got });

        }
        return listOfTest;
    }

    public static void main(String[] args) {
        List<String[]> testCase = generateTestCase();
        List<String[]> randomTestCase = generateRandomTestCase(8);

        testCase.addAll(randomTestCase);

        int counter = 0;
        for (String[] test : testCase) {
            counter++;
            if (test[0].equals(test[1])) {
                System.out.printf("# %d PASSING\n", counter);
            } else {
                System.out.printf("# %d FAILED\n", counter);
                System.out.printf("> EXPECTED %s\n", test[0]);
                System.out.printf("> GOT %s\n", test[1]);
            }
        }
    }
}
