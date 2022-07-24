println("✨ Starting tests")

_args = ARGS
dirpath = ""

if length(_args) > 0
    global dirpath = _args[1]
else
    global dirpath = "./tests"
end

failedresults = 0

mappedext = Dict(
    "c" => ".c",
    "cpp" => ".cpp",
    "java" => ".java",
    "javascript" => ".js",
    "php" => ".php",
    "python" => ".py",
)

mappedcmd = Dict(
    "c" => "gcc -Wall -Wextra -Werror -O2 -std=c99 -pedantic -o code {file} && ./code && rm code",
    "cpp" => "g++ -Wall -Wextra -Werror -O2 -std=c++17 -pedantic -o code {file} && ./code && rm code",
    "java" => "java {file}",
    "javascript" => "node {file}",
    "php" => "php {file}",
    "python" => "python3 {file}"
)

twinkle = "Twinkle twinkle little star\nHow I wonder what you are\nUp above the world so high\nLike a diamond in the sky\nTwinkle twinkle little star\nHow I wonder what you are"

function normalassertion(result::String)
    # split each string by new line
    results = split(strip(result), "\n")

    ok = false

    # for every line of results, check whether
    # it contains "PASSED"
    for line in results
        if contains(strip(line), "PASSING")
            ok = true
        end
    end

    ok
end


mkpath(joinpath(dirpath, "tmp"))

for walkedpath in walkdir(dirpath)
    if length(walkedpath) < 3
        println(walkedpath)
    end

    (root, dirs, files) = walkedpath

    if root == dirpath || root == joinpath(dirpath, "tmp")
        continue
    end

    if length(files) == 0
        continue
    end

    language = split(root, "/")[3]
    questionnumber = split(root, "/")[4]

    execcmd = get(mappedcmd, language, "")

    println("⌛ Testing ", language, " - ", questionnumber)

    open(joinpath(language, questionnumber * get(mappedext, language, "")), "r") do questionio
        questioncontent = read(questionio, String)
        for file in files
            open(joinpath(root, file), "r") do assertio
                assertcontent = read(assertio, String)

                executecontent = replace(questioncontent, "_REPLACE_ME_WITH_SOLUTION_" => assertcontent)
                executecontent = replace(executecontent, "_REPLACE_ME_WITH_DIRECTIVES_" => "")

                executepath = joinpath(dirpath, "tmp", file)
                touch(executepath)
                open(executepath, "w") do executeio
                    write(executeio, strip(executecontent))
                end

                executecommand = replace(execcmd, "{file}" => executepath)

                result = read(`sh -c $executecommand`, String)

                if questionnumber == "question1"
                    if strip(result) == twinkle
                        println("   ✅ Test passed: " * file)
                    else
                        println("   ❌ Test failed: " * file)
                        println(result)
                        global failedresults += 1
                    end
                else
                    if normalassertion(result)
                        println("   ✅ Test passed: " * file)
                    else
                        println("   ❌ Test failed: " * file)
                        println(result)
                        global failedresults += 1
                    end
                end

                rm(executepath; force=true)
            end
        end
    end

    print("\n")
end

if failedresults > 0
    println("❌ Tests failed: ", failedresults)
    exit(1)
end

println("✅ All tests passed")
exit(0)