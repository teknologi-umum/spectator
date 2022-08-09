<?php

_REPLACE_ME_WITH_DIRECTIVES_

_REPLACE_ME_WITH_SOLUTION_

function main(): void {
    $test_cases = [
        [
            "arguments" => "isSameNumber(100, 212)",
            "got" => isSameNumber(100, 212),
            "expected" => false
        ],
        [
            "arguments" => "isSameNumber(25, 25)",
            "got" => isSameNumber(25, 25),
            "expected" => true
        ],
    ];

    for ($i = 0; $i < 4; $i++) {
        $a = rand(0, 9999);
        $b = rand(0, 9999);
        $expected = $a === $b;
        $got = isSameNumber($a, $b);
        $arguments = "isSameNumber($a, $b)";
        array_push($test_cases, [
            "got" => $got,
            "expected" => $expected,
            "arguments" => $arguments
        ]);
    }

    for ($i = 0; $i < 4; $i++) {
        $a = rand(0, 9999);
        $expected = true;
        $got = isSameNumber($a, $a);
        $arguments = "isSameNumber($a, $a)";
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