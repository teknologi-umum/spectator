<?php

_REPLACE_ME_WITH_DIRECTIVES_

_REPLACE_ME_WITH_SOLUTION_

function main(): void {
    function is_celcius (string $s): bool { return $s === "Celcius"; }
    function is_fahrenheit (string $s): bool { return  $s === "Farenheit"; }
    function is_kelvin (string $s): bool { return  $s === "Kelvin"; }
    function to_fixed_of_two (float $n): number { return number_format($n, 2, ".", ""); }

    $temperatures = ["Celcius", "Farenheit", "Kelvin"];
    $test_cases = [
        [
            "got" => calculateTemperature(100, "Celcius", "Fahrenheit"),
            "expected" => 212
        ],
        [
            "got" => calculateTemperature(212, "Fahrenheit", "Kelvin"),
            "expected" => 373.15
        ],
        [
            "got" => calculateTemperature(0, "Celcius", "Kelvin"),
            "expected" => 273.15
        ],
        [
            "got" => calculateTemperature(0, "Celcius", "Fahrenheit"),
            "expected" => 32
        ],
        [
            "got" => calculateTemperature(0, "Kelvin", "Fahrenheit"),
            "expected" => -459.67
        ],
    ];

    $working_answer = function(float $n, string $a, string $b): float {
        if (is_celcius($a) && is_fahrenheit($b)) return ($n * 9 / 5) + 32;
        if (is_celcius($a) && is_kelvin($b)) return $n + 273.15;
        if (is_fahrenheit($a) && is_celcius($b)) return ($n - 32) * 5 / 9;
        if (is_fahrenheit($a) && is_kelvin($b)) return ($n - 32) * 5 / 9 + 273.15;
        if (is_kelvin($a) && is_celcius($b)) return $n - 273.15;
        if (is_kelvin($a) && is_fahrenheit($b)) return ($n - 273.15) * 9 / 5 + 32;
        return $n;
    };

    for ($i = 0; $i < 5; $i++) {
        $from = $temperatures[rand(0, count($temperatures) - 1)];
        $to = $temperatures[rand(0, count($temperatures) - 1)];
        $n = rand(-500, 500);
        $expected = $working_answer($n, $from, $to);
        $got = calculateTemperature($n, $from, $to);
        array_push($test_cases, [
            "got" => $got,
            "expected" => $expected
        ]);
    }

    for ($i = 0; $i < count($test_cases); $i++) {
        $test = $test_cases[$i];

        if (to_fixed_of_two($test["got"]) === to_fixed_of_two($test["expected"])) {
            echo "# ". $i + 1 . " PASSING";
        } else {
            echo "# " . $i + 1 . " FAILED";
            echo "> EXPECTED " . to_fixed_of_two($test["expected"]);
            echo "> GOT " . to_fixed_of_two($test["got"]);
        }
    }
}

main();
