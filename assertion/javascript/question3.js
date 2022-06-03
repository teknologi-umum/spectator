_REPLACE_ME_WITH_DIRECTIVES_

_REPLACE_ME_WITH_SOLUTION_

function main() {
    const testCases = [
        {
            got: isSameNumber(100, 212),
            expected: false
        },
        {
            got: isSameNumber(25, 25),
            expected: true
        }
    ];

    const randomNumber = (a, b) => Math.floor(Math.random() * (b - a + 1) + a);
    for (let i = 0; i < 4; i++) {
        // Create 4 random tests
        const a = randomNumber(0, 9999);
        const b = randomNumber(0, 9999);
        const expected = a === b;
        const got = isSameNumber(a, b);
        testCases.push({ got, expected });
    }

    for (let i = 0; i < 4; i++) {
        // Create 4 random tests
        const a = randomNumber(0, 9999);
        const expected = true;
        const got = isSameNumber(a, a);
        testCases.push({ got, expected });
    }

    // Test the thang
    for (let i = 0; i < testCases.length; i++) {
        const test = testCases[i];

        if (test.got === test.expected) {
            console.log(`# ${i+1} PASSING`);
        } else {
            console.log(`# ${i+1} FAILED`);
            console.log(`> EXPECTED ${test.expected}`);
            console.log(`> GOT ${test.got}`);
        }
    }
}

main();
