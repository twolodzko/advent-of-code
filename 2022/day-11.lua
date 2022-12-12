require "common"

local example1 = [[
Monkey 0:
  Starting items: 79, 98
  Operation: new = old * 19
  Test: divisible by 23
    If true: throw to monkey 2
    If false: throw to monkey 3

Monkey 1:
  Starting items: 54, 65, 75, 74
  Operation: new = old + 6
  Test: divisible by 19
    If true: throw to monkey 2
    If false: throw to monkey 0

Monkey 2:
  Starting items: 79, 60, 97
  Operation: new = old * old
  Test: divisible by 13
    If true: throw to monkey 1
    If false: throw to monkey 3

Monkey 3:
  Starting items: 74
  Operation: new = old + 3
  Test: divisible by 17
    If true: throw to monkey 0
    If false: throw to monkey 1
]]

Monkey = {}

function Monkey:new()
    local monkey = { items = {}, inspections = 0 }
    setmetatable(monkey, self)
    self.__index = self
    self.__tostring = function(obj)
        return string.format(
            "Monkey %d\nnew = %s %s %s\n{%s}",
            obj.id,
            obj.x,
            obj.op,
            obj.y,
            table.concat(obj.items, ", "))
    end
    return monkey
end

function Monkey:exec(item)
    self.inspections = self.inspections + 1

    local x = self.x == "old" and item or self.x
    local y = self.y == "old" and item or self.y
    local new
    if self.op == "+" then
        new = x + y
    else
        new = x * y
    end
    new = new // 3
    if (new % self.isdivby) == 0 then
        return self.iftrue, new
    else
        return self.iffalse, new
    end
end

function Monkey:receive(item)
    table.insert(self.items, item)
end

local function parsemonkeys(lines, i, monkeys)
    if i >= #lines then
        return monkeys
    end
    if lines[i] == "" then
        i = i + 1
    end

    local monkey = Monkey:new()
    -- we want to start counting at 1 like Lua
    monkey.id = tonumber(string.match(lines[i], "Monkey (%d+):")) + 1
    local items = string.match(lines[i + 1], "Starting items: (.*)")
    for item in string.gmatch(items, "(%d+)") do
        table.insert(monkey.items, tonumber(item))
    end
    local x, op, y = string.match(lines[i + 2], "Operation: new = ([%d%l]+) ([+*]) ([%d%l]+)")
    monkey.x = tonumber(x) or x
    monkey.y = tonumber(y) or y
    monkey.op = op
    monkey.isdivby = tonumber(string.match(lines[i + 3], "Test: divisible by (%d+)"))
    monkey.iftrue = tonumber(string.match(lines[i + 4], "If true: throw to monkey (%d+)")) + 1
    monkey.iffalse = tonumber(string.match(lines[i + 5], "If false: throw to monkey (%d+)")) + 1

    table.insert(monkeys, monkey)

    return parsemonkeys(lines, i + 6, monkeys)
end

local function parse(input)
    local lines = {}
    for line in string.gmatch(input, "[^\n]*") do
        table.insert(lines, line)
    end
    return parsemonkeys(lines, 1, {})
end

local function round(monkeys)
    for _, monkey in ipairs(monkeys) do
        for _, item in ipairs(monkey.items) do
            local i, val = monkey:exec(item)
            monkeys[i]:receive(val)
            monkey.items = {}
        end
    end
    return monkeys
end

do
    local monkeys = parse(example1)
    assert(equal(monkeys[1].items, { 79, 98 }))
    assert(equal(monkeys[2].items, { 54, 65, 75, 74 }))
    assert(equal(monkeys[3].items, { 79, 60, 97 }))
    assert(equal(monkeys[4].items, { 74 }))

    monkeys = round(monkeys)
    assert(equal(monkeys[1].items, { 20, 23, 27, 26 }))
    assert(equal(monkeys[2].items, { 2080, 25, 167, 207, 401, 1046 }))
    assert(equal(monkeys[3].items, {}))
    assert(equal(monkeys[4].items, {}))

    monkeys = round(monkeys)
    assert(equal(monkeys[1].items, { 695, 10, 71, 135, 350 }))
    assert(equal(monkeys[2].items, { 43, 49, 58, 55, 362 }))
    assert(equal(monkeys[3].items, {}))
    assert(equal(monkeys[4].items, {}))

    monkeys = round(monkeys)
    assert(equal(monkeys[1].items, { 16, 18, 21, 20, 122 }))
    assert(equal(monkeys[2].items, { 1468, 22, 150, 286, 739 }))
    assert(equal(monkeys[3].items, {}))
    assert(equal(monkeys[4].items, {}))

    monkeys = round(monkeys)
    assert(equal(monkeys[1].items, { 491, 9, 52, 97, 248, 34 }))
    assert(equal(monkeys[2].items, { 39, 45, 43, 258 }))
    assert(equal(monkeys[3].items, {}))
    assert(equal(monkeys[4].items, {}))
end

function problem1(input)
    local monkeys = parse(input)
    for _ = 1, 20 do
        monkeys = round(monkeys)
    end

    table.sort(monkeys, function(a, b)
        return a.inspections > b.inspections
    end)
    return monkeys[1].inspections * monkeys[2].inspections
end

assert(problem1(example1) == 10605)
print(problem1(readfile("data/day-11.txt")))



function problem2(input)
    local monkeys = parse(input)
    for _ = 1, 20 do
        monkeys = round(monkeys, 1)
    end

    table.sort(monkeys, function(a, b)
        return a.inspections > b.inspections
    end)
    return monkeys[1].inspections * monkeys[2].inspections
end
