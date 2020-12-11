# https://adventofcode.com/2020/day/1

# find the two entries that sum to 2020 and then multiply those two numbers together

function part1(list_of_numbers, total = 2020)
  for (i, x) in enumerate(list_of_numbers)
    y = total - x
    if y in list_of_numbers[i+1:end]
      return x * y
    end
  end
end

example = [1721, 979, 366, 299, 675, 1456]
@assert part1(example) == 514579

# Using the above example again, the three entries that sum to 2020 are
# 979, 366, and 675. Multiplying them together produces the answer, 241861950.

function part2(list_of_numbers)
  for (i, x) in enumerate(list_of_numbers)
    partial = 2020 - x
    tmp = filter(y -> y <= partial, list_of_numbers[i+1:end])
    y = part1(tmp, partial)
    if y !== nothing
      return x * y
    end
  end
end

@assert part2(example) == 241861950

function read_array(string::AbstractString)::Vector{Int}
  return map(x -> parse(Int, x), filter(x -> x != "", split(string, '\n')))
end

test = read_array(read("data/day-01.txt", String))
println("Part 1: $(part1(test))")
println("Part 2: $(part2(test))")

@assert part1(test) == 326211
@assert part2(test) == 131347190
