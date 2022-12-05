function readfile(filename)
    local f = assert(io.open(filename, "r"))
    local str = f:read("*all")
    f:close()
    return str
end

function lines(input)
    return string.gmatch(input, "[^\n]*")
end

function chars(input)
    return string.gmatch(input, ".")
end

function Set(list)
    local set = {}
    for _, l in ipairs(list) do
        set[l] = true
    end
    return set
end

function arrayfromstring(str)
    local arr = {}
    for ch in string.gmatch(str, ".") do
        table.insert(arr, ch)
    end
    return arr
end

function reverse(arr)
    local reversed = {}
    for i = #arr, 1, -1 do
        table.insert(reversed, arr[i])
    end
    return reversed
end
