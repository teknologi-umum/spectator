<?php

_REPLACE_ME_WITH_DIRECTIVES_

_REPLACE_ME_WITH_SOLUTION_

function main(): void {
    $test_cases = [
        [
            "arguments" => "mumble(\"abcd\")",
            "expected" => "A-Bb-Ccc-Dddd",
            "got" => mumble("abcd")
        ],
        [
            "arguments" => "mumble(\"RqaEzTy\")",
            "expected" => "R-Qq-Aaa-Eeee-Zzzzz-Tttttt-Yyyyyyy",
            "got" => mumble("RqaEzTy")
        ]
    ];

    function working_answer(string $s): string {
        $result = '';
        for ($i = 0; $i < strlen($s); $i++) {
            $result .= ucfirst(strtolower(str_repeat($s[$i], $i + 1)));
            if ($i < strlen($s) - 1) {
                $result .= "-";
            }
        }
        
        return $result;
    }

    $characters = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz";
    for ($i = 0; $i < 8; $i++) {
        $chars = "";
        $seed = rand(4, 20);
        for ($j = 0; $j < $seed; $j++) {
            $chars .= $characters[rand(0, strlen($characters) - 1)];
        }

        $expected = working_answer($chars);
        $got = mumble($chars);
        $arguments = "mumble($chars)";
        
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