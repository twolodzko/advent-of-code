
# https://adventofcode.com/2020/day/8

using Printf

example = "
nop +0
acc +1
jmp +4
acc +3
jmp -3
acc -99
acc +1
jmp -4
acc +6
"

struct Command
    name::AbstractString
    value::Integer

    Command(name, value) = new(name, value)

    function Command(string::AbstractString)
        name, value = split(strip(string), r"\s+")
        if !(name in ("nop", "acc", "jmp"))
            error("invalid command: $(name)")
        end
        new(String(name), parse(Int, value))
    end
end

cmd(obj::Command) = obj.name
val(obj::Command) = obj.value

function parser(string::AbstractString)
    commands = Command[]
    for row in split(string, '\n')
        if row == ""
            continue
        end
        push!(commands, Command(row))
    end
    return commands
end

function run_code(commands::Vector{Command}, max_loops=1000, verbose=false)
    exec_count = zeros(Int, length(commands))
    status = false
    accumulator = 0
    offset = 1

    verbose && print("       <start>     =>  ")
    while true
        verbose && println("$(accumulator)")

        if offset == length(commands) + 1
            status = true
            break
        elseif offset > length(commands)
            verbose && println("<error>")
            break
        end
        exec_count[offset] += 1
        command = commands[offset]
        verbose && @printf("%3.0f (%d): %s %+4.0f  =>  ", offset, exec_count[offset], cmd(command), val(command))

        if exec_count[offset] > max_loops
            verbose && println("<break>")
            break
        end

        if cmd(command) == "nop"
            offset += 1
            continue
        end

        if cmd(command) == "acc"
            accumulator += val(command)
            offset += 1
        elseif cmd(command) == "jmp"
            offset += val(command)
        else
            verbose && println("<error>")
            break
        end
    end
    verbose && println()

    return accumulator, status
end

@assert run_code(parser(example), 1) == (5, false)

code_that_finishes = "
nop +0
acc +10
jmp +3
nop +0
acc +100
acc -5
"

@assert run_code(parser(code_that_finishes), 1) == (5, true)

function part1(code)
    commands = parser(code)
    accumulator, status = run_code(commands, 1)
    return accumulator
end

function swap(command::Command)
    if cmd(command) == "acc"
        error("invalid command: $(command)")
    end
    name = cmd(command) == "nop" ? "jmp" : "nop"
    return Command(name, val(command))
end

function part2(code, verbose=false)
    commands_orig = parser(code)
    command_names = reverse([cmd(command) for command in commands_orig])
    start_search = 1
    success = false
    accumulator = nothing

    while !success
        commands = copy(commands_orig)
        idx = findfirst(x -> x in ("nop", "jmp"), command_names[start_search:end])
        idx += start_search - 1
        # because we search starting from the back
        commands[end-idx+1] = swap(commands[end-idx+1])
        accumulator, success = run_code(commands, 2)
        start_search = idx + 1

        verbose && @printf("Progress: %.2f%%\n", idx / length(commands) * 100)
    end

    return accumulator
end

@assert part2(example) == 8

test = read("data/day-08.txt", String)
println("Part 1: $(part1(test))")
println("Part 2: $(part2(test))")

@assert part1(test) == 1832
@assert part2(test) == 662
