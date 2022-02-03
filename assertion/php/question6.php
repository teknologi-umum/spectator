<?php

_REPLACE_ME_WITH_DIRECTIVES_

_REPLACE_ME_WITH_SOLUTION_

function main(): void {
    $test_cases = [
        [
            "got" => calculateGrade([73, 67, 38, 33]),
            "expected" => [75, 67, 40, 33]
        ]
    ];

    function working_answer(array $arr) {
        $out = [];
        for ($i = 0; $i < count($arr); $i++) {
            $grade = $arr[$i];

            if ($grade >= 38 && $grade % 5 !== 0 && $grade % 5 >= 3) {
                array_push($out, $grade + (5 - $grade % 5));
                continue;
            }
            array_push($out, $grade);
        }

        return out;
    }

    for ($i = 0; $i < 9; $i++) {
        $input = [];
        $arr_len = rand(4, 20);

        for ($j = 0; $j < $arr_len; $j++) {
            array_push($input, rand(0, 100));
        }

        $expected = workingAnswer($input);
        $got = calculateGrade($input);
        array_push($test_cases, [
            "got" => $got,
            "expected" => $expected
        ]);
    }

    for ($i = 0; $i < count($test_cases); $i++) {
        $test = $test_cases[$i];

        if (join(", ", $test["got"]) === join(", ", $test["expected"])) {
            echo "# ". $i + 1 . " PASSING"; 
        } else {
            echo "# " . $i + 1 . " FAILED";
            echo "> EXPECTED " . join(", ", $test["expected"]);
            echo "> GOT " . join(", ", $test["got"]);
        }
    }
}

main();