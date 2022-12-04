require "common"

local example1 = [[
2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8
]]

Range = {}

function Range:new(lower, upper)
    local range = { lower = lower, upper = upper }
    setmetatable(range, self)
    self.__index = self
    return range
end

function Range:size()
    return self.upper - self.lower + 1
end

function Range:contains(other)
    return self.lower <= other.lower
        and self.upper >= other.upper
end

function Range:has(value)
    return self.lower <= value
        and self.upper >= value
end

local function parse(line)
    local l1, l2, u1, u2
    local r1, r2
    l1, u1, l2, u2 = string.match(line, "(%d+)-(%d+),(%d+)-(%d+)")
    r1 = Range:new(tonumber(l1), tonumber(u1))
    r2 = Range:new(tonumber(l2), tonumber(u2))
    return r1, r2
end

function problem1(input)
    local count = 0
    local r1, r2
    for line in lines(input) do
        if line ~= "" then
            r1, r2 = parse(line)
            if r1:contains(r2) or r2:contains(r1) then
                count = count + 1
            end
        end
    end
    return count
end

assert(problem1(example1) == 2)
print(problem1(readfile("data/day-04.txt")))

function problem2(input)
    local count = 0
    local r1, r2
    for line in lines(input) do
        if line ~= "" then
            r1, r2 = parse(line)
            if r1:has(r2.lower) or r1:has(r2.upper) or r2:has(r1.lower) or r2:has(r1.upper) then
                count = count + 1
            end
        end
    end
    return count
end

assert(problem2(example1) == 4)
print(problem2(readfile("data/day-04.txt")))
