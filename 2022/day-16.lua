require "common"

local example1 = [[
Valve AA has flow rate=0; tunnels lead to valves DD, II, BB
Valve BB has flow rate=13; tunnels lead to valves CC, AA
Valve CC has flow rate=2; tunnels lead to valves DD, BB
Valve DD has flow rate=20; tunnels lead to valves CC, AA, EE
Valve EE has flow rate=3; tunnels lead to valves FF, DD
Valve FF has flow rate=0; tunnels lead to valves EE, GG
Valve GG has flow rate=0; tunnels lead to valves FF, HH
Valve HH has flow rate=22; tunnel leads to valve GG
Valve II has flow rate=0; tunnels lead to valves AA, JJ
Valve JJ has flow rate=21; tunnel leads to valve II
]]

local function parse(input)
    local valves = {}
    for line in lines(input) do
        if line ~= "" then
            local id, flowrate, leads = string.match(line,
                "^Valve (%u+) has flow rate=(%d+); tunnels? leads? to valves? ([%u,%s]+)$")
            local valve = { id = id, flowrate = flowrate, leads = {} }
            for next in string.gmatch(leads, "(%u+)") do
                table.insert(valve.leads, next)
            end
            valves[id] = valve
        end
    end
    return valves
end

do
    local valves = parse(example1)
    local count = 0
    for key, valve in pairs(valves) do
        assert(key == valve.id)
        count = count + 1
    end
    assert(count == 10)
    assert(#valves["AA"].leads == 3)
    assert(#valves["JJ"].leads == 1)
end

local function solve(valves)
    local results = {}

    local function step(valve, time, pressure, open)
        time = time - 1
        for v in pairs(open) do
            pressure = pressure + valves[v].flowrate
        end

        if time == 0 then
            table.insert(results, pressure)
            return
        end

        if not open[valve] then
            local newopen = {}
            for v in pairs(open) do
                newopen[v] = true
            end
            newopen[valve] = true
            return step(valve, time, pressure, newopen)
        end

        for _, v in pairs(valves[valve].leads) do
            step(v, time, pressure, open)
        end
    end

    step("AA", 30, 0, {})

    table.sort(results)
    return results[#results]
end

local function problem1(input)
    local valves = parse(input)
    return solve(valves)
end

assert(problem1(example1) == 1651)
