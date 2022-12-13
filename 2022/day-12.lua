require "common"

local example1 = [[
Sabqponm
abcryxxl
accszExk
acctuvwj
abdefghi
]]

function Node(i, j)
    local node = { row = i, col = j }
    setmetatable(node, {
        __eq = function(x, y)
            return x.row == y.row and x.col == y.col
        end,
        __tostring = function(o)
            return string.format("{%d, %d}", o.row, o.col)
        end
    })
    return node
end

Grid = {}

function Grid:new()
    local grid = { grid = { {} } }
    setmetatable(grid, self)
    self.__index = self
    return grid
end

function Grid:set(node, value)
    if not self.grid[node.row] then
        self.grid[node.row] = {}
    end
    self.grid[node.row][node.col] = value
end

function Grid:get(node)
    if self.grid[node.row] then
        return self.grid[node.row][node.col]
    end
end

function Grid:neighbours(node)
    local function isvalid(current, next)
        current = self:get(current)
        next = self:get(next)
        return next and next <= current + 1
    end

    local nodes = {}
    for _, neighbour in ipairs({
        Node(node.row - 1, node.col),
        Node(node.row + 1, node.col),
        Node(node.row, node.col - 1),
        Node(node.row, node.col + 1)
    }) do
        if isvalid(node, neighbour) then
            table.insert(nodes, neighbour)
        end
    end
    return nodes
end

function Grid:show()
    for i = 1, #self.grid do
        local out = ""
        for j = 1, #self.grid[i] do
            out = out .. string.format(" %3s", self.grid[i][j])
        end
        print(out)
    end
end

local function toelevation(char)
    return string.byte(char) - 96
end

local function parse(input)
    local grid = Grid:new()
    local i = 1
    local j = 1
    local start = {}
    local final = {}
    for char in chars(input) do
        if char == "\n" then
            i = i + 1
            j = 0
        elseif char == "S" then
            start = Node(i, j)
            grid:set(start, toelevation("a"))
        elseif char == "E" then
            final = Node(i, j)
            grid:set(final, toelevation("z"))
        else
            local value = toelevation(char)
            grid:set(Node(i, j), value)
        end
        j = j + 1
    end
    return grid, start, final
end

do
    local grid, start, final = parse(example1)
    local expected = {
        { 1, 1, 2, 17, 16, 15, 14, 13 },
        { 1, 2, 3, 18, 25, 24, 24, 12 },
        { 1, 3, 3, 19, 26, 26, 24, 11 },
        { 1, 3, 3, 20, 21, 22, 23, 10 },
        { 1, 2, 4, 5, 6, 7, 8, 9 },
    }
    for i = 1, #expected do
        for j = 1, #expected[i] do
            assert(grid:get(Node(i, j)) == expected[i][j])
        end
    end

    assert(start == Node(1, 1))
    assert(final == Node(3, 6))

    local neighbours = grid:neighbours(Node(1, 1))
    for i, expected in ipairs({ Node(2, 1), Node(1, 2) }) do
        assert(neighbours[i] == expected)
    end

    local neighbours = grid:neighbours(Node(1, 3))
    for i, expected in ipairs({ Node(2, 3), Node(1, 2) }) do
        assert(neighbours[i] == expected)
    end

    local neighbours = grid:neighbours(Node(2, 3))
    for i, expected in ipairs({ Node(1, 3), Node(3, 3), Node(2, 2) }) do
        assert(neighbours[i] == expected)
    end
end

local function bfs(grid, start, final)
    local dist = Grid:new()
    dist:set(start, 0)

    local queue = { start }
    repeat
        local current = table.remove(queue, 1)
        if current == final then
            break
        end
        for _, next in ipairs(grid:neighbours(current)) do
            if not dist:get(next) then
                local cost = dist:get(current) + 1
                dist:set(next, cost)
                table.insert(queue, next)
            end
        end
    until #queue == 0

    return dist:get(final)
end

function problem1(input)
    local grid, start, final = parse(input)
    return bfs(grid, start, final)
end

do
    local grid = parse(readfile("data/day-12.txt"))
    assert(#grid.grid == 41)
    for i = 1, #grid.grid do
        assert(#grid.grid[i] == 172)
    end
end

assert(problem1(example1) == 31)
print(problem1(readfile("data/day-12.txt")))

function problem2(input)
    local grid, _, final = parse(input)
    local best = math.huge
    local starts = {}
    for i = 1, #grid.grid do
        for j = 1, #grid.grid[1] do
            local node = Node(i, j)
            if grid:get(node) == 1 then
                table.insert(starts, node)
            end
        end
    end
    for _, start in ipairs(starts) do
        local solution = bfs(grid, start, final)
        if solution and solution < best then
            best = solution
        end
    end
    return best
end

assert(problem2(example1) == 29)
print(problem2(readfile("data/day-12.txt")))
