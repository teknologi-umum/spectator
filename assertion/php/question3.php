<?php

_REPLACE_ME_WITH_DIRECTIVES_

_REPLACE_ME_WITH_SOLUTION_

function main(): void {
    $test_cases = [
        [
            "got" => isSameNumber(100, 212),
            "expected" => false
        ],
        [
            "got" => isSameNumber(25, 25),
            "expected" => true
        ],
    ];

    for ($i = 0; $i < 4; $i++) {
        $a = rand(0, 9999);
        $b = rand(0, 9999);
        $expected = $a === $b;
        $got = isSameNumber($a, $b);
        array_push($test_cases, [
            "got" => $got,
            "expected" => $expected
        ]);
    }

    for ($i = 0; $i < 4; $i++) {
        $a = rand(0, 9999);
        $expected = true;
        $got = isSameNumber($a, $a);
        array_push($test_cases, [
            "got" => $got,
            "expected" => $expected
        ]);
    }

    for ($i = 0; $i < count($test_cases); $i++) {
        $test = $test_cases[$i];

        if ($test["got"] === $test["expected"]) {
            echo "# ". $i + 1 . " PASSING"; 
        } else {
            echo "# " . $i + 1 . " FAILED";
            echo "> EXPECTED " . $test["expected"];
            echo "> GOT " . $test["got"];
        }
    }
}

main();