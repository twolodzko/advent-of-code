# https://adventofcode.com/2020/day/24

const sides = Symbol[:N, :NE, :SE, :S, :SW, :NW]

function parse_line(line::AbstractString)
    coords = Symbol[]

    i = 1
    while i <= length(line)
        d = string(line[i])
        if d in ("s", "n")
            i += 1
            d *= line[i]
        end
        push!(coords, Symbol(uppercase(d)))
        i += 1
    end

    return coords
end

@assert all(parse_line("esenee") .== [:E, :SE, :NE, :E])

example = "
sesenwnenenewseeswwswswwnenewsewsw
neeenesenwnwwswnenewnwwsewnenwseswesw
seswneswswsenwwnwse
nwnwneseeswswnenewneswwnewseswneseene
swweswneswnenwsewnwneneseenw
eesenwseswswnenwswnwnwsewwnwsene
sewnenenenesenwsewnenwwwse
wenwwweseeeweswwwnwwe
wsweesenenewnwwnwsenewsenwwsesesenwne
neeswseenwwswnwswswnw
nenwswwsewswnenenewsenwsenwnesesenew
enewnwewneswsewnwswenweswnenwsenwsw
sweneswneswneneenwnewenewwneswswnese
swwesenesewenwneswnwwneseswwne
enesenwswwswneneswsenwnewswseenwsese
wnwnesenesenenwwnenwsewesewsesesew
nenewswnwewswnenesenwnesewesw
eneswnwswnwsenenwnwnwwseeswneewsenese
neswnwewnwnwseenwseesewsenwsweewe
wseweeenwnesenwwwswnew
"

function part1(input)
    
    for row in split(input, '\n', keepempty=false)
        coords = parse_line(row)
    end

    return 
end

@assert part1(example) == 10



# test = 
# println("Part 1: $(result1 = part1(test))")
# println("Part 2: $(result2 = part2(test))")

# @assert result1 == 
# @assert result2 == 