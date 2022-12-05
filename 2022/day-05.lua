require "common"

local example1 = [[
    [D]
[N] [C]
[Z] [M] [P]
 1   2   3

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2
]]

local function parsecrates(crates)
    local piles = {}
    for i = 1, #crates - 1 do
        local j = 1
        while true do
            local s = string.sub(crates[i], j, j + 3)
            if s == "" then
                break
            end
            local c = (j - 1) // 4 + 1
            local m = string.match(s, "%[(%u)%]")
            if m then
                if not piles[c] then
                    piles[c] = {}
                end
                table.insert(piles[c], m)
            end
            j = j + 4
        end
    end
    for i = 1, #piles do
        piles[i] = reverse(piles[i])
    end
    return piles
end

local function parsemoves(moves)
    local parsed = {}
    for i = 1, #moves do
        local u, f, t = string.match(moves[i], "move (%d+) from (%d+) to (%d+)")
        if u then
            table.insert(parsed, { units = tonumber(u), from = tonumber(f), to = tonumber(t) })
        end
    end
    return parsed
end

local function parse(input)
    local crates = {}
    local moves = {}
    local container = crates
    for line in lines(input) do
        if line == "" then
            container = moves
        else
            table.insert(container, line)
        end

    end
    return parsecrates(crates), parsemoves(moves)
end

function problem1(input)
    local crates, moves = parse(input)
    for _, move in ipairs(moves) do
        for i = 1, move.units do
            local n = #crates[move.from]
            table.insert(crates[move.to], crates[move.from][n])
            crates[move.from][n] = nil
        end
    end
    local out = ""
    for _, pile in ipairs(crates) do
        out = out .. (pile[#pile] or "")
    end
    return out
end

assert(problem1(example1) == "CMZ")
print(problem1(readfile("data/day-05.txt")))

function problem2(input)
    local crates, moves = parse(input)
    for _, move in ipairs(moves) do
        local n = #crates[move.from]
        for i = move.units - 1, 0, -1 do
            local j = n - i
            table.insert(crates[move.to], crates[move.from][j])
            crates[move.from][j] = nil
        end
    end
    local out = ""
    for _, pile in ipairs(crates) do
        out = out .. (pile[#pile] or "")
    end
    return out
end

assert(problem2(example1) == "MCD")
print(problem2(readfile("data/day-05.txt")))
