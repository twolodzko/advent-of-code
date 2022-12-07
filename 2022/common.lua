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

function merge(t1, t2, check)
    local out = {}
    for k, v in pairs(t1) do
        out[k] = v
    end
    for k, v in pairs(t2) do
        if check and t1[k] then
            error(string.format("duplicated key: %s", k))
        end
        out[k] = v
    end
    return out
end
