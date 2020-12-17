
# https://adventofcode.com/2020/day/10

# Each of your joltage adapters is rated for a specific output joltage (your puzzle input).
# Any given adapter can take an input 1, 2, or 3 jolts lower than its rating and still produce
# its rated output joltage.

# In addition, your device has a built-in joltage adapter rated for 3 jolts higher than the
# highest-rated adapter in your bag. (If your adapter list were 3, 9, and 6, your device's
# built-in adapter would be rated for 12 jolts.)

# Treat the charging outlet near your seat as having an effective joltage rating of 0.

function read_array(string::AbstractString)::Vector{Int}
    return map(x -> parse(Int, x), split(string, '\n', keepempty=false))
end

example1 = read_array("
16
10
15
5
1
11
7
19
6
12
4
")

example2 = read_array("
28
33
18
42
31
14
46
20
48
47
24
23
49
45
19
38
39
11
1
32
25
35
8
17
7
9
4
2
34
10
3
")

# What is the number of 1-jolt differences multiplied by the number of 3-jolt differences?

function part1(numbers)
    sorted = sort(numbers)
    sorted = [0; sorted; sorted[end] + 3]
    differences = diff(sorted)
    @assert minimum(differences) == 1
    @assert maximum(differences) == 3

    return sum(differences .== 1) * sum(differences .== 3)
end

@assert part1(example2) == 220

# What is the total number of distinct ways you can arrange the adapters to connect
# the charging outlet to your device?

"""
Cache for holding the partial solution, it maps: node => solutions_count.
"""
Cache = Dict{Int,Int}

"""
Use memoized shortcuts when traversing the tree. On the way, update the cache
in place. Returns the number of paths from the beggining of the sorted list
of nodes.
"""
function count_solutions!(adapters, cache::Cache=Cache())
    socket = adapters[1]
    if socket in keys(cache)
        return cache[socket]
    end

    solutions_found = 0
    for i = 2:4
        if i > length(adapters)
            break
        end
        if adapters[i] <= (socket + 3)
            if i == length(adapters)
                solutions_found += 1
            else
                # updates cache in place, will be available for next iteration
                solutions_found += count_solutions!(adapters[i:end], cache)
            end
        end
    end
    cache[socket] = solutions_found
    return solutions_found
end

@assert count_solutions!([0; sort(example1); 22], Cache()) == 8
@assert count_solutions!([0; sort(example2); 52], Cache()) == 19208

"""
Traverse the tree starting from the back, to build the memoization
cache.
"""
function init_cache(adapters, step_size=0.05)
    n = length(adapters)
    step = Int(max(1, round(n * step_size)))
    cache = Cache()

    pos = n
    while pos >= n / 3
        pos -= step
        count_solutions!(adapters[pos:end], cache)
    end
    return cache
end

function part2(numbers)
    sorted = sort(numbers)
    full_sequence = [0; sorted; sorted[end] + 3]
    cache = init_cache(full_sequence)
    return count_solutions!(full_sequence, cache)
end

@assert part2(example1) == 8
@assert part2(example2) == 19208

test = read_array(read("data/day-10.txt", String))
println("Part 1: $(part1(test))")
println("Part 2: $(part2(test))")

part1(test) == 2738
part2(test) == 74049191673856

println()
print("Example 1: ")
@time @assert part2(example1) == 8
print("Example 2: ")
@time @assert part2(example2) == 19208
print("Part 2:    ")
@time @assert part2(test) == 74049191673856
