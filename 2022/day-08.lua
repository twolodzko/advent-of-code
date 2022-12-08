require "common"

local example1 = [[
30373
25512
65332
33549
35390
]]

local function parse(input)
    local i = 1
    local data = {}
    for line in lines(input) do
        if line ~= "" then
            table.insert(data, {})
            for char in chars(line) do
                table.insert(data[i], tonumber(char))
            end
            i = i + 1
        end
    end
    return data
end

assert(parse(example1)[1][1] == 3)
assert(parse(example1)[2][3] == 5)
assert(parse(example1)[5][4] == 9)
assert(parse(example1)[2][6] == nil)
assert(parse(example1)[6] == nil)

local function isblocked(map, row, col, d)
    local i = row + d.i
    local j = col + d.j
    local height = map[row][col]
    while map[i] and map[i][j] do
        if map[i][j] >= height then
            return true
        end
        i = i + d.i
        j = j + d.j
    end
    return false
end

local function isvisible(map, row, col)
    local ds = {
        { i = 0, j = -1 },
        { i = 0, j = 1 },
        { i = -1, j = 0 },
        { i = 1, j = 0 }
    }
    for _, d in ipairs(ds) do
        if not isblocked(map, row, col, d) then
            return true
        end
    end
    return false
end

do
    local map = assert(parse(example1))
    assert(isvisible(map, 1, 1) == true)
    assert(isvisible(map, 1, 5) == true)
    assert(isvisible(map, 5, 5) == true)
    assert(isvisible(map, 3, 5) == true)

    assert(isvisible(map, 2, 2) == true)
    assert(isvisible(map, 2, 3) == true)
    assert(isvisible(map, 2, 4) == false)
    assert(isvisible(map, 3, 2) == true)
    assert(isvisible(map, 3, 3) == false)
    assert(isvisible(map, 3, 4) == true)
    assert(isvisible(map, 4, 2) == false)
    assert(isvisible(map, 4, 3) == true)
    assert(isvisible(map, 4, 4) == false)
end

local function countvisible(map)
    local count = 0
    for i = 1, #map do
        for j = 1, #map[1] do
            if isvisible(map, i, j) then
                count = count + 1
            end
        end
    end
    return count
end

function problem1(input)
    local map = parse(input)
    return countvisible(map)
end

assert(problem1(example1) == 21)
print(problem1(readfile("data/day-08.txt")))

local function visibletrees(map, row, col, d)
    local count = 0
    local i = row + d.i
    local j = col + d.j
    local height = map[row][col]
    while map[i] and map[i][j] do
        count = count + 1
        if map[i][j] >= height then
            return count
        end
        i = i + d.i
        j = j + d.j
    end
    return count
end

local function scenicscore(map, row, col)
    local score = 1
    local ds = {
        { i = 0, j = -1 },
        { i = 0, j = 1 },
        { i = -1, j = 0 },
        { i = 1, j = 0 }
    }
    for _, d in ipairs(ds) do
        score = score * visibletrees(map, row, col, d)
    end
    return score
end

do
    local map = assert(parse(example1))
    assert(scenicscore(map, 2, 3) == 4)
    assert(scenicscore(map, 4, 3) == 8)
end

local function bestscenicscore(map)
    local best = 0
    local score
    for i = 1, #map do
        for j = 1, #map[1] do
            score = scenicscore(map, i, j)
            if score > best then
                best = score
            end
        end
    end
    return best
end

function problem2(input)
    local map = parse(input)
    return bestscenicscore(map)
end

assert(problem2(example1) == 8)
print(problem2(readfile("data/day-08.txt")))
