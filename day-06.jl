
# https://adventofcode.com/2020/day/6

function part1(input)
  groups = map(x -> Set(replace(x, r"[^a-z]" => "")), split(input, "\n\n"))
  return sum(map(length, groups))
end

example = "
abc

a
b
c

ab
ac

a
a
a
a

b
"

@assert part1(example) == 11

function part2(input)
  groups = split(input, "\n\n")
  groups = map(s -> filter(ch -> ch != "", split(s, '\n')), groups)
  common_items = map(g -> reduce(intersect, g), groups)
  return sum(map(length, common_items))
end

@assert part2(example) == 6

test = read("data/day-06.txt", String)
println("Part 1: $(part1(test))")
println("Part 2: $(part2(test))")

@assert part1(test) == 6530
@assert part2(test) == 3323
