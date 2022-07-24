// `mumble` is a function that accepts an argument as string and returns a
// string as its output.
function mumble(input) {
    let mumble = "";
    for (let i = 0; i < input.length; i++) {
        mumble += input[i].toUpperCase() + input[i].toLowerCase().repeat(i);

        if (i < input.length - 1) {
            mumble += "-";
        }
    }

    return mumble;
}

