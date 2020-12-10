
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

function count_solutions(adapters, start=0)
    solutions_found = 0
    for i in 1:3
        if i > length(adapters)
            break
        end
        if adapters[i] <= (start + 3)
            if i == length(adapters)
                solutions_found += 1
            else
                solutions_found += count_solutions(adapters[(i + 1):end], adapters[i])
            end
        end
    end
    return solutions_found
end

@time @assert count_solutions([sort(example1); 22]) == 8
@time @assert count_solutions([sort(example2); 52]) == 19208

function parallel_count_solutions(adapters, start=0)
    adapters_left = length(adapters)
    if adapters_left == 0
        return 0
    end
    return @distributed (+) for i in 1:min(3, adapters_left)
        if adapters[i] <= (start + 3)
            if i == adapters_left
                1
            else
                parallel_count_solutions(adapters[(i + 1):end], adapters[i])
            end
        else
            0
        end
    end
end

@time @assert parallel_count_solutions([sort(example1); 22]) == 8
@time @assert parallel_count_solutions([sort(example2); 52]) == 19208

function part2(numbers)
    sorted = sort(numbers)
    return count_solutions([sorted; sorted[end] + 3])
end

test = read_array(read("data/day-10.txt", String))
println("Part 1: $(part1(test))")
# println("Part 2: $(part2(test))")

@assert part1(test) == 2738
# @assert part2(test) == 3340942
