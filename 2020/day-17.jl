# https://adventofcode.com/2020/day/17

# Each cube only ever considers its neighbors: any of the 26 other cubes where any of their
# coordinates differ by at most 1. For example, given the cube at x=1,y=2,z=3, its neighbors
# include the cube at x=2,y=2,z=2, the cube at x=0,y=2,z=3, and so on.
#
# During a cycle, all cubes simultaneously change their state according to the following rules:
#
#  * If a cube is active and exactly 2 or 3 of its neighbors are also active, the cube remains
#    active. Otherwise, the cube becomes inactive.
#  * If a cube is inactive but exactly 3 of its neighbors are active, the cube becomes active.
#    Otherwise, the cube remains inactive.

import Base: ==, iterate, length

example = "
.#.
..#
###
"

Coords = Tuple{Vararg{Int}}

mutable struct Grid
    coords::Vector{Coords}
end

(==)(x::Grid, y::Grid) = all(x.coords .== y.coords)
iterate(grid::Grid, i...) = iterate(grid.coords, i...)
length(grid::Grid) = length(grid.coords)

function read_init(string::AbstractString, other_dims::Tuple{Vararg{Int}}=(1,))
    rows = split(string, '\n', keepempty=false)
    coords = Coords[]
    for (i, row) in enumerate(rows)
        for (j, ch) in enumerate(collect(row))
            if ch == '#'
                # using rotated coordinates, so they are in same order
                # as when looking at the ASCII examples
                push!(coords, (other_dims..., i, j))
            end
        end
    end
    return Grid(coords)
end

@assert read_init(example) == Grid([(1, 1, 2), (1, 2, 3), (1, 3, 1), (1, 3, 2), (1, 3, 3)])
@assert length(read_init(example)) == 5

function isneighbour(x::Coords, y::Coords)
    return !all(x .== y) && all(abs.(x .- y) .<= 1)
end

@assert isneighbour((1, 2, 3), (2, 2, 2))
@assert isneighbour((1, 2, 3), (0, 2, 3))
@assert !isneighbour((1, 2, 3), (3, 4, 1))

function find_neighbours(point::Coords, grid::Grid)
    return filter(x -> isneighbour(point, x), grid.coords)
end

@assert find_neighbours((1, 1, 1), read_init(example)) == Coords[(1, 1, 2)]

function cycle!(grid::Grid)
    new_coords = Coords[]
    for point in grid.coords
        # explore neighbourhood of every activated point
        for neighbourhood in Iterators.product([-1:1 for _ = 1:length(point)]...)
            new_point = point .+ neighbourhood
            if new_point in new_coords
                continue
            end
            neighbours_count = length(find_neighbours(new_point, grid))
            if new_point in grid.coords
                if 2 <= neighbours_count <= 3
                    push!(new_coords, new_point)
                end
            else
                if neighbours_count == 3
                    push!(new_coords, new_point)
                end
            end
        end
    end
    grid.coords = new_coords
    return grid
end

function boot!(init::Grid, num_cycles=6)
    for _ = 1:num_cycles
        cycle!(init)
    end
    return init
end

@assert length(cycle!(read_init(example))) == 11
@assert length(cycle!(cycle!(read_init(example)))) == 21

function part1(input)
    init = read_init(input)
    return length(boot!(init))
end

@assert part1(example) == 112

function part2(input)
    init = read_init(input, (1, 1))
    return length(boot!(init))
end

@assert part2(example) == 848

test = read("data/day-17.txt", String)
println("Part 1: $(result1 = part1(test))")
println("Part 2: $(result2 = part2(test))")

@assert result1 == 304
@assert result2 == 1868
