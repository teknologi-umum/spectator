{0}

function main() {
    const testCases = [
        {
            expected: "A-Bb-Ccc-Dddd",
            got: accum("abcd")
        },
        {
            expected: "R-Qq-Aaa-Eeee-Zzzzz-Tttttt-Yyyyyyy",
            got: accum("RqaEzTy")
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
        const got = accum(chars);
        testCases.push({ expected, got });
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
