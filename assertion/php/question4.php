<?php

_REPLACE_ME_WITH_DIRECTIVES_

_REPLACE_ME_WITH_SOLUTION_

function main(): void {
    $test_cases = [
        [
            "arguments" => "findHeaterPower(100212373)",
            "got" => findHeaterPower(100212373),
            "expected" => 19
        ]
    ];

    function working_answer(int $n) {
        $s = strval($n);
        $sum = 0;
        for ($i = 0; $i < strlen($s); $i++) {
            $sum += intval($s[$i]);
        }

        return $sum;
    }

    for ($i = 0; $i < 9; $i++) {
        $n = rand(1000000000, 9999999999);
        $expected = working_answer($n);
        $got = findHeaterPower($n);
        $arguments = "findHeaterPower($n)";
        array_push($test_cases, [
            "got" => $got,
            "expected" => $expected,
            "arguments" => $arguments
        ]);
    }

    for ($i = 0; $i < count($test_cases); $i++) {
        $test = $test_cases[$i];

        if ($test["got"] === $test["expected"]) {
            echo "# ". $i + 1 . " PASSING\n"; 
            echo "> ARGUMENTS " . $test["arguments"] . "\n";
            echo "> EXPECTED " . $test["expected"] . "\n";
            echo "> GOT " . $test["got"] . "\n";
        } else {
            echo "# " . $i + 1 . " FAILED\n";
            echo "> ARGUMENTS " . $test["arguments"] . "\n";
            echo "> EXPECTED " . $test["expected"] . "\n";
            echo "> GOT " . $test["got"] . "\n";
        }
    }
}

main();