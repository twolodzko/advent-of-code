# https://adventofcode.com/2020/day/15

example = "0,3,6"

function parse_input(string)
    return map(x -> parse(Int, x), split(string, ',', keepempty=false))
end

struct Stack
    values::Vector{Integer}
    counts::Vector{Integer}

    Stack(values, counts) = new(values, counts)

    function Stack(numbers)
        @assert allunique(numbers)
        return new(numbers, collect(1:length(numbers)))
    end
end

function findnth(value::T, arr::Vector{T}, n::Integer) where {T}
    count = 0
    for i in 1:length(arr)
        if arr[i] == value
            count += 1
            if count == n
                return i
            end
        end
    end
    return nothing
end

@assert findnth(1, [1, 2, 3, 1, 2, 3], 2) == 4
@assert findnth(1, [1, 2, 3, 1, 1, 1], 2) == 4
@assert findnth(2, [1, 2, 3, 1, 1, 1], 2) === nothing

function turn!(stack::Stack)
    pos = findfirst(x -> x == stack.values[1], stack.values[2:end])
    number = isnothing(pos) ? 0 : stack.counts[pos]

    pushfirst!(stack.values, number)
    pushfirst!(stack.counts, 0)
    stack.counts .+= 1

    next = findnth(number, stack.values, 3)
    if !isnothing(next)
        deleteat!(stack.values, next)
        deleteat!(stack.counts, next)
    end

    return stack
end

(function ()
    stack = Stack([6, 3, 0])
    for (i, expected) in enumerate([0, 3, 3, 1, 0, 4, 0])
        number = stack.values[1]
        turn!(stack)
        @assert stack.values[1] == expected
    end
end)()

function create_sequence(starting_sequence, turns)
    # stack = Stack(starting_sequence)
    last_number = starting_sequence[1]
    history = starting_sequence[2:end]
    for i = (length(history) + 2):turns
        pos = findfirst(x -> x == last_number, history)
        pushfirst!(history, last_number)
        last_number = isnothing(pos) ? 0 : pos

        # turn!(stack)
        # if stack.values[1] != last_number
        #     return i, stack, last_number, history
        # end
    end
    pushfirst!(history, last_number)
    return history
end

function part1(string, turns=2020)
    starting_sequence = reverse(parse_input(string))
    return create_sequence(starting_sequence, turns)[1]
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
