# Assertion

Contains file to assert user's input code to a parsable format that will be read by the
backend service.

For anything other than the Hello world (question 0) and the Twinkle-twinlke little star (question 1)
question, the assertion consist of two phase: the pre-defined test cases and random test case.
The random test cases are generated with the correct answer of each test case.

## Format

For each test case, will have the same exact format of:

```
# (number) (PASSING|FAILED)
> ARGUMENTS (string)
> EXPECTED (string)
> GOT (string)
```

Example format:

```
# 1 PASSING
> ARGUMENTS calculateGrade([73, 67, 38, 33])
> EXPECTED [75, 67, 40, 33]
> GOT [75, 67, 40, 33]
# 2 FAILED
> ARGUMENTS calculateGrade([0])
> EXPECTED [0]
> GOT [34]
```

## Unit tests

We've ran into some problem in which the assertion couldn't handle edge cases in the case of
invalid input by user that could be parsed. One example being if the user returned "None" type
on one of the Python test. Hence, we would need to create unit tests to validate the assertion
tests on each question for every language.

For each tests, give a `_failing` suffix on the file name to indicate that the file is expected
to be failing without any other compile/runtime error. Failing test means tests that output a
`# (number) FAILED` text on the output.

The test runner is written in [Julia](https://julialang.org), and it requires each language's
compiler/runtime to be installed:
- Julia 1.7+ (to run the test runner)
- Python 3.10
- Node.js 16
- Java 17 (OpenJDK would be fine)
- PHP 8.1
- GCC 9

To run the test, run:

```sh
julia tests/runner.jl ./tests
```