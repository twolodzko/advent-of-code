require "common"

local function split(list)
    return list[1], { table.unpack(list, 2, #list) }
end

do
    local h, t = split({})
    assert(h == nil)
    assert(#t == 0)

    local h, t = split({ 1 })
    assert(h == 1)
    assert(#t == 0)

    local h, t = split({ 1, 2, 3 })
    assert(h == 1)
    assert(#t == 2)
end

local function compare(left, right)
    if type(left) == "number" and type(right) == "number" then
        if left < right then
            return true
        elseif left > right then
            return false
        else
            return nil
        end
    elseif type(left) == "table" and type(right) == "table" then
        if #left == 0 and #right == 0 then
            return nil
        elseif #left == 0 then
            return true
        elseif #right == 0 then
            return false
        end

        local x, y
        x, left = split(left)
        y, right = split(right)

        local result = compare(x, y)
        if result ~= nil then
            return result
        end

        return compare(left, right)
    elseif type(left) == "number" then
        return compare({ left }, right)
    elseif type(right) == "number" then
        return compare(left, { right })
    else
        error("somethign went wrong")
    end
end

local function parse(input)
    if string.match(input, "[^%d,%[%]]") then
        error(string.format("invalid input: %s", input))
    end
    input = input
        :gsub("%[", "{")
        :gsub("%]", "}")
    local env = {}
    assert(load("result = " .. input, input, "t", env))()
    return env.result
end

function problem1(input)
    local total = 0
    local i = 1
    local left, right
    local index = 1
    for line in lines(input) do
        if i % 3 == 1 then
            left = parse(line)
        elseif i % 3 == 2 then
            right = parse(line)
            if compare(left, right) ~= false then
                total = total + index
            end
            index = index + 1
        else
            left, right = nil, nil
        end
        i = i + 1
    end
    return total
end

assert(problem1(readfile("data/day-13-example.txt")) == 13)
print(problem1(readfile("data/day-13.txt")))

local function insertsort(list, comp)
    for i = 1, #list do
        local j = i
        while j > 1 and comp(list[j], list[j - 1]) do
            list[j], list[j - 1] = list[j - 1], list[j]
            j = j - 1
        end
    end
end

do
    local list = {}
    insertsort(list, function(a, b)
        return a < b
    end)
    assert(equal(list, {}))

    local list = { 1 }
    insertsort(list, function(a, b)
        return a < b
    end)
    assert(equal(list, { 1 }))

    local list = { 3, 1, 2 }
    insertsort(list, function(a, b)
        return a < b
    end)
    assert(equal(list, { 1, 2, 3 }))
end

function problem2(input)
    local packets = { { { 2 } }, { { 6 } } }
    for line in lines(input) do
        if line ~= "" and line ~= "\n" then
            local packet = parse(line)
            table.insert(packets, packet)
        end
    end

    -- table.sort won't work, so I needed a replacement
    insertsort(packets, function(a, b)
        return compare(a, b) ~= false
    end)

    local decoderkey = 1
    for index, packet in ipairs(packets) do
        if type(packet) == "table"
            and #packet == 1
            and type(packet[1]) == "table"
            and #packet[1] == 1
            and (packet[1][1] == 2 or packet[1][1] == 6)
        then
            decoderkey = decoderkey * index
        end
    end
    return decoderkey
end

assert(problem2(readfile("data/day-13-example.txt")) == 140)
print(problem2(readfile("data/day-13.txt")))
