require "common"

local example1 = [[
1,1,1
2,1,1
]]

local function Point(x, y, z)
    local point = { x = x, y = y, z = z }
    setmetatable(point, {
        __eq = function(a, b)
            return a.x == b.x and a.y == b.y and a.z == b.z
        end,
        __add = function(a, b)
            return Point(a.x + b.x, a.y + b.y, a.z + b.z)
        end,
        __sub = function(a, b)
            return Point(a.x - b.x, a.y - b.y, a.z - b.z)
        end,
        __mul = function(a, b)
            return Point(a.x * b.x, a.y * b.y, a.z * b.z)
        end,
        __tostring = function(o)
            return string.format("Point{%d, %d, %d}", o.x, o.y, o.z)
        end
    })
    return point
end

Grid = {}

function Grid:new()
    local grid = { grid = {} }
    setmetatable(grid, self)
    self.__index = self
    return grid
end

function Grid:add(point)
    if not self.grid[point.x] then
        self.grid[point.x] = {}
    end
    if not self.grid[point.x][point.y] then
        self.grid[point.x][point.y] = {}
    end
    self.grid[point.x][point.y][point.z] = true
end

function Grid:get(point)
    if self.grid[point.x] and self.grid[point.x][point.y] then
        return self.grid[point.x][point.y][point.z]
    end
end

function Grid:size()
    local count = 0
    for x, _ in pairs(self.grid) do
        for y, _ in pairs(self.grid[x]) do
            for z, _ in pairs(self.grid[x][y]) do
                count = count + 1
            end
        end
    end
    return count
end

local function parse(input)
    local grid = Grid:new()
    local points = {}
    for line in lines(input) do
        if line ~= "" then
            local x, y, z = string.match(line, "(%d+),(%d+),(%d+)")
            local point = Point(tonumber(x), tonumber(y), tonumber(z))
            grid:add(point)
            table.insert(points, point)
        end
    end
    return grid, points
end

do
    local grid, points = parse(example1)
    assert(#points == 2)
    assert(grid:size() == 2)
    assert(grid:get(Point(1, 1, 1)))
    assert(grid:get(Point(2, 1, 1)))
    assert(not grid:get(Point(1, 1, 2)))

    for _, point in ipairs(points) do
        assert(grid:get(point))
    end
end

local function sides(grid, point)
    local count = 0
    for _, d in ipairs({
        Point(1, 0, 0),
        Point(0, 1, 0),
        Point(0, 0, 1),
        Point(-1, 0, 0),
        Point(0, -1, 0),
        Point(0, 0, -1),
    }) do
        if not grid:get(point + d) then
            count = count + 1
        end
    end
    return count
end

function problem1(input)
    local grid, points = parse(input)
    local count = 0
    for _, point in ipairs(points) do
        count = count + sides(grid, point)
    end
    return count
end

assert(problem1(example1) == 10)

local example2 = [[
2,2,2
1,2,2
3,2,2
2,1,2
2,3,2
2,2,1
2,2,3
2,2,4
2,2,6
1,2,5
3,2,5
2,1,5
2,3,5
]]

assert(problem1(example2) == 64)
print(problem1(readfile("data/day-18.txt")))

function problem2(input)
    -- local grid, points = parse(input)
    -- local count = 0
    -- for _, point in ipairs(points) do
    --     count = count + sides(grid, point)
    -- end
    -- return count
end

assert(problem2(example2) == 58)


local function outersides(grid, point)
    local count = 0
    for _, d in ipairs({
        Point(1, 0, 0),
        Point(0, 1, 0),
        Point(0, 0, 1),
        Point(-1, 0, 0),
        Point(0, -1, 0),
        Point(0, 0, -1),
    }) do
        if not grid:get(point + d) then
            count = count + 1
        end
    end
    return count
end
