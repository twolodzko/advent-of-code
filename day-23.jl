# https://adventofcode.com/2020/day/23

using ProgressMeter

example = "389125467"

include("LinkedList.jl")

function play(cups, rounds::Int)
    max = maximum(cups)

    head = tolinkedlist(cups)
    tail = last(head)
    tail.next = head

    p = Progress(rounds)
    for i in 1:rounds
        current = head.value

        picked = peek(head.next, 3)
        head.next = head.next.next.next.next

        destination = current - 1
        destination = destination < 1 ? max : destination
        while destination in picked
            destination -= 1
            destination = destination < 1 ? max : destination
        end

        insertafter!(head, destination, picked)

        head = head.next
        next!(p)
    end

    n = length(cups)
    return peek(head, n + (n - 1))[n:(n + n - 1)]
end

@assert all(play(Int[3, 8, 9, 1, 2, 5, 4, 6, 7], 10) .== [5, 8, 3, 7, 4, 1, 9, 2, 6])

function part1(input; rounds::Integer=100, verbose=false)
    cups = parse.(Int, split(input, ""))

    cups = play(cups, rounds)

    pos = findfirst(cups .== 1)
    n = length(cups)
    idx = [mod1(pos + i, n) for i in 1:(n - 1)]
    return join(cups[idx])
end

@assert part1(example, rounds=10, verbose=false) == "92658374"
@assert part1(example) == "67384529"

function part2(input; rounds=10_000_000, size=1_000_000, verbose=false)
    cups = parse.(Int, split(input, ""))
    max = maximum(cups)
    for x in (max + 1):size
        push!(cups, x)
    end
    n = length(cups)

    cups = play(cups, rounds)

    pos = findfirst(cups .== 1)
    return (*)([cups[mod1(pos + i, n)] for i in 1:2]...)
end

@assert part2(example) == 149245887792

test = "784235916"
println("Part 1: $(result1 = part1(test))")
# println("Part 2: $(result2 = part2(test))")

@assert result1 == "53248976"
# @assert result2 ==