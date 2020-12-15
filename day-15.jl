# https://adventofcode.com/2020/day/15

example = "0,3,6"

function parse_input(string)
    return map(x -> parse(Int, x), split(string, ',', keepempty=false))
end

function create_sequence(starting_sequence, turns)
    previous = starting_sequence[1]
    history = starting_sequence[2:end]
    for _ in 1:(turns - length(starting_sequence))
        pos = findfirst(x -> x == previous, history)
        pushfirst!(history, previous)
        previous = isnothing(pos) ? 0 : pos
    end
    pushfirst!(history, previous)
    return history
end

function static_memory_sequence(starting_sequence, turns)
    n = max(turns + 1, maximum(starting_sequence))
    memory = zeros(Int32, n)
    mask = Integer[]

    for (i, prev) in enumerate(starting_sequence[2:end])
        memory[prev + 1] = i
        push!(mask, prev + 1)
    end

    prev = starting_sequence[1]
    for _ in 1:(turns - length(starting_sequence))
        pos = memory[prev + 1]
        if pos == 0
            push!(mask, prev + 1)
        end
        memory[prev + 1] = 0
        memory[mask] .+= 1
        prev = pos
    end
    return prev
end

function part1(string, turns=2020)
    starting_sequence = reverse(parse_input(string))
    return static_memory_sequence(starting_sequence, turns)
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

# @assert part2("0,3,6") == 175594
# @assert part2("1,3,2") == 2578
# @assert part2("2,1,3") == 3544142
# @assert part2("1,2,3") == 261214
# @assert part2("2,3,1") == 6895259
# @assert part2("3,2,1") == 18
# @assert part2("3,1,2") == 362

test = "0,13,16,17,1,10,6"
println("Part 1: $(part1(test))")
# println("Part 2: $(part2(test))")

@assert part1(test) == 276
# @assert part2(test) == 4288986482164
