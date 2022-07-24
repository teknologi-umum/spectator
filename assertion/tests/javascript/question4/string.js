// `findHeaterPower` is function that accepts an argument as long and
// returns an integer as its output.
function findHeaterPower(input) {
    let power = 0;

    for (let i = 0; i < input.toString().length; i++) {
        power += Number((input.toString())[i]);
    }

    return power.toString();
}