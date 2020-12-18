# https://adventofcode.com/2020/day/15

example = "0,3,6"

function read_array(string::AbstractString)::Vector{Int}
    return parse.(Int, split(string, ',', keepempty=false))
end

"""
Use array as a lookup table since keys are integers and we know the
upper bound. Access to array's elements is O(1).
"""
function memoize_sequence(starting_sequence, turns)
    n = max(turns + 1, maximum(starting_sequence))
    memory = zeros(Int32, n) .- (2 * n)

    for (i, prev) in enumerate(starting_sequence[1:(end - 1)])
        memory[prev + 1] = i
    end

    prev = starting_sequence[end]
    pos = 0

    for i = length(starting_sequence):(turns - 1)
        pos = memory[prev + 1]
        memory[prev + 1] = i
        prev = pos > 0 ? i - pos : 0
    end
    return prev
end

function part1(string, turns=2020)
    starting_sequence = read_array(string)
    return memoize_sequence(starting_sequence, turns)
end

@assert part1(example, 10) == 0
@assert part1(example) == 436
@assert part1("1,3,2") == 1
@assert part1("2,1,3") == 10
@assert part1("1,2,3") == 27
@assert part1("2,3,1") == 78
@assert part1("3,2,1") == 438
@assert part1("3,1,2") == 1836

part2(string) = part1(string, 30000000)

@assert part2("0,3,6") == 175594
@assert part2("1,3,2") == 2578
@assert part2("2,1,3") == 3544142
@assert part2("1,2,3") == 261214
@assert part2("2,3,1") == 6895259
@assert part2("3,2,1") == 18
@assert part2("3,1,2") == 362

test = "0,13,16,17,1,10,6"
println("Part 1: $(result1 = part1(test))")
println("Part 2: $(result2 = part2(test))")

@assert result1 == 276
@assert result2 == 31916
