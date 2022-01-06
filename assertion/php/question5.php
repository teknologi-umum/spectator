<?php

{0}

function main(): void {
    $test_cases = [
        [
            "expected" => "A-Bb-Ccc-Dddd",
            "got" => mumble("abcd")
        ],
        [
            "expected" => "R-Qq-Aaa-Eeee-Zzzzz-Tttttt-Yyyyyyy",
            "got" => mumble("RqaEzTy")
        ]
    ];

    function working_answer(string $s): string {
        $result = '';
        for ($i = 0; $i < strlen($s); $i++) {
            $result .= ucfirst(strtolower(str_repeat($s[$i], $i + 1)));
            if ($i < strlen($s)) {
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