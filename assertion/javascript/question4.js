{0}

function main() {
    const testCases = [
        {
            got: findHeaterPower(100212373),
            expected: 19
        }
    ];

    const workingAnswer = (n) => {
        const s = String(n);
        let sum = 0;
        for (let i = 0; i < s.length; i++) {
          sum += Number(s[i]);
        }
        return sum;
    };
    const randomNumber = (a, b) => Math.floor(Math.random() * (b - a + 1) + a);
    for (let i = 0; i < 9; i++) {
        const n = randomNumber(1000000000, 9999999999);
        const expected = workingAnswer(n);
        const got = findHeaterPower(n);
        testCases.push({ got, expected });
    }

    for (let i = 0; i < testCases.length; i++) {
        const test = testCases[i];
        if (test.got === test.expected) {
            console.log(`# ${i+1} PASSING`);
        } else {
            console.log(`# ${i+1} FAILED`);
            console.log(`> EXPECTED ${test.expected}`);
            console.log(`> GOT ${test.got}`)
        }
    }
}

main();
