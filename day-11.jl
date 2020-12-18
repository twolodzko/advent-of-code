
# https://adventofcode.com/2020/day/11

# The following rules are applied to every seat simultaneously:
#
#   * If a seat is empty (L) and there are no occupied seats adjacent to it, the seat
#     becomes occupied.
#   * If a seat is occupied (#) and four or more seats adjacent to it are also occupied,
#     the seat becomes empty.
#   * Otherwise, the seat's state does not change.
#
# Floor (.) never changes; seats don't move, and nobody sits on the floor.

function decode_seat(char::Char)
    if char == 'L'
        return :Empty
    elseif char == '.'
        return :Floor
    elseif char == '#'
        return :Occupied
    else
        error("invalid input: $(char)")
    end
end

function read_layout(string::AbstractString)
    rows = map(strip, split(string, '\n', keepempty=false))
    layout = Array{Symbol,2}(undef, length(rows), length(rows[1]))

    for (i, row) in enumerate(rows)
        for (j, char) in enumerate(row)
            layout[i, j] = decode_seat(char)
        end
    end
    return layout
end

@assert read_layout("L.#\n..L\n##.") == [:Empty :Floor :Occupied; :Floor :Floor :Empty; :Occupied :Occupied :Floor]

function count_adjecent_occupied_seats(layout, x, y)
    n, k = size(layout)
    occupied_seats = 0
    for i = (x - 1):(x + 1)
        if i < 1 || i > n
            continue
        end
        for j = (y - 1):(y + 1)
            if i == x && j == y || j < 1 || j > k
                continue
            end
            if layout[i, j] == :Occupied
                occupied_seats += 1
            end
        end
    end
    return occupied_seats
end

@assert count_adjecent_occupied_seats(read_layout("...\n...\n..."), 2, 2) == 0
@assert count_adjecent_occupied_seats(read_layout(".#.\n..#\n#.."), 2, 2) == 3
@assert count_adjecent_occupied_seats(read_layout(".#.\n.LL\n#.."), 2, 2) == 2
@assert count_adjecent_occupied_seats(read_layout("#####\n#...#\n#...#\n#...#\n#####"), 3, 3) == 0
@assert count_adjecent_occupied_seats(read_layout(".#.\n.#.\n..#"), 1, 1) == 2
@assert count_adjecent_occupied_seats(read_layout(".#.\n.##\n..."), 1, 3) == 3

function new_state(layout, x, y)
    occupied_seats = count_adjecent_occupied_seats(layout, x, y)
    if layout[x, y] == :Empty && occupied_seats == 0
        return :Occupied
    elseif layout[x, y] == :Occupied && occupied_seats >= 4
        return :Empty
    end
    return layout[x, y]
end

@assert new_state(read_layout("...\n...\n..."), 2, 2) == :Floor
@assert new_state(read_layout("...\n.L.\n..."), 2, 2) == :Occupied
@assert new_state(read_layout("...\n.L#\n..."), 2, 2) == :Empty
@assert new_state(read_layout("...\n.#.\n..."), 2, 2) == :Occupied
@assert new_state(read_layout(".#.\n###\n.#."), 2, 2) == :Empty
@assert new_state(read_layout(".#.\n##L\n.#."), 2, 2) == :Occupied
@assert new_state(read_layout("#.#\n.#.\n#.#"), 2, 2) == :Empty
@assert new_state(read_layout("###\n###\n###"), 2, 2) == :Empty
@assert new_state(read_layout("L..\n...\n..."), 1, 1) == :Occupied
@assert new_state(read_layout("##.\n##.\n..."), 1, 1) == :Occupied

function apply_rules(layout)
    n, k = size(layout)
    updated = Array{Symbol,2}(undef, n, k)
    for i = 1:n
        for j = 1:k
            updated[i, j] = new_state(layout, i, j)
        end
    end
    return updated
end

example1 = split(read("data/day-11-example1.txt", String), "\n\n", keepempty=false)

@assert apply_rules(read_layout(example1[1])) == read_layout(example1[2])
@assert apply_rules(read_layout(example1[2])) == read_layout(example1[3])
@assert apply_rules(read_layout(example1[3])) == read_layout(example1[4])
@assert apply_rules(read_layout(example1[4])) == read_layout(example1[5])
@assert apply_rules(read_layout(example1[5])) == read_layout(example1[6])
# Equilibrium?!
@assert apply_rules(read_layout(example1[6])) == read_layout(example1[6])

# How many seats end up occupied?

function part1(string)
    layout = read_layout(string)
    while true
        updated_layout = apply_rules(layout)
        if layout == updated_layout
            break
        end
        layout = updated_layout
    end
    return sum(layout .== :Occupied)
end

@assert part1(example1[1]) == 37

