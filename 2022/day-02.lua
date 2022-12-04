require "common"

local example1 = [[
A Y
B X
C Z
]]

local beats = { rock = "scissors", paper = "rock", scissors = "paper" }
local codes1 = { A = "rock", B = "paper", C = "scissors" }
local codes2 = { X = "rock", Y = "paper", Z = "scissors" }
local shape = { rock = 1, paper = 2, scissors = 3 }

local function outcome(m1, m2)
    if m1 == m2 then
        return 3
    elseif beats[m2] == m1 then
        return 6
    else
        return 0
    end
end

local function score(m1, m2)
    return shape[m2] + outcome(m1, m2)
end

function problem1(input)
    local p1, p2
    local acc = 0
    for line in lines(input) do
        if line ~= "" then
            p1 = codes1[line:sub(1, 1)]
            p2 = codes2[line:sub(3, 3)]
            acc = acc + score(p1, p2)
        end
    end
    return acc
end

assert(problem1(example1), 15)
print(problem1(readfile("data/day-02.txt")))

local gameend = { X = "lose", Y = "draw", Z = "win" }
local isbeatenby = { rock = "paper", paper = "scissors", scissors = "rock" }

local function strategy(m1, result)
    local m2
    if result == "win" then
        return isbeatenby[m1]
    elseif result == "lose" then
        return beats[m1]
    else
        return m1
    end
end

function problem2(input)
    local p1, p2, result
    local acc = 0
    for line in lines(input) do
        if line ~= "" then
            p1 = codes1[line:sub(1, 1)]
            result = gameend[line:sub(3, 3)]
            p2 = strategy(p1, result)
            acc = acc + score(p1, p2)
        end
    end
    return acc
end

assert(problem2(example1), 12)
print(problem2(readfile("data/day-02.txt")))
