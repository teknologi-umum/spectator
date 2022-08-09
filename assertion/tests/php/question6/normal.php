// `calculateGrade` is a function that accepts an array of integer as its input and
// returns an array of integer as its output
function calculateGrade($grades) {
    $out = [];
    for ($i = 0; $i < count($grades); $i++) {
        $grade = $grades[$i];

        if ($grade >= 38 && $grade % 5 !== 0 && $grade % 5 >= 3) {
            array_push($out, $grade + (5 - $grade % 5));
            continue;
        }
        array_push($out, $grade);
    }

    return $out;
}