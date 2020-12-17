
# https://adventofcode.com/2020/day/8

# The boot code is represented as a text file with one instruction per line of text.
# Each instruction consists of an operation (`acc`, `jmp`, or `nop`) and an argument
# (a signed number like `+4` or `-20`).
#
#  * `acc` increases or decreases a single global value called the accumulator by the
#    value given in the argument. For example, `acc +7` would increase the accumulator
#    by 7. The accumulator starts at 0. After an acc instruction, the instruction
#    immediately below it is executed next.
#  * `jmp` jumps to a new instruction relative to itself. The next instruction to execute
#    is found using the argument as an offset from the jmp instruction; for example,
#    `jmp +2` would skip the next instruction, `jmp +1` would continue to the instruction
#    immediately below it, and `jmp -20` would cause the instruction 20 lines above to be
#    executed next.
#  * `nop` stands for No OPeration - it does nothing. The instruction immediately below
#    it is executed next.

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

const commands = Dict("nop" => :nop, "acc" => :acc, "jmp" => :jmp)

struct Command
    name::Symbol
    value::Integer

    Command(name, value) = new(name, value)

    function Command(string::AbstractString)
        name, value = split(strip(string), r"\s+")
        new(commands[name], parse(Int, value))
    end
end

cmd(obj::Command) = obj.name
val(obj::Command) = obj.value

function parser(string::AbstractString)
    program = Command[]
    for row in split(string, '\n', keepempty=false)
        push!(program, Command(row))
    end
    return program
end

function run_code(program::Vector{Command}, max_loops=1000, verbose=false)
    exec_count = zeros(Int, length(program))
    status = false
    accumulator = 0
    offset = 1

    verbose && print("       <start>     =>  ")
    while true
        verbose && println("$(accumulator)")

        if offset == length(program) + 1
            status = true
            break
        elseif offset > length(program)
            verbose && println("<error>")
            break
        end
        exec_count[offset] += 1
        command = program[offset]
        verbose && @printf("%3.0f (%d): %s %+4.0f  =>  ", offset, exec_count[offset], cmd(command), val(command))

        if exec_count[offset] > max_loops
            verbose && println("<break>")
            break
        end

        if cmd(command) == :nop
            offset += 1
            continue
        end

        if cmd(command) == :acc
            accumulator += val(command)
            offset += 1
        elseif cmd(command) == :jmp
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

# Run your copy of the boot code. Immediately before any instruction is executed a second
# time, what value is in the accumulator?

function part1(code)
    program = parser(code)
    accumulator, _ = run_code(program, 1)
    return accumulator
end

# Fix the program so that it terminates normally by changing exactly one `jmp` (to `nop`)
# or `nop` (to `jmp`). What is the value of the accumulator after the program terminates?

function swap(command::Command)
    if cmd(command) == :acc
        error("invalid command: $(command)")
    end
    name = cmd(command) == :nop ? :jmp : :nop
    return Command(name, val(command))
end

function part2(code, verbose=false)
    program_orig = parser(code)
    command_names = reverse([cmd(command) for command in program_orig])
    start_search = 1
    success = false
    accumulator = nothing

    while !success
        program = copy(program_orig)
        idx = findfirst(x -> x in (:nop, :jmp), command_names[start_search:end])
        idx += start_search - 1
        # because we search starting from the back
        program[end-idx+1] = swap(program_orig[end-idx+1])
        accumulator, success = run_code(program, 2)
        start_search = idx + 1

        verbose && @printf("Progress: %.2f%%\n", idx / length(program) * 100)
    end

    return accumulator
end

@assert part2(example) == 8

test = read("data/day-08.txt", String)
println("Part 1: $(result1 = part1(test))")
println("Part 2: $(result2 = part2(test))")

@assert result1 == 1832
@assert result2 == 662
