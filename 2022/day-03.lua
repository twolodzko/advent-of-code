require "common"

local example1 = [[
vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw
]]

local function split(str)
    return str:sub(1, #str / 2), str:sub(#str / 2 + 1, #str)
end

local function getrepeated(str)
    local head, tail = split(str)
    local seen = {}
    for char in chars(head) do
        seen[char] = true
    end
    for char in chars(tail) do
        if seen[char] then
            return char
        end
    end
end

assert(getrepeated("vJrwpWtwJgWrhcsFMMfFFhFp"), "p")

local function score(chr)
    if chr == string.lower(chr) then
        return string.byte(chr) - 96
    else
        return string.byte(chr) - 38
    end
end

assert(score("a"), 1)
assert(score("z"), 26)
assert(score("A"), 27)
assert(score("Z"), 52)

function problem1(input)
    local acc = 0
    for line in lines(input) do
        if line ~= "" then
            local chr = getrepeated(line)
            acc = acc + score(chr)
        end
    end
    return acc
end

assert(problem1(example1) == 157)
print(problem1(readfile("data/day-03.txt")))

local function repeatedingroup(group)
    local seen1 = {}
    local seen2 = {}
    for char in chars(group[1]) do
        seen1[char] = true
    end
    for char in chars(group[2]) do
        seen2[char] = true
    end
    for char in chars(group[3]) do
        if seen1[char] and seen2[char] then
            return char
        end
    end
end

function problem2(input)
    local acc = 0
    local group = {}
    for line in lines(input) do
        if line ~= "" then
            table.insert(group, line)
            if #group == 3 then
                local chr = repeatedingroup(group)
                acc = acc + score(chr)
                group = {}
            end
        end
    end
    return acc
end

assert(problem2(example1) == 70)
print(problem2(readfile("data/day-03.txt")))
