_REPLACE_ME_WITH_DIRECTIVES_

_REPLACE_ME_WITH_SOLUTION_

function main() {
    const testCases = [
        {
            got: calculateGrade([73, 67, 38, 33]),
            expected: [75, 67, 40, 33]
        }
    ];

    const workingAnswer = (arr) => {
        const out = [];
        for (let i = 0; i < arr.length; i++) {
            const grade = arr[i];

            if (grade >= 38 && grade % 5 !== 0 && grade % 5 >= 3) {
                out.push(grade + (5 - grade % 5));
                continue;
            }

            out.push(grade);
        }
        return out;
    }
    const randomNumber = (a, b) => Math.floor(Math.random() * (b - a + 1) + a);
    for (let i = 0; i < 9; i++) {
        const input = [];
        const arrLength = randomNumber(4, 20);
        for (let j = 0; j < arrLength; j++) {
            input.push(randomNumber(0, 100));
        }

        const expected = workingAnswer(input);
        const got = calculateGrade(input);

        testCases.push({ expected, got });
    }

    for (let i = 0; i < testCases.length; i++) {
        const test = testCases[i];
        if (test.expected.join(", ") === test.got.join(", ")) {
            console.log(`# ${i+1} PASSING`);
        } else {
            console.log(`# ${i+1} FAILED`);
            console.log(`> EXPECTED ${test.expected.join(", ")}`);
            console.log(`> GOT ${test.got.join(", ")}`)
        }
    }
}

main();
