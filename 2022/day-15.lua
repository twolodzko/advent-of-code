require "common"

example1 = [[
Sensor at x=2, y=18: closest beacon is at x=-2, y=15
Sensor at x=9, y=16: closest beacon is at x=10, y=16
Sensor at x=13, y=2: closest beacon is at x=15, y=3
Sensor at x=12, y=14: closest beacon is at x=10, y=16
Sensor at x=10, y=20: closest beacon is at x=10, y=16
Sensor at x=14, y=17: closest beacon is at x=10, y=16
Sensor at x=8, y=7: closest beacon is at x=2, y=10
Sensor at x=2, y=0: closest beacon is at x=2, y=10
Sensor at x=0, y=11: closest beacon is at x=2, y=10
Sensor at x=20, y=14: closest beacon is at x=25, y=17
Sensor at x=17, y=20: closest beacon is at x=21, y=22
Sensor at x=16, y=7: closest beacon is at x=15, y=3
Sensor at x=14, y=3: closest beacon is at x=15, y=3
Sensor at x=20, y=1: closest beacon is at x=15, y=3
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
        __mul = function(a, b)
            return Point(a.x * b.x, a.y * b.y)
        end,
        __tostring = function(o)
            return string.format("Point{%d, %d}", o.x, o.y)
        end
    })
    return point
end

local function abs(p)
    return Point(math.abs(p.x), math.abs(p.y))
end

assert(abs(Point(1, 2)) == Point(1, 2))
assert(abs(Point(-1, 2)) == Point(1, 2))
assert(abs(Point(-1, -2)) == Point(1, 2))

local function dist(a, b)
    local absdiff = abs(a - b)
    return absdiff.x + absdiff.y
end

assert(dist(Point(8, 7), Point(8, 7)) == 0)
assert(dist(Point(8, 7), Point(2, 10)) == 9)
assert(dist(Point(8, 7), Point(-1, 7)) == 9)
assert(dist(Point(8, 7), Point(10, 0)) == 9)

Sensor = {}

function Sensor:new(position, beacon)
    local sensor = {
        position = position,
        closest = beacon,
        range = dist(position, beacon)
    }
    setmetatable(sensor, self)
    self.__index = self
    return sensor
end

function Sensor:yrange(y)
    local dy = math.abs(self.position.y - y)
    if dy <= self.range then
        local dx = self.range - dy
        return self.position.x - dx, self.position.x + dx
    end
end

function Sensor:sees(point)
    return dist(self.position, point) <= self.range
end

function parse(input)
    local sensors = {}
    for line in lines(input) do
        if line ~= "" then
            local sx, sy, bx, by = string.match(line,
                "Sensor at x=(-?%d+), y=(-?%d+): closest beacon is at x=(-?%d+), y=(-?%d+)")
            local beacon = Point(tonumber(bx), tonumber(by))
            local position = Point(tonumber(sx), tonumber(sy))
            table.insert(sensors, Sensor:new(position, beacon))
        end
    end
    return sensors
end

do
    local sensors = parse(example1)
    assert(#sensors == 14)
    assert(sensors[1].position == Point(2, 18))
    assert(sensors[14].closest == Point(15, 3))
end

local SENSOR = "S"
local BEACON = "B"

function makemap(sensors)
    local map = {}
    local p
    for _, sensor in ipairs(sensors) do
        p = sensor.position
        if not map[p.x] then
            map[p.x] = {}
        end
        map[p.x][p.y] = SENSOR

        p = sensor.closest
        if not map[p.x] then
            map[p.x] = {}
        end
        map[p.x][p.y] = BEACON
    end
    return map
end

function problem1(input, y)
    local sensors = parse(input)
    local map = makemap(sensors)

    local minx = math.maxinteger
    local maxx = math.mininteger
    for _, sensor in ipairs(sensors) do
        local lo, hi = sensor:yrange(y)
        if lo then
            if lo < minx then
                minx = lo
            end
            if hi > maxx then
                maxx = hi
            end
        end
    end

    local count = 0
    for x = minx, maxx do
        if not map[x] or map[x][y] ~= BEACON then
            for _, sensor in ipairs(sensors) do
                local lo, hi = sensor:yrange(y)
                if lo and x >= lo and x <= hi then
                    count = count + 1
                    break
                end
            end
        end
    end
    return count
end

assert(problem1(example1, 10) == 26)
print(problem1(readfile("data/day-15.txt"), 2000000))

function problem2(input)
    local sensors = parse(input)
    local map = makemap(sensors)
    local lo = 0
    local hi = 4000000
end