# Now, instead of considering just the eight immediately adjacent seats, consider the first
# seat in each of those eight directions. (...)
# Also, people seem to be more tolerant than you expected: it now takes five or more visible
# occupied seats for an occupied seat to become empty (rather than four or more from the
# previous rules). The other rules still apply (...)

function first_visible_seat(layout, x, y, dx, dy)
    n, k = size(layout)
    while true
        x += dx
        y += dy
        try
            if layout[x, y] in (:Empty, :Occupied)
                return layout[x, y]
            end
        catch BoundsError
            # we bumped the wall
            break
        end
    end
    return :Floor
end

@assert first_visible_seat(read_layout("L..\n...\n..."), 1, 1, +1, 0) == :Floor
@assert first_visible_seat(read_layout("L..\n...\n..."), 1, 1, 0, +1) == :Floor
@assert first_visible_seat(read_layout("L..\n...\n..."), 1, 1, +1, +1) == :Floor
@assert first_visible_seat(read_layout("L#.\n...\n..."), 1, 1, 0, +1) == :Occupied
@assert first_visible_seat(read_layout("LL#\n...\n..."), 1, 1, 0, +1) == :Empty
@assert first_visible_seat(read_layout("L#L\n...\n..."), 1, 1, 0, +1) == :Occupied
@assert first_visible_seat(read_layout("L.#\n...\n..."), 1, 1, 0, +1) == :Occupied
@assert first_visible_seat(read_layout("L..\n#..\n..."), 1, 1, +1, 0) == :Occupied
@assert first_visible_seat(read_layout("L..\n#..\nL.."), 1, 1, +1, 0) == :Occupied
@assert first_visible_seat(read_layout("L..\n...\n#.."), 1, 1, +1, 0) == :Occupied
@assert first_visible_seat(read_layout("L..\n.#.\n..."), 1, 1, +1, +1) == :Occupied
@assert first_visible_seat(read_layout("L..\n.#.\n..L"), 1, 1, +1, +1) == :Occupied
@assert first_visible_seat(read_layout("L..\n...\n..#"), 1, 1, +1, +1) == :Occupied
@assert first_visible_seat(read_layout("L..\n.L.\n..#"), 1, 1, +1, +1) == :Empty
@assert first_visible_seat(read_layout("#..\n...\n..L"), 3, 3, -1, -1) == :Occupied

function count_first_visible_seats(layout, x, y)
    n, k = size(layout)
    occupied_seats = 0

    for dx in (-1, 0, +1)
        for dy in (-1, 0, +1)
            dx == dy == 0 && continue
            if first_visible_seat(layout, x, y, dx, dy) == :Occupied
                occupied_seats += 1
            end
        end
    end
    return occupied_seats
end

@assert count_first_visible_seats(read_layout("#.#..\n.....\n.#L..\n.....\n...#."), 3, 3) == 3
@assert count_first_visible_seats(read_layout("L...#\n.....\n#....\n.#...\n....."), 1, 1) == 2

no_occupied_seats_example = "
.##.##.
#.#.#.#
##...##
...L...
##...##
#.#.#.#
.##.##.
"

@assert count_first_visible_seats(read_layout(no_occupied_seats_example), 4, 4) == 0
@assert count_first_visible_seats(read_layout("#.##.##.##\n#######.##"), 1, 3) == 5

function new_state_2(layout, x, y)
    occupied_seats = count_first_visible_seats(layout, x, y)
    if layout[x, y] == :Empty && occupied_seats == 0
        return :Occupied
    elseif layout[x, y] == :Occupied && occupied_seats >= 5
        return :Empty
    end
    return layout[x, y]
end

function apply_rules_2(layout)
    n, k = size(layout)
    updated = Array{Symbol,2}(undef, n, k)
    for i = 1:n
        for j = 1:k
            updated[i, j] = new_state_2(layout, i, j)
        end
    end
    return updated
end

example2 = split(read("data/day-11-example2.txt", String), "\n\n", keepempty=false)

@assert apply_rules_2(read_layout(example2[1])) == read_layout(example2[2])
@assert apply_rules_2(read_layout(example2[2])) == read_layout(example2[3])
@assert apply_rules_2(read_layout(example2[3])) == read_layout(example2[4])
@assert apply_rules_2(read_layout(example2[4])) == read_layout(example2[5])
@assert apply_rules_2(read_layout(example2[5])) == read_layout(example2[6])
# Equilibrium?!
@assert apply_rules_2(apply_rules_2(read_layout(example2[6]))) == apply_rules_2(read_layout(example2[6]))

function part2(string)
    layout = read_layout(string)
    while true
        updated_layout = apply_rules_2(layout)
        if layout == updated_layout
            break
        end
        layout = updated_layout
    end
    return sum(layout .== :Occupied)
end

@assert part2(example2[1]) == 26

test = read("data/day-11.txt", String)
println("Part 1: $(result1 = part1(test))")
println("Part 2: $(result2 = part2(test))")

@assert result1 == 2238
@assert result2 == 2013
