require "common"

local example1 = [[
noop
addx 3
addx -5
]]

local function signalstrength(cycle, value)
    if cycle >= 20 then
        if (cycle - 20) % 40 == 0 then
            return cycle * value
        end
    end
end

ClockCircuit = {}

function ClockCircuit:new()
    local clock = { cycle = 1, register = 1, strength = 0 }
    setmetatable(clock, self)
    self.__index = self
    return clock
end

function ClockCircuit:next()
    self.cycle = self.cycle + 1
    local strength = signalstrength(self.cycle, self.register)
    if strength then
        self.strength = self.strength + strength
    end
end

function ClockCircuit:exec(cmd, val)
    if cmd == "noop" then
        self:next()
    elseif cmd == "addx" then
        self:next()
        self.register = self.register + val
        self:next()
    else
        error(string.format("invalid command: %s", cmd))
    end
end

local function eval(clock, code)
    local cmd, val = string.match(code, "([a-z]+)%s+(-?[0-9]+)")
    if val then
        val = tonumber(val)
        clock:exec(cmd, val)
    else
        clock:exec(code)
    end
end

function problem1(input)
    local clock = ClockCircuit:new()
    for line in lines(input) do
        if line ~= "" then
            eval(clock, line)
        end
    end
    return clock.strength
end

assert(problem1(readfile("data/day-10-example.txt")) == 13140)
print(problem1(readfile("data/day-10.txt")))

CRT = ClockCircuit:new()

function CRT:next()
    self.cycle = self.cycle + 1
    local col = (self.cycle - 1) % 40
    local row = (self.cycle - 1) // 40 + 1
    if col >= self.register - 1 and col <= self.register + 1
    then
        if not self.rows then
            self.rows = {}
        end
        if not self.rows[row] then
            self.rows[row] = {}
        end
        self.rows[row][col] = true
    end
end

function CRT:tostring()
    local out = ""
    for i = 1, 6 do
        for j = 0, 39 do
            if self.rows[i][j] then
                out = out .. "#"
            else
                out = out .. "."
            end
        end
        out = out .. "\n"
    end
    return out
end

function problem2(input)
    local crt = CRT:new()
    for line in lines(input) do
        if line ~= "" then
            eval(crt, line)
        end
    end
    return crt:tostring()
end

print(problem2(readfile("data/day-10-example.txt")))
print(problem2(readfile("data/day-10.txt")))
