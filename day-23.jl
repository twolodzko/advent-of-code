# https://adventofcode.com/2020/day/23

using ProgressMeter
import Base: peek

example = "389125467"

movefirstback!(arr) = push!(arr, popfirst!(arr))

@assert all(movefirstback!([1, 2, 3, 4]) .== [2, 3, 4, 1])

function move(cups::Vector{Int}, pos, max; verbose=false)
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
    destination = destination < 1 ? max : destination
    while destination in picked
        destination -= 1
        destination = destination < 1 ? max : destination
    end
    verbose && println("destination: ", destination)

    # insert picked cups back
    d = findfirst(cups .== destination)
    for i in 1:3
        insert!(cups, d + i, picked[i])
    end

    # rotate to the previous position
    while findnext(cups .== current, pos) > pos
        movefirstback!(cups)
    end

    return cups
end

@assert all(move([3, 8, 9, 1, 2, 5, 4, 6, 7], 1, 9) .== [3, 2, 8, 9, 1, 5, 4, 6, 7])
@assert all(move([3, 2, 5, 4, 6, 7, 8, 9, 1], 3, 9) .== [7, 2, 5, 8, 9, 1, 3, 4, 6])
@assert all(move([7, 4, 1, 5, 8, 3, 9, 2, 6], 9, 9) .== [5, 7, 4, 1, 8, 3, 9, 2, 6])

function play(cups::Vector{Int}, rounds::Int; verbose=false)
    n = length(cups)
    max = maximum(cups)

    for i in 1:rounds
        verbose && println("-- move $i --")
        pos = mod1(i, n)
        cups = move(cups, pos, max, verbose=verbose)
        verbose && println()
    end
    return cups
end

@assert all(play(Int[3, 8, 9, 1, 2, 5, 4, 6, 7], 10) .== [5, 8, 3, 7, 4, 1, 9, 2, 6])

mutable struct LinkedList
    value::Any
    next::Union{LinkedList,Nothing}
end

function peek(list::LinkedList, n::Integer)
    out = Array{Int}(undef, n)
    for i in 1:n
        out[i] = list.value
        list = list.next
    end
    return out
end

function insertnext!(list::LinkedList, arr)
    head = list
    tail = list.next
    for x in arr
        head.next = LinkedList(x, nothing)
        head = head.next
    end
    head.next = tail
    return list
end

function play_ll(cups, rounds::Int)
    max = maximum(cups)

    cups = reverse(cups)
    tail = LinkedList(cups[1], nothing)
    head = tail
    for cup in cups
        head = LinkedList(cup, head)
    end
    tail.next = head

    for i in 1:rounds
        current = head.value

        picked = peek(head.next, 3)
        head.next = head.next.next.next.next

        destination = current - 1
        destination = destination < 1 ? max : destination
        while destination in picked
            destination -= 1
            destination = destination < 1 ? max : destination
            println(destination)
        end

        pos = head.next
        while pos.value != destination
            pos = pos.next
        end

        insertnext!(pos, picked)

        head = head.next
    end

    return peek(head, length(cups))
end






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

    @showprogress for i in 1:rounds
        verbose && println("-- move $i --")
        pos = mod1(i, n)
        cups = move(cups, pos, size, verbose=verbose)
        verbose && println()
    end

    pos = findfirst(cups .== 1)
    return (*)([cups[mod1(pos + i, n)] for i in 1:2]...)
end

# @assert part2(example) == 149245887792

test = "784235916"
println("Part 1: $(result1 = part1(test))")
# println("Part 2: $(result2 = part2(test))")

@assert result1 == "53248976"
# @assert result2 ==