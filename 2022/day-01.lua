require "common"

local example1 = [[
1000
2000
3000

4000

5000
6000

7000
8000
9000

10000
]]

function findmax(lines)
    local max = 0
    local acc = 0
    for line in string.gmatch(lines, "[^\n]*") do
        local num = tonumber(line)
        if num then
            acc = acc + num
        else
            if acc > max then
                max = acc
            end
            acc = 0
        end
    end
    return max
end

assert(findmax(example1) == 24000)
print(findmax(readfile("data/day-01.txt")))

function findtop3(input)
    local arr = {}
    local acc = 0
    for line in lines(input) do
        local num = tonumber(line)
        if num then
            acc = acc + num
        else
            table.insert(arr, acc)
            acc = 0
        end
    end
    table.sort(arr, function(x, y) return x > y end)

    local total = 0
    for i = 1, 3 do
        total = total + arr[i]
    end
    return total
end

assert(findtop3(example1) == 45000)
print(findtop3(readfile("data/day-01.txt")))
