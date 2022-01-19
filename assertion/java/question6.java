import java.util.ArrayList;
import java.util.List;

_REPLACE_ME_WITH_DIRECTIVES_

public class question6 {

    _REPLACE_ME_WITH_SOLUTION_

    private static ArrayList<Integer> workingAnswer(ArrayList<Integer> input) {
        ArrayList<Integer> result = new ArrayList<Integer>();
        for (int i = 0; i < input.size(); i++) {
            int grade = input.get(i).intValue();
            if (grade >= 38 && grade % 5 != 0 && grade % 5 >= 3) {
                int added = grade + (5 - grade % 5);
                result.add(added);
            } else {
                result.add(grade);
            }
        }
        return result;
    }

    private static List<String[]> generateTestCase() {
        List<String[]> listOfTest = new ArrayList<String[]>();

        ArrayList<Integer> answer = new ArrayList<Integer>();

        answer.add(75);
        answer.add(67);
        answer.add(40);
        answer.add(33);

        ArrayList<Integer> testInput = new ArrayList<Integer>();

        testInput.add(73);
        testInput.add(67);
        testInput.add(38);
        testInput.add(33);

        listOfTest.add(new String[] { answer.toString(), Grade.calculateGrade(testInput).toString() });

        return listOfTest;
    }

    private static List<String[]> generateRandomTestCase(int numberOfTC) {
        List<String[]> listOfTest = new ArrayList<String[]>();

        for (int i = 0; i < numberOfTC; i++) {
            int len = getRandomNumber(4, 20);
            ArrayList<Integer> input = genArray(len);

            String expected = workingAnswer(input).toString();
            String got = Grade.calculateGrade(input).toString();

            listOfTest.add(new String[] { expected, got });

        }
        return listOfTest;
    }

    private static int getRandomNumber(int min, int max) {
        return (int) Math.floor(Math.random() * (max - min + 1) + min);
    }

    private static ArrayList<Integer> genArray(int len) {
        ArrayList<Integer> inner = new ArrayList<Integer>();

        for(int i = 0; i < len; i++) {
            inner.add(getRandomNumber(10, 100));
        }

        return inner;
    }
    public static void main(String[] args) {
        List<String[]> testCase = generateTestCase();
        List<String[]> randomTestCase = generateRandomTestCase(9);

        testCase.addAll(randomTestCase);

        int counter = 0;
        for (String[] test : testCase) {
            counter++;
            if (test[0].equals(test[1])) {
                System.out.printf("# %d PASSING\n", counter);
                System.out.printf("> EXPECTED %s\n", test[0]);
                System.out.printf("> GOT %s\n", test[1]);
            } else {
                System.out.printf("# %d FAILED\n", counter);
                System.out.printf("> EXPECTED %s\n", test[0]);
                System.out.printf("> GOT %s\n", test[1]);
            }
        }
    }

}
