// `mumble` is a function that accepts an argument as string and returns a
// string as its output.
function mumble($input) {
    $result = '';
    for ($i = 0; $i < strlen($input); $i++) {
        $result .= ucfirst(strtolower(str_repeat($input[$i], $i + 1)));
        if ($i < strlen($input) - 1) {
            $result .= "-";
        }
    }
    
    return $result;
}