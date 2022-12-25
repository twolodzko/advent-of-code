require "common"

local example1 = [[
root: pppw + sjmn
dbpl: 5
cczh: sllz + lgvd
zczc: 2
ptdq: humn - dvpt
dvpt: 3
lfqf: 4
humn: 5
ljgn: 2
sjmn: drzm * dbpl
sllz: 4
pppw: cczh / lfqf
lgvd: ljgn * ptdq
drzm: hmdt - zczc
hmdt: 32
]]


-- input = input
--         :gsub("%[", "{")
--         :gsub("%]", "}")
--     local env = {}
--     assert(load("result = " .. input, input, "t", env))()

local function tofunc(str)
    local name, body = string.match(str, "([^:]+): (.+)")
    body = string.gsub(body, "(%a+)", "%1()")
    return string.format("function %s() return %s end", name, body)
end

assert(tofunc("hmdt: 32") == "function hmdt() return 32 end")
assert(tofunc("lgvd: ljgn * ptdq") == "function lgvd() return ljgn() * ptdq() end")

function problem1(input)
    local env = {}
    for line in lines(input) do
        if line ~= "" then
            load(tofunc(line), line, "t", env)()
        end
    end
    return env.root()
end

assert(problem1(example1) == 152)
assert(root == nil)

print(problem1(readfile("data/day-21.txt")))

local function tofunc2(str)
    local name, body = string.match(str, "([^:]+): (.+)")
    if name == "root" then
        body = string.gsub(body, "(%w+) [*/+-] (%w+)", "%1 - %2")
    elseif name == "humn" then
        body = "VALUE"
    end
    body = string.gsub(body, "(%l+)", "%1()")
    return string.format("function %s() return %s end", name, body)
end

assert(tofunc2("root: pppw + sjmn") == "function root() return pppw() - sjmn() end")
assert(tofunc2("humn: 5") == "function humn() return VALUE end")



function problem2(input)
    local env = {}
    for line in lines(input) do
        if line ~= "" then
            load(tofunc2(line), line, "t", env)()
        end
    end

    local function func(x)
        env.VALUE = x
        return env.root()
    end

    local result = 301
    assert(func(result) == 0)

    return result
end

assert(problem2(example1) == 301)
print(problem2(readfile("data/day-21.txt")))
