
# https://adventofcode.com/2020/day/10

using Distributed

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

function parallel_count_solutions(adapters)
    adapters_left = length(adapters)
    if adapters_left == 0
        return 0
    end
    return @distributed (+) for i in 2:min(4, adapters_left)
        if adapters[i] <= (adapters[1] + 3)
            if i == adapters_left
                1
            else
                parallel_count_solutions(adapters[i:end])
            end
        else
            0
        end
    end
end

@time @assert parallel_count_solutions([0; sort(example1); 22]) == 8
@time @assert parallel_count_solutions([0; sort(example2); 52]) == 19208

function count_solutions(adapters, cache::Dict{Int, Int} = Dict{Int, Int}())
    solutions_found = 0
    for i in 2:4
        if i > length(adapters)
            break
        end
        println("+$(i) => $(adapters[i]) \t $(cache)")
        if adapters[i] in keys(cache)
            solutions_found += cache[adapters[i]]
        elseif adapters[i] <= (adapters[1] + 3)
            if i == length(adapters)
                solutions_found += 1
            else
                partial_count = count_solutions(adapters[i:end], cache)
                cache[adapters[i]] = partial_count
                solutions_found += partial_count
            end
        end
    end
    return solutions_found
end

@time @assert count_solutions([0; sort(example1); 22], Dict{Int, Int}()) == 8
@time @assert count_solutions([0; sort(example2); 52], Dict{Int, Int}()) == 19208

function init_cache(adapters)
    n = length(adapters)
    cache = Dict{Int, Int}()
    for i in 20:-2:2
        pos = n - div(n, i)
        if pos == n
            continue
        end
        count_solutions(adapters[pos:end], cache)
    end
    return cache
end

function part2(numbers)
    sorted = sort(numbers)
    sorted = [0; sorted; sorted[end] + 3]
    cache = init_cache(sorted)
    return count_solutions(sorted, cache)
end

@time @assert part2(example1) == 8
@time @assert part2(example2) == 19208

test = read_array(read("data/day-10.txt", String))
println("Part 1: $(part1(test))")
# println("Part 2: $(part2(test))")

@assert part1(test) == 2738
# @assert part2(test) == 3340942
