
# https://adventofcode.com/2020/day/10

function read_array(string::AbstractString)::Vector{Int}
    return map(x -> parse(Int, x), filter(x -> x != "", split(string, '\n')))
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

function part1(numbers)
    sorted = sort(numbers)
    sorted = [0; sorted; sorted[end] + 3]
    differences = diff(sorted)
    @assert minimum(differences) == 1
    @assert maximum(differences) == 3

    return sum(differences .== 1) * sum(differences .== 3)
end

@assert part1(example2) == 220

"""
Traverse the tree using recurrsion directly.
"""
function count_solutions(adapters)
    solutions_found = 0
    for i in 2:4
        if i > length(adapters)
            break
        end
        if adapters[i] <= (adapters[1] + 3)
            if i == length(adapters)
                solutions_found += 1
            else
                solutions_found += count_solutions(adapters[i:end])
            end
        end
    end
    return solutions_found
end

@time @assert count_solutions([0; sort(example1); 22]) == 8
@time @assert count_solutions([0; sort(example2); 52]) == 19208

"""
Cache for holding the partial solution, it maps: node => solutions_count.
"""
Cache = Dict{Int, Int}

"""
Use memoized shortcuts when traversing the tree. On the way, update the cache
in place. Returns the number of paths from the beggining of the sorted list
of nodes.
"""
function count_solutions!(adapters, cache::Cache = Cache())
    socket = adapters[1]
    if socket in keys(cache)
        return cache[socket]
    end

    solutions_found = 0
    for i in 2:4
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

@time @assert count_solutions!([0; sort(example1); 22], Cache()) == 8
@time @assert count_solutions!([0; sort(example2); 52], Cache()) == 19208

"""
Traverse the tree starting from the back, to build the memoization
cache.
"""
function init_cache(adapters, step=0.05)
    n = length(adapters)
    step = Int(max(1, round(n * step)))
    cache = Cache()

    pos = n
    while pos >= div(n, 3)
        pos = pos - step
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

@time @assert part2(example1) == 8
@time @assert part2(example2) == 19208

test = read_array(read("data/day-10.txt", String))
println("Part 1: $(part1(test))")
println("Part 2: $(part2(test))")

@assert part1(test) == 2738
@assert part2(test) == 74049191673856
