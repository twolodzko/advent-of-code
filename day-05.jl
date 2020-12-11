# https://adventofcode.com/2020/day/5

# Instead of zones or groups, this airline uses binary space partitioning to seat people.
# A seat might be specified like FBFBBFFRLR, where F means "front", B means "back", L
# means "left", and R means "right".
#
# The first 7 characters will either be F or B; these specify exactly one of the 128 rows
# on the plane (numbered 0 through 127). Each letter tells you which half of a region the
# given seat is in. Start with the whole list of rows; the first letter indicates whether
# the seat is in the front (0 through 63) or the back (64 through 127). The next letter
# indicates which half of that region the seat is in, and so on until you're left with
# exactly one row.
#
# For example, consider just the first seven characters of FBFBBFFRLR:
#
#   Start by considering the whole range, rows 0 through 127.
#   F means to take the lower half, keeping rows 0 through 63.
#   B means to take the upper half, keeping rows 32 through 63.
#   F means to take the lower half, keeping rows 32 through 47.
#   B means to take the upper half, keeping rows 40 through 47.
#   B keeps rows 44 through 47.
#   F keeps rows 44 through 45.
#   The final F keeps the lower of the two, row 44.

function code_to_int(code)
    code = replace(code, r"F|L" => '0')
    code = replace(code, r"B|R" => '1')
    return parse(Int, code, base=2)
end

@assert code_to_int("BFFFBBFRRR") == 567
@assert code_to_int("FFFBBBFRRR") == 119
@assert code_to_int("BBFFBBFRLL") == 820
@assert code_to_int("FFFFFFFFLL") == 0
@assert code_to_int("BBBBBBBBRR") == 1023

function part1(codes)
    highest = 0
    for code in split(codes, '\n')
        if code == ""
            continue
        end
        highest = max(highest, code_to_int(code))
    end
    return highest
end

example = "
FFFBBFFFLL
BFFBBFBBRR
BBBBBBBBRR
FFFFFFFFLL
"

@assert part1(example) == 1023

function find_missing(numbers)
    return findfirst(diff(numbers) .> 1)
end

@assert find_missing([1, 2, 3, 4, 6, 7, 8]) == 4

function part2(codes)
    seat_numbers = sort([code_to_int(code) for code in split(codes, '\n') if code != ""])
    idx = find_missing(seat_numbers)
    return seat_numbers[idx] + 1
end

test = read("data/day-05.txt", String)
println("Part 1: $(part1(test))")
println("Part 2: $(part2(test))")

@assert part1(test) == 913
@assert part2(test) == 717
