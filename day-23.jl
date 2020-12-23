# https://adventofcode.com/2020/day/23

using ProgressMeter

example = "389125467"

movefirstback!(arr) = push!(arr, popfirst!(arr))

@assert all(movefirstback!([1, 2, 3, 4]) .== [2, 3, 4, 1])

function move(cups, pos; verbose=false)
    verbose && println("cups: ", join(map(x -> x[1] == pos ? "($(x[2]))" : "$(x[2]) ", enumerate(cups)), " "))

    current = cups[pos]

    # pick cups
    picked = Int[]
    i = pos + 1
    for _ in 1:3
        i = i > length(cups) ? 1 : i
        push!(picked, popat!(cups, i))
    end
    verbose && println("pick up: ", join(picked, ", "))

    # pick destination cup
    destination = current - 1
    while !(destination in cups)
        destination -= 1
        destination = destination < 1 ? maximum(cups) : destination
    end
    verbose && println("destination: ", destination)

    # insert picked cups back
    d = findfirst(cups .== destination)
    for i in 1:3
        insert!(cups, d + i, picked[i])
    end

    # rotate to the previous position
    while (findfirst(cups .== current) - pos) > 0
        movefirstback!(cups)
    end

    return cups
end

@assert all(move([3, 8, 9, 1, 2, 5, 4, 6, 7], 1) .== [3, 2, 8, 9, 1, 5, 4, 6, 7])
@assert all(move([3, 2, 5, 4, 6, 7, 8, 9, 1], 3) .== [7, 2, 5, 8, 9, 1, 3, 4, 6])
@assert all(move([7, 4, 1, 5, 8, 3, 9, 2, 6], 9) .== [5, 7, 4, 1, 8, 3, 9, 2, 6])

function part1(input; rounds=100, verbose=false)
    cups = parse.(Int, split(input, ""))
    n = length(cups)

    for i in 1:rounds
        verbose && println("-- move $i --")
        pos = mod1(i, n)
        cups = move(cups, pos, verbose=verbose)
        verbose && println()
    end

    pos = findfirst(cups .== 1)
    idx = [mod1(pos + i, n) for i in 1:(n - 1)]
    return join(cups[idx])
end

@assert part1(example, rounds=10, verbose=false) == "92658374"
@assert part1(example) == "67384529"

function part2(input; rounds=10_000_000, size=1_000_000, verbose=false)
    cups = parse.(Int, split(input, ""))
    biggest = maximum(cups)
    for x in (biggest + 1):size
        push!(cups, x)
    end
    n = length(cups)

    @showprogress for i in 1:rounds
        verbose && println("-- move $i --")
        pos = mod1(i, n)
        cups = move(cups, pos, verbose=verbose)
        verbose && println()
    end

    pos = findfirst(cups .== 1)
    return (*)([cups[mod1(pos + i, n)] for i in 1:2]...)
end

# @assert part2(example) == 149245887792

test = "784235916"
println("Part 1: $(result1 = part1(test))")
println("Part 2: $(result2 = part2(test))")

@assert result1 == "53248976"
# @assert result2 ==