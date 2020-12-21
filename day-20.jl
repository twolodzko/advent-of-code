# https://adventofcode.com/2020/day/20

function parse_input(input)
    tiles = Dict()
    for chunk in split(input, "\n\n", keepempty=false)
        rows = split(chunk, '\n', keepempty=false)
        id = parse(Int, match(r"Tile (\d+):", popfirst!(rows)).captures[1])
        tile = zeros(Bool, (length(rows), length(rows[2])))
        for (i, row) in enumerate(rows)
            tile[i, :] .= collect(row) .== '#'
        end
        tiles[id] = tile
    end
    return tiles
end

"""
Sides:
```
     1
   +---+
 4 |   | 2
   +---+
     3
```
"""
function get_sides(tiles)
    sides = Dict()
    for (id, tile) in tiles
        sides[id] = [
            tile[1, :][:],
            tile[:, end][:],
            tile[end, :][:],
            tile[:, 1][:],
        ]
    end
    return sides
end

function find_matches(tiles)
    sides = get_sides(tiles)
    matches = Dict()
    for (x_id, x_sides) in sides
        for (y_id, y_sides) in sides
            if x_id == y_id
                continue
            end

            for i in 1:4
                for j in 1:4
                    reversed = all(x_sides[i] .== reverse(y_sides[j]))
                    if all(x_sides[i] .== y_sides[j]) || reversed
                        if x_id in keys(matches)
                            push!(matches[x_id], (y_id, i, reversed))
                        else
                            matches[x_id] = [(y_id, i, reversed)]
                        end
                    end
                end
            end
        end
    end
    return matches
end

function part1(input)
    tiles = parse_input(input)
    matches = find_matches(tiles)
    corners = [k for (k, v) in matches if length(v) == 2]
    @assert length(corners) == 4
    return (*)(corners...)
end

example = read("data/day-20-example.txt", String)

@assert part1(example) == 20899048083289

function flip(arr::Array{T,2}; dims=1) where {T}
    n, k = size(arr)
    out = Array{T,2}(undef, n, k)
    if dims == 1
        for i in 1:n
            out[n - i + 1, :] = arr[i, :]
        end
    elseif dims == 2
        for j in 1:k
            out[:, k - j + 1, :] = arr[:, j]
        end
    else
        return flip(flip(arr, dims=1), dims=2)
    end
    return out
end

@assert flip([1 2; 3 4], dims=1) == [3 4; 1 2]
@assert flip([1 2; 3 4], dims=2) == [2 1; 4 3]
@assert flip([1 2; 3 4], dims=(1, 2)) == [4 3; 2 1]

test = read("data/day-20.txt", String)
println("Part 1: $(result1 = part1(test))")
# println("Part 2: $(result2 = part2(test))")

@assert result1 == 107399567124539
# @assert result2 ==