require "common"

local function different(str, pos, num)
    local seen = {}
    for i = pos - num + 1, pos do
        local chr = str:sub(i, i)
        if seen[chr] then
            return false
        end
        seen[chr] = true
    end
    return true
end

assert(different("abcd", 4, 4))
assert(not different("abca", 4, 4))
assert(not different("aacd", 4, 4))

local function findpos(str, num)
    for i = num, #str do
        if different(str, i, num) then
            return i
        end
    end
end

function problem1(str)
    return findpos(str, 4)
end

assert(problem1("bvwbjplbgvbhsrlpgdmjqwftvncz"), 5)
assert(problem1("nppdvjthqldpwncqszvftbrmjlhg"), 6)
assert(problem1("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"), 10)
assert(problem1("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"), 11)

print(problem1(readfile("data/day-06.txt")))

function problem2(str)
    return findpos(str, 14)
end

assert(problem2("mjqjpqmgbljsphdztnvjfqwrcgsmlb"), 19)
assert(problem2("bvwbjplbgvbhsrlpgdmjqwftvncz"), 23)
assert(problem2("nppdvjthqldpwncqszvftbrmjlhg"), 23)
assert(problem2("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"), 29)
assert(problem2("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"), 26)

print(problem2(readfile("data/day-06.txt")))
