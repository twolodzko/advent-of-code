require "common"

local example1 = [[
R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2
]]

Point = {}

function Point:new(tab)
    local point = {
        x = tab and tab[1] or 0,
        y = tab and tab[2] or 0
    }
    setmetatable(point, self)
    self.__index = self
    self.__add = function(p1, p2)
        return Point:new { p1.x + p2.x, p1.y + p2.y }
    end
    self.__sub = function(p1, p2)
        return Point:new { p1.x - p2.x, p1.y - p2.y }
    end
    self.__eq = function(p1, p2)
        return p1.x == p2.x and p1.y == p2.y
    end
    self.__tostring = function(obj)
        return string.format("Point{%d, %d}", obj.x, obj.y)
    end
    return point
end

local function sign(x)
    if x < 0 then
        return -1
    elseif x > 0 then
        return 1
    else
        return 0
    end
end

function Point:sign()
    return Point:new { sign(self.x), sign(self.y) }
end

function Point:abs()
    return Point:new { math.abs(self.x), math.abs(self.y) }
end

function Point:touches(point)
    local dif = (self - point):abs()
    return dif.x <= 1 and dif.y <= 1
end

Rope = {}

function Rope:new(head, tail)
    local rope = { head = Point:new(head), tail = Point:new(tail), history = {} }
    setmetatable(rope, self)
    self.__index = self
    self.__eq = function(r1, r2)
        return r1.head == r2.head and r1.tail == r2.tail
    end
    self.__tostring = function(obj)
        return string.format("Rope{%s, %s}", obj.head, obj.tail)
    end
    return rope
end

local function trunc(x)
    if math.abs(x) > 1 then
        return sign(x)
    else
        return 0
    end
end

function Rope:move(direction, steps)
    local moves = {
        ["U"] = Point:new { 0, 1 },
        ["D"] = Point:new { 0, -1 },
        ["R"] = Point:new { 1, 0 },
        ["L"] = Point:new { -1, 0 },
    }
    local dif
    self.history[tostring(self.tail)] = true
    for _ = 1, steps do
        self.head = self.head + moves[direction]
        dif = (self.head - self.tail)
        if not self.head:touches(self.tail) then
            self.tail = self.tail + dif:sign()
        else
            dif = Point:new { trunc(dif.x), trunc(dif.y) }
            self.tail = self.tail + dif
        end
        self.history[tostring(self.tail)] = true
    end
end

do
    local rope = Rope:new(Point:new(), Point:new())
    assert(rope == rope)
    assert(rope == Rope:new(Point:new(), Point:new()))
    rope:move("R", 1)
    assert(rope == Rope:new({ 1, 0 }, { 0, 0 }))
    rope:move("R", 3)
    assert(rope == Rope:new({ 4, 0 }, { 3, 0 }))
    rope:move("U", 1)
    assert(rope == Rope:new({ 4, 1 }, { 3, 0 }))
    rope:move("U", 1)
    assert(rope == Rope:new({ 4, 2 }, { 4, 1 }))
    rope:move("U", 2)
    assert(rope == Rope:new({ 4, 4 }, { 4, 3 }))
    rope:move("L", 1)
    assert(rope == Rope:new({ 3, 4 }, { 4, 3 }))
    rope:move("L", 1)
    assert(rope == Rope:new({ 2, 4 }, { 3, 4 }))
    rope:move("L", 1)
    assert(rope == Rope:new({ 1, 4 }, { 2, 4 }))
    rope:move("D", 1)
    assert(rope == Rope:new({ 1, 3 }, { 2, 4 }))
    rope:move("R", 4)
    assert(rope == Rope:new({ 5, 3 }, { 4, 3 }))
    rope:move("D", 1)
    assert(rope == Rope:new({ 5, 2 }, { 4, 3 }))
    rope:move("L", 5)
    assert(rope == Rope:new({ 0, 2 }, { 1, 2 }))
    rope:move("R", 2)
    assert(rope == Rope:new({ 2, 2 }, { 1, 2 }))
end

function problem1(input)
    local rope = Rope:new(Point:new(), Point:new())
    for line in lines(input) do
        if line ~= "" then
            local direction, steps = string.match(line, "(%u+)%s(%d+)")
            steps = tonumber(steps)
            rope:move(direction, steps)
        end
    end
    local count = 0
    for _, _ in pairs(rope.history) do
        count = count + 1
    end
    return count
end

assert(problem1(example1) == 13)
print(problem1(readfile("data/day-09.txt")))

LongRope = {}

function LongRope:new(length)
    local rope = { rope = {}, history = {} }
    for _ = 1, 10 do
        table.insert(rope.rope, Point:new())
    end
    setmetatable(rope, self)
    self.__index = self
    return rope
end

function LongRope:move(direction, steps)
    local moves = {
        ["U"] = Point:new { 0, 1 },
        ["D"] = Point:new { 0, -1 },
        ["R"] = Point:new { 1, 0 },
        ["L"] = Point:new { -1, 0 },
    }
    local dif
    self.history[tostring(self.rope[#self.rope])] = true
    for _ = 1, steps do
        self.rope[1] = self.rope[1] + moves[direction]
        for i = 2, #self.rope do
            dif = (self.rope[i - 1] - self.rope[i])
            if not self.rope[i - 1]:touches(self.rope[i]) then
                self.rope[i] = self.rope[i] + dif:sign()
            else
                dif = Point:new { trunc(dif.x), trunc(dif.y) }
                self.rope[i] = self.rope[i] + dif
            end
        end
        self.history[tostring(self.rope[#self.rope])] = true
    end
end

function problem2(input)
    local rope = LongRope:new(10)
    for line in lines(input) do
        if line ~= "" then
            local direction, steps = string.match(line, "(%u+)%s(%d+)")
            steps = tonumber(steps)
            rope:move(direction, steps)
        end
    end
    local count = 0
    for _, _ in pairs(rope.history) do
        count = count + 1
    end
    return count
end

local example2 = [[
R 5
U 8
L 8
D 3
R 17
D 10
L 25
U 20
]]

assert(problem2(example1) == 1)
assert(problem2(example2) == 36)
print(problem2(readfile("data/day-09.txt")))
