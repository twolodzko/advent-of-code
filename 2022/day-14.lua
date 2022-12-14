require "common"

local example1 = [[
498,4 -> 498,6 -> 496,6
503,4 -> 502,4 -> 502,9 -> 494,9
]]

local function Point(x, y)
    local point = { x = x, y = y }
    setmetatable(point, {
        __eq = function(a, b)
            return a.x == b.x and a.y == b.y
        end,
        __add = function(a, b)
            return Point(a.x + b.x, a.y + b.y)
        end,
        __sub = function(a, b)
            return Point(a.x - b.x, a.y - b.y)
        end,
        __tostring = function(o)
            return string.format("Point{%d, %d}", o.x, o.y)
        end
    })
    return point
end

local function parse(input)
    local paths = {}
    for line in lines(input) do
        if line ~= "" then
            local path = {}
            for x, y in string.gmatch(line, "(%d+),(%d+)") do
                table.insert(path, Point(tonumber(x), tonumber(y)))
            end
            table.insert(paths, path)
        end
    end
    return paths
end

assert(#parse(example1) == 2)
assert(equal(parse(example1)[1], { Point(498, 4), Point(498, 6), Point(496, 6) }))
assert(equal(parse(example1)[2], { Point(503, 4), Point(502, 4), Point(502, 9), Point(494, 9) }))

local function sign(num)
    if num == 0 then
        return 0
    elseif num < 0 then
        return -1
    else
        return 1
    end
end

local function interpolate(p1, p2)
    if p1 == p2 then
        return { p1 }
    end
    local dif = p2 - p1
    local step = Point(sign(dif.x), sign(dif.y))
    local path = { p1 }
    repeat
        local next = path[#path] + step
        table.insert(path, next)
    until next == p2
    return path
end

assert(equal(interpolate(Point(0, 0), Point(0, 0)), { Point(0, 0) }))
assert(equal(interpolate(Point(0, 0), Point(0, 2)), { Point(0, 0), Point(0, 1), Point(0, 2) }))
assert(equal(interpolate(Point(2, 0), Point(0, 0)), { Point(2, 0), Point(1, 0), Point(0, 0) }))

local ROCK = "#"
local SAND = "o"
local START = Point(500, 0)

Map = {}

function Map:new()
    local map = {
        map = {},
        minx = math.maxinteger,
        maxx = math.mininteger,
        miny = 0,
        maxy = math.mininteger,
        floor = math.huge
    }
    setmetatable(map, self)
    self.__index = self
    return map
end

function Map:updatelimits(point)
    if point.x < self.minx then
        self.minx = point.x
    elseif point.x > self.maxx then
        self.maxx = point.x
    end
    if point.y < self.miny then
        self.miny = point.y
    elseif point.y > self.maxy then
        self.maxy = point.y
    end
end

function Map:set(point, value)
    self:updatelimits(point)
    if not self.map[point.x] then
        self.map[point.x] = {}
    end
    self.map[point.x][point.y] = value
end

function Map:get(point)
    if point.y >= self.floor then
        return ROCK
    end

    if self.map[point.x] then
        return self.map[point.x][point.y]
    end
end

function Map:tostring()
    local str = ""
    for y = self.miny, self.maxy do
        local line = ""
        for x = self.minx, self.maxx do
            local value = self:get(Point(x, y))
            if Point(x, y) == START then
                line = line .. "+"
            elseif value == nil then
                line = line .. "."
            else
                line = line .. value
            end
        end
        str = str .. line .. "\n"
    end
    return str
end

local function generatemap(paths)
    local map = Map:new()
    for _, path in ipairs(paths) do
        local p1 = table.remove(path, 1)
        while #path ~= 0 do
            local p2 = table.remove(path, 1)
            for _, p in ipairs(interpolate(p1, p2)) do
                map:set(p, ROCK)
            end
            p1 = p2
        end
    end
    return map
end

do
    local map = generatemap(parse(example1))
    assert(map:get(Point(498, 3)) == nil)
    assert(map:get(Point(497, 4)) == nil)
    assert(map:get(Point(498, 4)) == ROCK)
    assert(map:get(Point(498, 6)) == ROCK)
    assert(map:get(Point(498, 7)) == nil)
    assert(map:get(Point(500, 9)) == ROCK)
    assert(map:get(Point(500, 9)) == ROCK)
    assert(map:get(Point(500, 1)) == nil)
    assert(map:get(Point(503, 4)) == ROCK)
    assert(map:get(Point(504, 4)) == nil)

    assert(map.minx == 494)
    assert(map.maxx == 503)
    assert(map.miny == 0)
    assert(map.maxy == 9)
end

local function sandfall(map)
    local position = START
    local down = Point(0, 1)
    local left = Point(-1, 1)
    local right = Point(1, 1)

    while true do
        if position.y > map.maxy then
            return false
        end

        if map:get(position + down) == nil then
            position = position + down
        else
            if map:get(position + left) == nil then
                position = position + left
            elseif map:get(position + right) == nil then
                position = position + right
            else
                if map:get(position) ~= nil then
                    error("this space is taken")
                end
                map:set(position, SAND)
                return true
            end
        end
    end
end

do
    local map = generatemap(parse(example1))

    sandfall(map)
    assert(map:get(Point(500, 7)) == nil)
    assert(map:get(Point(500, 8)) == SAND)
    assert(map:get(Point(499, 8)) == nil)
    assert(map:get(Point(501, 8)) == nil)

    sandfall(map)
    assert(map:get(Point(500, 7)) == nil)
    assert(map:get(Point(500, 8)) == SAND)
    assert(map:get(Point(499, 8)) == SAND)
    assert(map:get(Point(501, 8)) == nil)
end

do
    local map = generatemap(parse(example1))

    for _ = 1, 24 do
        sandfall(map)
    end

    local result = map:tostring()
    local expected = [[
......+...
..........
......o...
.....ooo..
....#ooo##
...o#ooo#.
..###ooo#.
....oooo#.
.o.ooooo#.
#########.
]]
    assert(result == expected)

end

do
    local map = generatemap(parse(example1))

    for _ = 1, 100 do
        sandfall(map)
    end

    local result = map:tostring()
    local expected = [[
......+...
..........
......o...
.....ooo..
....#ooo##
...o#ooo#.
..###ooo#.
....oooo#.
.o.ooooo#.
#########.
]]
    assert(result == expected)
end

function problem1(input)
    local map = generatemap(parse(input))
    local counter = 0
    while sandfall(map) do
        counter = counter + 1
    end
    return counter
end

assert(problem1(example1) == 24)
print(problem1(readfile("data/day-14.txt")))

local function sandfall2(map)
    local position = START
    local down = Point(0, 1)
    local left = Point(-1, 1)
    local right = Point(1, 1)

    while true do
        if map:get(position + down) == nil then
            position = position + down
        else
            if map:get(position + left) == nil then
                position = position + left
            elseif map:get(position + right) == nil then
                position = position + right
            else
                if map:get(position) ~= nil then
                    error("this space is taken")
                end
                map:set(position, SAND)
                return position.y > 0
            end
        end
    end
end

function problem2(input)
    local map = generatemap(parse(input))
    map.floor = map.maxy + 2
    local counter = 0
    while sandfall2(map) do
        counter = counter + 1
    end
    assert(map:get(START) == SAND)
    return counter + 1
end

assert(problem2(example1) == 93)
print(problem2(readfile("data/day-14.txt")))
