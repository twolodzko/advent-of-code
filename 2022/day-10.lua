require "common"

ClockCircuit = {}

function ClockCircuit:new(comp)
    local clock = { cycle = 1, register = 1, comp = comp }
    setmetatable(clock, self)
    self.__index = self
    return clock
end

function ClockCircuit:next()
    self.cycle = self.cycle + 1
    self.comp:update(self)
end

function ClockCircuit:exec(cmd, arg)
    self:next()
    if cmd == "addx" then
        self.register = self.register + arg
        self:next()
    end
end

function ClockCircuit:eval(cmd)
    if cmd == "noop" then
        self:exec(cmd)
    else
        local arg
        cmd, arg = string.match(cmd, "([a-z]+)%s+(-?[0-9]+)")
        self:exec(cmd, tonumber(arg))
    end
end

StrengthCalculator = {}

function StrengthCalculator:new()
    local calc = { strength = 0 }
    setmetatable(calc, self)
    self.__index = self
    return calc
end

function StrengthCalculator:update(clock)
    if (clock.cycle - 20) % 40 == 0 then
        local strength = clock.cycle * clock.register
        if strength then
            self.strength = self.strength + strength
        end
    end
end

function StrengthCalculator:result()
    return self.strength
end

function problem1(input)
    local clock = ClockCircuit:new(StrengthCalculator:new())
    for line in lines(input) do
        if line ~= "" then
            clock:eval(line)
        end
    end
    return clock.comp:result()
end

assert(problem1(readfile("data/day-10-example.txt")) == 13140)
print(problem1(readfile("data/day-10.txt")))

Display = {}

function Display:new()
    local disp = { rows = {} }
    setmetatable(disp, self)
    self.__index = self
    return disp
end

function Display:update(clock)
    local col = (clock.cycle - 1) % 40
    local row = (clock.cycle - 1) // 40
    if col >= clock.register - 1 and col <= clock.register + 1
    then
        if not self.rows[row] then
            self.rows[row] = {}
        end
        self.rows[row][col] = true
    end
end

function Display:result()
    local out = ""
    for i = 0, 5 do
        for j = 0, 39 do
            out = out .. (self.rows[i][j] and "#" or ".")
        end
        out = out .. "\n"
    end
    return out
end

function problem2(input)
    local clock = ClockCircuit:new(Display:new())
    for line in lines(input) do
        if line ~= "" then
            clock:eval(line)
        end
    end
    return clock.comp:result()
end

print(problem2(readfile("data/day-10-example.txt")))
print(problem2(readfile("data/day-10.txt")))
