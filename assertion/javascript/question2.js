_REPLACE_ME_WITH_DIRECTIVES_

_REPLACE_ME_WITH_SOLUTION_

function main() {
  const testCases = [
    {
      arguments: "calculateTemperature(100, \"Celcius\", \"Fahrenheit\")",
      got: calculateTemperature(100, "Celcius", "Fahrenheit"),
      expected: 212
    },
    {
      arguments: "calculateTemperature(212, \"Fahrenheit\", \"Kelvin\")",
      got: calculateTemperature(212, "Fahrenheit", "Kelvin"),
      expected: 373.15
    },
    {
      arguments: "calculateTemperature(0, \"Celcius\", \"Kelvin\")",
      got: calculateTemperature(0, "Celcius", "Kelvin"),
      expected: 273.15
    },
    {
      arguments: "calculateTemperature(0, \"Celcius\", \"Fahrenheit\")",
      got: calculateTemperature(0, "Celcius", "Fahrenheit"),
      expected: 32
    },
    {
      arguments: "calculateTemperature(0, \"Kelvin\", \"Fahrenheit\")",
      got: calculateTemperature(0, "Kelvin", "Fahrenheit"),
      expected: -459.67
    }
  ];

  const workingAnswer = (n, a, b) => {
    if (a === "Celcius" && b === "Fahrenheit") {
      return (n * 9 / 5) + 32;
    } else if (a === "Celcius" && b == "Kelvin") {
      return n + 273.15;
    } else if (a === "Fahrenheit" && b === "Celcius") {
      return (n - 32) * 5 / 9;
    } else if (a === "Fahrenheit" && b === "Kelvin") {
      return (n - 32) * 5 / 9 + 273.15;
    } else if (a === "Kelvin" && b === "Celcius") {
      return n - 273.15;
    } else if (a === "Kelvin" && b === "Fahrenheit") {
      return (n - 273.15) * 9 / 5 + 32;
    } else {
      return n;
    }
  }

  const randomNumber = (a, b) => Math.floor(Math.random() * (b - a + 1) + a);
  const temperatures = ["Celcius", "Fahrenheit", "Kelvin"]
  for (let i = 0; i < 5; i++) {
    // Create 5 random tests
    const from = temperatures[randomNumber(0, temperatures.length - 1)];
    const to = temperatures[randomNumber(0, temperatures.length - 1)];
    const n = randomNumber(0, 1000);
    const arguments = `calculateTemperature(${n}, "${from}", "${to}")`;
    const expected = workingAnswer(n, from, to);
    const got = calculateTemperature(n, from, to);
    testCases.push({ got, expected, arguments });
  }

  // Test the thang
  for (let i = 0; i < testCases.length; i++) {
    const test = testCases[i];

    if (typeof test.got === "string" && Number.isNaN(Number(test.got))) {
      console.log(`# ${i + 1} FAILED`);
      console.log(`> ARGUMENTS ${test.arguments}`);
      console.log(`> EXPECTED ${test.expected.toFixed(2)}`);
      console.log(`> GOT ${Number(test.got).toFixed(2)}`);
      continue;
    }

    if (Number(test.got).toFixed(2) === test.expected.toFixed(2)) {
      console.log(`# ${i + 1} PASSING`);
      console.log(`> ARGUMENTS ${test.arguments}`);
      console.log(`> EXPECTED ${test.expected.toFixed(2)}`);
      console.log(`> GOT ${Number(test.got).toFixed(2)}`);
    } else {
      console.log(`# ${i + 1} FAILED`);
      console.log(`> ARGUMENTS ${test.arguments}`);
      console.log(`> EXPECTED ${test.expected.toFixed(2)}`);
      console.log(`> GOT ${Number(test.got).toFixed(2)}`);
    }
  }
}

main();
