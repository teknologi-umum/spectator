// `findHeaterPower` is function that accepts an argument as long and
// returns an integer as its output.
function findHeaterPower($input) {
    $s = strval($input);
    $sum = 0;
    for ($i = 0; $i < strlen($s); $i++) {
        $sum += intval($s[$i]);
    }

    return $sum;
}