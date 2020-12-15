# https://adventofcode.com/2020/day/3

# You start on the open square (.) in the top-left corner and need to reach the
# bottom (below the bottom-most row on your map).

# The toboggan can only follow a few specific slopes (you opted for a cheaper model
# that prefers rational numbers); start by counting all the trees you would encounter
# for the slope right 3, down 1:

# From your starting position at the top-left, check the position that is right 3
# and down 1. Then, check the position that is right 3 and down 1 from there, and
# so on until you go past the bottom of the map.

function replace_at(string, ind, char)
    tmp = collect(string)
    tmp[ind] = char
    return String(tmp)
end

function part1(patch, right, down = 1)
    trees_count = 0
    position = 1
    time_to_move = 1

    for row in split(patch, '\n', keepempty = false)
        if time_to_move == 1
            if row[position] == '#'
                trees_count += 1
            end
            position = mod1(position + right, length(row))
            time_to_move = down
        else
            time_to_move -= 1
        end
    end
    return trees_count
end

example = "
..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#
"

@assert part1(example, 3) == 7

# Determine the number of trees you would encounter if, for each of the following
# slopes, you start at the top-left corner and traverse the map all the way
# to the bottom:
#
#  Right 1, down 1.
#  Right 3, down 1. (This is the slope you already checked.)
#  Right 5, down 1.
#  Right 7, down 1.
#  Right 1, down 2.
#
# In the above example, these slopes would find 2, 7, 3, 4, and 2 tree(s) respectively;
# multiplied together, these produce the answer 336.

function part2(patch, right, down)
    result = 1
    for (r, d) in zip(right, down)
        result *= part1(patch, r, d)
    end
    return result
end

@assert part2(example, [1 3 5 7 1], [1 1 1 1 2]) == 336

test = read("data/day-03.txt", String)
println("Part 1: $(part1(test, 3))")
println("Part 2: $(part2(test, [1 3 5 7 1], [1 1 1 1 2]))")

@assert part1(test, 3) == 203
@assert part2(test, [1 3 5 7 1], [1 1 1 1 2]) == 3316272960
