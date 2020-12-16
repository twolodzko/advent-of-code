# https://adventofcode.com/2020/day/13

example = "
939
7,13,x,x,59,x,31,19
"

function read_input(string)
    earliest_timestamp, bus_numbers = split(string, '\n', keepempty = false)
    earliest_timestamp = parse(Int, earliest_timestamp)
    bus_numbers = map(x -> tryparse(Int, x), split(bus_numbers, ','))
    return earliest_timestamp, bus_numbers
end

@assert read_input(example) == (939, [7, 13, nothing, nothing, 59, nothing, 31, 19])

function closest_arrival(bus_number, timestamp)
    divisor, reminder = divrem(timestamp, bus_number)
    if reminder == 0
        return 0
    else
        return bus_number * (divisor + 1) - timestamp
    end
end

function part1(input)
    earliest_timestamp, bus_numbers = read_input(input)
    bus_numbers = convert(Vector{Int}, filter(x -> !isnothing(x), bus_numbers))
    closest_arrivals = closest_arrival.(bus_numbers, earliest_timestamp)
    i = argmin(closest_arrivals)
    return closest_arrivals[i] * bus_numbers[i]
end

@assert part1(example) == 295

"""
Solve Chinese reminder theorem problem using the inverse modulo algorithm

See:
* https://en.wikipedia.org/wiki/Chinese_remainder_theorem
* https://www.geeksforgeeks.org/chinese-remainder-theorem-set-2-implementation/
* https://rosettacode.org/wiki/Chinese_remainder_theorem
"""
function chinese_reminder(modulus, reminder)
    N = prod(modulus)
    Ni = div.(N, modulus)
    i = invmod.(Ni, modulus)
    return mod(sum(reminder .* Ni .* i), N)
end

@assert chinese_reminder([3, 4, 5], [2, 3, 1]) == 11

function part2(input)
    _, bus_numbers = read_input(input)

    nonmissing = map(x -> !isnothing(x), bus_numbers)
    indexes = collect(0:length(bus_numbers)-1)
    indexes = indexes[nonmissing]
    bus_numbers = bus_numbers[nonmissing]
    reminders = maximum(indexes) .- indexes

    return chinese_reminder(bus_numbers, reminders) - maximum(reminders)
end

@assert part2(example) == 1068781
@assert part2("0\n17,x,13,19") == 3417
@assert part2("0\n67,7,59,61") == 754018
@assert part2("0\n67,x,7,59,61") == 779210
@assert part2("0\n67,7,x,59,61") == 1261476
@assert part2("0\n1789,37,47,1889") == 1202161486

test = read("data/day-13.txt", String)
println("Part 1: $(part1(test))")
println("Part 2: $(part2(test))")

@assert part1(test) == 4207
@assert part2(test) == 725850285300475
