require "common"

local example1 = [[
1=-0-2
12111
2=0=
21
2=01
111
20012
112
1=-1=
1-12
12
1=
122
]]

local function todigit(chr)
    if chr == "-" then
        return -1
    elseif chr == "=" then
        return -2
    else
        return tonumber(chr)
    end
end

local function fromsnafu(str)
    local num = 0
    for i = #str, 1, -1 do
        local chr = str:sub(i, i)
        local pow = #str - i
        num = num + 5 ^ pow * todigit(chr)
    end
    return num
end

assert(fromsnafu("1") == 1)
assert(fromsnafu("2") == 2)
assert(fromsnafu("1=") == 3)
assert(fromsnafu("1-") == 4)
assert(fromsnafu("10") == 5)
assert(fromsnafu("11") == 6)
assert(fromsnafu("12") == 7)
assert(fromsnafu("2=") == 8)
assert(fromsnafu("2-") == 9)
assert(fromsnafu("20") == 10)
assert(fromsnafu("1=0") == 15)
assert(fromsnafu("1-0") == 20)
assert(fromsnafu("1=11-2") == 2022)
assert(fromsnafu("1-0---0") == 12345)
assert(fromsnafu("1121-1110-1=0") == 314159265)

local function tosnafu(num)
    local str = ""
    repeat
        local rem = math.floor(num % 5)
        num = num // 5
        if rem <= 2 then
            str = tostring(rem) .. str
        else
            -- (rem - 5) is in {-2, -1}
            if rem == 3 then
                str = "=" .. str
            else
                str = "-" .. str
            end
            num = num + 1
        end
    until num == 0
    return str
end

assert(tosnafu(1) == "1")
assert(tosnafu(2) == "2")
assert(tosnafu(3) == "1=")
assert(tosnafu(4) == "1-")
assert(tosnafu(5) == "10")
assert(tosnafu(6) == "11")
assert(tosnafu(7) == "12")
assert(tosnafu(8) == "2=")
assert(tosnafu(9) == "2-")
assert(tosnafu(10) == "20")
assert(tosnafu(15) == "1=0")
assert(tosnafu(20) == "1-0")
assert(tosnafu(2022) == "1=11-2")
assert(tosnafu(12345) == "1-0---0")
assert(tosnafu(314159265) == "1121-1110-1=0")
assert(tosnafu(4890) == "2=-1=0")

function problem1(input)
    local total = 0
    for line in lines(input) do
        if line ~= "" then
            local num = fromsnafu(line)
            total = total + num
        end
    end
    return tosnafu(total)
end

assert(problem1(example1) == "2=-1=0")
print(problem1(readfile("data/day-25.txt")))
