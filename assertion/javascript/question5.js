_REPLACE_ME_WITH_DIRECTIVES_

_REPLACE_ME_WITH_SOLUTION_

function main() {
    const testCases = [
        {
            arguments: `mumble("abcd")`,
            expected: "A-Bb-Ccc-Dddd",
            got: mumble("abcd")
        },
        {
            arguments: `got: mumble("RqaEzTy")`,
            expected: "R-Qq-Aaa-Eeee-Zzzzz-Tttttt-Yyyyyyy",
            got: mumble("RqaEzTy")
        }
    ]

    const workingAnswer = (str) => {
        let result = "";
        for (let i = 0; i < str.length; i++) {
            const c = str[i];
            result += c.toUpperCase() + c.toLowerCase().repeat(i);
            if (i < str.length - 1) {
                result += "-"
            }
        }
        return result;
    }

    const randomNumber = (a, b) => Math.floor(Math.random() * (b - a + 1) + a);
    const characters = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz";
    for (let i = 0; i < 8; i++) {
        let chars = "";
        const seed = randomNumber(4, 20);
        for (let j = 0; j < seed; j++) {
            chars += characters[randomNumber(0, characters.length - 1)];
        }

        const expected = workingAnswer(chars);
        const got = mumble(chars);
        const arguments = `mumble("${chars}")`;
        testCases.push({ expected, got, arguments });
    }

    for (let i = 0; i < testCases.length; i++) {
        const test = testCases[i];
        if (test.got === test.expected) {
            console.log(`# ${i + 1} PASSING`);
            console.log(`> ARGUMENTS ${test.arguments}`);
            console.log(`> EXPECTED ${test.expected}`);
            console.log(`> GOT ${test.got}`);
        } else {
            console.log(`# ${i + 1} FAILED`);
            console.log(`> ARGUMENTS ${test.arguments}`);
            console.log(`> EXPECTED ${test.expected}`);
            console.log(`> GOT ${test.got}`);
        }
    }
}

main();
