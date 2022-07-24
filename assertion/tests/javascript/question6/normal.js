// calculateGrade is a function that accepts an array of integer as its input
// and returns an array of integer as its output
function calculateGrade(grade) {
    let grades = [];
    for (let i = 0; i < grade.length; i++) {
        if (grade[i] >= 38 && grade[i] % 5 !== 0 && grade[i] % 5 >= 3) {
            grades.push(grade[i] + (5 - grade[i] % 5));
            continue;
        }

        grades.push(grade[i]);
    }

    return grades;
}