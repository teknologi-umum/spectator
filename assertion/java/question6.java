import java.util.ArrayList;
import java.util.List;

_REPLACE_ME_WITH_DIRECTIVES_

public class question6 {

    _REPLACE_ME_WITH_SOLUTION_

    private static ArrayList<Integer> workingAnswer(ArrayList<Integer> input) {
        var result = new ArrayList<Integer>();
        for (int i = 0; i < input.size(); i++) {
            var grade = input.get(i).intValue();
            if (grade >= 38 && grade % 5 != 0 && grade % 5 >= 3) {
                int added = grade + (5 - grade % 5);
                result.add(added);
            } else {
                result.add(grade);
            }
        }
        return result;
    }

    private static List<Object[]> generateTestCase() {
        var listOfTest = new ArrayList<Object[]>();

        var answer = new ArrayList<Integer>();

        answer.add(75);
        answer.add(67);
        answer.add(40);
        answer.add(33);

        var testInput = new ArrayList<Integer>();

        testInput.add(73);
        testInput.add(67);
        testInput.add(38);
        testInput.add(33);

        listOfTest.add(new Object[] { answer.toString(), Grade.calculateGrade(testInput).toString(), new QuestionAttribute(testInput) });

        return listOfTest;
    }

    private static List<Object[]> generateRandomTestCase(int numberOfTC) {
        var listOfTest = new ArrayList<Object[]>();

        for (int i = 0; i < numberOfTC; i++) {
            var len = getRandomNumber(4, 20);
            var input = genArray(len);

            var expected = workingAnswer(input).toString();
            var got = Grade.calculateGrade(input).toString();

            listOfTest.add(new Object[] { expected, got, new QuestionAttribute(input)  });

        }
        return listOfTest;
    }

    private static int getRandomNumber(int min, int max) {
        return (int) Math.floor(Math.random() * (max - min + 1) + min);
    }

    private static ArrayList<Integer> genArray(int len) {
        var inner = new ArrayList<Integer>();

        for(int i = 0; i < len; i++) {
            inner.add(getRandomNumber(10, 100));
        }

        return inner;
    }
    public static void main(String[] args) {
        var testCase = generateTestCase();
        var randomTestCase = generateRandomTestCase(9);

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
            System.out.printf("> ARGUMENTS calculateGrade(%s)\n", argument.input);
            System.out.printf("> EXPECTED %s\n", test[0]);
            System.out.printf("> GOT %s\n", test[1]);
        }
    }
}

class QuestionAttribute {
    public List<Integer> input;

    public QuestionAttribute(List<Integer> input) {
        this.input = input;
    }
}