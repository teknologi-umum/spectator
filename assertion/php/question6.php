<?php

_REPLACE_ME_WITH_DIRECTIVES_

_REPLACE_ME_WITH_SOLUTION_

function main(): void {
    $test_cases = [
        [
            "arguments" => "calculateGrade([73, 67, 38, 33])",
            "got" => calculateGrade([73, 67, 38, 33]),
            "expected" => [75, 67, 40, 33]
        ]
    ];

    function working_answer(array $arr): array {
        $out = [];
        for ($i = 0; $i < count($arr); $i++) {
            $grade = $arr[$i];

            if ($grade >= 38 && $grade % 5 !== 0 && $grade % 5 >= 3) {
                array_push($out, $grade + (5 - $grade % 5));
                continue;
            }
            array_push($out, $grade);
        }

        return $out;
    }

    for ($i = 0; $i < 9; $i++) {
        $input = [];
        $arr_len = rand(4, 20);

        for ($j = 0; $j < $arr_len; $j++) {
            array_push($input, rand(0, 100));
        }

        $expected = working_answer($input);
        $got = calculateGrade($input);
        $arguments = "calculateGrade(".json_encode($input).")";
        array_push($test_cases, [
            "got" => $got,
            "expected" => $expected,
            "arguments" => $arguments
        ]);
    }

    for ($i = 0; $i < count($test_cases); $i++) {
        $test = $test_cases[$i];
        
        if (is_null($test["got"])) {
            echo "# " . $i + 1 . " FAILED\n";
            echo "> ARGUMENTS " . $test["arguments"] . "\n";
            echo "> EXPECTED " . json_encode($test["expected"]) . "\n";
            echo "> GOT null\n";
            continue;
        }

        if (join(", ", $test["got"]) === join(", ", $test["expected"])) {
            echo "# ". $i + 1 . " PASSING\n";
            echo "> ARGUMENTS " . $test["arguments"] . "\n";
            echo "> EXPECTED " . json_encode($test["expected"]) . "\n";
            echo "> GOT " . json_encode($test["got"]) . "\n";
        } else {
            echo "# " . $i + 1 . " FAILED\n";
            echo "> ARGUMENTS " . $test["arguments"] . "\n";
            echo "> EXPECTED " . json_encode($test["expected"]) . "\n";
            echo "> GOT " . json_encode($test["got"]) . "\n";
        }
    }
}

main();