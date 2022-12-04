function readfile(filename)
    local f = assert(io.open(filename, "r"))
    local str = f:read("*all")
    f:close()
    return str
end

function lines(input)
    return string.gmatch(input, "[^\n]*")
end
